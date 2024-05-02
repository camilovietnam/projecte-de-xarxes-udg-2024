package function

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("[%s] %s", serverName, string(input))))
}
