package main

import "swag/api"

func main() {
	// GetUser godoc
	// @Summary Get all users
	// @Description Get a list of users
	// @Tags User
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router /users [get]

	api.InitService()
}
