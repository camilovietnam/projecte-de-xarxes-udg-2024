package function

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type httpResponse struct {
	Server string
	Body   string
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var (
		input      []byte
		serverName = os.Getenv("SERVER_NAME")
	)

	if r.Body != nil {
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		input = body
	}

	_json, _ := json.Marshal(httpResponse{
		Server: serverName,
		Body:   string(input),
	})

	w.Header().Add("Content-type", "applcation/json")
	w.Write(_json)
}
