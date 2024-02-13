package dto

type AmortizationTable struct {
	Amount   float64          `json:"amount"`
	Interest float32          `json:"interest"`
	Periods  int              `json:"periods"`
	Payments []*PaymentDetail `json:"payments"`
}

type PaymentDetail struct {
	Installment     int     `json:"installment"`
	Principal       float64 `json:"principal"`
	InterestAmount  float64 `json:"interestAmount"`
	PaymentDetail   float64 `json:"paymentDetail"`
	RemainingAmount float64 `json:"remainingAmount"`
}
