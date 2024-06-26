// Code generated by ent, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Read fetches the ent.Compartment identified by a given url-parameter from the
// database and returns it to the client.
func (h *CompartmentHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Compartment.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching compartments from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("compartments rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment1151786357Views(es), w)
}

// Read fetches the ent.Fridge identified by a given url-parameter from the
// database and returns it to the client.
func (h *FridgeHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Fridge.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching fridges from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("fridges rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2716213877Views(es), w)
}

// Read fetches the ent.Item identified by a given url-parameter from the
// database and returns it to the client.
func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Item.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching items from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("items rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewItem1509516544Views(es), w)
}
