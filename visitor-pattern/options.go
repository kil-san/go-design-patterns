package visitor

import (
	"github.com/go-design-patterns/model"
)

type Option func(*model.Food)

func WithCount(n int) Option {
	return func(f *model.Food) {
		f.count = n
	}
}

func WithFarmRate(n int) Option {
	return func(f *model.Food) {
		f.farmRate = n
	}
}

func WithSellRate(n int) Option {
	return func(f *model.Food) {
		f.sellRate = n
	}
}
