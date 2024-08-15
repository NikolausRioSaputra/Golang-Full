package main

import (
	"github.com/gin-gonic/gin"
	_ "swaggerping/docs" // Ganti dengan path ke package docs yang akan dihasilkan oleh swag init
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"net/http"
)

// User represents the user model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`	
}

// HTTPError represents the error model
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}


// @title           Test Operation API
// @version         1.0
// @description     Ini adalah dokumentasi Swagger untuk API Test Operation di Golang.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

// @tag.name        tests
// @tag.description Operations related to testing

func main() {
	// Set router
	r := gin.Default()

	// Register Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define routes
	r.GET("/test/:id", GetTest)
	r.POST("/test", CreateTest)

	// Run server
	r.Run(":8080") // Jalankan server pada port 8080
}

// @Summary      Get Test
// @Description  Mengambil data test berdasarkan ID.
// @Tags         tests
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "Test ID"
// @Success      200  {object}  Test
// @Failure      404  {object}  HTTPError
// @Router       /test/{id} [get]
func GetTest(c *gin.Context) {
	id := c.Param("id")
	// Logika untuk mengambil data test
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "Sample Test"})
}

// Test represents the test model
type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// @Summary      Create Test
// @Description  Membuat test baru.
// @Tags         tests
// @Accept       json
// @Produce      json
// @Param        test  body      Test     true  "Test Info"
// @Success      201   {object}  Test
// @Failure      400   {object}  HTTPError
// @Router       /test [post]
func CreateTest(c *gin.Context) {
	var test Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logika untuk membuat test baru
	c.JSON(http.StatusCreated, test)
	
}

