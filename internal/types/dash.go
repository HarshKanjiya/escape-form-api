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
