//	@title			RancoDev GBMS API Docs
//	@version		1.0
//	@description	Gin swagger

//	@contact.name	cody chen
//	@contact.url	https://gbms.codychen.me/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// schemes http
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Ranco-dev/gbms/api/controllers"
	invade "github.com/Ranco-dev/gbms/api/controllers/invade"
	payment "github.com/Ranco-dev/gbms/api/controllers/payment"
	"github.com/Ranco-dev/gbms/pkg/config"

	"github.com/Ranco-dev/gbms/pkg/db"
	"github.com/Ranco-dev/gbms/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Ranco-dev/gbms/docs"
)

var releaseMode string = "debug"

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	var conf = config.GetConfig()
	var port = "9480"

	log.InitLogger("./logs/gbms.log", "debug", 1, 5, 30)
	defer log.Sugar.Sync()
	log.Logger.Info("Start GBMS Server")

	dbName := conf.GetString("db.database")
	dbHost := conf.GetString("db.host")
	dbPort := conf.GetString("db.port")
	dbUser := conf.GetString("db.user")
	dbPassword := conf.GetString("db.pass")

	if err := db.PgDBConnection(dbHost, dbPort, dbUser, dbPassword, dbName); err != nil {
		log.Sugar.Fatal(err.Error())
	}

	r := setupRouter()
	if mode := gin.Mode(); mode == gin.DebugMode {
		// url := ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:%s/swagger/doc.json", conf.GetString("server.port")))
		// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	}

	if releaseMode == "release" {
		port = conf.GetString("server.port")
	}
	r.Run("127.0.0.1:" + port)
}

func setupRouter() (router *gin.Engine) {
	if releaseMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	// gin.DisableConsoleColor()

	router = gin.New()

	router.Use(log.GinLogger(log.Logger), log.GinRecovery(log.Logger, true))
	router.Use(gin.Recovery())

	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	// corsConf.AllowHeaders = []string{
	// 	"authorization",
	// 	"content-type",
	// }
	router.Use(cors.New(corsConf))

	routeURL(router)

	return
}

func routeURL(r *gin.Engine) {
	r.GET("api/v1/ping", controllers.PING)
	r.GET("api/v1/ping2", controllers.PING)

	//payment request
	r.GET("api/v1/pymReq/:uid/:amount", payment.PymReq)
	r.GET("api/v1/pymCheck/:tid", payment.PymCheck)

	//invade log luke
	r.GET("api/v1/invade/:value", invade.InvadeLog)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status": 0,
			"msg":    "API Not Found",
		})
	})
}
