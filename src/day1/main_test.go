package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFuelCalc(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, fuelCalc(12), "12/3 == 4 - 2 == 2")
	assert.Equal(2, fuelCalc(14), "14/3 == 4 - 2 == 2")
	assert.Equal(654, fuelCalc(1969), "1969/3 == 656 - 2 == 654")
	assert.Equal(33583, fuelCalc(100756), "100756/3 == 33585 - 2 == 33583")
}

func TestFuelTotal(t *testing.T) {
	assert := assert.New(t)
	
	assert.Equal(5, fuelTotal([]int{1, 1, 1, 1, 1}))
}

func TestFuelForFuel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(966, fuelForFuel(1969), "654 + 216 + 70 + 21 + 5 = 966")
	assert.Equal(50346, fuelForFuel(100756), "33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346")
}