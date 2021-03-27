package main

import (
	factory "github.com/go-design-patterns/factory-pattern"
	generator "github.com/go-design-patterns/generator-pattern"
	visitor "github.com/go-design-patterns/visitor-pattern"
)

func main() {
	visitor.VisitorDemo()
	generator.GeneratorDemo()
	factory.FactoryDemo()
}
