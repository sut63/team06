package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// AppointmentResults holds the schema definition for the AppointmentResults entity.
type AppointmentResults struct {
	ent.Schema
}

// Fields of the AppointmentResults.
func (AppointmentResults) Fields() []ent.Field {
	return []ent.Field{
		field.String("causeAppoint").NotEmpty(),
		field.String("advice").NotEmpty(),
		field.Time("dateAppoint"),
		field.Time("timeAppoint"),
		field.Time("addtimeSave").Default(time.Now),
		field.Int("hourBeforeAppoint").Range(0, 3).Default(0),
		field.Int("minuteBeforeAppoint").Range(0, 59).Default(0),
	}
}

// Edges of the AppointmentResults.
func (AppointmentResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appointmentResultsToPatient", Patient.Type).Ref("PatientToAppointmentResults").Unique(),
		edge.From("appointmentResultsToNurse", Nurse.Type).Ref("NurseToAppointmentResults").Unique(),
		edge.From("appointmentResultsToDoctor", Doctor.Type).Ref("DoctorToAppointmentResults").Unique(),
		edge.From("appointmentResultsToRoom", Room.Type).Ref("RoomToAppointmentResults").Unique(),
	}
}
