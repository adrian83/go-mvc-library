package dal

import (
	"domain/author/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "authors"

	firstName = "firstName"
	lastName  = "lastName"

	id = "_id"
)

type AuthorDal interface {
	Add(author model.Author) (model.Author, error)
	GetAuthors() ([]model.Author, error)
	Update(author model.AuthorUpdate) error
	Delete(authorID string) error
	GetAuthor(authorID string) (model.Author, error)
	FindAuthorsByIDs(authorIDs []string) ([]model.Author, error)
}

func NewAuthorMongoDal(database *mgo.Database) *AuthorMongoDal {
	return &AuthorMongoDal{
		database:   database,
		collection: database.C(collectionName),
	}
}

type AuthorMongoDal struct {
	database   *mgo.Database
	collection *mgo.Collection
}

func (d AuthorMongoDal) Add(author model.Author) (model.Author, error) {
	entity := dtoToEntity(author)
	return entityToDto(entity), d.collection.Insert(entity)
}

func (d AuthorMongoDal) GetAuthors() ([]model.Author, error) {
	authorsEntities := make([]model.AuthorEntity, 0)
	err := d.collection.Find(nil).All(&authorsEntities)
	return entitiesToDtos(authorsEntities), err
}

func (d AuthorMongoDal) Update(author model.AuthorUpdate) error {
	dict := make(map[string]interface{})
	dict[firstName] = author.FirstName
	dict[lastName] = *author.LastName
	return d.collection.Update(bson.M{id: bson.ObjectIdHex(author.ID)}, dict)
}

func (d AuthorMongoDal) Delete(authorID string) error {
	return d.collection.RemoveId(bson.ObjectIdHex(authorID))
}

func (d AuthorMongoDal) GetAuthor(authorID string) (model.Author, error) {
	authorEntity := new(model.AuthorEntity)
	err := d.collection.FindId(bson.ObjectIdHex(authorID)).One(authorEntity)
	return entityToDto(*authorEntity), err
}

func (d AuthorMongoDal) FindAuthorsByIDs(authorIDs []string) ([]model.Author, error) {
	ids := make([]bson.ObjectId, 0)
	for _, authorID := range authorIDs {
		ids = append(ids, bson.ObjectIdHex(authorID))
	}
	authorsEntities := make([]model.AuthorEntity, 0)
	err := d.collection.Find(bson.M{id: map[string]interface{}{"$in": ids}}).All(&authorsEntities)
	return entitiesToDtos(authorsEntities), err
}

func dtoToEntity(dto model.Author) model.AuthorEntity {
	ID := bson.NewObjectId()
	return model.AuthorEntity{ID: ID, FirstName: dto.FirstName, LastName: dto.LastName}
}

func entityToDto(entity model.AuthorEntity) model.Author {
	return model.Author{ID: entity.ID.Hex(), FirstName: entity.FirstName, LastName: entity.LastName}
}

func entitiesToDtos(entities []model.AuthorEntity) []model.Author {
	authors := make([]model.Author, len(entities))
	for i, entity := range entities {
		authors[i] = entityToDto(entity)
	}
	return authors
}
