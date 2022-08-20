package services

import (
	"context"
	"errors"
	"book.com/sina-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutherServiceimpl struct {
	authercollection *mongo.Collection
	ctx              context.Context
}
func NewAutherservice(authercollection *mongo.Collection, ctx context.Context) AutherService {
	return &AutherServiceimpl{
		authercollection: authercollection,
		ctx:              ctx,
	}
}

// AddBook implements AutherService
func (a *AutherServiceimpl) AddBook(book *models.Book) error {
  _,err:=a.authercollection.InsertOne(a.ctx,book)
  return err
}

// GetAllAutherBook implements AutherService
func (a *AutherServiceimpl) GetAllAutherBook(auther *string) ([]*models.Book, error) {
   var books []*models.Book
   cursor,err:=a.authercollection.Find(a.ctx,bson.D{bson.E{Key:"book_auther.first_name",Value:auther}})	
   if err != nil {
	return nil, err
   }
   for cursor.Next(a.ctx){
	var book models.Book
	err:=cursor.Decode(&book)
	if err != nil {
		return nil, err
	}
	books = append(books, &book)
   }
   if err:=cursor.Err();err!=nil{
	return nil,err
   }
   cursor.Close(a.ctx)
   if len(books)==0{
	return nil,errors.New("documents not found")
   }
    return books,nil
}

// RemoveBook implements AutherService
func (a *AutherServiceimpl) RemoveBook(bn *string,an *string) error {
	filter:=bson.D{
		{Key: "$and",
		Value: bson.A{
			bson.D{{Key: "book_name",Value: bn}},
			bson.D{{Key: "book_auther.first_name",Value: an}},
			},
		},
	}
	result,_:=a.authercollection.DeleteOne(a.ctx,filter)
	if result.DeletedCount!=1{
		return errors.New("documents not found")
	}
	

	return nil
}

// UpdateBook implements AutherService
func (a *AutherServiceimpl) UpdateBook(book *models.Book) error {
	filter:=bson.D{
		{Key: "$and",
		Value: bson.A{
			bson.D{{Key: "book_name",Value: book.Bname}},
			bson.D{{Key: "book_auther.first_name",Value:book.Auther.Fname}},

		    },

		},
	}
    update:=bson.D{
		{Key: "$set",
		Value: bson.D{
			bson.E{Key: "book_name",Value: book.Bname},
			bson.E{Key: "book_price",Value: book.Price},
			bson.E{Key: "book_auther",Value: book.Auther},
	        },
        },
    }
	_,err:=a.authercollection.UpdateOne(a.ctx,filter,update)
	if err!=nil{
		return err
	}
	return nil
}

