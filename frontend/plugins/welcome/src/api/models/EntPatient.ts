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
    EntPatientEdges,
    EntPatientEdgesFromJSON,
    EntPatientEdgesFromJSONTyped,
    EntPatientEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatient
 */
export interface EntPatient {
    /**
     * AddedTime holds the value of the "added_time" field.
     * @type {string}
     * @memberof EntPatient
     */
    addedTime?: string;
    /**
     * Age holds the value of the "age" field.
     * @type {number}
     * @memberof EntPatient
     */
    age?: number;
    /**
     * DrugAllergy holds the value of the "drugAllergy" field.
     * @type {string}
     * @memberof EntPatient
     */
    drugAllergy?: string;
    /**
     * 
     * @type {EntPatientEdges}
     * @memberof EntPatient
     */
    edges?: EntPatientEdges;
    /**
     * HospitalNumber holds the value of the "hospitalNumber" field.
     * @type {string}
     * @memberof EntPatient
     */
    hospitalNumber?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatient
     */
    id?: number;
    /**
     * PatientName holds the value of the "patientName" field.
     * @type {string}
     * @memberof EntPatient
     */
    patientName?: string;
    /**
     * PersonalID holds the value of the "personalID" field.
     * @type {number}
     * @memberof EntPatient
     */
    personalID?: number;
}

export function EntPatientFromJSON(json: any): EntPatient {
    return EntPatientFromJSONTyped(json, false);
}

export function EntPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'added_time') ? undefined : json['added_time'],
        'age': !exists(json, 'age') ? undefined : json['age'],
        'drugAllergy': !exists(json, 'drugAllergy') ? undefined : json['drugAllergy'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientEdgesFromJSON(json['edges']),
        'hospitalNumber': !exists(json, 'hospitalNumber') ? undefined : json['hospitalNumber'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'patientName': !exists(json, 'patientName') ? undefined : json['patientName'],
        'personalID': !exists(json, 'personalID') ? undefined : json['personalID'],
    };
}

export function EntPatientToJSON(value?: EntPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time': value.addedTime,
        'age': value.age,
        'drugAllergy': value.drugAllergy,
        'edges': EntPatientEdgesToJSON(value.edges),
        'hospitalNumber': value.hospitalNumber,
        'id': value.id,
        'patientName': value.patientName,
        'personalID': value.personalID,
    };
}

