package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	utils2 "backend/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	DB *database.DB
}

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.JsonResponse{data=[]models.Todo}
// @Success 204 "No Content"
// @Router /todos [get]
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.DB.GetTodos()
	if err != nil {
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	if len(todos) == 0 {
		utils2.RespondNoContent(w, "No todos found")
		return
	}

	utils2.RespondJSON(w, http.StatusOK, utils2.JsonResponse{
		Status: "success",
		Data:   todos,
	})
}

// GetTodo godoc
// @Summary Get a single todo
// @Description Get a todo by its ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} utils.JsonResponse{data=models.Todo}
// @Failure 404 {object} utils.JsonResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils2.BadRequestError(err).LogAndRespond(w)
		return
	}

	todo, err := h.DB.GetTodo(id)
	if err != nil {
		if err == database.ErrNotFound {
			utils2.RespondError(w, http.StatusNotFound, "Todo not found")
			return
		}
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	utils2.RespondJSON(w, http.StatusOK, utils2.JsonResponse{
		Status: "success",
		Data:   todo,
	})
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body models.CreateTodoRequest true "New todo"
// @Success 201 {object} utils.JsonResponse{data=models.Todo}
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils2.BadRequestError(err).LogAndRespond(w)
		return
	}

	todo := models.Todo{
		Title:     req.Title,
		Completed: req.Completed,
		CreatedAt: time.Now(), // Set the current time
	}

	if err := h.DB.CreateTodo(&todo); err != nil {
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	utils2.RespondJSON(w, http.StatusCreated, utils2.JsonResponse{
		Status:  "success",
		Message: "Todo created successfully",
		Data:    todo,
	})
}

// UpdateTodo godoc
// @Summary Update an existing todo
// @Description Update an existing todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body models.UpdateTodoRequest true "Updated todo"
// @Success 200 {object} utils.JsonResponse{data=models.Todo}
// @Failure 404 {object} utils.JsonResponse
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils2.BadRequestError(err).LogAndRespond(w)
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils2.BadRequestError(err).LogAndRespond(w)
		return
	}

	todo, err := h.DB.GetTodo(id)
	if err != nil {
		if err == database.ErrNotFound {
			utils2.RespondError(w, http.StatusNotFound, "Todo not found")
			return
		}
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	if req.Title != "" {
		todo.Title = req.Title
	}
	todo.Completed = req.Completed

	if err := h.DB.UpdateTodo(&todo); err != nil {
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	utils2.RespondJSON(w, http.StatusOK, utils2.JsonResponse{
		Status:  "success",
		Message: "Todo updated successfully",
		Data:    todo,
	})
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} utils.JsonResponse
// @Failure 404 {object} utils.JsonResponse
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils2.BadRequestError(err).LogAndRespond(w)
		return
	}

	if err := h.DB.DeleteTodo(id); err != nil {
		if err == database.ErrNotFound {
			utils2.RespondError(w, http.StatusNotFound, "Todo not found")
			return
		}
		utils2.InternalServerError(err).LogAndRespond(w)
		return
	}

	utils2.RespondJSON(w, http.StatusOK, utils2.JsonResponse{
		Status:  "success",
		Message: "Todo deleted successfully",
	})
}
