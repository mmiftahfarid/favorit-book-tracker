package repository

import (
	"favorit-book-tracker/domain/entity"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	GetAllFavoritBook(data any, conditions ...any) error
	Create(value any) error
}

type databaseRepositoryImpl struct {
	database *gorm.DB
}

func NewDatabaseRepository(database *gorm.DB) databaseRepositoryImpl {
	return databaseRepositoryImpl{
		database: database,
	}
}

func (this databaseRepositoryImpl) GetAllFavoritBook(data any, conditions ...any) error {
	var favoritBook []entity.Book
	result := this.database.Find(&favoritBook)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error fatching book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (this databaseRepositoryImpl) Create(value any) error {
	result := this.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating book:: %v", result.Error))
		return result.Error
	}

	return nil
}
