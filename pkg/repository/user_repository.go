package repository

import (
	"context"
	"errors"

	"aas.dev/pkg/models/types"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) interfaces.UserRepository {
	return &UserRepo{collection: db.Collection(types.USERS)}
}

func NewPendingUserRepo(db *mongo.Database) interfaces.UserRepository {
	return &UserRepo{collection: db.Collection(types.PENDING_USERS)}
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	userDoc := bson.M{
		"email":      user.UserLogin.Email,
		"password":   user.UserLogin.Password,
		"name":       user.Name,
		"employeeId": user.EmployeeId,
		"weeklyPlan": user.WeeklyPlan,
		"createdAt":  user.CreatedAt,
		"updatedAt":  user.UpdatedAt,
		"department": user.Department,
	}
	_, err := repo.collection.InsertOne(context.Background(), userDoc)
	return err
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}
