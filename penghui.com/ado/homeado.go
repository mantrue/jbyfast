package ado

import "penghui.com/model"

func UserAdo() ([]model.User, error) {
	user, err := model.UserModel()
	return user, err
}
