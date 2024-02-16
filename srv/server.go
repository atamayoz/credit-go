package srv

import (
	"log"
	"time"

	"github.com/atamayoz/credit-go/ent"
	"github.com/atamayoz/credit-go/infrastructure/db"
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

func initializeSimulatorHandler(client *ent.Client) handlers.SimulatorHandler {
	simulatorService := services.NewSimulatorService(client)
	return handlers.NewSimulationHandler(simulatorService)
}

func StartServer() {
	// Connect to de db
	client, err := db.Connect()

	if err != nil {
		log.Fatalf("failed connecting to DB: %v", err)
	}

	defer client.Close()

	router := initializeEngine()
	simulatorHandler := initializeSimulatorHandler(client)

	initializeCreditGroup(router, simulatorHandler)

	router.Run(":8081")
}

func initializeCreditGroup(router *gin.Engine, simulatorHandler handlers.SimulatorHandler) {

	const CREDIT_BASE_PATH = "/credit"
	const SIMULATOR_PATH = "/payment/simulator"
	const AMORTIZATION_PATH = "/payment/amortization"

	creditGroup := router.Group(CREDIT_BASE_PATH)
	creditGroup.GET(SIMULATOR_PATH, simulatorHandler.GetMonthlyPayment)
	creditGroup.GET(AMORTIZATION_PATH, simulatorHandler.GetAmortizationTable)
}
