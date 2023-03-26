package facade

import "fmt"

//外观模式
type CarModel struct{}
func NewCarModel() *CarModel{
	return &CarModel{}
}
func (c *CarModel) SetModel(){
	fmt.Println("carmodel - set Model")
}

type CarEngine struct {}
func NewCarEngine() *CarEngine{
	return &CarEngine{}
}
func (c *CarEngine) SetEngine(){
	fmt.Println("carengine --setengine")
}

type CarBody struct{}
func NewCarBody() *CarBody{
	return &CarBody{}
}
func (c *CarBody) SetCarBody(){
	fmt.Println("Carbody---setcarbody")
}

type CarFacade struct {
	model CarModel
	engine CarEngine
	body CarBody
}
func NewCarFacade() *CarFacade{
	return &CarFacade{
		model: CarModel{},
		body: CarBody{},
		engine: CarEngine{},
	}
}
func (c *CarFacade)CreateCompleteCar(){
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}