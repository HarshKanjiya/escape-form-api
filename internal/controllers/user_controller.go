package controllers

import (
	"strconv"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	userService *services.UserService
	config      *config.Config
}

func NewUserController(userService *services.UserService, cfg *config.Config) *UserController {
	return &UserController{
		userService: userService,
		config:      cfg,
	}
}

// @Router /auth/register [post]
func (ctrl *UserController) Register(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	// Validate request
	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	// Create user
	user, err := ctrl.userService.Create(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")

		if utils.IsDuplicateError(err) {
			return utils.Conflict(c, "User with this email already exists")
		}

		return utils.InternalServerError(c, "Failed to create user")
	}

	return utils.Created(c, "User registered successfully", user.ToResponse())
}

// @Router /auth/login [post]
func (ctrl *UserController) Login(c *fiber.Ctx) error {
	var req models.LoginRequest

	// Validate request
	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	// Authenticate user
	user, err := ctrl.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("Authentication failed")
		return utils.Unauthorized(c, "Invalid email or password")
	}

	// Generate JWT tokens
	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
		ctrl.config.JWT.Secret,
		ctrl.config.JWT.Expiry,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return utils.InternalServerError(c, "Failed to generate authentication token")
	}

	refreshToken, err := utils.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
		ctrl.config.JWT.Secret,
		ctrl.config.JWT.RefreshExpiry,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate refresh token")
		return utils.InternalServerError(c, "Failed to generate refresh token")
	}

	response := models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User:         user.ToResponse(),
	}

	return utils.Success(c, "Login successful", response)
}

// @Router /users [get]
func (ctrl *UserController) GetAll(c *fiber.Ctx) error {
	var pagination models.PaginationQuery

	if err := c.QueryParser(&pagination); err != nil {
		pagination = models.GetDefaultPagination()
	}

	users, paginationResp, err := ctrl.userService.GetAll(pagination)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch users")
		return utils.InternalServerError(c, "Failed to fetch users")
	}

	// Convert users to response format
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	return utils.SuccessWithMeta(c, "Users fetched successfully", userResponses, paginationResp)
}

// @Router /users/{id} [get]
func (ctrl *UserController) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequest(c, "Invalid user ID")
	}

	user, err := ctrl.userService.GetByID(uint(id))
	if err != nil {
		if utils.IsNotFoundError(err) {
			return utils.NotFound(c, "User not found")
		}
		log.Error().Err(err).Uint64("id", id).Msg("Failed to fetch user")
		return utils.InternalServerError(c, "Failed to fetch user")
	}

	return utils.Success(c, "User fetched successfully", user.ToResponse())
}

// @Router /users/{id} [put]
func (ctrl *UserController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequest(c, "Invalid user ID")
	}

	var req models.UpdateUserRequest
	if err := utils.ValidateRequest(c, &req); err != nil {
		return err
	}

	user, err := ctrl.userService.Update(uint(id), req)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return utils.NotFound(c, "User not found")
		}
		if utils.IsDuplicateError(err) {
			return utils.Conflict(c, "Email already in use")
		}
		log.Error().Err(err).Uint64("id", id).Msg("Failed to update user")
		return utils.InternalServerError(c, "Failed to update user")
	}

	return utils.Success(c, "User updated successfully", user.ToResponse())
}

// @Router /users/{id} [delete]
func (ctrl *UserController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.BadRequest(c, "Invalid user ID")
	}

	err = ctrl.userService.Delete(uint(id))
	if err != nil {
		if utils.IsNotFoundError(err) {
			return utils.NotFound(c, "User not found")
		}
		log.Error().Err(err).Uint64("id", id).Msg("Failed to delete user")
		return utils.InternalServerError(c, "Failed to delete user")
	}

	return utils.Success(c, "User deleted successfully", nil)
}

// @Router /users/search [get]
func (ctrl *UserController) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return utils.BadRequest(c, "Search query is required")
	}

	var pagination models.PaginationQuery
	if err := c.QueryParser(&pagination); err != nil {
		pagination = models.GetDefaultPagination()
	}

	users, paginationResp, err := ctrl.userService.Search(query, pagination)
	if err != nil {
		log.Error().Err(err).Str("query", query).Msg("Failed to search users")
		return utils.InternalServerError(c, "Failed to search users")
	}

	// Convert users to response format
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	return utils.SuccessWithMeta(c, "Search completed successfully", userResponses, paginationResp)
}
