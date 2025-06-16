package main

import (
        // STD
        //"os"
        //"time"
	"sync"

        // 3PPs
        "github.com/gin-gonic/gin"
        ginSwagger "github.com/swaggo/gin-swagger"
        swaggerFiles "github.com/swaggo/files"
        "github.com/Depado/ginprom"
        "github.com/gin-contrib/cors"

        // Our Stuff
        "cribbage/logutil"
        "cribbage/docs"
        "cribbage/api"
        "cribbage/core"

)

// VARS
const version = "0.1.2"
var component_name = "cribbage-svc-main"
var log = logutil.InitLogger(component_name)

// Global State for now
var (
	mu   sync.Mutex
	game *core.Game
)




// @title MCP Explorer - MCP Server APIs
// @version 0.0.1
// @description APIs for MCP Server Instantiation, Configuration and Handling
// @BasePath /


func main() {
        log.Infof("DASMLAB WhatsNew Service - Starting %s", component_name)
        docs.SwaggerInfo.Version = version

        // Set gin Prod mode
        gin.SetMode(gin.ReleaseMode)


        // Primary App Router
        mainRouter := gin.Default()

        // Allow CORS
        mainRouter.Use(cors.Default()) // Allows all - Depado rocks!

        // Metrics (out of band) Router
        metricsRouter := gin.Default()

        // ginprim hooks
        p := ginprom.New(
                ginprom.Engine(metricsRouter),
                ginprom.Subsystem("gin"),
                        ginprom.Path("/metrics"),
        )

        // Wrap our mainRouter
        mainRouter.Use(p.Instrument())

        // Add our Custom Metrics - turned off for now
        //metrics.RegisterCustomMetrics()

        // Add Swagger UI Route
        mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

        // Init Routes
        //initializeRoutes(mainRouter)
	mainRouter.GET("/deal", api.DealHandler)
	mainRouter.GET("/status", api.StatusHandler)
	mainRouter.POST("/reset", api.ResetHandler)

        // Launch metricsROuter
        go func() {
                log.Infof("Starting metrics server on :9201")
                if err := metricsRouter.Run(":9201"); err != nil {
                        log.Fatalf("Metrics Server Error: %v", err)
                }
        }()

        // Launch MainRouter


        log.Info("Start main Server listening on :8001")
        if err := mainRouter.Run(":8001"); err != nil {
                log.Fatalf("Main Server Error: %v", err)
        }

}

