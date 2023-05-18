package people_api

import (
	"charlotte_backend/config"
	"charlotte_backend/entities"
	"charlotte_backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		peopleModel := models.PeopleModel{
			Db: db,
		}
		peoples, err2 := peopleModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, peoples)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		peopleModel := models.PeopleModel{
			Db: db,
		}
		peoples, err2 := peopleModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, peoples)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var people entities.People
	err := json.NewDecoder(request.Body).Decode(&people)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		peopleModel := models.PeopleModel{
			Db: db,
		}
		err2 := peopleModel.Create(&people)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, people)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var people entities.People
	err := json.NewDecoder(request.Body).Decode(&people)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		peopleModel := models.PeopleModel{
			Db: db,
		}
		_, err2 := peopleModel.Update(&people)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, people)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		peoplesModel := models.PeopleModel{
			Db: db,
		}
		RowsAffected, err2 := peoplesModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, map[string]int64{
				"RowsAffected": RowsAffected,
			})
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
