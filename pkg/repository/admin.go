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

func (repo *AdminRepo) UpdateAdminById(admin *models.Admin) error {
	// adminBson, _ := bson.Marshal(admin)
	updateFields := bson.M{
		"$set": bson.M{
			"isEmailApproved": admin.IsEmailApproved,
			// Add more fields as needed
		},
	}
	filterData := bson.M{
		"email": admin.Email,
	}
	_, err := repo.collection.UpdateOne(context.TODO(), filterData, updateFields)
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

func (repo *AdminRepo) GetAdminById(id primitive.ObjectID) (*models.Admin, error) {
	var admin *models.Admin
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&admin)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("admin user not found")
		}
		log.Println("Error fetching admin by id:", id, "Error:", err)
		return nil, err
	}
	return admin, nil
}

func (repo *AdminRepo) GetAdmins() (*[]models.Admin, error) {
	var adminDocs []models.Admin // Use a non-pointer slice here
	// cursor, err := repo.collection.Find(context.Background(), bson.M{"email": email})
	cursor, err := repo.collection.Find(context.Background(), bson.M{})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("admin user not found")
		}
		// log.Println("Error fetching admin by email:", email, "Error:", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	// Pass a pointer to the slice (not a pointer to a pointer)
	if err := cursor.All(context.Background(), &adminDocs); err != nil {
		log.Println("Error decoding admins:", err)
		return nil, err
	}

	return &adminDocs, nil // Return a pointer to the slice
}
