package api

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

var openAPISpec = H{
	"openapi": "3.0.3",
	"info": H{
		"title":       "Golang Starter Kit API",
		"version":     "1.0.0",
		"description": "Simple API running on Vercel with g4vercel.",
	},
	"servers": []H{
		{"url": "/"},
	},
	"paths": H{
		"/": H{
			"get": H{
				"summary":     "Health check endpoint",
				"operationId": "getRootMessage",
				"responses": H{
					"200": H{
						"description": "Success",
						"content": H{
							"application/json": H{
								"schema": H{
									"type": "object",
									"properties": H{
										"message": H{"type": "string", "example": "hello"},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

const swaggerHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Swagger Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    html, body { margin: 0; padding: 0; }
    #swagger-ui { max-width: 1100px; margin: 0 auto; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: '/swagger.json',
      dom_id: '#swagger-ui',
      deepLinking: true,
      presets: [SwaggerUIBundle.presets.apis],
      layout: 'BaseLayout'
    });
  </script>
</body>
</html>`

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"message": "hello",
		})
	})

	server.GET("/swagger.json", func(c *Context) {
		c.JSON(http.StatusOK, openAPISpec)
	})

	server.GET("/swagger", func(c *Context) {
		c.SetHeader("Content-Type", "text/html; charset=utf-8")
		c.Data(http.StatusOK, []byte(swaggerHTML))
	})

	server.Handle(w, r)
}
