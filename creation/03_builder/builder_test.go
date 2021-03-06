package builder

import (
	"fmt"
	"testing"
)

func TestBuilderCar(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("sky").Speed(120).Engine("audi")
	car := builder.Build()
	if car.Speed() != 120 {
		t.Fatalf("Builder1 fail expect 120 ,but get %d", car.Speed())
	}
	if car.Brand() != "sky" {
		t.Fatalf("Builder1 fail expect sky ,but get %s", car.Brand())
	}

	fmt.Println(car.Speed())
	fmt.Println(car.Brand())
}

func TestBuilderCarMore(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("land").Speed(110).Engine("bmw")
	builder.Engine("man made").Brand("panda").Wheel(15)
	car := builder.Build()

	fmt.Println(car.Speed())
	fmt.Println(car.Brand())
	car.Brief()
}
