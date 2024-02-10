package usecase

import (
	"favorit-book-tracker/domain/entity"
	"favorit-book-tracker/domain/repository"
)

type BookUsecase interface {
	GetFavoritBook() (*entity.Book, error)
	SaveFavoritBook(action entity.Book) error
}

type bookUsecaseImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewBookUsecase(databaseRepository repository.DatabaseRepository) bookUsecaseImpl {
	return bookUsecaseImpl{
		databaseRepository: databaseRepository,
	}
}

func (this *bookUsecaseImpl) GetFavoritBook() (*entity.Book, error) {
	var book entity.Book
	if err := this.databaseRepository.GetAllFavoritBook(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (this *bookUsecaseImpl) SaveFavoritBook(action entity.Book) error {
	if err := this.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}
