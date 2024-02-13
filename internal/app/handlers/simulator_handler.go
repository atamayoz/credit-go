package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/atamayoz/credit-go/internal/app/services"
	"github.com/gin-gonic/gin"
)

// SimulationHandler interface define the signature of the credit simulation
type SimulatorHandler interface {
	GetMonthlyPayment(c *gin.Context)
	GetAmortizationTable(c *gin.Context)
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

func (simulator *simulatorHandler) GetMonthlyPayment(c *gin.Context) {
	// In this part I get the query params
	amount, interest, periods := extractMandatoryParams(c)

	payment, err := simulator.service.GetMonthlyPayment(amount, interest, periods)

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, errors.New("error in monthly payment"))
	}

	c.JSON(http.StatusOK, payment)
}

func extractMandatoryParams(c *gin.Context) (float64, float64, int) {
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
	return amount, interest, periods
}

func (simulator *simulatorHandler) GetAmortizationTable(c *gin.Context) {
	// In this part I get the query params
	amount, interest, periods := extractMandatoryParams(c)

	table, err := simulator.service.GetAmortizationTable(amount, interest, periods)

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, errors.New("error in monthly payment"))
	}

	c.JSON(http.StatusOK, table)

}
