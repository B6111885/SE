/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDepartmentEdges,
    EntDepartmentEdgesFromJSON,
    EntDepartmentEdgesFromJSONTyped,
    EntDepartmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDepartment
 */
export interface EntDepartment {
    /**
     * DepartmentName holds the value of the "Department_Name" field.
     * @type {string}
     * @memberof EntDepartment
     */
    departmentName?: string;
    /**
     * 
     * @type {EntDepartmentEdges}
     * @memberof EntDepartment
     */
    edges?: EntDepartmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDepartment
     */
    id?: number;
}

export function EntDepartmentFromJSON(json: any): EntDepartment {
    return EntDepartmentFromJSONTyped(json, false);
}

export function EntDepartmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDepartment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'departmentName': !exists(json, 'Department_Name') ? undefined : json['Department_Name'],
        'edges': !exists(json, 'edges') ? undefined : EntDepartmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDepartmentToJSON(value?: EntDepartment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Department_Name': value.departmentName,
        'edges': EntDepartmentEdgesToJSON(value.edges),
        'id': value.id,
    };
}


