package auth

import (
    "net/http"
    "fmt"
    "github.com/danielgtaylor/huma/v2"
	"skillspark/internal/config"
)

func AuthMiddleware(api huma.API, cfg *config.Supabase) func(ctx huma.Context, next func(huma.Context)) {
    skipPaths := map[string]bool{
        "/auth/login":    true,
        "/auth/register": true,
        "/health":        true,
    }

    return func(ctx huma.Context, next func(huma.Context)) {
        if skipPaths[ctx.Operation().Path] {
            next(ctx)
            return
        }

        cookie, err := huma.ReadCookie(ctx, "jwt")
        if err != nil || cookie.Value == "" {
            huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token Not Found")
            return
        }

        req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/v1/user", cfg.URL), nil)
        if err != nil {
            huma.WriteErr(api, ctx, http.StatusInternalServerError, "Failed to create request")
            return
        }

        req.Header.Set("Authorization", "Bearer "+cookie.Value)
        req.Header.Set("apikey", cfg.ServiceRoleKey)

        res, err := Client.Do(req)
        if err != nil {
            huma.WriteErr(api, ctx, http.StatusInternalServerError, "Failed to validate token")
            return
        }
        defer res.Body.Close()

        if res.StatusCode != http.StatusOK {
            huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid/Expired Token")
            return
        }

        next(ctx)
    }
}
