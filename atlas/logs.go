package atlas

import (
	"context"
	"fmt"
	"time"

	"github.com/javiertlopez/numeral"

	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection keeps the collection name
const Collection = "logs"

// log model for mongodb
type log struct {
	ID        string    `bson:"_id"`
	Count     int       `bson:"count,omitempty"`
	ImageKey  string    `bson:"image_key,omitempty"`
	DeviceID  string    `bson:"device_id,omitempty"`
	ShapeID   string    `bson:"shape_id,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (r *repository) CreateLog(ctx context.Context, countLog numeral.Log) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	time := time.Now()

	uuid := guuid.New().String()

	insert := &log{
		ID:        uuid,
		ImageKey:  countLog.ImageKey,
		DeviceID:  countLog.DeviceID,
		ShapeID:   countLog.ShapeID,
		CreatedAt: time,
		UpdatedAt: time,
	}

	_, err := collection.InsertOne(ctx, insert)
	if err != nil {
		return numeral.Log{}, err
	}

	return insert.toModel(), nil
}

// GetByID retrieves a log with the ID
func (r *repository) GetByID(ctx context.Context, id string) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	var response log

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&response)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return numeral.Log{}, fmt.Errorf("not found")
		}

		return numeral.Log{}, err
	}

	return response.toModel(), nil
}

func (r *repository) UpdateLog(ctx context.Context, countLog numeral.Log) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	time := time.Now()

	update := &log{
		Count:     countLog.Count,
		UpdatedAt: time,
	}

	_, err := collection.UpdateByID(ctx, countLog.ID, update)
	if err != nil {
		return numeral.Log{}, err
	}

	response, err := r.GetByID(ctx, countLog.ID)
	if err != nil {
		return numeral.Log{}, err
	}

	return response, nil
}

func (l log) toModel() numeral.Log {
	return numeral.Log{
		ID:        l.ID,
		Count:     l.Count,
		ImageKey:  l.ImageKey,
		DeviceID:  l.DeviceID,
		ShapeID:   l.ShapeID,
		Timestamp: l.CreatedAt.UTC().Unix(),
	}
}
