package handler

import (
	"de.robinscloud.admin-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	ServiceName	string `json:"serviceName"`
}

type ServiceStatus struct {
	ServiceName string	`json:"serviceName"`
	IsOnline	bool	`json:"isOnline"`
}

func CheckServiceStatus(c *gin.Context) {
	var s Service
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid JSON")
		return
	}
	status, err := services.CheckServiceStatus(s.ServiceName)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error while checking service status")
		return
	}
	var service = ServiceStatus{
		ServiceName: s.ServiceName,
		IsOnline: status,
	}
	c.JSON(http.StatusOK, service)
}