package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)

	// create data to decode
	err := decoder.Decode(result)
	IfErrorPanic(err)
}

func WriteResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	IfErrorPanic(err)
}
