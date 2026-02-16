package translations

import (
	"context"
)

type TranslationInterface interface {
	GetTranslation(ctx context.Context, input string) (*string, error)
}
