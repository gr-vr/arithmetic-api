package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type response struct {
	Success bool `json:"success"`
	ErrCode int  `json:"errCode,omitempty"`
	Value   int  `json:"value,omitempty"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var res response
	argOne, argTwo, err := convToInt(r)
	if err != nil {
		res = response{
			Success: false,
			ErrCode: 500,
		}
	} else {
		res = response{
			Success: true,
			Value:   argOne + argTwo,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func Sub(w http.ResponseWriter, r *http.Request) {
	var res response
	argOne, argTwo, err := convToInt(r)
	if err != nil {
		res = response{
			Success: false,
			ErrCode: 500,
		}
	} else {
		res = response{
			Success: true,
			Value:   argOne - argTwo,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func Mul(w http.ResponseWriter, r *http.Request) {
	var res response
	argOne, argTwo, err := convToInt(r)
	if err != nil {
		res = response{
			Success: false,
			ErrCode: 500,
		}
	} else {
		res = response{
			Success: true,
			Value:   argOne * argTwo,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func Div(w http.ResponseWriter, r *http.Request) {
	var res response
	argOne, argTwo, err := convToInt(r)
	if err != nil || argTwo == 0 {
		res = response{
			Success: false,
			ErrCode: 500,
		}
	} else {
		res = response{
			Success: true,
			Value:   argOne / argTwo,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func convToInt(r *http.Request) (int, int, error) {
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	return a, b, err
}
