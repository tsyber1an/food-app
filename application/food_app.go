package application

import (
	"food-app/domain/entity"
	"food-app/domain/repository"
)

type foodApp struct {
	fr repository.FoodRepository
}

var _ FoodAppInterface = &foodApp{}

type FoodAppInterface interface {
	SaveFood(*entity.Food) (*entity.Food, map[string]string)
	GetFood(uint64) (*entity.Food, error)
}

func (f *foodApp) SaveFood(food *entity.Food) (*entity.Food, map[string]string) {
	return f.fr.SaveFood(food)
}
func (f *foodApp) GetFood(foodId uint64) (*entity.Food, error) {
	return f.fr.GetFood(foodId)
}
