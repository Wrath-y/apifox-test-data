package base

type Base struct {
	Token   string  `json:"token"`
	BodyI   BodyI   `json:"-"`
	Body    string  `json:"body"`
	ExpectI ExpectI `json:"-"`
	Expect  string  `json:"expect"`
}

type BodyI interface {
	GetBody() string
}

type ExpectI interface {
	GetExpect() string
}
