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
    EntRightToTreatment,
    EntRightToTreatmentFromJSON,
    EntRightToTreatmentFromJSONTyped,
    EntRightToTreatmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRightToTreatmentTypeEdges
 */
export interface EntRightToTreatmentTypeEdges {
    /**
     * Type holds the value of the type edge.
     * @type {Array<EntRightToTreatment>}
     * @memberof EntRightToTreatmentTypeEdges
     */
    type?: Array<EntRightToTreatment>;
}

export function EntRightToTreatmentTypeEdgesFromJSON(json: any): EntRightToTreatmentTypeEdges {
    return EntRightToTreatmentTypeEdgesFromJSONTyped(json, false);
}

export function EntRightToTreatmentTypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRightToTreatmentTypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': !exists(json, 'type') ? undefined : ((json['type'] as Array<any>).map(EntRightToTreatmentFromJSON)),
    };
}

export function EntRightToTreatmentTypeEdgesToJSON(value?: EntRightToTreatmentTypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': value.type === undefined ? undefined : ((value.type as Array<any>).map(EntRightToTreatmentToJSON)),
    };
}


