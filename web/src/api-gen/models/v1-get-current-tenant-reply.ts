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

import { V1TenantInfo } from './v1-tenant-info';

/**
 *
 * @export
 * @interface V1GetCurrentTenantReply
 */
export interface V1GetCurrentTenantReply {
  /**
   *
   * @type {V1TenantInfo}
   * @memberof V1GetCurrentTenantReply
   */
  tenant?: V1TenantInfo;
  /**
   *
   * @type {boolean}
   * @memberof V1GetCurrentTenantReply
   */
  isHost?: boolean;
}