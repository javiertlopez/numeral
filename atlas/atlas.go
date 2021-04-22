package atlas

import (
	"github.com/javiertlopez/numeral"

	"go.mongodb.org/mongo-driver/mongo"
)

// repository struct holds the MongoDB client
type repository struct {
	mongo *mongo.Database
}

// New returns a repository instance
func New(
	m *mongo.Database,
) numeral.Repository {
	return &repository{
		mongo: m,
	}
}
