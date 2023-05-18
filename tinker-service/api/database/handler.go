package database

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	// "github.com/docker/docker/integration-cli/cli"
	"github.com/fatih/structs"
	"github.com/saint-rivers/tinker/app/container"
	request "github.com/saint-rivers/tinker/common"
)

func CreateService() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req request.DatabaseServiceRequest
		json.NewDecoder(r.Body).Decode(&req)

		// row := InsertDatabaseRequest(db, &req)

		// if row.Err() != nil {
		// 	var errorResponse map[string]string = errorCreateDatabaseService
		// 	json.NewEncoder(w).Encode(errorResponse)
		// 	log.Println(row.Err().Error())
		// 	return
		// }

		c := container.CreateDatabaseContainer(&req)
		err := container.StartDatabaseContainer(c.ID)
		if err != nil {
			log.Panic("unable to start container")
		}

		response := structs.Map(req)
		response["containerId"] = c.ID
		json.NewEncoder(w).Encode(response)
	}
}

func getContainerQuery(q *url.Values) (string, string) {
	status := q.Get("status")
	cid := q.Get("container")
	if status == "" {
		log.Panic("no status provided")
	}
	if cid == "" {
		log.Panic("container id not provided")
	}

	return status, cid
}

func ListContainers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containers, err := container.ListContainers()
		if err != nil {
			log.Fatal("unable to list containers")
		}
		json.NewEncoder(w).Encode(containers)
	}
}

func ManageService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		status, cid := getContainerQuery(&q)

		var err error
		var message string

		switch status {
		case "up":
			err = container.StartDatabaseContainer(cid)
			if err != nil {
				message = "unable to start container"
				json.NewEncoder(w).Encode(ContainerError(message))
				return
			}
			success := ContainerSuccess(cid, status)
			json.NewEncoder(w).Encode(success)

		case "down":
			err = container.StopDatabaseContainer(cid)
			if err != nil {
				message = "unable to stop container"
				json.NewEncoder(w).Encode(ContainerError(message))
				return
			}
			success := ContainerSuccess(cid, status)
			json.NewEncoder(w).Encode(success)
		}

	}
}
