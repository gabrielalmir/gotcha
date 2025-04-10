package services

import (
	"context"
	"time"

	"gotcha/src/db"
	"gotcha/src/models"
	"gotcha/src/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type URLService struct {
	database  *db.Database
	snowflake *utils.Snowflake
}

func NewURLService(database *db.Database) (*URLService, error) {
	snowflake, err := utils.NewSnowflake(1)
	if err != nil {
		return nil, err
	}

	return &URLService{
		database:  database,
		snowflake: snowflake,
	}, nil
}

func (s *URLService) CreateShortURL(ctx context.Context, originalURL string) (*models.URL, error) {
	if !utils.IsValidURL(originalURL) {
		return nil, utils.ErrInvalidURL
	}

	id := s.snowflake.Generate()
	shortURL := utils.ToBase62(id)

	url := &models.URL{
		Original:  originalURL,
		Short:     shortURL,
		CreatedAt: time.Now(),
	}

	collection := s.database.GetCollection("urls")
	_, err := collection.InsertOne(ctx, url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetByShortURL(ctx context.Context, shortURL string) (*models.URL, error) {
	var url models.URL
	collection := s.database.GetCollection("urls")

	err := collection.FindOne(ctx, bson.M{"short": shortURL}).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
