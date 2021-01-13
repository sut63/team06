/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntTriageResult,
    EntTriageResultFromJSON,
    EntTriageResultFromJSONTyped,
    EntTriageResultToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDepartmentEdges
 */
export interface EntDepartmentEdges {
    /**
     * DepartmentToTriageResult holds the value of the departmentToTriageResult edge.
     * @type {Array<EntTriageResult>}
     * @memberof EntDepartmentEdges
     */
    departmentToTriageResult?: Array<EntTriageResult>;
}

export function EntDepartmentEdgesFromJSON(json: any): EntDepartmentEdges {
    return EntDepartmentEdgesFromJSONTyped(json, false);
}

export function EntDepartmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDepartmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'departmentToTriageResult': !exists(json, 'departmentToTriageResult') ? undefined : ((json['departmentToTriageResult'] as Array<any>).map(EntTriageResultFromJSON)),
    };
}

export function EntDepartmentEdgesToJSON(value?: EntDepartmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'departmentToTriageResult': value.departmentToTriageResult === undefined ? undefined : ((value.departmentToTriageResult as Array<any>).map(EntTriageResultToJSON)),
    };
}


