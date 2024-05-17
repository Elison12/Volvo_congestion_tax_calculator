package main

import (
	"congestion-calculator/calculator"
	"fmt"
	"time"
)

// "congestion-calculator/calculator"
// "fmt"
// // "time"
// "time"

func main() {
	// err := calculator.IsTollFreeDate(time.Time{})
	// fmt.Println(err)

	t := time.Date(2013, 5, 17, 8, 45, 0, 0, time.UTC)
	// dates := []time.Time{
	// 	time.Date(2013, 5, 17, 6, 0, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 6, 30, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 7, 0, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 8, 30, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 9, 0, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 15, 0, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 16, 0, 0, 0, time.Local),
	// 	// time.Date(2013, 5, 17, 17, 0, 0, 0, time.Local),
	// }

	// motocicleta := calculator.Motorbike{}
	// fee := calculator.GetTax(calculator.Car{}, dates)
	// tollfee := calculator.GetTollFee(t, calculator.Motorbike{})
	fmt.Println("a taxa a se pagar é:", calculator.GetTollFee(t, calculator.Van{}))
	// fmt.Println(motocicleta.GetVehicleType())
	// fmt.Println(calculator.IsTollFreeVehicle(vehicle))
}
