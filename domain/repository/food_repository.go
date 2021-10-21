package repository

import "food-app/domain/entity"

type FoodRepository interface {
	SaveFood(*entity.Food) (*entity.Food, map[string]string)
	GetFood(uint64) (*entity.Food, error)
}
