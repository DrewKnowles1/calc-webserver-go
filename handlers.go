package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type calculation struct {
	FirstNumber  float64 `json: "firstnumber"`
	SecondNumber float64 `json: "secondnumber"`
	Operator     string  `json: "operator"`
}

func (app application) calculate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		message := fmt.Sprintf("method %v not supported", r.Method)
		http.Error(w, message, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body malformed, could not be read", http.StatusBadRequest)
		return
	}

	var calc calculation

	err = json.Unmarshal(body, &calc)
	if err != nil {
		http.Error(w, "body malformed, could not be read", http.StatusBadRequest)
		return
	}

	err = verifyCalculation(calc, w)
	if err != nil {
		return
	}

	switch calc.Operator {
	case "+":
		answer := calc.FirstNumber + calc.SecondNumber
		json.NewEncoder(w).Encode(answer)
	case "-":
		answer := calc.FirstNumber - calc.SecondNumber
		json.NewEncoder(w).Encode(answer)
	case "*":
		answer := calc.FirstNumber * calc.SecondNumber
		json.NewEncoder(w).Encode(answer)

	case "/":
		answer := calc.FirstNumber / calc.SecondNumber
		json.NewEncoder(w).Encode(answer)
	case "":
		http.Error(w, "case statement: operator cannot be blank", http.StatusBadRequest)
	}

}
