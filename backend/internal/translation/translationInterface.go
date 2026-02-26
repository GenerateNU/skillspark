package translations

import (
	"context"
)

type TranslationInterface interface {
	CallTranslateAPI(ctx context.Context, srcInputs []*string, AcceptLanguage string) (map[string]*string, error)
}
