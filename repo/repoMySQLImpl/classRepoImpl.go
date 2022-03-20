package repoMySqlImpl

import (
	"context"
	"gorm.io/gorm"
	"student_score_sword/models"
	"student_score_sword/repo"
)

type classRepoMysql struct {
	DB *gorm.DB
}

func (c classRepoMysql) SingleInsert(ctx context.Context, class models.Class) (res models.Class, err error) {
	result := c.DB.WithContext(ctx).Create(class)
	if result.Error == nil {
		return class, nil
	}
	return models.Class{}, result.Error
}

func (c classRepoMysql) SimpleStructQuery(ctx context.Context, class models.Class) (res []models.Class, err error) {
	var classList []models.Class
	err = c.DB.WithContext(ctx).Debug().Where(&class).Find(&classList).Error
	return classList, err

}

func NewClassRepoMysql() repo.IClassRepo {
	conn, _ := GetConn()
	return &classRepoMysql{conn}
}
