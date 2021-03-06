package dto

import "gorm.io/datatypes"

// EventCreateParams - input params for event create post request
type EventCreateParams struct {
	Email       string         `json:"email" binding:"required,email"`
	Environment string         `json:"environment" binding:"required"`
	Component   string         `json:"component" binding:"required"`
	Message     string         `json:"message" binding:"required"`
	Data        datatypes.JSON `json:"data"`
}

// EventFilterParams - holder for filter query params
type EventFilterParams struct {
	Email       string `form:"email"`
	Environment string `form:"environment"`
	Component   string `form:"component"`
	Message     string `form:"message"`
	From        string `form:"from"`
}
