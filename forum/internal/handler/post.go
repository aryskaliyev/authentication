package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lincoln.boris/forum/models"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodPost {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var input models.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(err.Error())
		return
	}

	if validationErrors := h.validatePost(input); validationErrors != nil {
		h.respondWithJSON(w, http.StatusUnprocessableEntity, validationErrors)
		return
	}

	post_id, err := h.services.Post.Create(input)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"post_id": post_id,
	})
}

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodGet {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	posts, err := h.services.Post.GetAll()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, posts)
}

func (h *Handler) getPostById(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodGet {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var post_id int

	post_id, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil || post_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(http.StatusText(http.StatusBadRequest))
		return
	}

	input, err := h.services.Post.GetById(post_id)
	if err == models.ErrNoRecord {
		h.handleError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		h.errorLog(err.Error())
		return
	} else if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, input)
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodPost {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var post_id int

	post_id, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil || post_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(http.StatusText(http.StatusBadRequest))
		return
	}

	var input models.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(err.Error())
		return
	}

	if validationErrors := h.validatePost(input); validationErrors != nil {
		h.respondWithJSON(w, http.StatusUnprocessableEntity, validationErrors)
		return
	}

	err = h.services.Post.Update(post_id, input)
	if err == models.ErrNoRecord {
		h.handleError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
                h.errorLog(err.Error())
                return
	} else if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"post_id": post_id,
	})
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodPost {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	post_id, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil || post_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(http.StatusText(http.StatusBadRequest))
		return
	}

	err = h.services.Post.Delete(post_id)
	if err == models.ErrNoRecord {
		h.handleError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
                h.errorLog(err.Error())
                return
	} else if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, map[string]interface{}{"Status": "ok"})
}
