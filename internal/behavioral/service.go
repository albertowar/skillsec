package behavioral

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

type Service struct {
	Model llms.Model
}

func NewService(provider, apiKey, modelName, baseURL string) (*Service, error) {
	var model llms.Model
	var err error

	switch provider {
	case "google":
		model, err = googleai.New(context.Background(), googleai.WithAPIKey(apiKey), googleai.WithDefaultModel(modelName))
	case "openai":
		model, err = openai.New(openai.WithToken(apiKey), openai.WithModel(modelName), openai.WithBaseURL(baseURL))
	default:
		return nil, nil // Not configured
	}

	if err != nil {
		return nil, err
	}
	return &Service{Model: model}, nil
}

func (s *Service) Test(ctx context.Context, systemPrompt, userMessage string) (string, error) {
	if s == nil || s.Model == nil {
		return "Behavioral testing not configured.", nil
	}
	
	resp, err := s.Model.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, systemPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, userMessage),
	})
	if err != nil {
		return "", err
	}
	if resp == nil || len(resp.Choices) == 0 {
		return "", nil
	}
	return resp.Choices[0].Content, nil
}
