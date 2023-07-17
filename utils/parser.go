package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Parser(rq *http.Request, conttent interface{}){
  if body, err := io.ReadAll(rq.Body); err == nil {
    if err := json.Unmarshal([]byte(body), conttent); err == nil {
      return
    }
  }
}