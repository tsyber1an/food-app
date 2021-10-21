package persistence

import (
	"errors"
	"food-app/domain/entity"
	"food-app/domain/repository"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

type FoodRepo struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepo {
	return &FoodRepo{db}
}

//FoodRepo implements the repository.FoodRepository interface
var _ repository.FoodRepository = &FoodRepo{}

func (r *FoodRepo) SaveFood(food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	food.FoodImage = os.Getenv("DO_SPACES_URL") + food.FoodImage

	err := r.db.Debug().Create(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "food title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return food, nil
}

func (r *FoodRepo) GetFood(id uint64) (*entity.Food, error) {
	var food entity.Food
	err := r.db.Debug().Where("id = ?", id).Take(&food).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &food, nil
}
