// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/treatmenttype"
)

// TreatmentType is the model entity for the TreatmentType schema.
type TreatmentType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TreatmentTypeQuery when eager-loading is set.
	Edges TreatmentTypeEdges `json:"edges"`
}

// TreatmentTypeEdges holds the relations/edges for other nodes in the graph.
type TreatmentTypeEdges struct {
	// TreatmentTypeToDiagnosis holds the value of the TreatmentTypeToDiagnosis edge.
	TreatmentTypeToDiagnosis []*Diagnosis
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TreatmentTypeToDiagnosisOrErr returns the TreatmentTypeToDiagnosis value or an error if the edge
// was not loaded in eager-loading.
func (e TreatmentTypeEdges) TreatmentTypeToDiagnosisOrErr() ([]*Diagnosis, error) {
	if e.loadedTypes[0] {
		return e.TreatmentTypeToDiagnosis, nil
	}
	return nil, &NotLoadedError{edge: "TreatmentTypeToDiagnosis"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TreatmentType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case treatmenttype.FieldID:
			values[i] = &sql.NullInt64{}
		case treatmenttype.FieldType:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type TreatmentType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TreatmentType fields.
func (tt *TreatmentType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case treatmenttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case treatmenttype.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Type", values[i])
			} else if value.Valid {
				tt.Type = value.String
			}
		}
	}
	return nil
}

// QueryTreatmentTypeToDiagnosis queries the "TreatmentTypeToDiagnosis" edge of the TreatmentType entity.
func (tt *TreatmentType) QueryTreatmentTypeToDiagnosis() *DiagnosisQuery {
	return (&TreatmentTypeClient{config: tt.config}).QueryTreatmentTypeToDiagnosis(tt)
}

// Update returns a builder for updating this TreatmentType.
// Note that you need to call TreatmentType.Unwrap() before calling this method if this TreatmentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TreatmentType) Update() *TreatmentTypeUpdateOne {
	return (&TreatmentTypeClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TreatmentType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TreatmentType) Unwrap() *TreatmentType {
	tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TreatmentType is not a transactional entity")
	}
	tt.config.driver = tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TreatmentType) String() string {
	var builder strings.Builder
	builder.WriteString("TreatmentType(")
	builder.WriteString(fmt.Sprintf("id=%v", tt.ID))
	builder.WriteString(", Type=")
	builder.WriteString(tt.Type)
	builder.WriteByte(')')
	return builder.String()
}

// TreatmentTypes is a parsable slice of TreatmentType.
type TreatmentTypes []*TreatmentType

func (tt TreatmentTypes) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}
