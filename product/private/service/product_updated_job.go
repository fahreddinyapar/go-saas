package service

import (
	"context"
	"fmt"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-saas/kit/pkg/job"
	stripe2 "github.com/go-saas/kit/pkg/stripe"
	"github.com/go-saas/kit/product/private/biz"
	"github.com/go-saas/saas"
	"github.com/hibiken/asynq"
	"github.com/samber/lo"
	"github.com/stripe/stripe-go/v76"
	stripeclient "github.com/stripe/stripe-go/v76/client"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
	"time"
)

const (
	JobTypeProductUpdated = "product" + ":" + "product_updated"
)

func NewProductUpdatedTask(prams *biz.ProductUpdatedJobPram) (*asynq.Task, error) {
	payload, err := protojson.Marshal(prams)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(JobTypeProductUpdated, payload, asynq.Queue(string(biz.ConnName))), nil
}

func NewProductUpdatedTaskHandler(repo biz.ProductRepo, client *stripeclient.API) *job.Handler {
	return job.NewHandlerFunc(JobTypeProductUpdated, func(ctx context.Context, t *asynq.Task) error {
		msg := &biz.ProductUpdatedJobPram{}
		if err := protojson.Unmarshal(t.Payload(), msg); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}
		//change to product tenant
		ctx = saas.NewCurrentTenant(ctx, msg.TenantId, "")
		product, err := repo.Get(ctx, msg.ProductId)
		if err != nil {
			return err
		}
		if product == nil {
			return fmt.Errorf("can not find %s: %w", msg.ProductId, asynq.SkipRetry)
		}
		if product.Version.String != msg.ProductVersion {
			return fmt.Errorf("product:%s version mismatch, should be:%s, got:%s: %w", msg.ProductId, msg.ProductVersion, product.Version.String, asynq.SkipRetry)
		}
		group, ctx := errgroup.WithContext(ctx)
		//sync with stripe
		group.Go(func() error {
			return syncWithStripe(ctx, client, repo, product, msg.ProductVersion)
		})
		return group.Wait()

	})
}

