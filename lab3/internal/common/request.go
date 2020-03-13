package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ParseBody(w http.ResponseWriter, r *http.Request, object interface{}) bool {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, errors.New("failed to read request body"), http.StatusInternalServerError)
		return false
	}
	err = json.Unmarshal(bytes, object)
	if err != nil {
		RespondError(w, errors.New("failed to unmarshal request"), http.StatusBadRequest)
		return false
	}
	return true
}

func ParamInt(w http.ResponseWriter, r *http.Request, name string) (int, bool) {
	param := r.URL.Query().Get(name)
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		RespondError(w, errors.New(fmt.Sprintf("param %s must be integer", name)), http.StatusBadRequest)
		return 0, false
	}
	return paramInt, true
}

func MustParam(r *http.Request, w http.ResponseWriter, name string) (string, bool) {
	param := r.URL.Query().Get(name)
	if param == "" {
		RespondError(w, errors.New(fmt.Sprintf("param %s must present in query", name)), http.StatusBadRequest)
		return param, false
	}
	return param, true
}

func Param(r *http.Request, name string) (string, bool) {
	param := r.URL.Query().Get(name)
	if param == "" {
		return param, false
	}
	return param, true
}
