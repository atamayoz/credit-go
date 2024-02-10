package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/atamayoz/credit-go/ent"
)

type SimulatorService interface {
	GetSimulation(amount float64, interest float64, periods int) (float64, error)
}

type simulatorService struct {
	client *ent.Client
}

func NewSimulatorService(client *ent.Client) SimulatorService {
	return &simulatorService{
		client: client,
	}
}

func (s *simulatorService) GetSimulation(amount float64, interest float64, periods int) (float64, error) {

	// Here I calculate the PMT periodic payment (annuity payment).
	// PMT = (PV * r * (1 + r)^n) / ((1 + r)^n - 1)
	convInterest, err := roundToFourDecimals(interest / 100)

	if err != nil {
		return 0.0, err
	}

	pow := math.Pow((1 + convInterest), float64(periods))
	pmt := (amount * convInterest * pow) / (pow - 1)

	// Save the Simulation
	ctx := context.Background()

	err = s.client.Simulation.Create().
		SetAmount(amount).
		SetInterest(convInterest).
		SetMonthlyPayment(pmt).
		SetPeriods(float64(periods)).
		SetCreatedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		log.Println("An error occurred saving the simulation", err.Error())
		return 0.0, err
	}

	return pmt, nil
}

func roundToFourDecimals(number float64) (float64, error) {
	roundedStr := fmt.Sprintf("%.4f", number)
	return strconv.ParseFloat(roundedStr, 64)
}
