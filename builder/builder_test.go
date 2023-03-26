package builder

import (
	"fmt"
	"testing"
)

//制造者模式
func TestNewConcreteBuilder(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewDirector(&builder)
	director.Construct()
	result := builder.GetResult()
	fmt.Println(result.Built)
}
