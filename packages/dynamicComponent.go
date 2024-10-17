package packages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/nox/content"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	openAIEndpoint = "https://api.openai.com/v1/chat/completions"
)

var (
	startCodeBlockRegex = regexp.MustCompile(`(?m)^` + "```" + `(jsx|tsx)\s*$`)
	endCodeBlockRegex   = regexp.MustCompile(`(?m)^` + "```" + `\s*$`)
)

type OpenAIRequest struct {
	Model            string    `json:"model"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	PresencePenalty  float64   `json:"presence_penalty"`
	Messages         []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type PromptData struct {
	Name          string
	Description   string
	TitleCaseName string
}

func DynamicComponent(name, description string) {
	fmt.Println("✨ Creating a new dynamic component... Hold tight, magic is happening! 📄✨")

	openAIKey, err := GetOpenAIKey()
	if err != nil {
		fmt.Println("Please set the OpenAI key using 'nox --key'")
		return
	}

	titleCaser := cases.Title(language.English)
	titleCaseName := titleCaser.String(name)

	promptData := PromptData{
		Name:          name,
		Description:   description,
		TitleCaseName: titleCaseName,
	}

	prompt, err := executeTemplate(content.PromptTemplate, promptData)
	if err != nil {
		fmt.Printf("Error generating prompt: %v\n", err)
		return
	}

	openAIRequest := OpenAIRequest{
		Model:            "gpt-4o",
		Temperature:      0,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Messages: []Message{
			{Role: "system", Content: content.SystemPrompt},
			{Role: "user", Content: prompt},
		},
	}

	generatedCode, err := callOpenAI(openAIRequest, openAIKey)
	if err != nil {
		fmt.Printf("Error generating component: %v\n", err)
		return
	}

	if generatedCode == "" {
		fmt.Println("No component generated. Please try again.")
		return
	}

	generatedCode = removeCodeMarkers(generatedCode)

	componentPath := filepath.Join("client", "src", "pages", name, "modules", "index.tsx")

	if err := os.WriteFile(componentPath, []byte(generatedCode), 0644); err != nil {
		fmt.Printf("error writing generated code to file")
	}

	fmt.Printf("Congratulations! The component has been updated as per the description. 🎉🚀\n")
	fmt.Println("Happy coding! 🎉🚀")
}

func executeTemplate(tmpl string, data interface{}) (string, error) {
	t, err := template.New("prompt").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func callOpenAI(request OpenAIRequest, apiKey string) (string, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error creating request body: %w", err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to OpenAI: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", fmt.Errorf("error parsing OpenAI response: %w", err)
	}

	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", nil
}

func removeCodeMarkers(code string) string {
	code = startCodeBlockRegex.ReplaceAllString(code, "")

	code = endCodeBlockRegex.ReplaceAllString(code, "")

	return strings.TrimSpace(code)
}
