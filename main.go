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
	rowsData := make([]*base.ApiFoxRowData, 0)
	rowsData = append(rowsData, &base.ApiFoxRowData{
		Token:  "efdd45da273b99550d3a9b00b1a3be8c", // token ttl已经设置为-1 user: wrath
		Body:   new(couponlist.Body).SetBody([]int{15934, 15897, 15878}, "CN9999999", "210126288755961791"),
		Expect: new(couponlist.Expect).SetExpect(1, []string{"210126288755961791"}),
	})
	Done(rowsData, "apifox_import_data.json")

	// 方式二 导入写好的json文件
	rowsData = make([]*base.ApiFoxRowData, 0)
	for _, v := range new(base.Row).ReadFile("testing_data.json") {
		rowsData = append(rowsData, &base.ApiFoxRowData{
			Token:  v.Token,
			Body:   v,
			Expect: v,
		})
	}
	Done(rowsData, "apifox_import_data.json")
}

func Done(rowsData []*base.ApiFoxRowData, exportFilename string) {
	finalData := make([]*base.ApiFoxImportJson, 0, len(rowsData))

	for _, v := range rowsData {
		finalData = append(finalData, &base.ApiFoxImportJson{
			Token:  v.Token,
			Body:   v.Body.GetBody(),
			Expect: v.Expect.GetExpect(),
		})
	}

	jsonBytes, err := json.Marshal(finalData)
	if err != nil {
		panic(err)
	}

	GenerateFile(exportFilename, jsonBytes)
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
