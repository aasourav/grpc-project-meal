package repository

import (
	"context"
	"time"

	"aas.dev/pkg/models/types"
	models "aas.dev/pkg/models/verification"

	"aas.dev/pkg/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VerificationRepo struct {
	collection *mongo.Collection
}

func NewVerificationRepo(db *mongo.Database, isIndexed bool) interfaces.VerifiactionRepository {
	expires := time.Now().Add(time.Second * 120).Unix()
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"createdAt": 1},
		Options: options.Index().SetExpireAfterSeconds(int32(expires)),
	}

	if isIndexed {
		db.Collection(types.VERIFICATION).Indexes().CreateOne(context.TODO(), indexModel)
	}

	return &VerificationRepo{collection: db.Collection(types.VERIFICATION)}
}

func (repo *VerificationRepo) CreateVerificationRepo(verificationData *models.Verification) error {
	data := bson.M{
		"email":     verificationData.Email,
		"createdAt": verificationData.CreatedAt,
	}
	_, err := repo.collection.InsertOne(context.TODO(), data)
	return err
}
