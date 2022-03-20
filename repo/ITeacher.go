package repo

import (
	"context"
	"student_score_sword/models"
)

type ITeacherRepo interface {
	SimpleStructQuery(ctx context.Context, class models.Teacher) (res []models.Teacher, err error)
	SingleInsert(ctx context.Context, teacher models.Teacher) (res models.Teacher, err error)
}
