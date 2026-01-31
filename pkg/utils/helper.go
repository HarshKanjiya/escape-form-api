package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetIsoDateTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format("2006-01-02T15:04:05Z07:00")
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GetUserId(ctx *fiber.Ctx) (string, bool) {
	id := ctx.Locals("user_id", "").(string)
	if id == "" {
		return "", false
	}
	return id, true
}
