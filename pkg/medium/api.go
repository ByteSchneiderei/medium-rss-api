package medium

import (
	"net/http"

	"github.com/go-chi/render"

	errPkg "github.com/ByteSchneiderei/medium-rss-api/pkg/error"
)

// Handler handles http request and return response
func (m *Medium) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := m.Fetch()
		if err != nil {
			render.Render(w, r, errPkg.RenderError(err, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))
			return
		}
		render.Render(w, r, response)
	})
}

// Render renders the JSON
func (*Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
