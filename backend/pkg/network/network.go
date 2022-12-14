package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tech-db-forum/app/models"
)

func WriteErrorResponse(w http.ResponseWriter, modelError *models.ModelError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(modelError.ErrorCode)
	marshalBody, err := json.Marshal(&modelError)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(marshalBody)
}

func WriteResponse(w http.ResponseWriter, code int, body interface{ MarshalJSON() ([]byte, error) }) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	marshalBody, err := body.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(marshalBody) == "null" {
		w.Write([]byte("[]"))
		return
	}
	w.Write(marshalBody)
}
