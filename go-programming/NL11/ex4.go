package main

import (
	"errors"
	"fmt"
	"log"
)

const lat = "50.2289 N"
const long = "99.4656 W"

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{lat: lat, long: long, err: errors.New("Less than 0")}
	}
	return 42, nil
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

}
