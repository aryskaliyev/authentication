package handler

import (
	"net/http"

	"lincoln.boris/forum/internal/service"
	"lincoln.boris/forum/pkg/logger"
	"lincoln.boris/forum/pkg/validator"
)

type Handler struct {
	services  *service.Service
	logger    *logger.Logger
	validator *validator.Validator
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services:  services,
		logger:    logger.NewLogger(),
		validator: validator.NewValidator(),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(h.getAllPosts))
	mux.Handle("/posts", http.HandlerFunc(h.getPostById))
	mux.Handle("/posts/create", http.HandlerFunc(h.createPost))
	mux.Handle("/posts/update", http.HandlerFunc(h.updatePost))
	mux.Handle("/posts/delete", http.HandlerFunc(h.deletePost))

	mux.Handle("/categories/all", http.HandlerFunc(h.getAllCategories))
	mux.Handle("/categories/create", http.HandlerFunc(h.createCategory))
	mux.Handle("/categories", http.HandlerFunc(h.getCategoryById))
	mux.Handle("/categories/update", http.HandlerFunc(h.updateCategory))
	mux.Handle("/categories/delete", http.HandlerFunc(h.deleteCategory))

	return mux
}
