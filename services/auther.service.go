package services

import "book.com/sina-apis/models"

type AutherService interface {
	AddBook(*models.Book) error
	RemoveBook(*string,*string)error
	UpdateBook(*models.Book)error
	GetAllAutherBook(*string)([]*models.Book,error)
}