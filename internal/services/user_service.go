package services

import (
	"fmt"

	"github.com/HarshKanjiya/escape-form-api/internal/database"
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/gorm"
)

// UserService handles business logic for user operations
type UserService struct{}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{}
}

// Create creates a new user with hashed password
func (s *UserService) Create(req models.CreateUserRequest) (*models.User, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = "user"
	}

	// Create user model
	user := models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      role,
		IsActive:  true,
	}

	// Save to database
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByID retrieves a user by their ID
func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email address
func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users with pagination
func (s *UserService) GetAll(pagination models.PaginationQuery) ([]models.User, *models.PaginationResponse, error) {
	pagination.Normalize()
	
	var users []models.User
	var totalItems int64

	// Count total items
	if err := database.DB.Model(&models.User{}).Count(&totalItems).Error; err != nil {
		return nil, nil, err
	}

	// Get paginated users
	query := database.DB.
		Limit(pagination.PageSize).
		Offset(pagination.GetOffset()).
		Order(fmt.Sprintf("%s %s", pagination.SortBy, pagination.Order))

	if err := query.Find(&users).Error; err != nil {
		return nil, nil, err
	}

	// Calculate total pages
	totalPages := int(totalItems) / pagination.PageSize
	if int(totalItems)%pagination.PageSize != 0 {
		totalPages++
	}

	paginationResp := &models.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	return users, paginationResp, nil
}

// Update updates a user's information
func (s *UserService) Update(id uint, req models.UpdateUserRequest) (*models.User, error) {
	user, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	updates := make(map[string]interface{})
	
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	// Perform update
	if err := database.DB.Model(user).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Reload user to get updated values
	if err := database.DB.First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete soft deletes a user by ID
func (s *UserService) Delete(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Authenticate verifies user credentials and returns the user if valid
func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verify password
	if !utils.VerifyPassword(user.Password, password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("user account is inactive")
	}

	return user, nil
}

// Search searches for users by name or email
func (s *UserService) Search(query string, pagination models.PaginationQuery) ([]models.User, *models.PaginationResponse, error) {
	pagination.Normalize()
	
	var users []models.User
	var totalItems int64

	searchPattern := "%" + query + "%"
	
	// Count total items
	countQuery := database.DB.Model(&models.User{}).Where(
		"email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?",
		searchPattern, searchPattern, searchPattern,
	)
	if err := countQuery.Count(&totalItems).Error; err != nil {
		return nil, nil, err
	}

	// Get paginated results
	dbQuery := database.DB.
		Where("email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?",
			searchPattern, searchPattern, searchPattern).
		Limit(pagination.PageSize).
		Offset(pagination.GetOffset()).
		Order(fmt.Sprintf("%s %s", pagination.SortBy, pagination.Order))

	if err := dbQuery.Find(&users).Error; err != nil {
		return nil, nil, err
	}

	// Calculate total pages
	totalPages := int(totalItems) / pagination.PageSize
	if int(totalItems)%pagination.PageSize != 0 {
		totalPages++
	}

	paginationResp := &models.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	return users, paginationResp, nil
}
