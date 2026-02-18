package translations

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type TranslateResponse struct {
	DestinationText string `json:"destination-text"`
}

func (t *TranslateClient) GetTranslation(ctx context.Context, input string) (*string, error) {
	var result TranslateResponse

	encodedInput := url.QueryEscape(input)
	apiUrl := os.Getenv("TRANSLATIONS_API_URL")
	url := apiUrl + encodedInput

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", `application/json`)
	resp, err := t.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result.DestinationText, nil

}
