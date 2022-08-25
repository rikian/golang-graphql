package repository

import (
	"errors"
	"golang/graphql/app/config"
	"golang/graphql/app/entity"
	"golang/graphql/app/graph/model"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	GormUser             *entity.User
	GormUsers            []entity.User
	GrapResponseLogin    *model.ResponseLogin
	GrapResponseRegister *model.ResponseRegister
	Product              *model.Product
	Products             []model.Product
	User                 *model.User
	Users                []*model.User
}

func (u *UserRepo) SelectUser(input *model.Login) (*model.ResponseLogin, error) {
	var db gorm.DB = *config.DB
	u.GormUser = &entity.User{}

	err := db.Model(u.GormUser).
		Preload("UserStatus").
		Preload("Products").
		Where("user_email = ? AND user_password = ?", input.UserEmail, input.UserPassword).
		Find(u.GormUser).Error

	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if u.GormUser.UserId == "" {
		return nil, errors.New("data not found")
	}

	status := 200
	msg := "ok"

	u.GrapResponseLogin = &model.ResponseLogin{
		User: &model.User{
			UserID:      &u.GormUser.UserId,
			UserEmail:   &u.GormUser.UserEmail,
			UserName:    &u.GormUser.UserName,
			UserImage:   &u.GormUser.UserImage,
			UserStatus:  &u.GormUser.UserStatus.Status,
			CreatedDate: &u.GormUser.CreatedDate,
			LastUpdate:  &u.GormUser.LastUpdate,
		},
		Status:  &status,
		Message: &msg,
	}

	for i := 0; i < len(u.GormUser.Products); i++ {
		u.GrapResponseLogin.User.Products = append(u.GrapResponseLogin.User.Products, &model.Product{
			UserID:       &u.GormUser.Products[i].UserId,
			ProductID:    &u.GormUser.Products[i].ProductId,
			ProductName:  &u.GormUser.Products[i].ProductName,
			ProductPrice: &u.GormUser.Products[i].ProductPrice,
			ProductStock: &u.GormUser.Products[i].ProductStock,
			CreatedDate:  &u.GormUser.Products[i].CreatedDate,
			LastUpdate:   &u.GormUser.Products[i].LastUpdate,
		})
	}

	return u.GrapResponseLogin, nil
}

func (u *UserRepo) SelectUsers(limit int) ([]*model.User, error) {
	var db gorm.DB = *config.DB

	u.GormUser = &entity.User{}
	u.GormUsers = []entity.User{}

	err := db.Model(u.GormUser).Preload("UserStatus").Preload("Products").Find(&u.GormUsers)

	if err.Error != nil {
		log.Print(err.Error.Error())
		return nil, errors.New("something wrong")
	}

	u.Users = []*model.User{}

	for i := 0; i < len(u.GormUsers); i++ {
		u.User = &model.User{
			UserID:      &u.GormUsers[i].UserId,
			UserEmail:   &u.GormUsers[i].UserEmail,
			UserName:    &u.GormUsers[i].UserName,
			UserImage:   &u.GormUsers[i].UserImage,
			UserStatus:  &u.GormUsers[i].UserStatus.Status,
			CreatedDate: &u.GormUsers[i].CreatedDate,
			LastUpdate:  &u.GormUsers[i].LastUpdate,
		}

		for j := 0; j < len(u.GormUsers[i].Products); j++ {
			u.User.Products = append(u.User.Products, &model.Product{
				UserID:       &u.GormUsers[i].Products[j].UserId,
				ProductID:    &u.GormUsers[i].Products[j].ProductId,
				ProductName:  &u.GormUsers[i].Products[j].ProductName,
				ProductPrice: &u.GormUsers[i].Products[j].ProductPrice,
				ProductStock: &u.GormUsers[i].Products[j].ProductStock,
				CreatedDate:  &u.GormUsers[i].Products[j].CreatedDate,
				LastUpdate:   &u.GormUsers[i].Products[j].LastUpdate,
			})
		}

		u.Users = append(u.Users, u.User)
	}

	return u.Users, nil
}

func (u *UserRepo) InsertUsert(input *model.Register) (*model.ResponseRegister, error) {
	var db gorm.DB = *config.DB
	var time string = time.Now().Format("20060102150405")

	u.GormUser = &entity.User{
		UserId:       uuid.New().String(),
		UserEmail:    input.UserEmail,
		UserName:     input.UserName,
		UserPassword: input.UserPassword,
		UserSession:  "12345",
		UserStatusId: 1,
		CreatedDate:  time,
		LastUpdate:   time,
	}

	// begin
	tx := db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Print(tx.Error.Error())
		return nil, tx.Error
	}

	// logic
	createUser := tx.Create(&u.GormUser)

	if createUser.Error != nil {
		log.Print(createUser.Error.Error())
		tx.Rollback()
		return nil, createUser.Error
	}

	// comit
	comit := tx.Commit()

	if comit.Error != nil {
		log.Print(comit.Error.Error())
		return nil, comit.Error
	}

	status := 200
	msg := "ok!"

	u.GrapResponseRegister = &model.ResponseRegister{
		Status:  &status,
		Message: &msg,
	}

	return u.GrapResponseRegister, nil
}
