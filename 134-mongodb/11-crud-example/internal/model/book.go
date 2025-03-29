package model

import (
	"crud-example/internal/db"
	apperr "crud-example/pkg/util/app_err"
	ctxhelper "crud-example/pkg/util/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn" bson:"isbn"`
	Title  string             `json:"title" bson:"title"`
	Author string             `json:"author" bson:"author"`
	Price  float64            `json:"price" bson:"price"`
}

func FetchAllBooks() ([]Book, error) {
	bks := []Book{}

	cl := getBooksCollection()

	ctx, cancel := ctxhelper.CtxWithNormalTimeout()
	defer cancel()

	cur, err := cl.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	// Manual
	// if cur.Next(ctx) {
	// 	bk := Book{}
	// 	err := cur.Decode(bk)
	// }
	// defer cur.Close(ctx)

	err = cur.All(ctx, &bks)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

func FetchBook(isbn string) (Book, error) {
	bk := Book{}

	cl := getBooksCollection()

	ctx, cancel := ctxhelper.CtxWithNormalTimeout()
	defer cancel()

	res := cl.FindOne(ctx, bson.M{"isbn": isbn})

	if err := res.Err(); err != nil {
		return bk, err
	}

	err := res.Decode(&bk)
	return bk, err
}

func InsertBook(book Book) (primitive.ObjectID, error) {
	cl := getBooksCollection()

	ctx, cancel := ctxhelper.CtxWithNormalTimeout()
	defer cancel()

	res, err := cl.InsertOne(ctx, book)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func UpdateBook(isbn string, book Book) (Book, error) {
	cl := getBooksCollection()

	ctx, cancel := ctxhelper.CtxWithNormalTimeout()
	defer cancel()

	bk := Book{}
	upOpts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := cl.FindOneAndUpdate(ctx, bson.M{"isbn": isbn}, bson.M{"$set": book}, upOpts)
	if err := res.Err(); err != nil {
		return bk, err
	}

	if err := res.Decode(&bk); err != nil {
		return bk, err
	}

	return bk, nil
}

func DeleteBook(isbn string) error {
	cl := getBooksCollection()

	ctx, cancel := ctxhelper.CtxWithNormalTimeout()
	defer cancel()

	res, err := cl.DeleteOne(ctx, bson.M{"isbn": isbn}, nil)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return apperr.ErrDeleteItemFailed
	}

	return nil
}

func getBooksCollection() *mongo.Collection {
	// Get collection
	return db.MongoClient.Database("bookstore").Collection("books")
}
