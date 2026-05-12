package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Endpoint: forge ps
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(s.Registry.GetAll())
	})

	// Endpoint: forge run
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			http.Error(w, "Missing agent name", http.StatusBadRequest)
			return
		}

		workspacePath := "./agents/" + name

		err := s.Manager.SpawnAgent(name, workspacePath)

		if err != nil {
			http.Error(w, "Failed to start agent: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Agent %s is spinning up....", name)
	})

	return http.ListenAndServe(":"+port, mux)

}
