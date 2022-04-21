package repository

import (
	"golangEx/connections"
	"golangEx/models"
)

type UserRepository interface {
	Save(models.User) (models.User, error)   // tao moi user
	FindAll() ([]models.User, error)         // tra ve tat ca user
	FindById(int) (models.User, error)       // tim bang id
	FindByEmail(string) (models.User, error) // tim bang email
	Update(int, models.User) (int, error)    // thay doi user
	Delete(int) (bool, error)                // delete
}

type UsersCRUD struct{}

func (UsersCRUD) Save(u models.User) (models.User, error) {
	var err error
	// query
	err = connections.DB.Debug().Model(&models.User{}).Create(&u).Error
	return u, err
}

func (UsersCRUD) FindAll() ([]models.User, error) {
	var users = []models.User{}
	var err error
	err = connections.DB.Debug().Limit(100).Find(&users).Error
	return users, err
}

func (UsersCRUD) FindById(i int) (models.User, error) {
	var u = models.User{}
	var err error
	err = connections.DB.Debug().Where("id=?", i).First(&u).Error
	return u, err
}
func (UsersCRUD) FindByEmail(s string) (models.User, error) {
	var u = models.User{}
	var err error
	err = connections.DB.Debug().Where("email=?", s).First(&u).Error
	return u, err
}

func (UsersCRUD) Update(i int, u models.User) (int, error) {
	var err error
	err = connections.DB.Debug().Model(&models.User{}).Where("id=?", i).Take(&models.User{}).UpdateColumns(map[string]interface{}{
		"Name":     u.Name,
		"Password": u.Password,
	}).Error
	return 1, err
}

func (UsersCRUD) Delete(i int) (bool, error) {
	var err error
	err = connections.DB.Debug().Where("id=?", i).Delete(&models.User{}).Error
	return true, err
}
