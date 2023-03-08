package base

type ApiFoxImportJson struct {
	Token  string `json:"token"`
	Body   string `json:"body"`
	Expect string `json:"expect"`
}

type ApiFoxRowData struct {
	Token  string
	Body   Body
	Expect Expect
}

type Body interface {
	GetBody() string
}

type Expect interface {
	GetExpect() string
}
