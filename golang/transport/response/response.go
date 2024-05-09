package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ResponseMessage string      `json:"responseMessage"`
	Answer          interface{} `json:"answer,omitempty"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	WriteResponse(w, code, Response{
		ResponseMessage: "error: " + message,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	WriteResponse(w, code, Response{
		ResponseMessage: "success",
		Answer:          payload,
	})
}

func WriteResponse(w http.ResponseWriter, code int, response Response) {
	respJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(respJSON)
}
