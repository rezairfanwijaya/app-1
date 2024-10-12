package handler

import (
	"log"
	"net/http"

	"github.com/rezairfanwijaya/app-1.git/model"
	"github.com/rezairfanwijaya/app-1.git/response"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp := response.Failed{
			Message: "method not allowed",
		}

		respJson, err := resp.ToJSON()
		if err != nil {
			log.Println("failed to marshal response")
			return
		}

		if _, err := w.Write(respJson); err != nil {
			log.Println("failed to write response")
			return
		}
		return
	}

	resp := response.Success{
		Data: model.Users,
	}

	respJson, err := resp.ToJSON()
	if err != nil {
		log.Println("failed to marshal response")
		return
	}

	if _, err := w.Write(respJson); err != nil {
		log.Println("failed to write response")
		return
	}
}
