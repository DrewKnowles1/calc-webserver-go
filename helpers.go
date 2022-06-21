package main

import (
	"fmt"
	"net/http"
)

func verifyCalculation(calc calculation, w http.ResponseWriter) error {
	if calc.Operator == "" {
		http.Error(w, "http error: operator cannot be blank", http.StatusBadRequest)
		err := fmt.Errorf("verifycalculation: operator cannot be blank")
		return err
	}
	if calc.Operator == "/" && calc.FirstNumber == 0 || calc.SecondNumber == 0 {
		http.Error(w, "dividing by zero is very naughty", http.StatusBadRequest)
		err := fmt.Errorf("dividing by zero is very naughty")
		return err
	}
	return nil
}
