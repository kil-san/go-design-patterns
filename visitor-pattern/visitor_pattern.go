package visitor

import (
	"fmt"

	"github.com/go-design-patterns/model"
)

func NewFood(name string, opts ...Option) *model.Food {
	f := &model.Food{
		name: name,
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
		withCount(3),
		withFarmRate(5),
	}

	f := NewFood("Yam", opts...)

	fmt.Printf("We have: %+v\n", f)
}
