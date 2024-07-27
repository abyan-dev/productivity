package handler

import (
	"fmt"
	"log/slog"

	"github.com/abyan-dev/productivity/pkg/model"
	"github.com/abyan-dev/productivity/pkg/response"
	"github.com/abyan-dev/productivity/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CreateTaskPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type UpdateTaskPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	IsComplete  bool   `json:"is_complete"`
}

func CreateTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	claims, ok := c.Locals("user").(jwt.MapClaims)
	if !ok {
		return response.Unauthorized(c, "Invalid user claims")
	}

	email, emailOk := claims["email"].(string)
	if !emailOk {
		return response.Unauthorized(c, "Invalid email claim")
	}

	requestPayload := CreateTaskPayload{}

	if err := c.BodyParser(&requestPayload); err != nil {
		return response.BadRequest(c, "Invalid request payload")
	}

	isEmailValid, emailVadFeedback := utils.ValidateEmail(email)
	if !isEmailValid {
		return response.BadRequest(c, emailVadFeedback)
	}

	isDueDateValid, dueDateValFeedback, dueDate := utils.ValidateTime(requestPayload.DueDate)
	if !isDueDateValid {
		return response.BadRequest(c, dueDateValFeedback)
	}

	task := model.Task{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		DueDate:     dueDate,
		IsComplete:  false,
		UserEmail:   email,
	}

	if err := db.Create(&task).Error; err != nil {
		return response.InternalServerError(c, "Failed to create task.")
	}

	return response.Created(c, "Successfully created task.")
}

func GetAllTasks(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	claims, ok := c.Locals("user").(jwt.MapClaims)
	if !ok {
		slog.Error("Invalid user claims")
		return response.Unauthorized(c, "Invalid user claims")
	}

	email, emailOk := claims["email"].(string)
	if !emailOk {
		slog.Error("Invalid email claim", "claims", claims)
		return response.Unauthorized(c, "Invalid email claim")
	}

	slog.Debug("Email claim extracted", slog.String("email", email))

	isEmailValid, emailValFeedback := utils.ValidateEmail(email)
	if !isEmailValid {
		slog.Error("Email validation failed", slog.String("email", email), slog.String("feedback", emailValFeedback))
		return response.BadRequest(c, emailValFeedback)
	}

	slog.Debug("Email validation succeeded", slog.String("email", email))

	var tasks []model.Task
	result := db.Where("user_email = ?", email).Find(&tasks)
	if result.Error != nil {
		slog.Error("Failed to retrieve tasks", slog.String("email", email), slog.String("error", result.Error.Error()))
		return response.InternalServerError(c, "Failed to retrieve tasks.")
	}

	slog.Debug("Tasks retrieved successfully", slog.String("email", email), slog.Int("task_count", len(tasks)))
	return response.Ok(c, fmt.Sprintf("Successfully retrieved all tasks for user %s", email), tasks)
}

func GetTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	var task model.Task
	result := db.First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return response.NotFound(c, "Task not found")
		}
		return response.InternalServerError(c, "Failed to retrieve task.")
	}

	return response.Ok(c, "Successfully retrieved task", task)
}

func UpdateTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	requestPayload := UpdateTaskPayload{}

	if err := c.BodyParser(&requestPayload); err != nil {
		return response.BadRequest(c, "Invalid request payload")
	}

	isDueDateValid, dueDateValFeedback, dueDate := utils.ValidateTime(requestPayload.DueDate)
	if !isDueDateValid {
		return response.BadRequest(c, dueDateValFeedback)
	}

	var task model.Task
	result := db.First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return response.NotFound(c, "Task not found")
		}
		return response.InternalServerError(c, "Failed to retrieve task.")
	}

	task.Title = requestPayload.Title
	task.Description = requestPayload.Description
	task.DueDate = dueDate
	task.IsComplete = requestPayload.IsComplete

	db.Save(&task)

	return response.Ok(c, "Successfully updated task", task)
}

func DeleteTask(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	result := db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return response.InternalServerError(c, "Failed to delete task.")
	}

	if result.RowsAffected == 0 {
		return response.NotFound(c, "Task not found")
	}

	return response.Ok(c, "Successfully deleted task.")
}