func syncWithStripe(ctx context.Context, client *stripeclient.API, repo biz.ProductRepo, product *biz.Product, version string) error {
	if client == nil {
		return nil
	}
	links, err := repo.GetSyncLinks(ctx, product)
	if err != nil {
		return err
	}
	stripeInfo, find := lo.Find(links, func(link biz.ProductSyncLink) bool {
		return string(biz.ProductManageProviderStripe) == link.ProviderName
	})
	var stripeProductId string
	if !find {
		// create stripe object
		params := mapBizProduct2Stripe(product)
		stripeProduct, err := client.Products.New(params)
		if err != nil {
			return err
		}
		stripeProductId = stripeProduct.ID
		t := time.Now()
		err = repo.UpdateSyncLink(ctx, product, &biz.ProductSyncLink{
			ProviderName: string(biz.ProductManageProviderStripe),
			ProviderId:   stripeProductId,
			LastSyncTime: &t,
		})
		if err != nil {
			return err
		}
	} else {
		stripeProductId = stripeInfo.ProviderId
		//update product if needed
		stripeProduct, err := client.Products.Get(stripeProductId, &stripe.ProductParams{})
		if stripeProduct.Metadata["version"] != version {
			klog.Infof("product_id:%s version:%s same with stipe, skip updates", product.ID.String(), version)
			return nil
		}
		params := mapBizProduct2Stripe(product)
		_, err = client.Products.Update(stripeProductId, params)
		if err != nil {
			return err
		}
	}

	// update price
	allKeys := lo.Map(product.Prices, func(t biz.Price, _ int) string {
		return t.ID.String()
	})
	priceIter := client.Prices.List(&stripe.PriceListParams{Product: &stripeProductId})
	if priceIter.Err() != nil {
		return priceIter.Err()
	}
	var stripePrices []*stripe.Price
	for priceIter.Next() {
		if priceIter.Err() != nil {
			return priceIter.Err()
		}
		stripePrices = append(stripePrices, priceIter.Price())
	}
	//delete  prices
	toDeactivate := lo.Filter(stripePrices, func(price *stripe.Price, _ int) bool {
		return !lo.Contains(allKeys, price.LookupKey)
	})
	for _, price := range toDeactivate {
		active := false
		_, err = client.Prices.Update(price.ID, &stripe.PriceParams{Active: &active})
		if err != nil {
			return err
		}
	}
	for _, price := range product.Prices {
		stripePrice, ok := lo.Find(stripePrices, func(p *stripe.Price) bool {
			return p.LookupKey == price.ID.String()
		})
		params := mapBizPrice2Stripe(stripeProductId, &price)
		if !ok {
			//create
			_, err = client.Prices.New(params)
			if err != nil {
				return err
			}
		} else {
			_, err = client.Prices.Update(stripePrice.ID, params)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func mapBizProduct2Stripe(product *biz.Product) *stripe.ProductParams {
	return &stripe.ProductParams{
		Active:      stripe.Bool(product.Active),
		Name:        stripe2.String(product.Title),
		Description: stripe2.String(product.Desc),
		Shippable:   stripe.Bool(product.NeedShipping),
		Metadata:    map[string]string{"version": product.Version.String},
		//TODO type
	}
}

func mapBizPrice2Stripe(stripeProductId string, price *biz.Price) *stripe.PriceParams {
	r := &stripe.PriceParams{
		BillingScheme: stripe2.String(string(price.BillingScheme)),
		Currency:      stripe2.String(strings.ToLower(price.CurrencyCode)),
		LookupKey:     stripe2.String(price.ID.String()),
		Product:       stripe2.String(stripeProductId),
		Tiers:         nil,
		TiersMode:     stripe2.String(string(price.TiersMode)),
		TransformQuantity: &stripe.PriceTransformQuantityParams{
			DivideBy: stripe.Int64(price.TransformQuantity.DivideBy),
			Round:    stripe2.String(string(price.TransformQuantity.Round)),
		},
		UnitAmount: stripe.Int64(price.DefaultAmount),
	}
	r.CurrencyOptions = lo.SliceToMap(price.CurrencyOptions, func(t biz.PriceCurrencyOption) (string, *stripe.PriceCurrencyOptionsParams) {
		return strings.ToLower(t.CurrencyCode), &stripe.PriceCurrencyOptionsParams{
			Tiers: lo.Map(t.Tiers, func(t biz.PriceCurrencyOptionTier, _ int) *stripe.PriceCurrencyOptionsTierParams {
				return &stripe.PriceCurrencyOptionsTierParams{
					FlatAmount: stripe.Int64(t.FlatAmount),
					UnitAmount: stripe.Int64(t.UnitAmount),
					UpTo:       stripe.Int64(t.UpTo),
				}
			}),
			UnitAmount: stripe.Int64(t.DefaultAmount),
		}
	})
	if price.Recurring != nil {
		r.Recurring = &stripe.PriceRecurringParams{
			AggregateUsage:  stripe2.String(string(price.Recurring.AggregateUsage)),
			Interval:        stripe2.String(string(price.Recurring.Interval)),
			IntervalCount:   stripe.Int64(price.Recurring.IntervalCount),
			TrialPeriodDays: stripe.Int64(price.Recurring.TrialPeriodDays),
			UsageType:       stripe2.String(string(price.Recurring.UsageType)),
		}
	}
	r.Tiers = lo.Map(price.Tiers, func(t biz.PriceTier, _ int) *stripe.PriceTierParams {
		return &stripe.PriceTierParams{
			FlatAmount: stripe.Int64(t.FlatAmount),
			UnitAmount: stripe.Int64(t.UnitAmount),
			UpTo:       stripe.Int64(t.UpTo),
		}
	})

	return r
}