package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type formRequest struct {
	FormId         string `json:"form_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	StockId        uint64 `json:"stock_id"`
	DestStockId    uint64 `json:"dest_stock_id"`
	DestUserId     uint16 `json:"dest_user_id"`
	CompanyId      uint16 `json:"company_id"`
	PartnerId      uint16 `json:"partner_id"`
	PartnerTypeId  uint16 `json:"partner_type_id"`
	OrganizationId uint16 `json:"partner_organization_id"`
	ContractId     uint32 `json:"partner_contract_id"`
	ContractType   uint8  `json:"partner_contract_type"`
	Reason         string `json:"reason,omitempty"`

	Data []formRequestItem `json:"detail"`
}

type formRequestItem struct {
	ContractItemId     uint64  `json:"partner_contract_item_id"`
	SpecId             int16   `json:"resource_spec_id"`
	StatusId           int16   `json:"resource_status_id"`
	Info10             string  `json:"info10,omitempty"`
	ContractItemAmount float64 `json:"partner_contract_item_amount"`
	DeliveryNumber     int16   `json:"partner_contract_item_delivery_number"`
	ResourceNumber     int16   `json:"resource_number"`
}

type formResp struct {
	Status  uint16 `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func TestResourceForm() {
	rt := &formRequest{
		FormId:         "11",
		UserId:         "12",
		StockId:        1000001,
		DestStockId:    9090909,
		DestUserId:     12,
		CompanyId:      2,
		PartnerId:      10194,
		PartnerTypeId:  3,
		OrganizationId: 10194,
		ContractId:     3967063,
		ContractType:   6,
		Reason:         "abc",
		Data: []formRequestItem{
			{
				ContractItemId:     0,
				SpecId:             2,
				StatusId:           3,
				Info10:             "abc",
				ContractItemAmount: 20,
				DeliveryNumber:     1,
				ResourceNumber:     1,
			},
		},
	}
	testResourceForm(rt, "成功用例")

	rt1 := &formRequest{}
	DeepCopy(rt, rt1)
	rt1.FormId = ""
	testResourceForm(rt1, "参数FormId传空")

	rt2 := &formRequest{}
	DeepCopy(rt, rt2)
	rt2.UserId = ""
	testResourceForm(rt2, "参数UserId传空")
}

func testResourceForm(rt *formRequest, v string) formResp {
	Logger.Infof("========== >>> C003 销售出库申请 [%s] start.", v)
	startTestResourceForm := time.Now()
	defer func() {
		Logger.Infof("========== >>> C003 销售出库申请 [%s] end. times %v.", v, time.Now().Sub(startTestResourceForm))
	}()

	formRt, _ := json.Marshal(&rt)

	Logger.Println("request params ", string(formRt))

	request, err := http.NewRequest("POST", url, bytes.NewReader(formRt))
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

	var resp formResp
	er := json.Unmarshal(body, &resp)
	if er != nil {
		Logger.Fatal(er)
	}
	//Logger.Info("resp ", resp)
	return resp
}
