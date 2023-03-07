package base

import (
	"encoding/json"
	"os"
)

type Row struct {
	Token  string      `json:"token"`
	Body   interface{} `json:"body"`
	Expect interface{} `json:"expect"`
}

// ReadFile 读取配置好的json文件
// json文件是一个对象数组,每个对象中必须有token、body、expect这三个字段
func (r *Row) ReadFile(filename string) []*Row {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var rows []*Row
	if err := json.Unmarshal(b, &rows); err != nil {
		panic(err)
	}

	return rows
}

func (r *Row) GetBody() string {
	var err error
	switch r.Body.(type) {
	case string:
		return r.Body.(string)
	default:
		r.Body, err = json.Marshal(r.Body)
		if err != nil {
			panic(err)
		}
		return string(r.Body.([]byte))
	}
}

func (r *Row) GetExpect() string {
	var err error
	switch r.Expect.(type) {
	case string:
		return r.Expect.(string)
	default:
		r.Expect, err = json.Marshal(r.Expect)
		if err != nil {
			panic(err)
		}
		return string(r.Expect.([]byte))
	}
}
