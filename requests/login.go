package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var AccessCode string

type LoginRequest struct {
	User      string `json:"user,omitempty"`
	Password  string `json:"password,omitempty"`
	LoginType string `json:"logintype,omitempty"`
	AccountId string `json:"accountId,omitempty"`
	TenantId  string `json:"tenantid,omitempty"`
}

type LoginResp struct {
	State     string `json:"state,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	Data      struct {
		AccessToken string `json:"access_token" json:"access_token,omitempty"`
	} `json:"data"`
}

func init() {
	Logger.Println("Login init")
	Login()

}

func Login() string {
	Logger.Println("========== >>> Login start.")
	start := time.Now()
	defer func() {
		Logger.Printf("========== >>> Login end. times %v ", time.Now().Sub(start))
	}()

	rt := &LoginRequest{
		User:      User,
		Password:  Password,
		LoginType: LoginType,
		AccountId: AccountId,
		TenantId:  TenantId,
	}
	rtJson, _ := json.Marshal(&rt)

	Logger.Printf("Login request params %s", string(rtJson))

	request, _ := http.NewRequest("POST", LoginUrl, bytes.NewReader(rtJson))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	Logger.Printf("resp is %v", string(body))

	var resp LoginResp
	err := json.Unmarshal(body, &resp)
	if err != nil {
		Logger.Fatal(err)
	}
	if resp.State != "success" {
		Logger.Infof("resp state is " + resp.ErrorMsg)
	}
	//Logger.Println(json.MarshalIndent(resp, "", ""))
	AccessCode = resp.Data.AccessToken
	return resp.Data.AccessToken
}
