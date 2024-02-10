package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/atamayoz/credit-go/internal/app/services"
	"github.com/gin-gonic/gin"
)

// SimulationHandler interface define the signature of the credit simulation
type SimulatorHandler interface {
	GetCreditSimulation(c *gin.Context)
}

// This is private struct with representing the simulation handler
type simulatorHandler struct {
	service services.SimulatorService
}

// As you can see here in the new functions I recurn an interface not a concrete implementation
func NewSimulationHandler(service services.SimulatorService) SimulatorHandler {
	return &simulatorHandler{
		service: service,
	}
}

func (simulator *simulatorHandler) GetCreditSimulation(c *gin.Context) {
	// In this part I get the query params
	amount, err := strconv.ParseFloat(c.Query("amount"), 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid amount"))
	}

	interest, err := strconv.ParseFloat(c.Query("interest"), 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid interest"))
	}

	periods, err := strconv.Atoi(c.Query("periods"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid periods"))
	}

	log.Println("amount: ", amount)
	log.Println("interest: ", interest)
	log.Println("periods: ", periods)

	simulation, err := simulator.service.GetSimulation(amount, interest, periods)

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, errors.New("error simulating"))
	}

	c.JSON(http.StatusOK, simulation)
}
