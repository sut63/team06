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
    EntDoctorEdges,
    EntDoctorEdgesFromJSON,
    EntDoctorEdgesFromJSONTyped,
    EntDoctorEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDoctor
 */
export interface EntDoctor {
    /**
     * DoctorName holds the value of the "doctorName" field.
     * @type {string}
     * @memberof EntDoctor
     */
    doctorName?: string;
    /**
     * DoctorPassword holds the value of the "doctorPassword" field.
     * @type {string}
     * @memberof EntDoctor
     */
    doctorPassword?: string;
    /**
     * DoctorUsername holds the value of the "doctorUsername" field.
     * @type {string}
     * @memberof EntDoctor
     */
    doctorUsername?: string;
    /**
     * 
     * @type {EntDoctorEdges}
     * @memberof EntDoctor
     */
    edges?: EntDoctorEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDoctor
     */
    id?: number;
}

export function EntDoctorFromJSON(json: any): EntDoctor {
    return EntDoctorFromJSONTyped(json, false);
}

export function EntDoctorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDoctor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctorName': !exists(json, 'doctorName') ? undefined : json['doctorName'],
        'doctorPassword': !exists(json, 'doctorPassword') ? undefined : json['doctorPassword'],
        'doctorUsername': !exists(json, 'doctorUsername') ? undefined : json['doctorUsername'],
        'edges': !exists(json, 'edges') ? undefined : EntDoctorEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDoctorToJSON(value?: EntDoctor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctorName': value.doctorName,
        'doctorPassword': value.doctorPassword,
        'doctorUsername': value.doctorUsername,
        'edges': EntDoctorEdgesToJSON(value.edges),
        'id': value.id,
    };
}


