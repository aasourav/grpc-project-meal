package repository

import (
	"context"

	"aas.dev/pkg/models/types"

	"aas.dev/pkg/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealRepo struct {
	collection *mongo.Collection
}

func NewMealRepo(db *mongo.Database) interfaces.MealRepository {
	return &MealRepo{collection: db.Collection(types.MEAL_CONSUME)}
}

func (repo *MealRepo) AddMeal(meal types.MealConsume) error {
	mealConsumeDoc := bson.M{
		"userId":      meal.UserId,
		"consumeDate": meal.ConsumeDate,
		"mealCount":   meal.MealCount,
		"createdAt":   meal.CreatedAt,
		"updatedAt":   meal.UpdatedAt,
		"modifiedBy":  meal.ModifiedBy,
	}
	_, err := repo.collection.InsertOne(context.TODO(), mealConsumeDoc)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MealRepo) GetMealByUser(meal types.MealConsume) (*types.MealConsume, error) {
	mealConsumeDoc := bson.M{
		"userId":      meal.UserId,
		"consumeDate": meal.ConsumeDate,
	}
	mealDoc := &types.MealConsume{}
	err := repo.collection.FindOne(context.TODO(), mealConsumeDoc).Decode(&mealDoc)
	if err != nil {
		return nil, err
	}
	return mealDoc, nil
}
