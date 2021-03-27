package visitor

import (
	"fmt"

	"github.com/kil-san/go-design-patterns/model"
)

func NewFood(name string, opts ...Option) *model.Food {
	f := &model.Food{
		Name: name,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(f)
		}
	}

	return f
}

// VisitorDemo test visitor pattern
func VisitorDemo() {
	fmt.Println("Visitor Demo.........")
	opts := []Option{
		WithCount(3),
		WithFarmRate(5),
	}

	f := NewFood("Yam", opts...)

	fmt.Printf("We have: %+v\n", f)
}
