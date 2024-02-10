package services

import (
	"fmt"
	"math"
	"strconv"
)

type SimulatorService interface {
	GetSimulation(amount float64, interest float64, periods int) (float64, error)
}

type simulatorService struct {
}

func NewSimulatorService() SimulatorService {
	return &simulatorService{}
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

	return pmt, nil
}

func roundToFourDecimals(number float64) (float64, error) {
	roundedStr := fmt.Sprintf("%.4f", number)
	return strconv.ParseFloat(roundedStr, 64)
}
