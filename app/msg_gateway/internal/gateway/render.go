package gateway

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type ConnectResponse struct {
	Code    int
	Message string
}

func renderJSON(w http.ResponseWriter, resp *ConnectResponse) {
	body, _ := jsoniter.Marshal(resp)
	_, _ = w.Write(body)
}
