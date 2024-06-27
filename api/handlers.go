package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeze322/wt-guided-weaponry/storage"
)

func (s *Server) handleListCategory(w http.ResponseWriter, r *http.Request) error {
	categories, err := s.mysql.ListCategory(r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, categories)
}

func (s *Server) handleWeaponParams(w http.ResponseWriter, r *http.Request) error {
	category := chi.URLParam(r, "category")
	params, err := s.mongo.WeaponParams(r.Context(), category)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, params)
}

func (s *Server) handleInsertWeapon(w http.ResponseWriter, r *http.Request) error {
	req := new(storage.Params)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.mongo.InsertWeapon(r.Context(), req); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, nil)
}

func (s *Server) handleUpdateWeapon(w http.ResponseWriter, r *http.Request) error {
	req := new(storage.Params)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.mongo.UpdateWeapon(r.Context(), req); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, nil)
}
