package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"orbit/internal/process"
)

type Server struct {
	Registry *Registry
	Manager  *process.Manager
}

func NewServer(reg *Registry, mgr *process.Manager) *Server {
	return &Server{
		Registry: reg,
		Manager:  mgr,
	}
}

func (s *Server) Start(port string) error {
	mux := http.NewServeMux()

	// Endpoint: orbit ps
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(s.Registry.GetAll())
	})

	// Endpoint: orbit run
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			http.Error(w, "Missing agent name", http.StatusBadRequest)
			return
		}

		workspacePath := "./agents/" + name
		if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
			workspacePath = "./" + name
		}

		err := s.Manager.SpawnAgent(name, workspacePath)

		if err != nil {
			http.Error(w, "Failed to start agent: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Agent %s is spinning up....", name)
	})

	// Endpoint: orbit stop
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			http.Error(w, "Missing agent name", http.StatusBadRequest)
			return
		}

		meta, ok := s.Registry.Get(name)
		if !ok {
			http.Error(w, fmt.Sprintf("Agent %s is not running", name), http.StatusNotFound)
			return
		}

		err := s.Manager.StopAgent(meta.PID)
		if err != nil {
			http.Error(w, "Failed to stop agent: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Agent %s stopped successfully", name)
	})

	return http.ListenAndServe(":"+port, mux)

}
