package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type apiHandler struct{}

func (a *apiHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetAll", "request", r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func (a *apiHandler) ListClusters(w http.ResponseWriter, r *http.Request) {
	slog.Info("ListClusters", "request", r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`
{
	"items": [
	  {
		  "name": "existing-cluster",
			"status": {
			  "aggregated": "STATE_HEALTHY",
				"creationTime": "2024-02-15T11:06:29Z",
				"error": null
			}
		}
	]
}
`))
}

func (a *apiHandler) CreateCluster(w http.ResponseWriter, r *http.Request) {
	slog.Info("CreateCluster", "request", r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`
{
  "name": "new-cluster",
	"status": {
	  "aggregated": "STATE_CREATING",
		"creationTime": "2024-06-09T11:06:29Z",
		"error": null
	}
}
`))
}

func (a *apiHandler) GetCluster(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetCluster", "request", r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`
{
  "name": "new-cluster",
	"status": {
	  "aggregated": "STATE_HEALTHY",
		"creationTime": "2024-06-09T11:06:29Z",
		"error": null
	}
}
`))
}

func main() {
	h := &apiHandler{}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1alpha1/projects/{project}/clusters", h.ListClusters)
	mux.HandleFunc("PUT /v1alpha1/projects/{project}/clusters/{name}", h.CreateCluster)
	mux.HandleFunc("GET /v1alpha1/projects/{project}/clusters/{name}", h.GetCluster)
	mux.HandleFunc("/v1alpha1/", h.GetAll)

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			slog.Info("GET /*", "request", req)
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to API home page!")
	})

	port := os.Getenv("PORT")
	if len(os.Getenv("PORT")) < 1 {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
