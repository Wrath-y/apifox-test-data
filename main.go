package main

import (
	"apifox-test-data/base_m"
	"apifox-test-data/settlement"
	"encoding/json"
	"fmt"
)

func main() {
	finalData := make([]*base_m.Base, 0)
	finalData = append(finalData, &base_m.Base{
		Token:   "951e8e31e96e599f678e17a1d07508d9",
		BodyI:   new(settlement.Body).SetBody([]int{15934, 15897, 15878}, "CN9999999", "210126288755961791"),
		ExpectI: new(settlement.Expect).SetExpect(1, []string{"210126288755961791"}),
	})
	r := new(base_m.ResultSet)
	r.Data = finalData

	for _, v := range finalData {
		v.Body = v.BodyI.GetBody()
		v.Expect = v.ExpectI.GetExpect()
	}

	jsonBytes, err := json.Marshal(finalData)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON string:", string(jsonBytes))
}
