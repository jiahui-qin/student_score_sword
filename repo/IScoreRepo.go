package repo

import (
	"context"
	"student_score_sword/models"
)

type IScoreRepo interface {
	SimpleStructQuery(ctx context.Context, score models.Score) (res []models.Score, err error)
	SingleInsert(ctx context.Context, score models.Score) (res models.Score, err error)
}
