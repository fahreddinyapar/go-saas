package apisix

import (
	"bytes"
	"encoding/json"
	"errors"
	klog "github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
	"strings"
)

var (
	ErrOption = errors.New("apisix endpoint and apiKey are required")
)

type AdminClient struct {
	endpoint string
	apiKey   string
	client   *http.Client
}

func NewAdminClient(endpoint string, apiKey string) (*AdminClient, error) {
	if len(endpoint) == 0 || len(apiKey) == 0 {
		return nil, ErrOption
	}
	if strings.HasSuffix(endpoint, "/") {
		endpoint = strings.TrimSuffix(endpoint, "/")
	}
	return &AdminClient{
		endpoint: endpoint,
		apiKey:   apiKey,
		client:   &http.Client{},
	}, nil
}

func (a *AdminClient) PutUpstream(id string, upstream *Upstream) error {
	j, err := json.Marshal(upstream)
	if err != nil {
		return err
	}
	klog.Infof("[apisix]  update upstream %s : %s", id, j)
	req, err := http.NewRequest(http.MethodPut, a.endpoint+"/apisix/admin/upstreams/"+id, bytes.NewReader(j))
	if err != nil {
		return err
	}
	_, err = a.do(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminClient) PutUpstreamStruct(id string, upstream *structpb.Struct) error {
	j, err := protojson.Marshal(upstream)
	if err != nil {
		return err
	}
	klog.Infof("[apisix]  update upstream %s : %s", id, j)
	req, err := http.NewRequest(http.MethodPut, a.endpoint+"/apisix/admin/upstreams/"+id, bytes.NewReader(j))
	if err != nil {
		return err
	}
	_, err = a.do(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminClient) PutRoute(id string, route *structpb.Struct) error {
	j, err := protojson.Marshal(route)
	if err != nil {
		return err
	}
	klog.Infof("[apisix]  update route %s : %s", id, j)
	req, err := http.NewRequest(http.MethodPut, a.endpoint+"/apisix/admin/routes/"+id, bytes.NewReader(j))
	if err != nil {
		return err
	}
	_, err = a.do(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminClient) PutGlobalRules(id string, route *structpb.Struct) error {
	j, err := protojson.Marshal(route)
	if err != nil {
		return err
	}
	klog.Infof("[apisix]  update global rules %s : %s", id, j)
	req, err := http.NewRequest(http.MethodPut, a.endpoint+"/apisix/admin/global_rules/"+id, bytes.NewReader(j))
	if err != nil {
		return err
	}
	_, err = a.do(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", a.apiKey)
	return a.client.Do(req)
}