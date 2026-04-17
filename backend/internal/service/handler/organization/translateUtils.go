package organization

import (
	"context"
	"skillspark/internal/models"
)

// translateAbout translates the user-provided About text into the opposite language.
// Returns (aboutEN, aboutTH) ready to be written to the DB columns. If the input is
// nil or empty, both returned pointers are nil.
func (h *Handler) translateAbout(ctx context.Context, about *string, acceptLanguage string) (*string, *string, error) {
	if about == nil || *about == "" {
		return nil, nil, nil
	}

	translated, err := h.TranslateClient.CallTranslateAPI(ctx, []*string{about}, acceptLanguage)
	if err != nil {
		return nil, nil, err
	}

	translatedText := translated[*about]

	switch acceptLanguage {
	case "th-TH":
		return translatedText, about, nil
	default:
		return about, translatedText, nil
	}
}

// buildCreateOrgDBInput constructs the DB-layer input by translating About and
// mapping the public body into the bilingual struct.
func buildCreateOrgDBInput(input *models.CreateOrganizationInput, aboutEN, aboutTH *string) *models.CreateOrganizationDBInput {
	return &models.CreateOrganizationDBInput{
		AcceptLanguage: input.AcceptLanguage,
		Body: models.CreateOrgDBBody{
			Name:       input.Body.Name,
			AboutEN:    aboutEN,
			AboutTH:    aboutTH,
			Active:     input.Body.Active,
			LocationID: input.Body.LocationID,
			Links:      input.Body.Links,
		},
	}
}

func buildUpdateOrgDBInput(input *models.UpdateOrganizationInput, aboutEN, aboutTH *string) *models.UpdateOrganizationDBInput {
	return &models.UpdateOrganizationDBInput{
		AcceptLanguage: input.AcceptLanguage,
		ID:             input.ID,
		Body: models.UpdateOrgDBBody{
			Name:       input.Body.Name,
			AboutEN:    aboutEN,
			AboutTH:    aboutTH,
			Active:     input.Body.Active,
			LocationID: input.Body.LocationID,
			Links:      input.Body.Links,
		},
	}
}
