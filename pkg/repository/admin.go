package repository

import (
	"context"
	"errors"
	"log"

	"aas.dev/pkg/models/types"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/admin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepo struct {
	collection *mongo.Collection
}

//	func NewUserRepo(db *mongo.Database) interfaces.UserRepository {
//		return &UserRepo{collection: db.Collection(types.USERS)}
//	}
func NewAdminRepo(db *mongo.Database) interfaces.AdminRepository {
	return &AdminRepo{collection: db.Collection(types.ADMINS)}
}

func (repo *AdminRepo) DeleteAdminById(id string) error {
	hexId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.collection.DeleteOne(context.TODO(), bson.M{"_id": hexId})
	if err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepo) CreateAdmin(admin *models.Admin) error {
	userDoc := bson.M{
		"email":              admin.Email,
		"password":           admin.Password,
		"name":               admin.Name,
		"employeeId":         admin.EmployeeId,
		"createdAt":          admin.CreatedAt,
		"updatedAt":          admin.UpdatedAt,
		"department":         admin.Department,
		"departmentInCharge": admin.DepartmentInCharge,
	}
	_, err := repo.collection.InsertOne(context.Background(), userDoc)
	return err
}

func (repo *AdminRepo) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin *models.Admin
	err := repo.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("admin user not found")
		}
		log.Println("Error fetching admin by email:", email, "Error:", err)
		return nil, err
	}
	return admin, nil
}
