package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Message struct {
	Message string `json:"message"`
}

func Echo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var message Message
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		_ = json.NewDecoder(r.Body).Decode(&message)
		fmt.Println(string(requestDump))
		fmt.Println(message.Message)
		json.NewEncoder(w).Encode(message)
	})

}
