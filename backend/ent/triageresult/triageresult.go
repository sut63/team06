// Code generated by entc, DO NOT EDIT.

package triageresult

import (
	"time"
)

const (
	// Label holds the string label denoting the triageresult type in the database.
	Label = "triage_result"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSymptom holds the string denoting the symptom field in the database.
	FieldSymptom = "symptom"
	// FieldTriageDate holds the string denoting the triagedate field in the database.
	FieldTriageDate = "triage_date"

	// EdgeTriageResultToUrgencyLevel holds the string denoting the triageresulttourgencylevel edge name in mutations.
	EdgeTriageResultToUrgencyLevel = "triageResultToUrgencyLevel"
	// EdgeTriageResultToDepartment holds the string denoting the triageresulttodepartment edge name in mutations.
	EdgeTriageResultToDepartment = "triageResultToDepartment"
	// EdgeTriageResultToNurse holds the string denoting the triageresulttonurse edge name in mutations.
	EdgeTriageResultToNurse = "triageResultToNurse"
	// EdgeTriageResultToPatient holds the string denoting the triageresulttopatient edge name in mutations.
	EdgeTriageResultToPatient = "triageResultToPatient"

	// Table holds the table name of the triageresult in the database.
	Table = "triage_results"
	// TriageResultToUrgencyLevelTable is the table the holds the triageResultToUrgencyLevel relation/edge.
	TriageResultToUrgencyLevelTable = "triage_results"
	// TriageResultToUrgencyLevelInverseTable is the table name for the UrgencyLevel entity.
	// It exists in this package in order to avoid circular dependency with the "urgencylevel" package.
	TriageResultToUrgencyLevelInverseTable = "urgency_levels"
	// TriageResultToUrgencyLevelColumn is the table column denoting the triageResultToUrgencyLevel relation/edge.
	TriageResultToUrgencyLevelColumn = "urgency_level_urgency_level_to_triage_result"
	// TriageResultToDepartmentTable is the table the holds the triageResultToDepartment relation/edge.
	TriageResultToDepartmentTable = "triage_results"
	// TriageResultToDepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	TriageResultToDepartmentInverseTable = "departments"
	// TriageResultToDepartmentColumn is the table column denoting the triageResultToDepartment relation/edge.
	TriageResultToDepartmentColumn = "department_department_to_triage_result"
	// TriageResultToNurseTable is the table the holds the triageResultToNurse relation/edge.
	TriageResultToNurseTable = "triage_results"
	// TriageResultToNurseInverseTable is the table name for the Nurse entity.
	// It exists in this package in order to avoid circular dependency with the "nurse" package.
	TriageResultToNurseInverseTable = "nurses"
	// TriageResultToNurseColumn is the table column denoting the triageResultToNurse relation/edge.
	TriageResultToNurseColumn = "nurse_nurse_to_triage_result"
	// TriageResultToPatientTable is the table the holds the triageResultToPatient relation/edge.
	TriageResultToPatientTable = "triage_results"
	// TriageResultToPatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	TriageResultToPatientInverseTable = "patients"
	// TriageResultToPatientColumn is the table column denoting the triageResultToPatient relation/edge.
	TriageResultToPatientColumn = "patient_patient_to_triage_result"
)

// Columns holds all SQL columns for triageresult fields.
var Columns = []string{
	FieldID,
	FieldSymptom,
	FieldTriageDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the TriageResult type.
var ForeignKeys = []string{
	"department_department_to_triage_result",
	"nurse_nurse_to_triage_result",
	"patient_patient_to_triage_result",
	"urgency_level_urgency_level_to_triage_result",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SymptomValidator is a validator for the "symptom" field. It is called by the builders before save.
	SymptomValidator func(string) error
	// DefaultTriageDate holds the default value on creation for the "triageDate" field.
	DefaultTriageDate func() time.Time
)
