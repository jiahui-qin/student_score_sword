package repo

import (
	"context"
	"student_score_sword/models"
)

type IStudentRepo interface {
	SimpleStructQuery(ctx context.Context, student models.Student) (res []models.Student, err error)
	SingleInsert(ctx context.Context, student models.Student) (res models.Student, err error)
}
