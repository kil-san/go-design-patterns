package factory

import (
	"fmt"

	visitor "github.com/go-design-patterns/visitor-pattern"
)

// FactoryDemo test factory pattern
func FactoryDemo() {
	fmt.Println("Factory Demo.........")
	yamOptions := []visitor.Option{
		visitor.WithCount(3),
		visitor.WithFarmRate(5),
	}

	cornOptions := []visitor.Option{
		visitor.WithCount(3),
		visitor.WithFarmRate(5),
	}

	factory := NewFoodFactory()

	yam := factory.CreateYam(yamOptions...)
	corn := factory.CreateCorn(cornOptions...)
	rice := factory.CreateRice()

	fmt.Printf("We have: %+v\n", yam)
	fmt.Printf("We have: %+v\n", corn)
	fmt.Printf("We have: %+v\n", rice)
}
