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
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntUrgencyLevel,
    EntUrgencyLevelFromJSON,
    EntUrgencyLevelFromJSONTyped,
    EntUrgencyLevelToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTriageResultEdges
 */
export interface EntTriageResultEdges {
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntTriageResultEdges
     */
    triageResultToDepartment?: EntDepartment;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntTriageResultEdges
     */
    triageResultToNurse?: EntNurse;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntTriageResultEdges
     */
    triageResultToPatient?: EntPatient;
    /**
     * 
     * @type {EntUrgencyLevel}
     * @memberof EntTriageResultEdges
     */
    triageResultToUrgencyLevel?: EntUrgencyLevel;
}

export function EntTriageResultEdgesFromJSON(json: any): EntTriageResultEdges {
    return EntTriageResultEdgesFromJSONTyped(json, false);
}

export function EntTriageResultEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTriageResultEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'triageResultToDepartment': !exists(json, 'triageResultToDepartment') ? undefined : EntDepartmentFromJSON(json['triageResultToDepartment']),
        'triageResultToNurse': !exists(json, 'triageResultToNurse') ? undefined : EntNurseFromJSON(json['triageResultToNurse']),
        'triageResultToPatient': !exists(json, 'triageResultToPatient') ? undefined : EntPatientFromJSON(json['triageResultToPatient']),
        'triageResultToUrgencyLevel': !exists(json, 'triageResultToUrgencyLevel') ? undefined : EntUrgencyLevelFromJSON(json['triageResultToUrgencyLevel']),
    };
}

export function EntTriageResultEdgesToJSON(value?: EntTriageResultEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'triageResultToDepartment': EntDepartmentToJSON(value.triageResultToDepartment),
        'triageResultToNurse': EntNurseToJSON(value.triageResultToNurse),
        'triageResultToPatient': EntPatientToJSON(value.triageResultToPatient),
        'triageResultToUrgencyLevel': EntUrgencyLevelToJSON(value.triageResultToUrgencyLevel),
    };
}


