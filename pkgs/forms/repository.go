package forms

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(Form) (Form, error)
	List(ListOptions) ([]Form, error)
	FindById(primitive.ObjectID) (*Form, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repositoryImpl{
		collection: collection,
	}
}

func (r *repositoryImpl) Create(form Form) (Form, error) {
	result, err := r.collection.InsertOne(context.TODO(), form)

	if err != nil {
		return Form{}, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	form.Id = insertedID

	return form, nil
}

func (r *repositoryImpl) FindById(id primitive.ObjectID) (*Form, error) {
	var form Form
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(context.TODO(), filter).Decode(&form)

	if err != nil {
		return nil, err
	}

	return &form, nil
}

type ListOptions struct {
	Skip  int64
	Limit int64
}

func (r *repositoryImpl) List(opts ListOptions) ([]Form, error) {
	filter := bson.M{}
	findOptions := options.Find()

	if opts.Skip > 0 {
		findOptions.SetSkip(opts.Skip)
	}

	if opts.Limit > 0 {
		findOptions.SetLimit(opts.Limit)
	}

	cursor, err := r.collection.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	var forms []Form

	for cursor.Next(context.Background()) {
		var form Form

		err := cursor.Decode(&form)

		if err != nil {
			return nil, err
		}

		forms = append(forms, form)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return forms, nil

}
