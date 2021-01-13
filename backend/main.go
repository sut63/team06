package main

import (
	"context"
	//"fmt"
	"log"
	//"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team06/app/controllers"
	_ "github.com/team06/app/docs"
	"github.com/team06/app/ent"
	//"github.com/team06/app/ent/department"
	//"github.com/team06/app/ent/nurse"
	//"github.com/team06/app/ent/patient"
	//"github.com/team06/app/ent/urgencylevel"
	//"github.com/team06/app/ent/doctor"
	//"github.com/team06/app/ent/proceduretype"
)

//Nurses Struct
type Nurses struct {
	Nurse []Nurse
}

//Nurse Struct
type Nurse struct {
	NurseName     string
	NurseUsername string
	NursePassword string
}

//TriageResults Struct
type TriageResults struct {
	TriageResult []TriageResults
}

//ProcedureTypes Struct
type ProcedureTypes struct {
	ProcedureType []ProcedureType
}

//ProcedureType Struct
type ProcedureType struct {
	ProcedureName string
}

//Doctors Struct
type Doctors struct {
	Doctor []Doctor
}

//Doctor Struct
type Doctor struct {
	DoctorName     string
	DoctorUsername string
	DoctorPassword string
}

//MedicalProcedures Struct
type MedicalProcedures struct {
	MedicalProcedure []MedicalProcedures
}

//Rooms Struct
type Rooms struct {
	Room []Room
}

//Room Struct
type Room struct {
	RoomName string
}

//AppointmentResults Struct
type AppointmentResultss struct {
	AppointmentResults []AppointmentResultss
}

//Hospitals struct
type Hospitals struct {
	Hospital []Hospital
}

//Hospital struct
type Hospital struct {
	HospitalName string
}

//RightToTreatmentTypes struct
type RightToTreatmentTypes struct {
	RightToTreatmentType []RightToTreatmentType
}

//RightToTreatmentType struct
type RightToTreatmentType struct {
	RightToTreatmentTypeName string
}

//RightToTreatments Struct
type RightToTreatments struct {
	RightToTreatment []RightToTreatments
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:medicalrecord.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewTriageResultController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewUrgencyLevelController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewProceduretypeController(v1, client)
	controllers.NewMedicalprocedureController(v1, client)
	controllers.NewAppointmentResultsController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewHospitalController(v1, client)
	controllers.NewRightToTreatmentController(v1, client)
	controllers.NewRightToTreatmentTypeController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewMedicalRecordController(v1, client)
	controllers.NewPrefixController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewBloodTypeController(v1, client)
	controllers.NewDiagnosisController(v1, client)
	controllers.NewTreatmenttypeController(v1, client)

	// Set Department Data
	departments := []string{"Anesthetics", "Breast Screening", "Cardiology",
		"Eae, nose and throat(ENT)", "Elderly services department", "Gastroenterology",
		"General Surgery", "Gynecology", "Hematology", "Neonatal Unit", "Neurology",
		"Nutrition and dietetics", "Obstetrics and gynecologt units", "Oncology",
		"Ophthalmology", "Orthopedics", "Physiotherapy", "Renal Unit", "Sexual Health",
		"Urology"}

	for _, d := range departments {
		client.Department.
			Create().
			SetDepartmentName(d).
			Save(context.Background())
	}

	// Set Nurses Data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"Somsri", "Somsri@nurse", "n1111"},
			Nurse{"Somwang", "Somwang@nurse", "n2222"},
		},
	}

	for _, n := range nurses.Nurse {
		client.Nurse.
			Create().
			SetNurseName(n.NurseName).
			SetNurseUsername(n.NurseUsername).
			SetNursePassword(n.NursePassword).
			Save(context.Background())
	}

	// Set UrgencyLevel Data
	urgencylevels := []string{"Normal", "Urgent", "Critical"}

	for _, u := range urgencylevels {
		client.UrgencyLevel.
			Create().
			SetUrgencyName(u).
			Save(context.Background())
	}

	// Set Doctor Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"Suchat", "schat", "12345g"},
			Doctor{"Suwit", "wittsu", "44355h"},
			Doctor{"Somkiat", "fickkm", "123452f"},
		},
	}

	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorName(d.DoctorName).
			SetDoctorUsername(d.DoctorUsername).
			SetDoctorPassword(d.DoctorPassword).
			Save(context.Background())
	}

	// Set ProcedureType Data
	proceduretypes := ProcedureTypes{
		ProcedureType: []ProcedureType{
			ProcedureType{"Endotracheal intubation"},
			ProcedureType{"Intercostal drainage"},
			ProcedureType{"Lumbar puncture"},
			ProcedureType{"Thoracentesis"},
			ProcedureType{"PCI"},
			ProcedureType{"Dystocia"},
		},
	}

	for _, pr := range proceduretypes.ProcedureType {
		client.ProcedureType.
			Create().
			SetProcedureName(pr.ProcedureName).
			Save(context.Background())
	}

	// Set Room Data
	rooms := Rooms{
		Room: []Room{
			Room{"ห้องตรวจ1"},
			Room{"ห้องตรวจ2"},
			Room{"ห้องตรวจ3"},
			Room{"ห้องตรวจ4"},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetRoomName(r.RoomName).
			Save(context.Background())
	}
	// Set Hospital Data
	hospitals := Hospitals{
		Hospital: []Hospital{
			Hospital{"Saint Louis Hospital"},
			Hospital{"Mahaesak Hospital"},
			Hospital{"Samitivej Hospital in Thonburi"},
			Hospital{"Charoenkrung Pracharak Hospital"},
			Hospital{"BNH Hospital"},
			Hospital{"Suranaree University of Technology Hospital"},
		},
	}

	for _, h := range hospitals.Hospital {
		client.Hospital.
			Create().
			SetHospitalName(h.HospitalName).
			Save(context.Background())
	}

	// Set RightToTreatmentType Data
	righttotreatmenttypes := RightToTreatmentTypes{
		RightToTreatmentType: []RightToTreatmentType{
			RightToTreatmentType{"สิทธิสวัสดิการการรักษาพยาบาลของข้าราชการ"},
			RightToTreatmentType{"สิทธิประกันสังคม"},
			RightToTreatmentType{"สิทธิประกันสุขภาพถ้วนหน้า"},
		},
	}

	for _, t := range righttotreatmenttypes.RightToTreatmentType {
		client.RightToTreatmentType.
			Create().
			SetTypeName(t.RightToTreatmentTypeName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
