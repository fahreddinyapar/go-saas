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

/**
 *
 * @export
 * @interface V1PermissionRequirement
 */
export interface V1PermissionRequirement {
  /**
   *
   * @type {string}
   * @memberof V1PermissionRequirement
   */
  namespace: string;
  /**
   *
   * @type {string}
   * @memberof V1PermissionRequirement
   */
  resource: string;
  /**
   *
   * @type {string}
   * @memberof V1PermissionRequirement
   */
  action: string;
}