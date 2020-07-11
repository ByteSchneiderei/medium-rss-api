package error

import (
	"net/http"

	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// ErrorResponse is an error response
type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render renders the JSON
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// RenderError renders an error as JSON and logs it
func RenderError(err error, message string, status int) render.Renderer {
	log.Error(err)
	errorText := err.Error()
	statusText := message
	if status == http.StatusInternalServerError {
		errorText = http.StatusText(http.StatusInternalServerError)
		statusText = http.StatusText(http.StatusInternalServerError)
	}
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     statusText,
		ErrorText:      errorText,
	}
}
