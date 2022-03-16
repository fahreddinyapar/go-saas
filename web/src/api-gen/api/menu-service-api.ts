/* tslint:disable */
/* eslint-disable */
/**
 * Saas Service
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import {
  DUMMY_BASE_URL,
  assertParamExists,
  setApiKeyToObject,
  setBasicAuthToObject,
  setBearerAuthToObject,
  setOAuthToObject,
  setSearchParams,
  serializeDataIfNeeded,
  toPathString,
  createRequestFunction,
} from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { RpcStatus } from '../models';
// @ts-ignore
import { V1CreateMenuRequest } from '../models';
// @ts-ignore
import { V1DeleteMenuReply } from '../models';
// @ts-ignore
import { V1GetAvailableMenusReply } from '../models';
// @ts-ignore
import { V1ListMenuReply } from '../models';
// @ts-ignore
import { V1ListMenuRequest } from '../models';
// @ts-ignore
import { V1Menu } from '../models';
// @ts-ignore
import { V1UpdateMenuRequest } from '../models';
/**
 * MenuServiceApi - axios parameter creator
 * @export
 */
export const MenuServiceApiAxiosParamCreator = function (configuration?: Configuration) {
  return {
    /**
     *
     * @param {V1CreateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceCreateMenu: async (
      body: V1CreateMenuRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('menuServiceCreateMenu', 'body', body);
      const localVarPath = `/v1/sys/menu`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceDeleteMenu: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('menuServiceDeleteMenu', 'id', id);
      const localVarPath = `/v1/sys/menu/{id}`.replace(`{${'id'}}`, encodeURIComponent(String(id)));
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceGetAvailableMenus: async (
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      const localVarPath = `/v1/sys/menus/available`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceGetMenu: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('menuServiceGetMenu', 'id', id);
      const localVarPath = `/v1/sys/menu/{id}`.replace(`{${'id'}}`, encodeURIComponent(String(id)));
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {Array<string>} [filterParentIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceListMenu: async (
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterParentIn?: Array<string>,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      const localVarPath = `/v1/sys/menus`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      if (pageOffset !== undefined) {
        localVarQueryParameter['pageOffset'] = pageOffset;
      }

      if (pageSize !== undefined) {
        localVarQueryParameter['pageSize'] = pageSize;
      }

      if (search !== undefined) {
        localVarQueryParameter['search'] = search;
      }

      if (sort) {
        localVarQueryParameter['sort'] = sort;
      }

      if (fields !== undefined) {
        localVarQueryParameter['fields'] = fields;
      }

      if (filterIdIn) {
        localVarQueryParameter['filter.idIn'] = filterIdIn;
      }

      if (filterNameIn) {
        localVarQueryParameter['filter.nameIn'] = filterNameIn;
      }

      if (filterParentIn) {
        localVarQueryParameter['filter.parentIn'] = filterParentIn;
      }

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {V1ListMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceListMenu2: async (
      body: V1ListMenuRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('menuServiceListMenu2', 'body', body);
      const localVarPath = `/v1/sys/menu/list`;
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceUpdateMenu: async (
      menuId: string,
      body: V1UpdateMenuRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'menuId' is not null or undefined
      assertParamExists('menuServiceUpdateMenu', 'menuId', menuId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('menuServiceUpdateMenu', 'body', body);
      const localVarPath = `/v1/sys/menu/{menu.id}`.replace(
        `{${'menu.id'}}`,
        encodeURIComponent(String(menuId)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceUpdateMenu2: async (
      menuId: string,
      body: V1UpdateMenuRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'menuId' is not null or undefined
      assertParamExists('menuServiceUpdateMenu2', 'menuId', menuId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('menuServiceUpdateMenu2', 'body', body);
      const localVarPath = `/v1/sys/menu/{menu.id}`.replace(
        `{${'menu.id'}}`,
        encodeURIComponent(String(menuId)),
      );
      // use dummy base URL string because the URL constructor only accepts absolute URLs.
      const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
      let baseOptions;
      if (configuration) {
        baseOptions = configuration.baseOptions;
      }

      const localVarRequestOptions = { method: 'PATCH', ...baseOptions, ...options };
      const localVarHeaderParameter = {} as any;
      const localVarQueryParameter = {} as any;

      localVarHeaderParameter['Content-Type'] = 'application/json';

      setSearchParams(localVarUrlObj, localVarQueryParameter);
      let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
      localVarRequestOptions.headers = {
        ...localVarHeaderParameter,
        ...headersFromBaseOptions,
        ...options.headers,
      };
      localVarRequestOptions.data = serializeDataIfNeeded(
        body,
        localVarRequestOptions,
        configuration,
      );

      return {
        url: toPathString(localVarUrlObj),
        options: localVarRequestOptions,
      };
    },
  };
};

/**
 * MenuServiceApi - functional programming interface
 * @export
 */
export const MenuServiceApiFp = function (configuration?: Configuration) {
  const localVarAxiosParamCreator = MenuServiceApiAxiosParamCreator(configuration);
  return {
    /**
     *
     * @param {V1CreateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceCreateMenu(
      body: V1CreateMenuRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1Menu>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceCreateMenu(
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceDeleteMenu(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1DeleteMenuReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceDeleteMenu(id, options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceGetAvailableMenus(
      options?: AxiosRequestConfig,
    ): Promise<
      (axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1GetAvailableMenusReply>
    > {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceGetAvailableMenus(
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceGetMenu(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1Menu>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceGetMenu(id, options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {Array<string>} [filterParentIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceListMenu(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterParentIn?: Array<string>,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListMenuReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceListMenu(
        pageOffset,
        pageSize,
        search,
        sort,
        fields,
        filterIdIn,
        filterNameIn,
        filterParentIn,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {V1ListMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceListMenu2(
      body: V1ListMenuRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListMenuReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceListMenu2(body, options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceUpdateMenu(
      menuId: string,
      body: V1UpdateMenuRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1Menu>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceUpdateMenu(
        menuId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async menuServiceUpdateMenu2(
      menuId: string,
      body: V1UpdateMenuRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1Menu>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.menuServiceUpdateMenu2(
        menuId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
  };
};

/**
 * MenuServiceApi - factory interface
 * @export
 */
export const MenuServiceApiFactory = function (
  configuration?: Configuration,
  basePath?: string,
  axios?: AxiosInstance,
) {
  const localVarFp = MenuServiceApiFp(configuration);
  return {
    /**
     *
     * @param {V1CreateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceCreateMenu(body: V1CreateMenuRequest, options?: any): AxiosPromise<V1Menu> {
      return localVarFp
        .menuServiceCreateMenu(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceDeleteMenu(id: string, options?: any): AxiosPromise<V1DeleteMenuReply> {
      return localVarFp
        .menuServiceDeleteMenu(id, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceGetAvailableMenus(options?: any): AxiosPromise<V1GetAvailableMenusReply> {
      return localVarFp
        .menuServiceGetAvailableMenus(options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceGetMenu(id: string, options?: any): AxiosPromise<V1Menu> {
      return localVarFp.menuServiceGetMenu(id, options).then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterNameIn]
     * @param {Array<string>} [filterParentIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceListMenu(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterNameIn?: Array<string>,
      filterParentIn?: Array<string>,
      options?: any,
    ): AxiosPromise<V1ListMenuReply> {
      return localVarFp
        .menuServiceListMenu(
          pageOffset,
          pageSize,
          search,
          sort,
          fields,
          filterIdIn,
          filterNameIn,
          filterParentIn,
          options,
        )
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {V1ListMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceListMenu2(body: V1ListMenuRequest, options?: any): AxiosPromise<V1ListMenuReply> {
      return localVarFp
        .menuServiceListMenu2(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceUpdateMenu(
      menuId: string,
      body: V1UpdateMenuRequest,
      options?: any,
    ): AxiosPromise<V1Menu> {
      return localVarFp
        .menuServiceUpdateMenu(menuId, body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @param {string} menuId
     * @param {V1UpdateMenuRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    menuServiceUpdateMenu2(
      menuId: string,
      body: V1UpdateMenuRequest,
      options?: any,
    ): AxiosPromise<V1Menu> {
      return localVarFp
        .menuServiceUpdateMenu2(menuId, body, options)
        .then((request) => request(axios, basePath));
    },
  };
};

/**
 * Request parameters for menuServiceCreateMenu operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceCreateMenuRequest
 */
export interface MenuServiceApiMenuServiceCreateMenuRequest {
  /**
   *
   * @type {V1CreateMenuRequest}
   * @memberof MenuServiceApiMenuServiceCreateMenu
   */
  readonly body: V1CreateMenuRequest;
}

/**
 * Request parameters for menuServiceDeleteMenu operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceDeleteMenuRequest
 */
export interface MenuServiceApiMenuServiceDeleteMenuRequest {
  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceDeleteMenu
   */
  readonly id: string;
}

/**
 * Request parameters for menuServiceGetMenu operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceGetMenuRequest
 */
export interface MenuServiceApiMenuServiceGetMenuRequest {
  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceGetMenu
   */
  readonly id: string;
}

/**
 * Request parameters for menuServiceListMenu operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceListMenuRequest
 */
export interface MenuServiceApiMenuServiceListMenuRequest {
  /**
   *
   * @type {number}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly pageOffset?: number;

  /**
   *
   * @type {number}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly pageSize?: number;

  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly search?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly sort?: Array<string>;

  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly fields?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly filterIdIn?: Array<string>;

  /**
   *
   * @type {Array<string>}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly filterNameIn?: Array<string>;

  /**
   *
   * @type {Array<string>}
   * @memberof MenuServiceApiMenuServiceListMenu
   */
  readonly filterParentIn?: Array<string>;
}

/**
 * Request parameters for menuServiceListMenu2 operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceListMenu2Request
 */
export interface MenuServiceApiMenuServiceListMenu2Request {
  /**
   *
   * @type {V1ListMenuRequest}
   * @memberof MenuServiceApiMenuServiceListMenu2
   */
  readonly body: V1ListMenuRequest;
}

/**
 * Request parameters for menuServiceUpdateMenu operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceUpdateMenuRequest
 */
export interface MenuServiceApiMenuServiceUpdateMenuRequest {
  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceUpdateMenu
   */
  readonly menuId: string;

  /**
   *
   * @type {V1UpdateMenuRequest}
   * @memberof MenuServiceApiMenuServiceUpdateMenu
   */
  readonly body: V1UpdateMenuRequest;
}

/**
 * Request parameters for menuServiceUpdateMenu2 operation in MenuServiceApi.
 * @export
 * @interface MenuServiceApiMenuServiceUpdateMenu2Request
 */
export interface MenuServiceApiMenuServiceUpdateMenu2Request {
  /**
   *
   * @type {string}
   * @memberof MenuServiceApiMenuServiceUpdateMenu2
   */
  readonly menuId: string;

  /**
   *
   * @type {V1UpdateMenuRequest}
   * @memberof MenuServiceApiMenuServiceUpdateMenu2
   */
  readonly body: V1UpdateMenuRequest;
}

/**
 * MenuServiceApi - object-oriented interface
 * @export
 * @class MenuServiceApi
 * @extends {BaseAPI}
 */
export class MenuServiceApi extends BaseAPI {
  /**
   *
   * @param {MenuServiceApiMenuServiceCreateMenuRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceCreateMenu(
    requestParameters: MenuServiceApiMenuServiceCreateMenuRequest,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceCreateMenu(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceDeleteMenuRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceDeleteMenu(
    requestParameters: MenuServiceApiMenuServiceDeleteMenuRequest,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceDeleteMenu(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceGetAvailableMenus(options?: AxiosRequestConfig) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceGetAvailableMenus(options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceGetMenuRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceGetMenu(
    requestParameters: MenuServiceApiMenuServiceGetMenuRequest,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceGetMenu(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceListMenuRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceListMenu(
    requestParameters: MenuServiceApiMenuServiceListMenuRequest = {},
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceListMenu(
        requestParameters.pageOffset,
        requestParameters.pageSize,
        requestParameters.search,
        requestParameters.sort,
        requestParameters.fields,
        requestParameters.filterIdIn,
        requestParameters.filterNameIn,
        requestParameters.filterParentIn,
        options,
      )
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceListMenu2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceListMenu2(
    requestParameters: MenuServiceApiMenuServiceListMenu2Request,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceListMenu2(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceUpdateMenuRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceUpdateMenu(
    requestParameters: MenuServiceApiMenuServiceUpdateMenuRequest,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceUpdateMenu(requestParameters.menuId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @param {MenuServiceApiMenuServiceUpdateMenu2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof MenuServiceApi
   */
  public menuServiceUpdateMenu2(
    requestParameters: MenuServiceApiMenuServiceUpdateMenu2Request,
    options?: AxiosRequestConfig,
  ) {
    return MenuServiceApiFp(this.configuration)
      .menuServiceUpdateMenu2(requestParameters.menuId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }
}