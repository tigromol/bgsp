package bus

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi"
	"grig/internal/common"
	"grig/internal/errors"
	"grig/internal/model"
	"net/http"
	"strconv"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.RespondError(w, errors.New("id must be integer"), http.StatusBadRequest)
		return
	}

	var student model.Student
	if !common.ParseBody(w, r, &student) {
		return
	}
	data := fmt.Sprintf(`{
		"from": "agb",
		"to": "dean",
		"subject": "ADD_ROW",
		"data": "{\"isBinariesChanged\":false,\"entityName\":\"student\",\"plainData\":{\"id\":%d,\"surname\":\"%s\",\"name\":\"%s\",\"secondName\":\"%s\",\"studyGroupId\":%d},\"binaryLinks\":{}}"
	}`, id, student.Surname, student.Name, student.SecondName, student.StudyGroup.ID)
	_, err = http.Post("http://up-lab1.mirea.ru/bus", "application/json", bytes.NewBufferString(data))
	if err != nil {
		common.RespondError(w, err, http.StatusBadRequest)
		return
	}
	common.RespondOK(w)
}
