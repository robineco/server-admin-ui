package services

import (
	"de.robinscloud.admin-backend/utils"
	"errors"
	"os"
)

type ServiceInterface interface {
	CheckServiceStatus(service string) (bool, error)
}

func CheckServiceStatus(service string) (bool , error) {
	var address = ""
	if service == "nextcloud"{
		address = os.Getenv("HOST_NEXTCLOUD")
	} else if service == "bitwarden" {
		address = os.Getenv("HOST_BITWARDEN")
	} else {
		return false, errors.New("unknown service")
	}
	status, err := utils.PingHost(address)
	if err != nil {
		return false, err
	}
	return status, nil
}
