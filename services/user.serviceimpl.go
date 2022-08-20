package services

import (
	"context"

	"book.com/sina-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceimpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}
func Newuserservice(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceimpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
// BuyBook implements UserService
func (us *UserServiceimpl) BuyBook(bn *string, an *string) (*models.Book, error) {
    var book *models.Book
	filter:=bson.D{
		{Key: "$and",
		Value: bson.A{
			bson.D{{Key: "book_name",Value: bn}},
			bson.D{{Key: "book_auther.first_name",Value: an}},
			},
		},
	}
	err:=us.usercollection.FindOne(us.ctx,filter).Decode(&book)
	return book,err
	
}

