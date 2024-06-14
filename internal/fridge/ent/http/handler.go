// Code generated by ent, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/masseelch/elk/internal/fridge/ent"
	"go.uber.org/zap"
)

// NewHandler returns a ready to use handler with all generated endpoints mounted.
func NewHandler(c *ent.Client, l *zap.Logger) chi.Router {
	r := chi.NewRouter()
	MountRoutes(c, l, r)
	return r
}

// MountRoutes mounts all generated routes on the given router.
func MountRoutes(c *ent.Client, l *zap.Logger, r chi.Router) {
	NewCompartmentHandler(c, l).MountRoutes(r)
	NewFridgeHandler(c, l).MountRoutes(r)
	NewItemHandler(c, l).MountRoutes(r)
}

// CompartmentHandler handles http crud operations on ent.Compartment.
type CompartmentHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewCompartmentHandler(c *ent.Client, l *zap.Logger) *CompartmentHandler {
	return &CompartmentHandler{
		client: c,
		log:    l.With(zap.String("handler", "CompartmentHandler")),
	}
}
func (h *CompartmentHandler) MountCreateRoute(r chi.Router) *CompartmentHandler {
	r.Post("/compartments", h.Create)
	return h
}
func (h *CompartmentHandler) MountReadRoute(r chi.Router) *CompartmentHandler {
	r.Get("/compartments/{id}", h.Read)
	return h
}
func (h *CompartmentHandler) MountUpdateRoute(r chi.Router) *CompartmentHandler {
	r.Patch("/compartments/{id}", h.Update)
	return h
}
func (h *CompartmentHandler) MountDeleteRoute(r chi.Router) *CompartmentHandler {
	r.Delete("/compartments/{id}", h.Delete)
	return h
}
func (h *CompartmentHandler) MountListRoute(r chi.Router) *CompartmentHandler {
	r.Get("/compartments", h.List)
	return h
}
func (h *CompartmentHandler) MountFridgeRoute(r chi.Router) *CompartmentHandler {
	r.Get("/compartments/{id}/fridge", h.Fridge)
	return h
}
func (h *CompartmentHandler) MountContentsRoute(r chi.Router) *CompartmentHandler {
	r.Get("/compartments/{id}/contents", h.Contents)
	return h
}
func (h *CompartmentHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountFridgeRoute(r).MountContentsRoute(r)
}

// FridgeHandler handles http crud operations on ent.Fridge.
type FridgeHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewFridgeHandler(c *ent.Client, l *zap.Logger) *FridgeHandler {
	return &FridgeHandler{
		client: c,
		log:    l.With(zap.String("handler", "FridgeHandler")),
	}
}
func (h *FridgeHandler) MountCreateRoute(r chi.Router) *FridgeHandler {
	r.Post("/fridges", h.Create)
	return h
}
func (h *FridgeHandler) MountReadRoute(r chi.Router) *FridgeHandler {
	r.Get("/fridges/{id}", h.Read)
	return h
}
func (h *FridgeHandler) MountUpdateRoute(r chi.Router) *FridgeHandler {
	r.Patch("/fridges/{id}", h.Update)
	return h
}
func (h *FridgeHandler) MountListRoute(r chi.Router) *FridgeHandler {
	r.Get("/fridges", h.List)
	return h
}
func (h *FridgeHandler) MountCompartmentsRoute(r chi.Router) *FridgeHandler {
	r.Get("/fridges/{id}/compartments", h.Compartments)
	return h
}
func (h *FridgeHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountListRoute(r).MountCompartmentsRoute(r)
}

// ItemHandler handles http crud operations on ent.Item.
type ItemHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewItemHandler(c *ent.Client, l *zap.Logger) *ItemHandler {
	return &ItemHandler{
		client: c,
		log:    l.With(zap.String("handler", "ItemHandler")),
	}
}
func (h *ItemHandler) MountCreateRoute(r chi.Router) *ItemHandler {
	r.Post("/items", h.Create)
	return h
}
func (h *ItemHandler) MountReadRoute(r chi.Router) *ItemHandler {
	r.Get("/items/{id}", h.Read)
	return h
}
func (h *ItemHandler) MountUpdateRoute(r chi.Router) *ItemHandler {
	r.Patch("/items/{id}", h.Update)
	return h
}
func (h *ItemHandler) MountDeleteRoute(r chi.Router) *ItemHandler {
	r.Delete("/items/{id}", h.Delete)
	return h
}
func (h *ItemHandler) MountListRoute(r chi.Router) *ItemHandler {
	r.Get("/items", h.List)
	return h
}
func (h *ItemHandler) MountCompartmentRoute(r chi.Router) *ItemHandler {
	r.Get("/items/{id}/compartment", h.Compartment)
	return h
}
func (h *ItemHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountCompartmentRoute(r)
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
