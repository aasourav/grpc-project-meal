package cronjobs

import (
	"fmt"
	"log"
	"time"

	"aas.dev/pkg/models/types"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"
	"github.com/robfig/cron/v3"
)

func MealAssignCj() {
	db := utils.MongoDatabase
	mealRepo := repository.NewMealRepo(db)
	userRepo := repository.NewUserRepo(db)
	userSvc := services.NewUserService(userRepo, nil)
	mealSvc := services.NewMealService(mealRepo)

	loc, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		log.Println("Error loading location: ", err.Error())
	}
	c := cron.New(cron.WithLocation(loc))

	_, err = c.AddFunc("* * * * *", func() {
		users, err := userSvc.GetAllUsers()
		if err != nil {
			log.Println("cronjob failed due to error getting users: ", err.Error())
			return
		}

		for _, user := range *users {
			now := time.Now().In(loc)
			// weekday start from 0 as Sunday
			mealConsumeFilter := types.MealConsume{
				UserId:      user.ID,
				ConsumeDate: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc),
			}
			mealConsumeByUser, _ := mealSvc.GetMealByUser(mealConsumeFilter)

			if mealConsumeByUser != nil {
				fmt.Println("cannot update meal by the cronjob. already assigned today")
				continue
			}

			weekDay := int(now.Weekday())
			mealConsume := types.MealConsume{
				UserId:      user.ID,
				MealCount:   utils.BoolToInt(user.WeeklyPlan[weekDay]),
				ConsumeDate: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc),
				CreatedAt:   now,
				UpdatedAt:   now,
				ModifiedBy:  types.SYSTEM_CRONJOB,
			}

			err = mealSvc.AddMeal(mealConsume)
			if err != nil {
				log.Println("Error adding meal consume cronjob: ", err.Error())
				continue
			}

			log.Printf("Cron job executed at: %s", time.Now().Format(time.RFC3339))
		}
	})

	if err != nil {
		log.Println("Error creating cron job: ", err.Error())
	}

	log.Println("Cron job started.")
	go c.Start()
}
