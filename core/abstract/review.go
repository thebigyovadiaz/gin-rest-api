package abstract

import (
	"context"

	"github.com/thebigyovadiaz/gin-rest-api/models"
)

type Review interface {
	AddReview(ctx context.Context, revParams models.ReviewParams) (bool, error)
	ListReview(ctx context.Context, bookId int64) ([]models.ReviewList, error)
}
