package routers

import (
	"io"
	"middleware-experience/constants"
	"middleware-experience/utils"

	ctrlMid "middleware-experience/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

var (
	MainSetupUrlPrefix string
	MainSetupAppsDebug string

	InternalRoutingHealthCheck string
	InternalRoutingVersion     string
)

type MiddlewareRouter struct {
	GinFunc  gin.HandlerFunc
	Router   *gin.Engine
	Tracer   opentracing.Tracer
	Reporter jaeger.Reporter
	Closer   io.Closer
	Err      error
}

func init() {
	MainSetupUrlPrefix = utils.EnvString("MainSetup.UrlPrefix", "")
	MainSetupAppsDebug = utils.EnvString("MainSetup.AppsDebug", "debug")
	InternalRoutingHealthCheck = utils.EnvString("InternalRouting.HealthCheck", "/healthCheck")
	InternalRoutingVersion = utils.EnvString("InternalRouting.Version", "/version")
}

func Server(listenAddress string) error {
	tracer, closer, err := utils.InitJaeger()
	if err != nil {
		utils.LogData("Routers", "Server", constants.LEVEL_LOG_WARNING, err.Error())
	}

	defer closer.Close()

	MiddlewareRouter := MiddlewareRouter{}
	MiddlewareRouter.Routers(tracer)

	err = utils.GinServerUp(listenAddress, MiddlewareRouter.Router)
	MiddlewareRouter.GinFunc = utils.OpenTracer([]byte("api-request-"))
	if err != nil {
		return err
	}

	return nil
}

func (MiddlewareRouter *MiddlewareRouter) Routers(tracer opentracing.Tracer) {
	gin.SetMode(MainSetupAppsDebug)

	router := gin.New()
	router.Use(gin.Recovery())

	ctrlMiddleware := ctrlMid.InitController(tracer)

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "access-token", "X-SIGNATURE", "X-TIMESTAMPS"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	api := router.Group(MainSetupUrlPrefix)
	apiMiddleware := api.Group("/middleware")

	//Account
	vAccAppsRouteGroup := apiMiddleware.Group("/e-accounting")
	{
		vAccAppsRouteGroup.POST("/login", ctrlMiddleware.CtrlLoginAuth)
		vAccAppsRouteGroup.POST("/profile/username", ctrlMiddleware.CtrlProfileUsername)
	}

	MiddlewareRouter.Router = router
}

// Close ..
func (trc *MiddlewareRouter) Close() {
	trc.Closer.Close()
}
