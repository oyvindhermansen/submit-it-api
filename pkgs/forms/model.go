package forms

import "go.mongodb.org/mongo-driver/bson/primitive"

type Form struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name                string             `bson:"name" json:"name" validate:"required"`
	InternalDescription string             `bson:"internalDescription" json:"internalDescription"`
	PublicDescription   string             `bson:"publicDescription" json:"publicDescription"`
	SuccessText         string             `bson:"successText" json:"successText"`
	ErrorText           string             `bson:"errorText" json:"errorText"`
	SubmitButtonText    string             `bson:"submitButtonText" json:"submitButtonText"`
}
