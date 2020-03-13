package student

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
	var student model.Student
	if !common.ParseBody(w, r, &student) {
		return
	}
	err := h.repository.CreateStudent(student)
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
	student, err := h.repository.GetStudentByID(id)
	if errors.Is(err, errors.ErrNotFound) {
		common.RespondError(w, errors.New("student with this id not found"), http.StatusNotFound)
		return
	}
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondJSON(w, http.StatusOK, student)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var students []model.Student
	var err error = nil
	_, exists := common.Param(r, "group_id")
	if !exists {
		students, err = h.repository.GetStudents()
		if err != nil {
			common.RespondInternal(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		groupID, ok := common.ParamInt(w, r, "group_id")
		if !ok {
			return
		}
		students, err = h.repository.GetStudentsByGroup(groupID)
		if err != nil {
			common.RespondInternal(w, err, http.StatusInternalServerError)
			return
		}
	}
	common.RespondJSON(w, http.StatusOK, students)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	if !common.ParseBody(w, r, &student) {
		return
	}
	err := h.repository.UpsertStudent(student)
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
	err = h.repository.DeleteStudent(id)
	if err != nil {
		common.RespondInternal(w, err, http.StatusInternalServerError)
		return
	}
	common.RespondOK(w)
}
