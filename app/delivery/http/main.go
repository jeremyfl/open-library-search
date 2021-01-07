package http

import "github.com/labstack/echo/v4"

// GenerateResponse generating HTTP response
// with proper format
func GenerateResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	response := map[string]interface{}{
		"status_code": statusCode,
		"message":     message,
		"data":        data,
	}

	return c.JSON(statusCode, response)
}
