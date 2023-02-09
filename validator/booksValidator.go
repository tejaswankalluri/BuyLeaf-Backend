package validator

import (
	"fiber-api/models"
)

func ValidateBook(book models.Book) []*ErrorResponse {
	err := GenValidate(book)
	return err
}
