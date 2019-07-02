package ado

import "penghui.com/model"

func UserAdo() ([]model.User, error) {
	user, err := model.UserModel()
	return user, err
}

func UserKeyId() (map[string]string, error) {
	userK, err := model.UserKeyId()
	data := map[string]string{"userInfo": userK}
	return data, err
}
