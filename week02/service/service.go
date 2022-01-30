package service

import "week02/dao"

func Query() (dao.Result, error) {
	return dao.Query()
}