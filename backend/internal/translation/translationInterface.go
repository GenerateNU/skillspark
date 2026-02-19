package translations

import (
	"context"
)

type TranslationInterface interface {
	GetTranslation(ctx context.Context, input string, sl string, dl string) (*string, error)
}
