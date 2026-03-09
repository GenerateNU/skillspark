package translations

import (
	"context"
	"sync"
)

func (t *TranslateClient) CallTranslateAPI(ctx context.Context, srcInputs []*string, AcceptLanguage string) (map[string]*string, error) {
	var sl string
	var dl string
	var wg sync.WaitGroup
	response := make(map[string]*string)
	errors := make(chan error, len(srcInputs))

	derefedInputs := CreateDerefedSlice(srcInputs)

	switch AcceptLanguage {
	case "th-TH":
		sl = "th"
		dl = "en"
	case "en-US":
		sl = "en"
		dl = "th"
	}

	for idx := range derefedInputs {
		text := derefedInputs[idx]
		wg.Add(1)
		go func() {
			defer wg.Done()
			output, err := t.GetTranslation(ctx, text, sl, dl)
			if err != nil {
				errors <- err
			}
			response[text] = output
		}()

	}

	wg.Wait()
	close(errors)
	for err := range errors {
		return nil, err
	}

	return response, nil
}

func CreateDerefedSlice(srcInputs []*string) []string {

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	var derefedInputs []string
	for idx := range srcInputs {
		derefedElement := deref(srcInputs[idx])
		derefedInputs = append(derefedInputs, derefedElement)
	}

	return derefedInputs
}
