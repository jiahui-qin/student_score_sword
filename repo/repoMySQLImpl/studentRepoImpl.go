package repoMySqlImpl

import (
	"context"
	"gorm.io/gorm"
	"student_score_sword/models"
	"student_score_sword/repo"
)

type studentRepoImpl struct {
	DB *gorm.DB
}

func (c studentRepoImpl) SingleInsert(ctx context.Context, student models.Student) (res models.Student, err error) {
	result := c.DB.WithContext(ctx).Create(student)
	if result.Error == nil {
		return student, nil
	}
	return models.Student{}, result.Error
}

func (c studentRepoImpl) SimpleStructQuery(ctx context.Context, student models.Student) (res []models.Student, err error) {
	var stuList []models.Student
	err = c.DB.WithContext(ctx).Debug().Where(&student).Find(&student).Error
	return stuList, err

}

func NewStudentRepoMysql() repo.IStudentRepo {
	conn, _ := GetConn()
	return &studentRepoImpl{conn}
}
