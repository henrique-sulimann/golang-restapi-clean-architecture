package mongodb

import (
	"context"
	"time"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/data/repository"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDB struct {
	collection *mongo.Collection
}

func NewUserRepositoryDB(mongoClient *mongo.Client, dbName string, collectionName string) repository.UserRepository {
	collection := mongoClient.Database(dbName).Collection(collectionName)
	return &UserRepositoryDB{collection: collection}
}

func (r *UserRepositoryDB) Create(ctx context.Context, game *entity.User) error {
	game.ID = primitive.NewObjectID().Hex()
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, game)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositoryDB) UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	user.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"name":      user.Name,
		"password":  user.Password,
		"updatedat": user.UpdatedAt,
	}})
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepositoryDB) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user *entity.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepositoryDB) DeleteUserByID(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepositoryDB) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	res, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var user entity.User
		err := res.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
