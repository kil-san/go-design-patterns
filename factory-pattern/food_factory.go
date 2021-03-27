package factory

import (
	"github.com/kil-san/go-design-patterns/model"
	visitor "github.com/kil-san/go-design-patterns/visitor-pattern"
)

type foodFactory struct {
}

func NewFoodFactory() *foodFactory {
	return &foodFactory{}
}

func (f *foodFactory) CreateYam(options ...visitor.Option) *model.Food {
	return visitor.NewFood("Yam", options...)
}

func (f *foodFactory) CreateCorn(options ...visitor.Option) *model.Food {
	return visitor.NewFood("Fish", options...)
}

func (f *foodFactory) CreateRice(options ...visitor.Option) *model.Food {
	return visitor.NewFood("Fish", options...)
}
