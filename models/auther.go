package models

type Auther struct {
	Fname        string   `json:"fname" bson:"first_name"`
	Lname        string   `json:"lname" bson:"last_name"`
	Age          int      `json:"age"   bson:"age"`
}
type Book struct {
	Bname        string        `json:"Bname"  bson:"book_name"`
	Price        int           `json:"price"  bson:"book_price"`
	Auther       Auther        `json:"auther" bson:"book_auther"`
	//Count        int           `json:"count"  bson:"Book_count"`
}