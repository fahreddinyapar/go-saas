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

import { V1UpdateRole } from './v1-update-role';

/**
 *
 * @export
 * @interface V1UpdateRoleRequest
 */
export interface V1UpdateRoleRequest {
  /**
   *
   * @type {V1UpdateRole}
   * @memberof V1UpdateRoleRequest
   */
  role?: V1UpdateRole;
  /**
   *
   * @type {string}
   * @memberof V1UpdateRoleRequest
   */
  updateMask?: string;
}