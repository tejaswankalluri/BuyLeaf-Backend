package validator

import "buyleaf/models"

func ValidateBook(book models.Book) []*ErrorResponse {
	err := GenValidate(book)
	return err
}
