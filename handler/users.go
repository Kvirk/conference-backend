package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/kvirk/conference-backend/models"
)

func users(router chi.Router) {
	router.Post("/signup", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	user.Password = userService.HashAndSalt(user.Password)

	if err := dbInstance.AddUser(user); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	userResponse, responseErr := user.ConvertToResponse()

	if responseErr != nil {
		render.Render(w, r, ErrorRenderer(responseErr))
		return
	}

	log.Println(userResponse)

	if err := render.Render(w, r, &userResponse); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
