package main

import "log"

var total int

func main() {
	moduleFuelReq := []int{}
	for _, mass := range modulesMass {
		moduleFuelReq = append(moduleFuelReq, fuelCalc(mass) + fuelForFuel(fuelCalc(mass)))
	}

	log.Println(fuelTotal(moduleFuelReq))
}

func fuelCalc(mass int) int {
	mass1 := mass / 3
	fuelReq := mass1 - 2
	if fuelReq < 1 {
		return 0
	}
	return fuelReq
}

func fuelTotal(moduleFuelReq []int) int {
	fuelTotal := 0
	for _, fuelReq := range moduleFuelReq {
		fuelTotal = fuelTotal + fuelReq
	}
	return fuelTotal
}

func fuelForFuel(mass int, total ...int) int {
	var acutalTotal int
	if len(total) > 1 {
		acutalTotal = total[0]
	} else {
		acutalTotal = 0
	}
	fuel := fuelCalc(mass)
	if fuel < 1 {
		return acutalTotal
	}
	return fuel + fuelForFuel(fuel, acutalTotal)
}