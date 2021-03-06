// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/department"
	"github.com/team06/app/ent/nurse"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/triageresult"
	"github.com/team06/app/ent/urgencylevel"
)

// TriageResult is the model entity for the TriageResult schema.
type TriageResult struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Height holds the value of the "height" field.
	Height float64 `json:"height,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
	// Pressure holds the value of the "pressure" field.
	Pressure float64 `json:"pressure,omitempty"`
	// Symptom holds the value of the "symptom" field.
	Symptom string `json:"symptom,omitempty"`
	// TriageDate holds the value of the "triageDate" field.
	TriageDate time.Time `json:"triageDate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TriageResultQuery when eager-loading is set.
	Edges                       TriageResultEdges `json:"edges"`
	department_triage_result    *int
	nurse_triage_result         *int
	patient_triage_result       *int
	urgency_level_triage_result *int
}

// TriageResultEdges holds the relations/edges for other nodes in the graph.
type TriageResultEdges struct {
	// UrgencyLevel holds the value of the urgencyLevel edge.
	UrgencyLevel *UrgencyLevel
	// Department holds the value of the department edge.
	Department *Department
	// Nurse holds the value of the nurse edge.
	Nurse *Nurse
	// Patient holds the value of the patient edge.
	Patient *Patient
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UrgencyLevelOrErr returns the UrgencyLevel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TriageResultEdges) UrgencyLevelOrErr() (*UrgencyLevel, error) {
	if e.loadedTypes[0] {
		if e.UrgencyLevel == nil {
			// The edge urgencyLevel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: urgencylevel.Label}
		}
		return e.UrgencyLevel, nil
	}
	return nil, &NotLoadedError{edge: "urgencyLevel"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TriageResultEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[1] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// NurseOrErr returns the Nurse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TriageResultEdges) NurseOrErr() (*Nurse, error) {
	if e.loadedTypes[2] {
		if e.Nurse == nil {
			// The edge nurse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nurse.Label}
		}
		return e.Nurse, nil
	}
	return nil, &NotLoadedError{edge: "nurse"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TriageResultEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[3] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TriageResult) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case triageresult.FieldHeight, triageresult.FieldWeight, triageresult.FieldPressure:
			values[i] = &sql.NullFloat64{}
		case triageresult.FieldID:
			values[i] = &sql.NullInt64{}
		case triageresult.FieldSymptom:
			values[i] = &sql.NullString{}
		case triageresult.FieldTriageDate:
			values[i] = &sql.NullTime{}
		case triageresult.ForeignKeys[0]: // department_triage_result
			values[i] = &sql.NullInt64{}
		case triageresult.ForeignKeys[1]: // nurse_triage_result
			values[i] = &sql.NullInt64{}
		case triageresult.ForeignKeys[2]: // patient_triage_result
			values[i] = &sql.NullInt64{}
		case triageresult.ForeignKeys[3]: // urgency_level_triage_result
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type TriageResult", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TriageResult fields.
func (tr *TriageResult) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case triageresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tr.ID = int(value.Int64)
		case triageresult.FieldHeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				tr.Height = value.Float64
			}
		case triageresult.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				tr.Weight = value.Float64
			}
		case triageresult.FieldPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pressure", values[i])
			} else if value.Valid {
				tr.Pressure = value.Float64
			}
		case triageresult.FieldSymptom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symptom", values[i])
			} else if value.Valid {
				tr.Symptom = value.String
			}
		case triageresult.FieldTriageDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field triageDate", values[i])
			} else if value.Valid {
				tr.TriageDate = value.Time
			}
		case triageresult.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field department_triage_result", value)
			} else if value.Valid {
				tr.department_triage_result = new(int)
				*tr.department_triage_result = int(value.Int64)
			}
		case triageresult.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field nurse_triage_result", value)
			} else if value.Valid {
				tr.nurse_triage_result = new(int)
				*tr.nurse_triage_result = int(value.Int64)
			}
		case triageresult.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field patient_triage_result", value)
			} else if value.Valid {
				tr.patient_triage_result = new(int)
				*tr.patient_triage_result = int(value.Int64)
			}
		case triageresult.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field urgency_level_triage_result", value)
			} else if value.Valid {
				tr.urgency_level_triage_result = new(int)
				*tr.urgency_level_triage_result = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUrgencyLevel queries the "urgencyLevel" edge of the TriageResult entity.
func (tr *TriageResult) QueryUrgencyLevel() *UrgencyLevelQuery {
	return (&TriageResultClient{config: tr.config}).QueryUrgencyLevel(tr)
}

// QueryDepartment queries the "department" edge of the TriageResult entity.
func (tr *TriageResult) QueryDepartment() *DepartmentQuery {
	return (&TriageResultClient{config: tr.config}).QueryDepartment(tr)
}

// QueryNurse queries the "nurse" edge of the TriageResult entity.
func (tr *TriageResult) QueryNurse() *NurseQuery {
	return (&TriageResultClient{config: tr.config}).QueryNurse(tr)
}

// QueryPatient queries the "patient" edge of the TriageResult entity.
func (tr *TriageResult) QueryPatient() *PatientQuery {
	return (&TriageResultClient{config: tr.config}).QueryPatient(tr)
}

// Update returns a builder for updating this TriageResult.
// Note that you need to call TriageResult.Unwrap() before calling this method if this TriageResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TriageResult) Update() *TriageResultUpdateOne {
	return (&TriageResultClient{config: tr.config}).UpdateOne(tr)
}

// Unwrap unwraps the TriageResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tr *TriageResult) Unwrap() *TriageResult {
	tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TriageResult is not a transactional entity")
	}
	tr.config.driver = tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TriageResult) String() string {
	var builder strings.Builder
	builder.WriteString("TriageResult(")
	builder.WriteString(fmt.Sprintf("id=%v", tr.ID))
	builder.WriteString(", height=")
	builder.WriteString(fmt.Sprintf("%v", tr.Height))
	builder.WriteString(", weight=")
	builder.WriteString(fmt.Sprintf("%v", tr.Weight))
	builder.WriteString(", pressure=")
	builder.WriteString(fmt.Sprintf("%v", tr.Pressure))
	builder.WriteString(", symptom=")
	builder.WriteString(tr.Symptom)
	builder.WriteString(", triageDate=")
	builder.WriteString(tr.TriageDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TriageResults is a parsable slice of TriageResult.
type TriageResults []*TriageResult

func (tr TriageResults) config(cfg config) {
	for _i := range tr {
		tr[_i].config = cfg
	}
}
