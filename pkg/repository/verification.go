package repository

import (
	"context"
	"errors"
	"log"

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
	// expires := time.Now().Add(time.Second * 10).Unix()
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"createdAt": 1},
		Options: options.Index().SetExpireAfterSeconds(types.VERIFICATION_EXPIRY_SECONDS),
	}

	if isIndexed {
		db.Collection(types.VERIFICATION).Indexes().CreateOne(context.TODO(), indexModel)
	}

	return &VerificationRepo{collection: db.Collection(types.VERIFICATION)}
}

func (repo *VerificationRepo) DeleteVeruficationByUserId(userId string) error {
	_, err := repo.collection.DeleteOne(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return err
	}
	return nil
}

func (repo *VerificationRepo) GetVerificationDocByUserId(userId string) (*models.Verification, error) {
	var verificationData *models.Verification
	err := repo.collection.FindOne(context.Background(), bson.M{"userId": userId}).Decode(&verificationData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("admin user not found")
		}
		log.Println("Error fetching verification data by userId:", userId, "Error:", err)
		return nil, err
	}
	return verificationData, nil
}

func (repo *VerificationRepo) CreateVerificationRepo(verificationData *models.Verification) error {
	data := bson.M{
		"userId":    verificationData.UserId,
		"email":     verificationData.Email,
		"createdAt": verificationData.CreatedAt,
	}
	_, err := repo.collection.InsertOne(context.TODO(), data)
	return err
}
