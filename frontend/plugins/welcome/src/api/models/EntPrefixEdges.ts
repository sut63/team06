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
    EntPatientDetail,
    EntPatientDetailFromJSON,
    EntPatientDetailFromJSONTyped,
    EntPatientDetailToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrefixEdges
 */
export interface EntPrefixEdges {
    /**
     * PatientDetails holds the value of the patient_details edge.
     * @type {Array<EntPatientDetail>}
     * @memberof EntPrefixEdges
     */
    patientDetails?: Array<EntPatientDetail>;
}

export function EntPrefixEdgesFromJSON(json: any): EntPrefixEdges {
    return EntPrefixEdgesFromJSONTyped(json, false);
}

export function EntPrefixEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrefixEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patientDetails': !exists(json, 'patientDetails') ? undefined : ((json['patientDetails'] as Array<any>).map(EntPatientDetailFromJSON)),
    };
}

export function EntPrefixEdgesToJSON(value?: EntPrefixEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patientDetails': value.patientDetails === undefined ? undefined : ((value.patientDetails as Array<any>).map(EntPatientDetailToJSON)),
    };
}


