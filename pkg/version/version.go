package version

import (
	"net/http"

	"github.com/go-chi/render"
)

// Version is a response indicating success or failure
type Version struct {
	Version string `json:"version"`
}

// New instantiates new Version
func New(version string) *Version {
	return &Version{Version: version}
}

// Handler returns the version response
func (v *Version) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, v)
	})
}

// Render renders the JSON
func (v *Version) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
