package domain

import (
	"favorit-book-tracker/domain/repository"
	"favorit-book-tracker/domain/usecase"
	"favorit-book-tracker/infrastructure"
)

type Domain struct {
	BookUsecase usecase.BookUsecase
}

func ConstructionDomain() Domain {
	databaseConn := infrastructure.NewDatabaseConnection()
	databaseRepository := repository.NewDatabaseRepository(databaseConn)
	bookUseCase := usecase.NewBookUsecase(databaseRepository)

	return Domain{
		BookUsecase: &bookUseCase,
	}
}
