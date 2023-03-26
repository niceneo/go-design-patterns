package abstract_factory

import "fmt"

type Lunch interface {
	Cook()
}
type Rise struct {}
func (r *Rise) Cook(){
	fmt.Println("Rise cook")
}

type Tomato struct {}
func (t *Tomato) Cook(){
	fmt.Println("Tomato cook")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func NewSimpleLunchFactory() LunchFactory{
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory) CreateFood() Lunch{
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch{
	return &Tomato{}
}