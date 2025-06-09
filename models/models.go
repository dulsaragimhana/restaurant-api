package models

type CreateOrderRequest struct {
	TableNumber string      `json:"table"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	Name      string     `json:"name"`
	Comment   string     `json:"comment"`
	UnitPrice int        `json:"unit_price"`
	Quantity  int        `json:"quantity"`
	Discounts []Discount `json:"discounts"`
	Modifiers []Modifier `json:"modifiers"`
}

type Discount struct {
	Name         string `json:"name"`
	IsPercentage bool   `json:"is_percentage"`
	Value        int    `json:"value"`
	Amount       int    `json:"amount"`
}

type Modifier struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Amount    int    `json:"amount"`
}

type OrderResponse struct {
	ID       string      `json:"id"`
	OpenedAt string      `json:"opened_at"`
	IsClosed bool        `json:"is_closed"`
	Table    string      `json:"table"`
	Items    []OrderItem `json:"items"`
	Totals   Totals      `json:"totals"`
}

type Totals struct {
	Discounts     int `json:"discounts"`
	Due           int `json:"due"`
	Tax           int `json:"tax"`
	ServiceCharge int `json:"service_charge"`
	Paid          int `json:"paid"`
	Tips          int `json:"tips"`
	Total         int `json:"total"`
}

type PaymentRequest struct {
	BillAmount float64 `json:"billAmount"`
	TipAmount  float64 `json:"tipAmount"`
	PaymentID  string  `json:"paymentId"`
}
