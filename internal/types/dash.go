package types

type MonthlySubmitData struct {
	Month      string `json:"month"`
	Unfinished int    `json:"Unfinished"`
	Completed  int    `json:"Completed"`
}

type FormAnalytics struct {
	ResponseCount      int                 `json:"responseCount"`
	AvgCompletionTime  int                 `json:"avgCompletionTime"`
	MinCompletionTime  int                 `json:"minCompletionTime"`
	MaxCompletionTime  int                 `json:"maxCompletionTime"`
	Opened             int                 `json:"opened"`
	Submitted          int                 `json:"submitted"`
	CompletionRate     int                 `json:"completionRate"`
	TodayResponseCount int                 `json:"todayResponseCount"`
	SubmitDataPoints   []MonthlySubmitData `json:"submitDataPoints"`
}

type PasswordRequest struct {
	ID         string `json:"id"`
	FormID     string `json:"formId"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	IsValid    bool   `json:"isValid"`
	UsableUpto int    `json:"usableUpto"`
	ExpireAt   string `json:"expireAt"`
}

type ActivePasswordResponse struct {
	ID         string `json:"id"`
	FormID     string `json:"formId"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	IsValid    bool   `json:"isValid"`
	ExpireAt   string `json:"expireAt"`
	CreatedAt  string `json:"createdAt"`
	UsableUpto int    `json:"usableUpto"`
}
