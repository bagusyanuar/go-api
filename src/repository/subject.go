package repository

import (
	"go-api/database"
	"go-api/src/model"
)

func CreateSubject(subject *model.Subject) (err error) {
	return database.CONFIG.Debug().Create(&subject).Error
}

func FindSubject(subject *[]model.Subject, param string) (s *[]model.Subject, err error) {
	if err = database.CONFIG.Debug().Where("name LIKE ?", "%"+param+"%").Find(&subject).Error; err != nil {
		return subject, err
	}
	return subject, nil
}

func FindSubjectByID(subject *model.Subject, id string) (s *model.Subject, err error) {
	if err = database.CONFIG.Debug().Where("id = ?", id).First(&subject).Error; err != nil {
		return subject, err
	}
	return subject, nil
}

func PatchSubjectByID(id uint, data map[string]interface{}) (err error) {
	if err = database.CONFIG.Debug().Model(model.Subject{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
