package atlas

import (
	"context"
	"fmt"

	"github.com/javiertlopez/numeral"
)

// Collection keeps the collection name
const Collection = "logs"

func (s *repository) CreateLog(ctx context.Context, countLog numeral.Log) (numeral.Log, error) {
	return numeral.Log{}, fmt.Errorf("not implemented")
}
