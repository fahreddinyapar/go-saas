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

import { V1Effect } from './v1-effect';

/**
 *
 * @export
 * @interface V1UpdateSubjectPermissionAcl
 */
export interface V1UpdateSubjectPermissionAcl {
  /**
   *
   * @type {string}
   * @memberof V1UpdateSubjectPermissionAcl
   */
  namespace: string;
  /**
   *
   * @type {string}
   * @memberof V1UpdateSubjectPermissionAcl
   */
  resource: string;
  /**
   *
   * @type {string}
   * @memberof V1UpdateSubjectPermissionAcl
   */
  action: string;
  /**
   *
   * @type {V1Effect}
   * @memberof V1UpdateSubjectPermissionAcl
   */
  effect?: V1Effect;
  /**
   *
   * @type {string}
   * @memberof V1UpdateSubjectPermissionAcl
   */
  tenantId?: string;
}