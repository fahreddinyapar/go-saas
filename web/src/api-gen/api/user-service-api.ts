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
import { V1CreateUserRequest } from '../models';
// @ts-ignore
import { V1GetUserRoleReply } from '../models';
// @ts-ignore
import { V1ListUsersRequest } from '../models';
// @ts-ignore
import { V1ListUsersResponse } from '../models';
// @ts-ignore
import { V1UpdateUserRequest } from '../models';
// @ts-ignore
import { V1User } from '../models';
/**
 * UserServiceApi - axios parameter creator
 * @export
 */
export const UserServiceApiAxiosParamCreator = function (configuration?: Configuration) {
  return {
    /**
     *
     * @summary authz: user.user,*,create
     * @param {V1CreateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceCreateUser: async (
      body: V1CreateUserRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('userServiceCreateUser', 'body', body);
      const localVarPath = `/v1/user`;
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
     * @summary authz: user.user,id,delete
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceDeleteUser: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('userServiceDeleteUser', 'id', id);
      const localVarPath = `/v1/user/{id}`.replace(`{${'id'}}`, encodeURIComponent(String(id)));
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
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceGetUser: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('userServiceGetUser', 'id', id);
      const localVarPath = `/v1/user/{id}`.replace(`{${'id'}}`, encodeURIComponent(String(id)));
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
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceGetUserRoles: async (
      id: string,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'id' is not null or undefined
      assertParamExists('userServiceGetUserRoles', 'id', id);
      const localVarPath = `/v1/user/{id}/roles`.replace(
        `{${'id'}}`,
        encodeURIComponent(String(id)),
      );
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
     * @summary authz: user.user,*,list
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterGenderIn]
     * @param {string} [filterBirthdayGte]
     * @param {string} [filterBirthdayLte]
     * @param {Array<string>} [filterRolesIdIn]
     * @param {Array<string>} [filterRolesNameIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceListUsers: async (
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterGenderIn?: Array<string>,
      filterBirthdayGte?: string,
      filterBirthdayLte?: string,
      filterRolesIdIn?: Array<string>,
      filterRolesNameIn?: Array<string>,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      const localVarPath = `/v1/users`;
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

      if (filterGenderIn) {
        localVarQueryParameter['filter.genderIn'] = filterGenderIn;
      }

      if (filterBirthdayGte !== undefined) {
        localVarQueryParameter['filter.birthdayGte'] =
          (filterBirthdayGte as any) instanceof Date
            ? (filterBirthdayGte as any).toISOString()
            : filterBirthdayGte;
      }

      if (filterBirthdayLte !== undefined) {
        localVarQueryParameter['filter.birthdayLte'] =
          (filterBirthdayLte as any) instanceof Date
            ? (filterBirthdayLte as any).toISOString()
            : filterBirthdayLte;
      }

      if (filterRolesIdIn) {
        localVarQueryParameter['filter.roles.idIn'] = filterRolesIdIn;
      }

      if (filterRolesNameIn) {
        localVarQueryParameter['filter.roles.nameIn'] = filterRolesNameIn;
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
     * @summary authz: user.user,*,list
     * @param {V1ListUsersRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceListUsers2: async (
      body: V1ListUsersRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'body' is not null or undefined
      assertParamExists('userServiceListUsers2', 'body', body);
      const localVarPath = `/v1/user/list`;
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
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceUpdateUser: async (
      userId: string,
      body: V1UpdateUserRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'userId' is not null or undefined
      assertParamExists('userServiceUpdateUser', 'userId', userId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('userServiceUpdateUser', 'body', body);
      const localVarPath = `/v1/user/{user.id}`.replace(
        `{${'user.id'}}`,
        encodeURIComponent(String(userId)),
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
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceUpdateUser2: async (
      userId: string,
      body: V1UpdateUserRequest,
      options: AxiosRequestConfig = {},
    ): Promise<RequestArgs> => {
      // verify required parameter 'userId' is not null or undefined
      assertParamExists('userServiceUpdateUser2', 'userId', userId);
      // verify required parameter 'body' is not null or undefined
      assertParamExists('userServiceUpdateUser2', 'body', body);
      const localVarPath = `/v1/user/{user.id}`.replace(
        `{${'user.id'}}`,
        encodeURIComponent(String(userId)),
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
 * UserServiceApi - functional programming interface
 * @export
 */
export const UserServiceApiFp = function (configuration?: Configuration) {
  const localVarAxiosParamCreator = UserServiceApiAxiosParamCreator(configuration);
  return {
    /**
     *
     * @summary authz: user.user,*,create
     * @param {V1CreateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceCreateUser(
      body: V1CreateUserRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1User>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceCreateUser(
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,id,delete
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceDeleteUser(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceDeleteUser(id, options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceGetUser(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1User>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceGetUser(id, options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceGetUserRoles(
      id: string,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1GetUserRoleReply>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceGetUserRoles(
        id,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,*,list
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterGenderIn]
     * @param {string} [filterBirthdayGte]
     * @param {string} [filterBirthdayLte]
     * @param {Array<string>} [filterRolesIdIn]
     * @param {Array<string>} [filterRolesNameIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceListUsers(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterGenderIn?: Array<string>,
      filterBirthdayGte?: string,
      filterBirthdayLte?: string,
      filterRolesIdIn?: Array<string>,
      filterRolesNameIn?: Array<string>,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListUsersResponse>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceListUsers(
        pageOffset,
        pageSize,
        search,
        sort,
        fields,
        filterIdIn,
        filterGenderIn,
        filterBirthdayGte,
        filterBirthdayLte,
        filterRolesIdIn,
        filterRolesNameIn,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,*,list
     * @param {V1ListUsersRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceListUsers2(
      body: V1ListUsersRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListUsersResponse>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceListUsers2(
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceUpdateUser(
      userId: string,
      body: V1UpdateUserRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1User>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceUpdateUser(
        userId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
    /**
     *
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    async userServiceUpdateUser2(
      userId: string,
      body: V1UpdateUserRequest,
      options?: AxiosRequestConfig,
    ): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1User>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.userServiceUpdateUser2(
        userId,
        body,
        options,
      );
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
    },
  };
};

/**
 * UserServiceApi - factory interface
 * @export
 */
export const UserServiceApiFactory = function (
  configuration?: Configuration,
  basePath?: string,
  axios?: AxiosInstance,
) {
  const localVarFp = UserServiceApiFp(configuration);
  return {
    /**
     *
     * @summary authz: user.user,*,create
     * @param {V1CreateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceCreateUser(body: V1CreateUserRequest, options?: any): AxiosPromise<V1User> {
      return localVarFp
        .userServiceCreateUser(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,id,delete
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceDeleteUser(id: string, options?: any): AxiosPromise<object> {
      return localVarFp
        .userServiceDeleteUser(id, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceGetUser(id: string, options?: any): AxiosPromise<V1User> {
      return localVarFp.userServiceGetUser(id, options).then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,id,get
     * @param {string} id
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceGetUserRoles(id: string, options?: any): AxiosPromise<V1GetUserRoleReply> {
      return localVarFp
        .userServiceGetUserRoles(id, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,*,list
     * @param {number} [pageOffset]
     * @param {number} [pageSize]
     * @param {string} [search]
     * @param {Array<string>} [sort]
     * @param {string} [fields]
     * @param {Array<string>} [filterIdIn]
     * @param {Array<string>} [filterGenderIn]
     * @param {string} [filterBirthdayGte]
     * @param {string} [filterBirthdayLte]
     * @param {Array<string>} [filterRolesIdIn]
     * @param {Array<string>} [filterRolesNameIn]
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceListUsers(
      pageOffset?: number,
      pageSize?: number,
      search?: string,
      sort?: Array<string>,
      fields?: string,
      filterIdIn?: Array<string>,
      filterGenderIn?: Array<string>,
      filterBirthdayGte?: string,
      filterBirthdayLte?: string,
      filterRolesIdIn?: Array<string>,
      filterRolesNameIn?: Array<string>,
      options?: any,
    ): AxiosPromise<V1ListUsersResponse> {
      return localVarFp
        .userServiceListUsers(
          pageOffset,
          pageSize,
          search,
          sort,
          fields,
          filterIdIn,
          filterGenderIn,
          filterBirthdayGte,
          filterBirthdayLte,
          filterRolesIdIn,
          filterRolesNameIn,
          options,
        )
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,*,list
     * @param {V1ListUsersRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceListUsers2(
      body: V1ListUsersRequest,
      options?: any,
    ): AxiosPromise<V1ListUsersResponse> {
      return localVarFp
        .userServiceListUsers2(body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceUpdateUser(
      userId: string,
      body: V1UpdateUserRequest,
      options?: any,
    ): AxiosPromise<V1User> {
      return localVarFp
        .userServiceUpdateUser(userId, body, options)
        .then((request) => request(axios, basePath));
    },
    /**
     *
     * @summary authz: user.user,id,update
     * @param {string} userId
     * @param {V1UpdateUserRequest} body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     */
    userServiceUpdateUser2(
      userId: string,
      body: V1UpdateUserRequest,
      options?: any,
    ): AxiosPromise<V1User> {
      return localVarFp
        .userServiceUpdateUser2(userId, body, options)
        .then((request) => request(axios, basePath));
    },
  };
};

/**
 * Request parameters for userServiceCreateUser operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceCreateUserRequest
 */
export interface UserServiceApiUserServiceCreateUserRequest {
  /**
   *
   * @type {V1CreateUserRequest}
   * @memberof UserServiceApiUserServiceCreateUser
   */
  readonly body: V1CreateUserRequest;
}

/**
 * Request parameters for userServiceDeleteUser operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceDeleteUserRequest
 */
export interface UserServiceApiUserServiceDeleteUserRequest {
  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceDeleteUser
   */
  readonly id: string;
}

/**
 * Request parameters for userServiceGetUser operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceGetUserRequest
 */
export interface UserServiceApiUserServiceGetUserRequest {
  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceGetUser
   */
  readonly id: string;
}

/**
 * Request parameters for userServiceGetUserRoles operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceGetUserRolesRequest
 */
export interface UserServiceApiUserServiceGetUserRolesRequest {
  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceGetUserRoles
   */
  readonly id: string;
}

/**
 * Request parameters for userServiceListUsers operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceListUsersRequest
 */
export interface UserServiceApiUserServiceListUsersRequest {
  /**
   *
   * @type {number}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly pageOffset?: number;

  /**
   *
   * @type {number}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly pageSize?: number;

  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly search?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly sort?: Array<string>;

  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly fields?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterIdIn?: Array<string>;

  /**
   *
   * @type {Array<string>}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterGenderIn?: Array<string>;

  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterBirthdayGte?: string;

  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterBirthdayLte?: string;

  /**
   *
   * @type {Array<string>}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterRolesIdIn?: Array<string>;

  /**
   *
   * @type {Array<string>}
   * @memberof UserServiceApiUserServiceListUsers
   */
  readonly filterRolesNameIn?: Array<string>;
}

/**
 * Request parameters for userServiceListUsers2 operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceListUsers2Request
 */
export interface UserServiceApiUserServiceListUsers2Request {
  /**
   *
   * @type {V1ListUsersRequest}
   * @memberof UserServiceApiUserServiceListUsers2
   */
  readonly body: V1ListUsersRequest;
}

/**
 * Request parameters for userServiceUpdateUser operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceUpdateUserRequest
 */
export interface UserServiceApiUserServiceUpdateUserRequest {
  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceUpdateUser
   */
  readonly userId: string;

  /**
   *
   * @type {V1UpdateUserRequest}
   * @memberof UserServiceApiUserServiceUpdateUser
   */
  readonly body: V1UpdateUserRequest;
}

/**
 * Request parameters for userServiceUpdateUser2 operation in UserServiceApi.
 * @export
 * @interface UserServiceApiUserServiceUpdateUser2Request
 */
export interface UserServiceApiUserServiceUpdateUser2Request {
  /**
   *
   * @type {string}
   * @memberof UserServiceApiUserServiceUpdateUser2
   */
  readonly userId: string;

  /**
   *
   * @type {V1UpdateUserRequest}
   * @memberof UserServiceApiUserServiceUpdateUser2
   */
  readonly body: V1UpdateUserRequest;
}

/**
 * UserServiceApi - object-oriented interface
 * @export
 * @class UserServiceApi
 * @extends {BaseAPI}
 */
export class UserServiceApi extends BaseAPI {
  /**
   *
   * @summary authz: user.user,*,create
   * @param {UserServiceApiUserServiceCreateUserRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceCreateUser(
    requestParameters: UserServiceApiUserServiceCreateUserRequest,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceCreateUser(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,id,delete
   * @param {UserServiceApiUserServiceDeleteUserRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceDeleteUser(
    requestParameters: UserServiceApiUserServiceDeleteUserRequest,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceDeleteUser(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,id,get
   * @param {UserServiceApiUserServiceGetUserRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceGetUser(
    requestParameters: UserServiceApiUserServiceGetUserRequest,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceGetUser(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,id,get
   * @param {UserServiceApiUserServiceGetUserRolesRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceGetUserRoles(
    requestParameters: UserServiceApiUserServiceGetUserRolesRequest,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceGetUserRoles(requestParameters.id, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,*,list
   * @param {UserServiceApiUserServiceListUsersRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceListUsers(
    requestParameters: UserServiceApiUserServiceListUsersRequest = {},
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceListUsers(
        requestParameters.pageOffset,
        requestParameters.pageSize,
        requestParameters.search,
        requestParameters.sort,
        requestParameters.fields,
        requestParameters.filterIdIn,
        requestParameters.filterGenderIn,
        requestParameters.filterBirthdayGte,
        requestParameters.filterBirthdayLte,
        requestParameters.filterRolesIdIn,
        requestParameters.filterRolesNameIn,
        options,
      )
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,*,list
   * @param {UserServiceApiUserServiceListUsers2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceListUsers2(
    requestParameters: UserServiceApiUserServiceListUsers2Request,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceListUsers2(requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,id,update
   * @param {UserServiceApiUserServiceUpdateUserRequest} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceUpdateUser(
    requestParameters: UserServiceApiUserServiceUpdateUserRequest,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceUpdateUser(requestParameters.userId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }

  /**
   *
   * @summary authz: user.user,id,update
   * @param {UserServiceApiUserServiceUpdateUser2Request} requestParameters Request parameters.
   * @param {*} [options] Override http request option.
   * @throws {RequiredError}
   * @memberof UserServiceApi
   */
  public userServiceUpdateUser2(
    requestParameters: UserServiceApiUserServiceUpdateUser2Request,
    options?: AxiosRequestConfig,
  ) {
    return UserServiceApiFp(this.configuration)
      .userServiceUpdateUser2(requestParameters.userId, requestParameters.body, options)
      .then((request) => request(this.axios, this.basePath));
  }
}
