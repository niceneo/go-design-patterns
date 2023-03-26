package facade

import "testing"
//外观模式
func TestNewCarFacade(t *testing.T) {
	car := CarFacade{}
	car.CreateCompleteCar()
}