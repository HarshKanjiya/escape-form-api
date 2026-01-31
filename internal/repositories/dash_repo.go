package repositories

import (
	"context"
	"math"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"gorm.io/gorm"
)

type IDashRepo interface {

	// OPS
	GetAnalytics(c context.Context, formId string) (*types.FormAnalytics, error)
	GetQuestions(ctx context.Context, formId string) ([]*models.Question, error)
	GetResponses(ctx context.Context, formId string) ([]*models.Response, error)

	// PASSWORD CONFIG
	GetPasswords(ctx context.Context, formId string) ([]*models.ActivePassword, error)
	CreatePassword(ctx context.Context, formId string, password *models.ActivePassword) (*models.ActivePassword, error)
	UpdatePassword(ctx context.Context, formId string, password *models.ActivePassword) (*models.ActivePassword, error)
	DeletePassword(ctx context.Context, passwordId string) error

	// FORM SETTINGS
	UpdateSecurity(ctx context.Context, formId string, body map[string]interface{}) (interface{}, error)
	UpdateSettings(ctx context.Context, formId string, body map[string]interface{}) (interface{}, error)
}

type DashRepo struct {
	db *gorm.DB
}

func NewDashRepo(db *gorm.DB) *DashRepo {
	return &DashRepo{
		db: db,
	}
}

func (r *DashRepo) GetAnalytics(ctx context.Context, formId string) (*types.FormAnalytics, error) {

	// Get all responses for this form
	var responses []*models.Response
	err := r.db.WithContext(ctx).Model(&models.Response{}).
		Where("formId = ? AND valid = ?", formId, true).
		Find(&responses).Error
	if err != nil {
		return nil, errors.Internal(err)
	}

	// // Calculate analytics
	responseCount := len(responses)
	opened := 0
	submitted := 0
	var completionTimes []float64

	for _, r := range responses {
		if r.StartedAt != nil {
			opened++
		}
		if r.SubmittedAt != nil {
			submitted++
		}

		// Calculate completion time if both timestamps exist
		if r.StartedAt != nil && r.SubmittedAt != nil {
			completionTime := r.SubmittedAt.Sub(*r.StartedAt).Seconds()
			completionTimes = append(completionTimes, completionTime)
		}
	}

	// // Calculate completion rate
	completionRate := 0
	if opened > 0 {
		completionRate = int(math.Round(float64(submitted) / float64(opened) * 100))
	}

	// // Calculate average, min, and max completion time
	avgCompletionTime := 0
	minCompletionTime := 0
	maxCompletionTime := 0

	if len(completionTimes) > 0 {
		sum := 0.0
		minTime := completionTimes[0]
		maxTime := completionTimes[0]

		for _, time := range completionTimes {
			sum += time
			if time < minTime {
				minTime = time
			}
			if time > maxTime {
				maxTime = time
			}
		}

		avgCompletionTime = int(math.Round(sum / float64(len(completionTimes))))
		minCompletionTime = int(math.Round(minTime))
		maxCompletionTime = int(math.Round(maxTime))
	}

	// // Calculate today's responses
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	todayResponseCount := 0
	for _, r := range responses {
		if r.StartedAt != nil && !r.StartedAt.Before(today) && r.StartedAt.Before(tomorrow) {
			todayResponseCount++
		}
	}

	// // Generate last 12 months data
	now := time.Now()
	monthNames := []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	submitDataPoints := make([]types.MonthlySubmitData, 0, 12)

	for i := 11; i >= 0; i-- {
		date := time.Date(now.Year(), now.Month()-time.Month(i), 1, 0, 0, 0, 0, now.Location())
		monthName := monthNames[date.Month()-1]
		nextMonth := date.AddDate(0, 1, 0)

		monthResponsesCompleted := 0
		monthResponsesUnfinished := 0

		for _, r := range responses {
			if r.StartedAt != nil && !r.StartedAt.Before(date) && r.StartedAt.Before(nextMonth) {
				if r.SubmittedAt != nil {
					monthResponsesCompleted++
				} else {
					monthResponsesUnfinished++
				}
			}
		}

		submitDataPoints = append(submitDataPoints, types.MonthlySubmitData{
			Month:      monthName,
			Unfinished: monthResponsesUnfinished,
			Completed:  monthResponsesCompleted,
		})
	}

	analytics := &types.FormAnalytics{
		ResponseCount:      responseCount,
		AvgCompletionTime:  avgCompletionTime,
		MinCompletionTime:  minCompletionTime,
		MaxCompletionTime:  maxCompletionTime,
		Opened:             opened,
		Submitted:          submitted,
		CompletionRate:     completionRate,
		TodayResponseCount: todayResponseCount,
		SubmitDataPoints:   submitDataPoints,
	}

	return analytics, nil
}

func (r *DashRepo) GetQuestions(ctx context.Context, formId string) ([]*models.Question, error) {
	var questions []*models.Question
	err := r.db.WithContext(ctx).Preload("Options").Model(&models.Question{}).
		Where("formId = ?", formId).
		Order("orderBy ASC").
		Find(&questions).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return questions, nil
}

func (r *DashRepo) GetResponses(ctx context.Context, formId string) ([]*models.Response, error) {
	// TODO: Implement GetResponses logic
	return nil, nil
}

func (r *DashRepo) GetPasswords(ctx context.Context, formId string) ([]*models.ActivePassword, error) {
	var passwords []*models.ActivePassword
	err := r.db.WithContext(ctx).Model(&models.ActivePassword{}).
		Where("formId = ?", formId).
		Find(&passwords).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return passwords, nil
}

func (r *DashRepo) CreatePassword(ctx context.Context, formId string, password *models.ActivePassword) (*models.ActivePassword, error) {

	err := r.db.WithContext(ctx).
		Model(&models.ActivePassword{}).
		Create(password).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return password, nil
}

func (r *DashRepo) UpdatePassword(ctx context.Context, formId string, password *models.ActivePassword) (*models.ActivePassword, error) {

	err := r.db.WithContext(ctx).
		Model(&models.ActivePassword{}).
		Where("id = ? AND formId = ?", password.ID, formId).
		Updates(map[string]interface{}{
			"password":   password.Password,
			"name":       password.Name,
			"usableUpto": password.UsableUpto,
			"expireAt":   password.ExpireAt,
			"isValid":    password.IsValid,
		}).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return password, nil
}

func (r *DashRepo) DeletePassword(ctx context.Context, passwordId string) error {

	err := r.db.WithContext(ctx).
		Model(&models.ActivePassword{}).
		Where("id = ?", passwordId).Delete(&models.ActivePassword{}).Error

	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *DashRepo) UpdateSecurity(ctx context.Context, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSecurity logic
	return nil, nil
}

func (r *DashRepo) UpdateSettings(ctx context.Context, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSettings logic
	return nil, nil
}
