package main

import (
	"apifox-test-data/base"
	"apifox-test-data/settlement/couponlist"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 方式一 直接写结构体
	finalData := make([]*base.Base, 0)
	finalData = append(finalData, &base.Base{
		Token:   "efdd45da273b99550d3a9b00b1a3be8c", // token ttl已经设置为-1 user: wrath
		BodyI:   new(couponlist.Body).SetBody([]int{15934, 15897, 15878}, "CN9999999", "210126288755961791"),
		ExpectI: new(couponlist.Expect).SetExpect(1, []string{"210126288755961791"}),
	})
	GetJson(finalData, "test_by_struct.json")

	// 方式二 导入写好的json文件
	finalData = make([]*base.Base, 0)
	for _, v := range new(base.Row).ReadFile("test.json") {
		finalData = append(finalData, &base.Base{
			Token:   v.Token,
			BodyI:   v,
			ExpectI: v,
		})
	}
	GetJson(finalData, "test_by_file.json")
}

func GetJson(finalData []*base.Base, filename string) {
	for _, v := range finalData {
		v.Body = v.BodyI.GetBody()
		v.Expect = v.ExpectI.GetExpect()
	}

	jsonBytes, err := json.Marshal(finalData)
	if err != nil {
		panic(err)
	}

	GenerateFile(filename, jsonBytes)
	fmt.Println("JSON string:", string(jsonBytes))
}

func GenerateFile(filename string, data []byte) {
	// 检查文件是否存在
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
			return
		}
		defer file.Close()

		// 写入数据
		_, err = file.Write(data)
		if err != nil {
			panic(err)
			return
		}

		return
	}
	// 文件存在，打开文件进行写入
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write(data)
	if err != nil {
		panic(err)
		return
	}
}
