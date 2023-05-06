package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	//request의 body를 byte slice로 읽어온다.
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//body를 json으로 unmarshal한다.
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
