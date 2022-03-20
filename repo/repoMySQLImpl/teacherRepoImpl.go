package repoMySqlImpl

import (
	"context"
	"gorm.io/gorm"
	"student_score_sword/models"
	"student_score_sword/repo"
)

type teacherRepoImpl struct {
	DB *gorm.DB
}

func (c teacherRepoImpl) SingleInsert(ctx context.Context, teacher models.Teacher) (res models.Teacher, err error) {
	result := c.DB.WithContext(ctx).Create(teacher)
	if result.Error != nil {
		return teacher, nil
	}
	return models.Teacher{}, result.Error
}

func (c teacherRepoImpl) SimpleStructQuery(ctx context.Context, teacher models.Teacher) (res []models.Teacher, err error) {
	var teacherList []models.Teacher
	err = c.DB.WithContext(ctx).Debug().Where(&teacher).Find(&teacherList).Error
	return teacherList, err

}

func NewTeacherRepoMysql() repo.ITeacherRepo {
	conn, _ := GetConn()
	return &teacherRepoImpl{conn}
}
