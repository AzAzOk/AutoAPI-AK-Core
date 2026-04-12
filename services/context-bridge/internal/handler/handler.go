package handler

import "net/http"

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/readyz", h.Readyz)
	mux.HandleFunc("/v1/sessions", h.CreateSession)
	mux.HandleFunc("/v1/sessions/", h.SessionByID) // GET + DELETE
	mux.HandleFunc("/v1/messages", h.SaveMessage)
	mux.HandleFunc("/v1/context", h.GetContext)
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func (h *Handler) Readyz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ready"}`))
}

func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request)  { w.WriteHeader(http.StatusNotImplemented) }
func (h *Handler) SessionByID(w http.ResponseWriter, r *http.Request)    { w.WriteHeader(http.StatusNotImplemented) }
func (h *Handler) SaveMessage(w http.ResponseWriter, r *http.Request)    { w.WriteHeader(http.StatusNotImplemented) }
func (h *Handler) GetContext(w http.ResponseWriter, r *http.Request)     { w.WriteHeader(http.StatusNotImplemented) }
