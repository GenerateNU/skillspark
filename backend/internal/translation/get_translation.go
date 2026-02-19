package translations

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TranslateResponse struct {
	DestinationText string `json:"destination-text"`
}

func (t *TranslateClient) GetTranslation(ctx context.Context, input string, sl string, dl string) (*string, error) {
	var result TranslateResponse

	encodedInput := url.QueryEscape(input)
	apiUrl := os.Getenv("TRANSLATIONS_API_URL")

	params := []string{"sl=", sl, "&dl=", dl, "&text="}
	paramSpecs := strings.Join(params, "")
	url := apiUrl + paramSpecs + encodedInput

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", `application/json`)
	resp, err := t.Client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &result.DestinationText, nil

}
