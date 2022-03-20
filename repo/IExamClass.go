package repo

import (
	"context"
	"student_score_sword/models"
)

type IExamRepo interface {
	SimpleStructQuery(ctx context.Context, class models.Exam) (res []models.Exam, err error)
	SingleInsert(ctx context.Context, teacher models.Exam) (res models.Exam, err error)
}
