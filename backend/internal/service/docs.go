package service

import "github.com/gofiber/fiber/v2"

// setupDocsRoutes configures the API documentation routes
func setupDocsRoutes(app *fiber.App, specPath string) {
	// Serve OpenAPI spec files
	app.Static("/api", specPath)

	// Scalar API Reference UI
	app.Get("/docs", scalarHandler())
}

// scalarHandler returns a Fiber handler that serves the Scalar API Reference UI
func scalarHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		html := `<!doctype html>
<html>
<head>
    <title>SkillSpark API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
    <script id="api-reference" data-url="/api/openapi.yaml"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>`
		c.Set("Content-Type", "text/html")
		return c.SendString(html)
	}
}
