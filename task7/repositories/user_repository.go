package repositories

import (
	"context"
	"task_manager/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &UserRepositoryImpl{collection: db.Collection("users")}
}

func (r *UserRepositoryImpl) Create(c context.Context, user *domain.User) (*domain.User, error) {
	user.ID = uuid.New().String()
	_, err := r.collection.InsertOne(c, bson.M{
		"_id":           user.ID,
		"username":      user.Username,
		"password_hash": user.PasswordHash,
		"role":          user.Role,
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, domain.ErrUsernameExists
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByID(c context.Context, id string) (*domain.User, error) {
	var result struct {
		ID           string        `bson:"_id" json:"id"`
		Username     string        `bson:"username" json:"username"`
		PasswordHash string        `bson:"password_hash" json:"-"`
		Role         domain.Role   `bson:"role" json:"role"`
	}
	err := r.collection.FindOne(c, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &domain.User{
		ID:           result.ID,
		Username:     result.Username,
		PasswordHash: result.PasswordHash,
		Role:         result.Role,
	}, nil
}

func (r *UserRepositoryImpl) GetByUsername(c context.Context, username string) (*domain.User, error) {
	var result struct {
		ID           string        `bson:"_id" json:"id"`
		Username     string        `bson:"username" json:"username"`
		PasswordHash string        `bson:"password_hash" json:"-"`
		Role         domain.Role   `bson:"role" json:"role"`
	}
	err := r.collection.FindOne(c, bson.M{"username": username}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &domain.User{
		ID:           result.ID,
		Username:     result.Username,
		PasswordHash: result.PasswordHash,
		Role:         result.Role,
	}, nil
}

func (r *UserRepositoryImpl) Update(c context.Context, user *domain.User) error {
	result, err := r.collection.UpdateOne(c, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"username":      user.Username,
			"password_hash": user.PasswordHash,
			"role":          user.Role,
		},
	})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *UserRepositoryImpl) CountUsers(c context.Context) (int64, error) {
	count, err := r.collection.CountDocuments(c, bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
