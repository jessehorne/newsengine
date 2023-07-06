package util

import (
	"encoding/json"
	"github.com/jessehorne/newsengine/types"
	"net/http"
)

func APIResponse(w http.ResponseWriter, code int, res types.APIResponse) {
	w.WriteHeader(code)

	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func APIResponseError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)

	ret := types.APIResponse{
		"msg": msg,
	}

	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
