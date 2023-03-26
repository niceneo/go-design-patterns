package factory

import "fmt"

type Restaurant interface {
	Getfood()
}

type Donglaishun struct {}

func (d *Donglaishun) Getfood(){
	fmt.Println("东来顺饭店")
}
type Qingfeng struct {}

func (q Qingfeng) Getfood(){
	fmt.Println("庆丰饭店")
}

func NewRestaurant(name string) Restaurant{
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}

