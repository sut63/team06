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
		field.Time("addtimeAppoint"),
		field.Time("addtimeSave").Default(time.Now),
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
