package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type StockRequest struct {
	CompanyId uint16      `json:"company_id,omitempty"`
	Data      []StockItem `json:"detail,omitempty"`
}

type StockItem struct {
	SpecId         int64 `json:"resource_spec_id,omitempty"`
	StatusId       int16 `json:"resource_status_id,omitempty"`
	ResourceNumber int16 `json:"resource_number,omitempty"`
}

type StockResp struct {
	Status  uint16 `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		CompanyId uint64 `json:"company_id"`
		StockId   uint64 `json:"stock_id"`
		StockName string `json:"stock_name"`
		StockCode string `json:"stock_code"`
		Detail    []struct {
			SpecId          uint64  `json:"resource_spec_id"`
			StatusId        uint64  `json:"resource_status_id"`
			ActualNumber    float64 `json:"actual_number"`
			StockLocationId uint64  `json:"stock_location_id"`
		}
	} `json:"data"`
}

func TestC021GetStock() {
	rt := &StockRequest{
		CompanyId: 5,
		Data: []StockItem{
			{
				SpecId:         123456,
				StatusId:       1,
				ResourceNumber: 20,
			},
		},
	}
	testC021GetStock(rt, "正常通过")
}

func testC021GetStock(rt *StockRequest, v string) StockResp {
	//accessCode := Loggerin()
	//Logger.Printf(" Loggerin success !!! %s \n\n", accessCode)

	Logger.Infof("========== >>> C021 销售出库申请获取仓库列表 [%s] start.", v)
	startTestRoTestC021GetStock := time.Now()
	defer func() {
		defer Logger.Printf("========== >>> C021 销售出库申请获取仓库列表 [%s] end. times %v ", v, time.Now().Sub(startTestRoTestC021GetStock))
	}()

	formRt, _ := json.Marshal(&rt)

	Logger.Println("request params ", string(formRt))

	request, err := http.NewRequest("POST", getStockUrl, bytes.NewReader(formRt))
	if err != nil {
		Logger.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access_token", AccessCode)
	request.Header.Set("language", "en")

	client := &http.Client{}

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	Logger.Printf("resp body is %s \n", string(body))

	var resp StockResp
	er := json.Unmarshal(body, &resp)
	if er != nil {
		Logger.Fatal(er)
	}
	//dd, _ := json.MarshalIndent(resp, "", "")
	//fmt.Println("resp indent is ", string(dd))
	return resp
}
