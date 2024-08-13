package api

import (
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func New() *spinhttp.Router {
	r := spinhttp.NewRouter()
	r.GET("/all_users", getAllUsers)
	r.GET("/user/:id", getUser)
	r.POST("/user/:username", createUser)
	r.DELETE("/user/:id", deleteUser)
	return r
}
