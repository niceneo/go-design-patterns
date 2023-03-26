package factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("d").Getfood()
	NewRestaurant("q").Getfood()
}
