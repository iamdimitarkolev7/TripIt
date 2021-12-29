package utils

import "tripit/models"

type ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseUser struct {
	Success bool        `json:"success"`
	Message models.User `json:"message"`
}
