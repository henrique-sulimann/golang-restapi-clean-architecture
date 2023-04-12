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

type GameRepositoryDB struct {
	collection *mongo.Collection
}

func NewGameRepositoryDB(mongoClient *mongo.Client, dbName string, collectionName string) repository.GameRepository {
	collection := mongoClient.Database(dbName).Collection(collectionName)
	return &GameRepositoryDB{collection: collection}
}

func (r *GameRepositoryDB) Create(ctx context.Context, game *entity.Game) (*entity.Game, error) {
	game.ID = primitive.NewObjectID().Hex()
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, game)
	if err != nil {
		return nil, err
	}
	return game, nil
}
func (r *GameRepositoryDB) UpdateGameByID(ctx context.Context, id string, game *entity.Game) (*entity.Game, error) {
	game.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"name":        game.Name,
		"description": game.Description,
		"updatedat":   game.UpdatedAt,
	}})
	if err != nil {
		return nil, err
	}
	return game, nil
}
func (r *GameRepositoryDB) FindByID(ctx context.Context, id string) (*entity.Game, error) {
	var game *entity.Game
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return game, nil
}
func (r *GameRepositoryDB) DeleteGameByID(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
func (r *GameRepositoryDB) FindAll(ctx context.Context) ([]*entity.Game, error) {
	var games []*entity.Game
	res, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var game entity.Game
		err := res.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return games, nil
}
