package models

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
)

type Auth struct {
	ID       int `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Auth) TableName() string {
	return "members"
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (int, error) {
	var auth Auth
	err := db.Debug().Select("id").Where(Auth{Email: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if auth.ID > 0 {
		return auth.ID, nil
	}

	return 0, nil
}

func CheckEmail(email string) int {
	var auth Auth
	err := db.Debug().Select("id").Where(Auth{Email: email}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0
	}
	if auth.ID > 0 {
		return auth.ID
	}

	return 0
}

func AddMember(data map[string]interface{}) int {
	username := RandInt(1000000000, 999999999)
	auth := Auth{
		Name: strconv.Itoa(username),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	if err := db.Create(&auth).Error; err != nil {
		return 0
	}
	return auth.ID
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

type Member struct {
	Auth
	AmountAll    float64
	AmountFrozen float64
	AmountUsed   float64
}

func GetMember(id int) (*Member, error) {
	var member Member
	err := db.Debug().Where("id = ?", id).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &member, nil
}

func UpdateMember(id int, data interface{}) error {
	if err := db.Model(&Member{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
