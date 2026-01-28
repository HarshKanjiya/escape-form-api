package repositories

import (
	"math"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DashRepo struct {
	q *query.Query
}

func NewDashRepo(db *gorm.DB) *DashRepo {
	return &DashRepo{
		q: query.Use(db),
	}
}

func (r *DashRepo) FetchAnalytics(ctx *fiber.Ctx, formId string) (*types.FormAnalytics, error) {
	// Get all responses for this form
	var responses []*models.Response
	responses, err := r.q.Response.WithContext(ctx.Context()).
		Where(r.q.Response.FormID.Eq(formId), r.q.Response.Valid.Is(true)).
		Select(
			r.q.Response.ID,
			r.q.Response.StartedAt,
			r.q.Response.SubmittedAt,
			r.q.Response.Status,
		).Find()
	if err != nil {
		return nil, err
	}

	// Calculate analytics
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

	// Calculate completion rate
	completionRate := 0
	if opened > 0 {
		completionRate = int(math.Round(float64(submitted) / float64(opened) * 100))
	}

	// Calculate average, min, and max completion time
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

	// Calculate today's responses
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	todayResponseCount := 0
	for _, r := range responses {
		if r.StartedAt != nil && !r.StartedAt.Before(today) && r.StartedAt.Before(tomorrow) {
			todayResponseCount++
		}
	}

	// Generate last 12 months data
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

func (r *DashRepo) GetQuestions(ctx *fiber.Ctx, formId string) ([]*models.Question, error) {
	// TODO: Implement GetQuestions logic
	return nil, nil
}

func (r *DashRepo) GetResponses(ctx *fiber.Ctx, formId string) ([]*models.Response, error) {
	// TODO: Implement GetResponses logic
	return nil, nil
}

func (r *DashRepo) GetPasswords(ctx *fiber.Ctx, formId string) ([]*models.ActivePassword, error) {
	var passwords []*models.ActivePassword
	passwords, err := r.q.ActivePassword.WithContext(ctx.Context()).
		Where(r.q.ActivePassword.FormID.Eq(formId)).Find()
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

func (r *DashRepo) CreatePassword(ctx *fiber.Ctx, formId string, password *types.ActivePasswordDto) (*models.ActivePassword, error) {
	var expireAt *time.Time
	if password.ExpireAt != "" {
		parsedTime, err := time.Parse("2006-01-02", password.ExpireAt)
		if err != nil {
			return nil, err
		}
		expireAt = &parsedTime
	}

	newPassword := &models.ActivePassword{
		ID:        password.ID,
		FormID:    formId,
		Password:  password.Password,
		Name:      password.Name,
		ExpireAt:  *expireAt,
		IsValid:   true,
		CreatedAt: time.Now(),
	}

	err := r.q.ActivePassword.WithContext(ctx.Context()).Create(newPassword)
	if err != nil {
		return nil, err
	}

	return newPassword, nil
}

func (r *DashRepo) UpdatePassword(ctx *fiber.Ctx, formId string, passwordId string, body map[string]interface{}) (*models.ActivePassword, error) {
	// TODO: Implement UpdatePassword logic
	return nil, nil
}

func (r *DashRepo) DeletePassword(ctx *fiber.Ctx, passwordId string) error {

	_, err := r.q.ActivePassword.WithContext(ctx.Context()).
		Where(r.q.ActivePassword.ID.Eq(passwordId)).
		Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *DashRepo) UpdateSecurity(ctx *fiber.Ctx, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSecurity logic
	return nil, nil
}

func (r *DashRepo) UpdateSettings(ctx *fiber.Ctx, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSettings logic
	return nil, nil
}
