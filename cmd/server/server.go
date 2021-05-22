package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/02amanag/p-02/config"
	"github.com/02amanag/p-02/internal/repository"
	"github.com/02amanag/p-02/internal/service"
	"github.com/02amanag/p-02/internal/usecase"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := serviceObject.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

var (
	ConfigService    = config.NewConfig(".env")
	repositoryObject = repository.NewRepositoryStruct()
	usecaseObject    = usecase.NewUsecaseStruct(*repositoryObject, ConfigService)
	serviceObject    = service.NewServiceStruct(usecaseObject)
)

func Run() {

	server := gin.Default()
	// server.StaticFS("file/profilePicture/*", gin.Dir("file/profilePicture/", false))
	server.LoadHTMLGlob("deploy/web/tempelate/*.html")
	// server.StaticFile("test.png", "file/profilePicture/test.png")

	server.Use(CORSMiddleware())

	server.GET("/health", serviceObject.Healthy)
	v1 := server.Group("/show")
	{
		v1.GET("/profile/:username", serviceObject.GetProfile) //complete profile
		// v1.GET("/userdetails/:username", serviceObject.GetUserDetails)
		v1.GET("/downloadresume/:username", serviceObject.DownloadResume)
		//download qr
		v1.GET("/downloadqr/:username", serviceObject.GetQr)
	}

	v2 := server.Group("/user")
	{
		v2.GET("/login", serviceObject.Login)
		v2.GET("/singup", serviceObject.SingUp)
		v2.GET("/getqr", TokenAuthMiddleware(), serviceObject.GenerateQr)

		v3 := server.Group("add")
		{
			//update, add , delete endpoint
			// Token Will be needed
			v3.POST("/profile", TokenAuthMiddleware(), serviceObject.AddProfile)
			v3.POST("/uploadresume", TokenAuthMiddleware(), serviceObject.UploadSingleFile)

		}
	}

	port, err := ConfigService.GetConfig("PORT")
	if err != nil {
		log.Fatal(err)
	}

	server.Run(":" + port)
}
