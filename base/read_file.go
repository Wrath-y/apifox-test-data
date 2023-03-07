package base

import (
	"encoding/json"
	"os"
)

type Row struct {
	Token  string `json:"token"`
	Body   string `json:"body"`
	Expect string `json:"expect"`
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
	return r.Body
}

func (r *Row) GetExpect() string {
	return r.Expect
}
