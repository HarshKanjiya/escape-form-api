package mapper

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

func ToEdgeResponse(edge *models.Edge) *types.EdgeResponse {
	return &types.EdgeResponse{
		ID:           edge.ID,
		FormID:       edge.FormID,
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		Condition:    edge.Condition,
	}
}

func ToQuestionResponse(question *models.Question) *types.QuestionResponse {
	optionResp := make([]*types.QueOptionResponse, len(question.Options))
	for i := range question.Options {
		optionResp[i] = ToQuestionOptionResponse(&question.Options[i])
	}

	return &types.QuestionResponse{
		ID:          question.ID,
		FormID:      question.FormID,
		Title:       question.Title,
		Placeholder: question.Placeholder,
		Description: question.Description,
		Required:    question.Required,
		Type:        question.Type,
		Metadata:    question.Metadata,
		PosX:        question.PosX,
		PosY:        question.PosY,
		SortOrder:   question.SortOrder,
		Options:     optionResp,
	}
}

func ToQuestionOptionResponse(option *models.QuestionOption) *types.QueOptionResponse {
	return &types.QueOptionResponse{
		ID:         option.ID,
		QuestionID: option.QuestionID,
		Label:      option.Label,
		Value:      option.Value,
		SortOrder:  option.SortOrder,
	}
}

func ToFormResponse(form *models.Form) *types.FormResponse {
	description := ""
	if form.Description != nil {
		description = *form.Description
	}
	logoURL := ""
	if form.LogoURL != nil {
		logoURL = *form.LogoURL
	}
	status := models.FormStatusDraft
	if form.Status != nil {
		status = *form.Status
	}
	uniqueSubdomain := ""
	if form.UniqueSubdomain != nil {
		uniqueSubdomain = *form.UniqueSubdomain
	}
	customDomain := ""
	if form.CustomDomain != nil {
		customDomain = *form.CustomDomain
	}

	edgeResp := make([]*types.EdgeResponse, len(form.Edges))
	for i := range form.Edges {
		edgeResp[i] = ToEdgeResponse(&form.Edges[i])
	}

	questionResp := make([]*types.QuestionResponse, len(form.Questions))
	for i := range form.Questions {
		questionResp[i] = ToQuestionResponse(&form.Questions[i])
	}

	return &types.FormResponse{
		ID:                  form.ID,
		Name:                form.Name,
		Description:         description,
		TeamID:              form.TeamID,
		ProjectID:           form.ProjectID,
		Theme:               form.Theme,
		LogoURL:             logoURL,
		MaxResponses:        form.MaxResponses,
		OpenAt:              utils.GetIsoDateTime(form.OpenAt),
		CloseAt:             utils.GetIsoDateTime(form.CloseAt),
		Status:              status,
		UniqueSubdomain:     uniqueSubdomain,
		CustomDomain:        customDomain,
		RequireConsent:      form.RequireConsent,
		AllowAnonymous:      form.AllowAnonymous,
		MultipleSubmissions: form.MultipleSubmissions,
		PasswordProtected:   form.PasswordProtected,
		AnalyticsEnabled:    form.AnalyticsEnabled,
		Metadata:            form.Metadata,
		Valid:               form.Valid,
		CreatedAt:           utils.GetIsoDateTime(form.CreatedAt),
		UpdatedAt:           utils.GetIsoDateTime(form.UpdatedAt),
		CreatedBy:           form.CreatedBy,
		FormPageType:        form.FormPageType,
		ResponseCount:       form.ResponseCount,
		Questions:           questionResp,
		Edges:               edgeResp,
	}
}

func ToActivePasswordResponse(password *models.ActivePassword) *types.ActivePasswordResponse {
	return &types.ActivePasswordResponse{
		ID:         password.ID,
		FormID:     password.FormID,
		Name:       password.Name,
		Password:   password.Password,
		IsValid:    password.IsValid,
		UsableUpto: password.UsableUpto,
		CreatedAt:  utils.GetIsoDateTime(&password.CreatedAt),
		ExpireAt:   utils.GetIsoDateTime(password.ExpireAt),
	}
}
