package couponlist

import "encoding/json"

type Body struct {
	OrderProduct []*OrderProduct `json:"order_product"`
	StoreSn      string          `json:"store_sn"`
	TicketCode   string          `json:"ticket_code"`
}

type OrderProduct struct {
	CartId     int    `json:"cart_id"`
	TicketCode string `json:"ticket_code"`
}

func (r *Body) SetBody(cardIds []int, storeSn, ticketCode string) *Body {
	for _, v := range cardIds {
		r.OrderProduct = append(r.OrderProduct, &OrderProduct{
			CartId:     v,
			TicketCode: "",
		})
	}
	r.StoreSn = storeSn
	r.TicketCode = ticketCode

	return r
}

func (r *Body) GetBody() string {
	bodyB, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(bodyB)
}
