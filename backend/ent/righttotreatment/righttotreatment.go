// Code generated by entc, DO NOT EDIT.

package righttotreatment

const (
	// Label holds the string label denoting the righttotreatment type in the database.
	Label = "right_to_treatment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"

	// EdgeHospital holds the string denoting the hospital edge name in mutations.
	EdgeHospital = "Hospital"
	// EdgeRightToTreatmentType holds the string denoting the righttotreatmenttype edge name in mutations.
	EdgeRightToTreatmentType = "RightToTreatmentType"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"

	// Table holds the table name of the righttotreatment in the database.
	Table = "right_to_treatments"
	// HospitalTable is the table the holds the Hospital relation/edge.
	HospitalTable = "right_to_treatments"
	// HospitalInverseTable is the table name for the Hospital entity.
	// It exists in this package in order to avoid circular dependency with the "hospital" package.
	HospitalInverseTable = "hospitals"
	// HospitalColumn is the table column denoting the Hospital relation/edge.
	HospitalColumn = "hospital_hospital"
	// RightToTreatmentTypeTable is the table the holds the RightToTreatmentType relation/edge.
	RightToTreatmentTypeTable = "right_to_treatments"
	// RightToTreatmentTypeInverseTable is the table name for the RightToTreatmentType entity.
	// It exists in this package in order to avoid circular dependency with the "righttotreatmenttype" package.
	RightToTreatmentTypeInverseTable = "right_to_treatment_types"
	// RightToTreatmentTypeColumn is the table column denoting the RightToTreatmentType relation/edge.
	RightToTreatmentTypeColumn = "right_to_treatment_type_type"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "right_to_treatments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "patient_patient_to_right_to_treatment"
)

// Columns holds all SQL columns for righttotreatment fields.
var Columns = []string{
	FieldID,
	FieldStartTime,
	FieldEndTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the RightToTreatment type.
var ForeignKeys = []string{
	"hospital_hospital",
	"patient_patient_to_right_to_treatment",
	"right_to_treatment_type_type",
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
