package auth

import (
	"github.com/danielgtaylor/huma/v2"
	"net/http"
	"skillspark/internal/config"
	"log/slog"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"errors"
)

var (
	ErrMissingToken   = errors.New("missing JWT token")
	ErrInvalidToken   = errors.New("invalid JWT token")
	ErrInvalidMethod  = errors.New("unexpected JWT signing method")
)

// SupabaseClaims represents the JWT claims from Supabase Auth
type SupabaseClaims struct {
	Sub           string                 `json:"sub"`
	Email         string                 `json:"email"`
	Phone         string                 `json:"phone"`
	Role          string                 `json:"role"`
	Aud           string                 `json:"aud"`
	AppMetadata   map[string]interface{} `json:"app_metadata"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	jwt.RegisteredClaims
}

// Verifier handles JWT verification
type Verifier struct {
	secret []byte
}

// NewVerifier creates a new JWT verifier
// If secret is empty, it reads from SUPABASE_JWT_SECRET env var
func NewVerifier(secret string) *Verifier {
	if secret == "" {
		secret = os.Getenv("SUPABASE_JWT_SECRET")
	}
	return &Verifier{secret: []byte(secret)}
}

// Verify validates a JWT token and returns the claims
func (v *Verifier) Verify(tokenString string) (*SupabaseClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SupabaseClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidMethod
		}
		return v.secret, nil
	})

	if err != nil {
		return nil, errors.Join(ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*SupabaseClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func AuthMiddleware(api huma.API, cfg *config.Supabase) func(ctx huma.Context, next func(huma.Context)) {
	skipPaths := map[string]bool{
		"/api/v1/auth/signup/guardian": true,
		"/api/v1/auth/signup/manager":  true,
		"/api/v1/auth/login/guardian":  true,
		"/api/v1/auth/login/manager":   true,
		"/api/v1/health":               true,
	}

	return func(ctx huma.Context, next func(huma.Context)) {
		if skipPaths[ctx.Operation().Path] {
			next(ctx)
			return
		}

		cookie, err := huma.ReadCookie(ctx, "jwt")
		if err != nil || cookie.Value == "" {
			err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token Not Found")
			if err != nil {
				slog.Error("Failed to write error", "err", err)
			}
			return
		}
		
		_, err = NewVerifier("").Verify(cookie.Value)
		if err != nil {
			err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid/Expired Token")
			if err != nil {
				slog.Error("Failed to write error", "err", err)
			}
			return
		}

		next(ctx)
	}
}
