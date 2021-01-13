package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tanapon395/playlist-video/controllers"
	"github.com/tanapon395/playlist-video/ent"
)

// Titles struct type
type Titles struct {
	Title []Title
}

// Title struct type
type Title struct {
	Titlename string
}

// Departments struct type
type Departments struct {
	Department []Department
}

// Department struct type
type Department struct {
	Departmentname string
}

// Genders struct type
type Genders struct {
	Gender []Gender
}

// Gender struct type
type Gender struct {
	Gendername string
}

// Brands struct type
type Brands struct {
	Brand []Brand
}

// Brand struct type
type Brand struct {
	Brandname string
}

// Typeproducts struct type
type Typeproducts struct {
	Typeproduct []Typeproduct
}

// Typeproduct struct type
type Typeproduct struct {
	Typeproductname string
}

// Fixcomtypes struct type
type Fixcomtypes struct {
	Fixcomtype []Fixcomtype
}

// Fixcomtype struct type
type Fixcomtype struct {
	Fixcomtypename string
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
	controllers.NewPersonalController(v1, client)
	controllers.NewTitleController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewProductController(v1, client)
	controllers.NewTypeproductController(v1, client)
	controllers.NewBrandController(v1, client)
	controllers.NewFixController(v1, client)
	controllers.NewFixcomtypeController(v1, client)
	controllers.NewAdminrepairController(v1, client)

	//Set Title Data
	titles := []string{"นาย", "นาง", "นางสาว", "ไม่ระบุ"}
	for _, t := range titles {
		client.Title.
			Create().
			SetTitlename(t).
			Save(context.Background())
	}

	//Set Department Data
	departments := []string{"แผนกแจ้งซ่อม", "แผนกบัญชี", "แผนกซ่อม", "แผนกอะไหล่", "แผนกส่วนบุคคล"}
	for _, d := range departments {
		client.Department.
			Create().
			SetDepartmentname(d).
			Save(context.Background())
	}

	//Set Gender Data
	genders := []string{"ชาย", "หญิง", "ไม่ระบุ"}
	for _, g := range genders {
		client.Gender.
			Create().
			SetGendername(g).
			Save(context.Background())
	}

	// Set Brands Data
	brands := Brands{
		Brand: []Brand{
			Brand{"Intel"},
			Brand{"AMD"},
			Brand{"Kingston"},
			Brand{"MSI"},
			Brand{"Western Digital"},
			Brand{"ASUS"},
			Brand{"ACER"},
			Brand{"DELL"},
			Brand{"LENOVO"},
		},
	}
	for _, b := range brands.Brand {
		client.Brand.
			Create().
			SetBrandname(b.Brandname).
			Save(context.Background())
	}

	// Set Typeproducts Data
	typeproducts := Typeproducts{
		Typeproduct: []Typeproduct{
			Typeproduct{"เมนบอร์ด"},
			Typeproduct{"พัดลมระบายความร้อน"},
			Typeproduct{"SSD"},
			Typeproduct{"HDD"},
			Typeproduct{"CPU"},
			Typeproduct{"RAM"},
		},
	}
	for _, t := range typeproducts.Typeproduct {
		client.Typeproduct.
			Create().
			SetTypeproductname(t.Typeproductname).
			Save(context.Background())
	}

	// Set Fixcomtypes Data
	fixcomtypes := Fixcomtypes{
		Fixcomtype: []Fixcomtype{
			Fixcomtype{"Personal Computer(PC)"},
			Fixcomtype{"Laptop"},
		},
	}
	for _, ft := range fixcomtypes.Fixcomtype {
		client.Fixcomtype.
			Create().
			SetFixcomtypename(ft.Fixcomtypename).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
