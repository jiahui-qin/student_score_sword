package repoMySqlImpl

import (
	"context"
	"gorm.io/gorm"
	"student_score_sword/models"
	"student_score_sword/repo"
)

type examRepoMysql struct {
	DB *gorm.DB
}

func (c examRepoMysql) SingleInsert(ctx context.Context, exam models.Exam) (res models.Exam, err error) {
	result := c.DB.WithContext(ctx).Create(exam)
	if result.Error == nil {
		return exam, nil
	}
	return models.Exam{}, result.Error
}

func NewExamRepoMysql() repo.IExamRepo {
	conn, _ := GetConn()
	return &examRepoMysql{conn}
}

func (c examRepoMysql) SimpleStructQuery(ctx context.Context, exam models.Exam) (res []models.Exam, err error) {
	var examList []models.Exam
	err = c.DB.WithContext(ctx).Debug().Where(&exam).Find(&examList).Error
	return examList, err

}
