package book

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "books"

	title   = "title"
	authors = "authors"

	id = "_id"
)

type BookDal interface {
	Add(book Book) (Book, error)
	GetBooks() ([]Book, error)
	Update(book BookUpdate) error
	Delete(bookID string) error
	GetBook(bookID string) (Book, bool, error)
}

func NewBookMongoDal(database *mgo.Database) *BookMongoDal {
	return &BookMongoDal{
		database:   database,
		collection: database.C(collectionName),
	}
}

type BookMongoDal struct {
	database   *mgo.Database
	collection *mgo.Collection
}

func (d BookMongoDal) Add(book Book) (Book, error) {
	book.ID = bson.NewObjectId()
	err := d.collection.Insert(book)
	return book, err
}

func (d BookMongoDal) GetBooks() ([]Book, error) {
	books := make([]Book, 0)
	err := d.collection.Find(nil).All(&books)
	return books, err
}

func (d BookMongoDal) Update(book BookUpdate) error {
	dict := make(map[string]interface{})
	if book.Title != nil {
		dict[title] = book.Title
	}
	if len(book.Authors) > 0 {
		dict[authors] = book.Authors
	}
	//dict[authors] = *book.LastName
	return d.collection.Update(bson.M{id: bson.ObjectIdHex(book.ID)}, dict)
}

func (d BookMongoDal) Delete(bookID string) error {
	return d.collection.RemoveId(bson.ObjectIdHex(bookID))
}

func (d BookMongoDal) GetBook(bookID string) (Book, bool, error) {
	book := new(Book)
	err := d.collection.FindId(bson.ObjectIdHex(bookID)).One(book)
	if err != nil && err == mgo.ErrNotFound {
		return *book, false, nil
	}
	return *book, true, err
}