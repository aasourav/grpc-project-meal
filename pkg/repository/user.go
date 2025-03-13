package repository

import (
	"context"
	"errors"
	"log"

	"aas.dev/pkg/models/types"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) interfaces.UserRepository {
	return &UserRepo{collection: db.Collection(types.USERS)}
}

func (repo *UserRepo) DeleteUserById(id string) error {
	hexId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.DeleteOne(context.TODO(), bson.M{"_id": hexId})
	if err != nil {
		return err
	}
	return nil
}

func NewPendingUserRepo(db *mongo.Database) interfaces.UserRepository {
	return &UserRepo{collection: db.Collection(types.PENDING_USERS)}
}

func (repo *UserRepo) CreateUser(user *models.User) error {

	userDoc := bson.M{
		"email":           user.Email,
		"password":        user.Password,
		"name":            user.Name,
		"employeeId":      user.EmployeeId,
		"isemailverified": user.IsEmailVerified,
		"isApproved":      user.IsApproved,
		"weeklyPlan":      user.WeeklyPlan,
		"createdAt":       user.CreatedAt,
		"updatedAt":       user.UpdatedAt,
		"department":      user.Department,
	}
	_, err := repo.collection.InsertOne(context.Background(), userDoc)
	return err
}

func (repo *UserRepo) UpdatePasswordById(user *models.User) error {
	updateFields := bson.M{
		"$set": bson.M{
			"password": user.Password,
			// Add more fields as needed
		},
	}
	hexId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}
	filterData := bson.M{
		"_id": hexId,
	}

	_, err = repo.collection.UpdateOne(context.TODO(), filterData, updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var userDoc models.User
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&userDoc)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &userDoc, err
}

func (repo *UserRepo) UpdateUserById(user *models.User) error {
	// adminBson, _ := bson.Marshal(admin)
	updateFields := bson.M{
		"$set": bson.M{
			"isEmailVerified": user.IsEmailVerified,
			// Add more fields as needed
		},
	}
	filterData := bson.M{
		"email": user.Email,
	}
	_, err := repo.collection.UpdateOne(context.TODO(), filterData, updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetUserById(id primitive.ObjectID) (*models.User, error) {
	var userDoc *models.User
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&userDoc)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		log.Println("Error fetching user by id:", id, "Error:", err)
		return nil, err
	}
	return userDoc, nil
}
