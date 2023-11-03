package handler

import (
	"lincoln.boris/forum/models"
	"lincoln.boris/forum/pkg/validator"
)

func (h *Handler) validatePost(post models.Post) map[string]interface{} {
	h.validator.AddRule(validator.MaxLength(100))
	h.validator.AddRule(validator.MinLength(5))
	h.validator.AddRule(validator.Required)

	errors := h.validator.Validate(map[string]interface{}{
		"title": post.Title,
		"body":  post.Body,
	})

	if len(errors) > 0 {
		return map[string]interface{}{
			"message": "validation failed",
			"data":    post,
			"errors":  errors,
		}
	}

	return nil
}

func (h *Handler) validateCategory(category models.Category) map[string]interface{} {
	h.validator.AddRule(validator.MaxLength(20))
	h.validator.AddRule(validator.Required)

	errors := h.validator.Validate(map[string]interface{}{
		"name": category.Name,
	})

	if len(errors) > 0 {
		return map[string]interface{}{
			"message": "validation failed",
			"data":    category,
			"errors":  errors,
		}
	}

	return nil
}
