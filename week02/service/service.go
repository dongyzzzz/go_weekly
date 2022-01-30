package service

import "week02/dao"

func Query() (dao.User, error) {
	return dao.Query()
}