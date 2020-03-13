package journal

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi"
	"grig/internal/common"
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.RespondError(w, errors.New("id must be integer"), http.StatusBadRequest)
		return
	}
	journal, err := h.repository.GetStudentByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		common.RespondError(w, errors.New("journal row with this id not found"), http.StatusNotFound)
		return
	}
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondJSON(w, http.StatusOK, journal)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var journals []model.Journal
	var err error = nil
	_, groupExists := common.Param(r, "group_id")
	_, studentExists := common.Param(r, "student_id")
	if studentExists {
		studentID, ok := common.ParamInt(w, r, "student_id")
		if !ok {
			return
		}
		journals, err = h.repository.GetJournalByStudent(studentID)
		if err != nil {
			common.RespondInternal(w, err, http.StatusInternalServerError)
			return
		}
	} else if groupExists {
		groupID, ok := common.ParamInt(w, r, "group_id")
		if !ok {
			return
		}
		journals, err = h.repository.GetJournalByStudyGroup(groupID)
		if err != nil {
			common.RespondInternal(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		journals, err = h.repository.GetJournals()
		if err != nil {
			common.RespondInternal(w, err, http.StatusInternalServerError)
			return
		}
	}
	common.RespondJSON(w, http.StatusOK, journals)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	var journal model.Journal
	if !common.ParseBody(w, r, &journal) {
		return
	}
	err := h.repository.UpsertJournal(journal)
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondOK(w)
}
