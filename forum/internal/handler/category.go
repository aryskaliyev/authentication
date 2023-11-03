package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"lincoln.boris/forum/models"
)

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodPost {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var input models.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(err.Error())
		return
	}

	if validationErrors := h.validateCategory(input); validationErrors != nil {
		h.respondWithJSON(w, http.StatusUnprocessableEntity, validationErrors)
		return
	}

	category_id, err := h.services.Category.Create(input)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"category_id": category_id,
	})
}

func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodGet {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	categories, err := h.services.Category.GetAll()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		h.errorLog(err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, categories)
}

func (h *Handler) getCategoryById(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodGet {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var category_id int

	category_id, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil || category_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(http.StatusText(http.StatusBadRequest))
		return
	}

	input, err := h.services.Category.GetById(category_id)
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


func (h *Handler) updateCategory(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

	if r.Method != http.MethodPost {
		h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
                h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
                return
	}

	var category_id int

	category_id, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil || category_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
                h.errorLog(http.StatusText(http.StatusBadRequest))
                return
	}

	var input models.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(err.Error())
		return
	}

	if validationErrors := h.validateCategory(input); validationErrors != nil {
		h.respondWithJSON(w, http.StatusUnprocessableEntity, validationErrors)
		return
	}

	err = h.services.Category.Update(category_id, input)
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
		"category_id": category_id,
	})
}

func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	h.infoLog(fmt.Sprintf("%v - %v %v %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

        if r.Method != http.MethodPost {
                h.handleError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
                h.errorLog(http.StatusText(http.StatusMethodNotAllowed))
                return
        }

	category_id, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil || category_id < 1 {
		h.handleError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		h.errorLog(http.StatusText(http.StatusBadRequest))
		return
	}

	err = h.services.Category.Delete(category_id)
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
