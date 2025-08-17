package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"

)

const geminiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponce struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func AskGemini(prompt string) (string, error) {
	apikey := "AIzaSyDvqPYipnjb5jAozUqdmcboOrNqSKSZUWE"
	// apikey := os.Getenv("GOOGLE_API_KEY")
	if apikey == "" {
		return "", fmt.Errorf("Gemini api is invalid")
	}

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	body, _ := json.Marshal(reqBody)


	url := fmt.Sprintf("%s?key=%s", geminiEndpoint, apikey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Gemini error %s", string(data))
	}

	var gemRes GeminiResponce
	if err := json.Unmarshal(data, &gemRes); err != nil {
		return "", err
	}
	if len(gemRes.Candidates) == 0 {
		return "", fmt.Errorf("no response from gemini")
	}

	return gemRes.Candidates[0].Content.Parts[0].Text, nil
}
