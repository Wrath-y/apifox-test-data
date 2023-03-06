package settlement

import "encoding/json"

type Expect struct {
	CouponsLength int       `json:"coupons.length"`
	Coupons       []*Coupon `json:"coupons"`
}

type Coupon struct {
	TicketCode string `json:"ticket_code"`
}

func (e *Expect) SetExpect(couponsLength int, codes []string) *Expect {
	e.CouponsLength = couponsLength
	for _, v := range codes {
		e.Coupons = append(e.Coupons, &Coupon{TicketCode: v})
	}
	return e
}

func (e *Expect) GetExpect() string {
	expectB, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(expectB)
}
