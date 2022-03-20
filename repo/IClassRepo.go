package repo

import (
	"context"
	models "student_score_sword/models"
)

type IClassRepo interface {
	SimpleStructQuery(ctx context.Context, class models.Class) (res []models.Class, err error)
	SingleInsert(ctx context.Context, class models.Class) (res models.Class, err error)
}
