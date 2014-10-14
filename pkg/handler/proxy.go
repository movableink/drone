package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/drone/drone/pkg/database"
)

type PRHandler struct {
	subdomain string
}

func NewPRHandler(subdomain string) *PRHandler {
	return &PRHandler{
		subdomain: subdomain,
	}
}

func (h *PRHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	build, err := database.GetRunningBuildByBranch(h.subdomain)
	if err != nil {
		return RenderText(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	println("Redirecting to port", build.Port)
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host: "localhost:" + build.Port,
	})

	proxy.ServeHTTP(w, r)
	return nil
}
