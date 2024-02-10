package srv

import (
	"time"

	"github.com/atamayoz/credit-go/internal/app/handlers"
	"github.com/atamayoz/credit-go/internal/app/services"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func initializeEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	return engine
}

func initializeSimulatorHandler() handlers.SimulatorHandler {
	simulatorService := services.NewSimulatorService()
	return handlers.NewSimulationHandler(simulatorService)
}

func StartServer() {
	router := initializeEngine()
	simulatorHandler := initializeSimulatorHandler()

	creditGroup := router.Group("/credit")
	creditGroup.GET("/simulator", simulatorHandler.GetCreditSimulation)

	router.Run(":8081")
}
