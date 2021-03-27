package main

import (
	factory "github.com/kil-san/go-design-patterns/factory-pattern"
	generator "github.com/kil-san/go-design-patterns/generator-pattern"
	visitor "github.com/kil-san/go-design-patterns/visitor-pattern"
)

func main() {
	visitor.VisitorDemo()
	generator.GeneratorDemo()
	factory.FactoryDemo()
}
