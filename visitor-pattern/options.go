package visitor

import (
	"github.com/kil-san/go-design-patterns/model"
)

type Option func(*model.Food)

func WithCount(n int) Option {
	return func(f *model.Food) {
		f.Count = n
	}
}

func WithFarmRate(n int) Option {
	return func(f *model.Food) {
		f.FarmRate = n
	}
}

func WithSellRate(n int) Option {
	return func(f *model.Food) {
		f.SellRate = n
	}
}
