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
    EntAppointmentResults,
    EntAppointmentResultsFromJSON,
    EntAppointmentResultsFromJSONTyped,
    EntAppointmentResultsToJSON,
    EntBloodType,
    EntBloodTypeFromJSON,
    EntBloodTypeFromJSONTyped,
    EntBloodTypeToJSON,
    EntDiagnosis,
    EntDiagnosisFromJSON,
    EntDiagnosisFromJSONTyped,
    EntDiagnosisToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntMedicalProcedure,
    EntMedicalProcedureFromJSON,
    EntMedicalProcedureFromJSONTyped,
    EntMedicalProcedureToJSON,
    EntPrefix,
    EntPrefixFromJSON,
    EntPrefixFromJSONTyped,
    EntPrefixToJSON,
    EntRightToTreatment,
    EntRightToTreatmentFromJSON,
    EntRightToTreatmentFromJSONTyped,
    EntRightToTreatmentToJSON,
    EntTriageResult,
    EntTriageResultFromJSON,
    EntTriageResultFromJSONTyped,
    EntTriageResultToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * 
     * @type {EntBloodType}
     * @memberof EntPatientEdges
     */
    bloodtype?: EntBloodType;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientEdges
     */
    gender?: EntGender;
    /**
     * PatientToAppointmentResults holds the value of the PatientToAppointmentResults edge.
     * @type {Array<EntAppointmentResults>}
     * @memberof EntPatientEdges
     */
    patientToAppointmentResults?: Array<EntAppointmentResults>;
    /**
     * PatientToDiagnosis holds the value of the PatientToDiagnosis edge.
     * @type {Array<EntDiagnosis>}
     * @memberof EntPatientEdges
     */
    patientToDiagnosis?: Array<EntDiagnosis>;
    /**
     * PatientToMedicalProcedure holds the value of the PatientToMedicalProcedure edge.
     * @type {Array<EntMedicalProcedure>}
     * @memberof EntPatientEdges
     */
    patientToMedicalProcedure?: Array<EntMedicalProcedure>;
    /**
     * PatientToRightToTreatment holds the value of the PatientToRightToTreatment edge.
     * @type {Array<EntRightToTreatment>}
     * @memberof EntPatientEdges
     */
    patientToRightToTreatment?: Array<EntRightToTreatment>;
    /**
     * 
     * @type {EntPrefix}
     * @memberof EntPatientEdges
     */
    prefix?: EntPrefix;
    /**
     * TriageResult holds the value of the triageResult edge.
     * @type {Array<EntTriageResult>}
     * @memberof EntPatientEdges
     */
    triageResult?: Array<EntTriageResult>;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bloodtype': !exists(json, 'Bloodtype') ? undefined : EntBloodTypeFromJSON(json['Bloodtype']),
        'gender': !exists(json, 'Gender') ? undefined : EntGenderFromJSON(json['Gender']),
        'patientToAppointmentResults': !exists(json, 'patientToAppointmentResults') ? undefined : ((json['patientToAppointmentResults'] as Array<any>).map(EntAppointmentResultsFromJSON)),
        'patientToDiagnosis': !exists(json, 'patientToDiagnosis') ? undefined : ((json['patientToDiagnosis'] as Array<any>).map(EntDiagnosisFromJSON)),
        'patientToMedicalProcedure': !exists(json, 'patientToMedicalProcedure') ? undefined : ((json['patientToMedicalProcedure'] as Array<any>).map(EntMedicalProcedureFromJSON)),
        'patientToRightToTreatment': !exists(json, 'patientToRightToTreatment') ? undefined : ((json['patientToRightToTreatment'] as Array<any>).map(EntRightToTreatmentFromJSON)),
        'prefix': !exists(json, 'Prefix') ? undefined : EntPrefixFromJSON(json['Prefix']),
        'triageResult': !exists(json, 'triageResult') ? undefined : ((json['triageResult'] as Array<any>).map(EntTriageResultFromJSON)),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bloodtype': EntBloodTypeToJSON(value.bloodtype),
        'gender': EntGenderToJSON(value.gender),
        'patientToAppointmentResults': value.patientToAppointmentResults === undefined ? undefined : ((value.patientToAppointmentResults as Array<any>).map(EntAppointmentResultsToJSON)),
        'patientToDiagnosis': value.patientToDiagnosis === undefined ? undefined : ((value.patientToDiagnosis as Array<any>).map(EntDiagnosisToJSON)),
        'patientToMedicalProcedure': value.patientToMedicalProcedure === undefined ? undefined : ((value.patientToMedicalProcedure as Array<any>).map(EntMedicalProcedureToJSON)),
        'patientToRightToTreatment': value.patientToRightToTreatment === undefined ? undefined : ((value.patientToRightToTreatment as Array<any>).map(EntRightToTreatmentToJSON)),
        'prefix': EntPrefixToJSON(value.prefix),
        'triageResult': value.triageResult === undefined ? undefined : ((value.triageResult as Array<any>).map(EntTriageResultToJSON)),
    };
}


