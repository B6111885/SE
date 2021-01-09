package main

import (
	"context"
	"log"

	"github.com/ADMIN/app/controller"
	_ "github.com/ADMIN/app/docs"
	"github.com/ADMIN/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
type Patients struct {
	Patient []Patient
}

type Patient struct {
	PatientName  string
	PatientAge  string
	PatientWeight  string
	PatientHeight  string
	PatientPrefix  string
	PatientGender  string
	PatientBlood  string
}

type Doctors struct {
	Doctor []Doctor
}

type Doctor struct {
	DoctorName string
	DoctorEmail string
	DoctorPassword string
	DoctorTel string
}

type Diseases struct {
	Disease []Disease
}

type Disease struct {
	DiseaseName string
}

type Departments struct {
	Department []Department
}

type Department struct {
	DepartmentName string
}

// @title SUT SA Example API Playlist Vidoe
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
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	v1 := router.Group("/api/v1")
	controllers.NewPatientController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewDiagnoseController(v1, client)
	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"สุมารี นามใจ","30","55","157","นาง","หญิง","B" },
			Patient{"สมชาติ มีใจ","28","67","178","นาย","ชาย","O" },
			Patient{"นานมี อารมณ์ดี","22","50","162","นางสาว","หญิง","A" },
		},
	}
	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PatientName).
			SetPatientAge(p.PatientAge).
			SetPatientWeight(p.PatientWeight).
			SetPatientHeight(p.PatientHeight).
			SetPatientPrefix(p.PatientPrefix).
			SetPatientGender(p.PatientGender).
			SetPatientBlood(p.PatientBlood).
			Save(context.Background())
	}
	// Set Doctor Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"Dr. James","James@gmail.com","James","0889977665"},
			Doctor{"Dr. John","John@gmail.com","John","0899177998"},
			Doctor{"Dr. Susan","Susan@gmail.com","Susan","0812345671"},
		},
	}
	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorName(d.DoctorName).
			SetDoctorEmail(d.DoctorEmail).
			SetDoctorPassword(d.DoctorPassword).
			SetDoctorTel(d.DoctorTel).
			Save(context.Background())
	}
	// Set Disease Data
	diseases := Diseases{
		Disease: []Disease{
			Disease{"แก้วหูอักเสบ"},
			Disease{"อัมพฤกษ์"},
			Disease{"โรคต้อหิน"},
			Disease{"โรคไส้ติ่งอักเสบ"},
			Disease{"โรคเส้นเลือดขอด"},
			Disease{"โรคทางเดินหายใจในเด็ก"},
		},
	}
	for _, di := range diseases.Disease {
		client.Disease.
			Create().
			SetDiseaseName(di.DiseaseName).
			Save(context.Background())
	}

	// Set Department Data
	departments := Departments{
		Department: []Department{
			Department{"แผนกกุมารเวชกรรม"},
			Department{"แผนกเวชศาสตร์"},
			Department{"แผนกจักษุ"},
			Department{"แผนกหู คอ จมูก"},
			Department{"แผนกศัลยกรรม"},
			Department{"แผนกจิตเวช"},
		},
	}
	for _, dep := range departments.Department {
		client.Department.
			Create().
			SetDepartmentName(dep.DepartmentName).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}