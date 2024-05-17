package calculator

import (
	"fmt"
	"time"
)

type TollFreeVehicles int

type timeRange struct {
	startHour, startMinute int
	endHour, endMinute     int
	fee                    int
}

const (
	Motorcycle TollFreeVehicles = iota
	Tractor
	Emergency
	Diplomat
	Foreign
	Military
)

var TabelaDeTaxas = []timeRange{
	{6, 0, 6, 29, 8},
	{6, 30, 6, 59, 13},
	{7, 0, 7, 59, 18},
	{8, 0, 8, 29, 13},
	{8, 30, 14, 59, 8},
	{15, 0, 15, 29, 13},
	{15, 30, 16, 59, 18},
	{17, 0, 17, 59, 13},
	{18, 0, 18, 29, 8},
}

func (tfv TollFreeVehicles) String() string {
	switch tfv {
	case Motorcycle:
		return "Motorcycle"
	case Tractor:
		return "Tractor"
	case Emergency:
		return "Emergency"
	case Foreign:
		return "Foreign"
	case Military:
		return "Military"
	default:
		return fmt.Sprintf("%d", int(tfv))
	}
}

func GetTax(vehicle Vehicle, dates []time.Time) int {
	intervalStart := dates[0]
	totalFee := 0
	for _, date := range dates {
		nextFee := GetTollFee(date, vehicle)
		tempFee := GetTollFee(date, vehicle)

		diffInNanos := date.UnixNano() - intervalStart.UnixNano()
		minutes := diffInNanos / 1000000 / 1000 / 60

		if minutes <= 60 {
			if totalFee > 0 {
				totalFee = totalFee - tempFee
			}
			if nextFee >= tempFee {
				tempFee = nextFee
			}
		} else {
			totalFee = totalFee + nextFee
		}
	}

	if totalFee > 60 {
		totalFee = 60
	}
	return totalFee
}

func IsTollFreeVehicle(v Vehicle) bool {
	switch v.getVehicleType() {
	case "Motorcycle", "Tractor", "Emergency", "Diplomat", "Foreign", "Military":
		return true
	default:
		return false
	}
}
func GetTollFee(t time.Time, v Vehicle) int {
	if IsTollFreeDate(t) || IsTollFreeVehicle(v) {
		return 0
	}

	hour, minute := t.Hour(), t.Minute()

	for _, tf := range TabelaDeTaxas {
		if (hour > tf.startHour || (hour == tf.startHour && minute >= tf.startMinute)) &&
			(hour < tf.endHour || (hour == tf.endHour && minute <= tf.endMinute)) {
			return tf.fee
		}
	}

	return 0
}

func IsTollFreeDate(date time.Time) bool {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	if year == 2013 {
		if month == 1 && day == 1 || month == 3 && (day == 28 || day == 29) || month == 4 && (day == 1 || day == 30) || month == 5 && (day == 1 || day == 8 || day == 9) || month == 6 && (day == 5 || day == 6 || day == 21) || month == 7 || month == 11 && day == 1 || month == 12 && (day == 24 || day == 25 || day == 26 || day == 31) {
			return true
		}
	}
	return false
}
