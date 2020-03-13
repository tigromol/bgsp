package group

import (
	"github.com/go-chi/chi"
	"grig/internal/common"
	"grig/internal/errors"
	"grig/internal/model"
	"grig/internal/repository"
	"net/http"
	"strconv"
)

type Handler struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Handler {
	return &Handler{repository:repository}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var group model.StudyGroup
	if !common.ParseBody(w, r, &group) {
		return
	}
	err := h.repository.CreateGroup(group)
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondOK(w)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.RespondError(w, errors.New("id must be integer"), http.StatusBadRequest)
		return
	}
	group, err := h.repository.GetGroupByID(id)
	if errors.Is(err, errors.ErrNotFound) {
		common.RespondError(w, errors.New("group with this id not found"), http.StatusNotFound)
		return
	}
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondJSON(w, http.StatusOK, group)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var groups []model.StudyGroup
	var err error = nil
	groups, err = h.repository.GetGroups()
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondJSON(w, http.StatusOK, groups)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	var group model.StudyGroup
	if !common.ParseBody(w, r, &group) {
		return
	}
	err := h.repository.UpsertGroup(group)
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondOK(w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.RespondError(w, errors.New("id must be integer"), http.StatusBadRequest)
		return
	}
	err = h.repository.DeleteGroup(id)
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondOK(w)
}
