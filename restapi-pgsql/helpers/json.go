package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json") //ini kita kasih untuk memberi tahu header bahwa bentuknya adalah json
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicError(err)
}