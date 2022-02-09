package test

import (
	"fmt"
	"github.com/chenyi/hello-go/model"
	"testing"
)

func TestCar(t *testing.T) {
	car := model.Car{Name: "Chen Yi"}
	fmt.Printf("the car is: %s", car.Name)
}
