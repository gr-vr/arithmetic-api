package router

import (
	"encoding/json"
	"github.com/satanCoffee/arithmetic-api/middleware"
	"log"
	"net/http/httptest"
	"testing"
)

type response struct {
	Success bool `json:"success"`
	ErrCode int  `json:"errCode,omitempty"`
	Value   int  `json:"value,omitempty"`
}

func TestHandler_Add(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/add?a=42&b=20", nil)
	w := httptest.NewRecorder()
	middleware.Add(w, r)
	res := response{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 62 {
		log.Fatalf("Invalid result %d expected %d", res.Value, 62)
	}
}

func TestHandler_Sub(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/sub?a=42&b=20", nil)
	w := httptest.NewRecorder()
	middleware.Sub(w, r)
	res := response{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 22 {
		log.Fatalf("Invalid result %d expected %d", res.Value, 22)
	}
}

func TestHandler_Mul(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/add?a=6&b=2", nil)
	w := httptest.NewRecorder()
	middleware.Mul(w, r)
	res := response{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 12 {
		log.Fatalf("Invalid result %d expected %d", res.Value, 12)
	}
}

func TestHandler_Div(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/add?a=6&b=2", nil)
	w := httptest.NewRecorder()
	middleware.Div(w, r)
	res := response{}
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Value != 3 {
		log.Fatalf("Invalid result %d expected %d", res.Value, 3)
	}
}
