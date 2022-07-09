package storage

import (
	"context"
	"log"
	"time"

	"github.com/Medveddo/team-booster/backend/models"
	"github.com/rs/zerolog"
	zero "github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ISkillStorage interface {
	Close() error
	GetSkills() ([]models.Skill, error)
	GetSkill(skillID primitive.ObjectID) (*models.Skill, error)
	InsertSkill(skill *models.Skill) error
	UpdateSkill(id primitive.ObjectID, newSkill *models.Skill) error
}

type skillStorage struct {
	client *mongo.Client
	logger zerolog.Logger
}

var _ ISkillStorage = &skillStorage{}

func NewSkillStorage() *skillStorage {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:mongo@localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	logger := zero.With().Str("module", "storage").Logger()
	logger.Info().
		Msg("successfully connected to MongoDB at localhost:27017")
	return &skillStorage{
		client: client,
		logger: logger,
	}
}

func (ss *skillStorage) InsertSkill(skill *models.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := ss.client.Database("tmbster").Collection("skills").InsertOne(ctx, *skill)
	if err != nil {
		panic(err)
	}
	return nil
}

func (ss *skillStorage) GetSkills() ([]models.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cur, err := ss.client.Database("tmbster").Collection("skills").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(ctx)
	var skills []models.Skill = make([]models.Skill, 0)
	for cur.Next(ctx) {
		var result models.Skill
		err := cur.Decode(&result)
		// fmt.Println(result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		skills = append(skills, result)
	}
	ss.logger.Debug().
		Str("action", "get_skills").
		Int("fetched_count", len(skills)).
		Send()
	return skills, nil
}

func (ss *skillStorage) GetSkill(id primitive.ObjectID) (*models.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cur, err := ss.client.Database("tmbster").Collection("skills").Find(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var skill models.Skill
	for cur.Next(ctx) {
		err := cur.Decode(&skill)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	ss.logger.Debug().
		Str("action", "get_skill").
		Str("id", id.Hex()).
		Send()
	return &skill, nil
}

func (ss *skillStorage) UpdateSkill(id primitive.ObjectID, newSkill *models.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	result, err := ss.client.
		Database("tmbster").
		Collection("skills").
		ReplaceOne(ctx, bson.M{"_id": id}, *newSkill)

	if err != nil {
		zero.Err(err).Msg("Skill was not updated")
		return err
	}

	ss.logger.Debug().
		Str("action", "update_skill").
		Str("id", id.Hex()).
		Int64("matched_count", result.MatchedCount).
		Int64("modified_count", result.ModifiedCount).
		Send()

	return nil
}

func (ss *skillStorage) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := ss.client.Disconnect(ctx); err != nil {
		panic(err)
	}
	return nil
}
