package main

import (
	"fmt"
	"net/http"
	"os"

	mitarbeiter "github.com/computerextra/local-horst-react-go/Mitarbeiter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	mode := os.Getenv("MODE")
	var port = ""
	port = os.Getenv("VITE_PORT")
	if len(port) < 1 {
		fmt.Println("No Port in .env specified. Using Port 8080 in Dev Mode")
		port = "8080"
	}

	r := gin.Default()

	if mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("MODE not set in .env, for Production Build use PROD as Value")
		gin.SetMode(gin.DebugMode)
	}

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// API Routes
	v1 := r.Group("/api/v1")
	{
		mitarbeiterRouter := v1.Group("/mitarbeiter")
		{
			mitarbeiterRouter.GET("/geburtstag", func(c *gin.Context) {
				c.JSON(http.StatusOK, "")
			})
			mitarbeiterRouter.GET("/all", func(c *gin.Context) {
				ma, err := mitarbeiter.GetMitarbeiter()
				if err != nil {
					c.JSON(http.StatusBadRequest, "")
					return
				}
				c.JSON(http.StatusOK, ma)
			})
		}
	}

	r.Run(fmt.Sprintf(":%s", port))

}
