package services

import "book.com/sina-apis/models"

type UserService interface {
	BuyBook(*string, *string) (*models.Book,error)
	

}