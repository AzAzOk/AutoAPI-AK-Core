package handler

import "net/http"

// Handler держит зависимости HTTP-обработчиков
type Handler struct {
	// service RouterService — добавить после реализации сервиса
}

func New() *Handler {
	return &Handler{}
}

// RegisterRoutes регистрирует все маршруты
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/readyz", h.Readyz)
	mux.HandleFunc("/v1/route", h.Route)
	mux.HandleFunc("/v1/status", h.Status)
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func (h *Handler) Readyz(w http.ResponseWriter, r *http.Request) {
	// TODO: проверить соединение с БД и Redis
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ready"}`))
}

func (h *Handler) Route(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
	w.WriteHeader(http.StatusNotImplemented)
}
