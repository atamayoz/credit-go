package services

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/atamayoz/credit-go/ent"
	"github.com/atamayoz/credit-go/internal/app/dto"
	"github.com/atamayoz/credit-go/internal/app/util/numbers"
)

type SimulatorService interface {
	GetMonthlyPayment(amount float64, interest float64, periods int) (float64, error)
	GetAmortizationTable(amount float64, interest float64, periods int) (*dto.AmortizationTable, error)
}

type simulatorService struct {
	client *ent.Client
}

func NewSimulatorService(client *ent.Client) SimulatorService {
	return &simulatorService{
		client: client,
	}
}

func (s *simulatorService) GetMonthlyPayment(amount float64, interest float64, periods int) (float64, error) {

	// Here I calculate the PMT periodic payment (annuity payment).
	// PMT = (PV * r * (1 + r)^n) / ((1 + r)^n - 1)
	convInterest, err := numbers.RoundToFourDecimals(interest / 100)

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

func (s *simulatorService) GetAmortizationTable(amount float64, interest float64, periods int) (*dto.AmortizationTable, error) {

	payment, err := s.GetMonthlyPayment(amount, interest, periods)

	if err != nil {
		return nil, err
	}

	remainingBalance := amount
	convInterest, err := numbers.RoundToFourDecimals(interest / 100)
	payment = numbers.RoundToTwoDecimal(payment)

	amortizationTable := dto.AmortizationTable{
		Amount:   amount,
		Interest: float32(convInterest),
		Periods:  periods,
		Payments: nil,
	}

	amortizationTable.Payments = make([]*dto.PaymentDetail, amortizationTable.Periods)

	if err != nil {
		return nil, err
	}

	for i := 0; i <= periods-1; i++ {
		interestAmount := numbers.RoundToTwoDecimal(remainingBalance * convInterest)
		principal := numbers.RoundToTwoDecimal(payment - interestAmount)
		remainingBalance = numbers.RoundToTwoDecimal(remainingBalance - principal)

		amortizationTable.Payments[i] = &dto.PaymentDetail{
			Installment:     i + 1,
			Principal:       principal,
			InterestAmount:  interestAmount,
			PaymentDetail:   payment,
			RemainingAmount: remainingBalance,
		}
	}

	return &amortizationTable, nil
}
