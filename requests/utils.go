package requests

import "encoding/json"

func DeepCopy(a, b interface{}) {
	byt, _ := json.Marshal(a)
	err := json.Unmarshal(byt, b)
	if err != nil {
		panic(err)
	}
}
