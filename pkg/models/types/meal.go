package types

import "time"

type MealConsume struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserId      string    `json:"userId" bson:"userId" validate:"required"`
	ConsumeDate time.Time `json:"consumeDate" bson:"consumeDate" validate:"required"`
	MealCount   int32     `json:"mealCount" bson:"mealCount" validate:"required"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}
