package apirequest

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetApiResponse(fileContent string) (string, error) {
	ctx := context.Background()

	// Access API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		return "", fmt.Errorf("error creating client: %v", err)
	}
	defer client.Close()

	// Access the gemini model using the gemini 1.5 flash
	model := client.GenerativeModel("gemini-1.5-flash")

	// Create the prompt for simplification
	prompt := fmt.Sprintf("Summarize the following text in 2-3 sentences:\n\n%s", fileContent)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("error generating content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	simplifiedContent, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	// Wrap the summary in HTML paragraph tags
	return fmt.Sprintf("<p>%s</p>", string(simplifiedContent)), nil
}