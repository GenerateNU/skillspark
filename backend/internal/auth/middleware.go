package auth

import (
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
	"skillspark/internal/config"
	"log/slog"
)

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

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/v1/user", cfg.URL), nil)
		if err != nil {
			err := huma.WriteErr(api, ctx, http.StatusInternalServerError, "Failed to create request")
			if err != nil {
				slog.Error("Failed to write error", "err", err)
			}
			return
		}

		req.Header.Set("Authorization", "Bearer "+cookie.Value)
		req.Header.Set("apikey", cfg.ServiceRoleKey)

		res, err := Client.Do(req)
		if err != nil {
			err := huma.WriteErr(api, ctx, http.StatusInternalServerError, "Failed to validate token")
			if err != nil {
				slog.Error("Failed to write error", "err", err)
			}
			return
		}
		err = res.Body.Close()
		if err != nil {
			slog.Error("Error closing response body: ", "err", err)
			return
		}

		if res.StatusCode != http.StatusOK {
			err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid/Expired Token")
			if err != nil {
				slog.Error("Failed to write error", "err", err)
			}
			return
		}

		next(ctx)
	}
}
