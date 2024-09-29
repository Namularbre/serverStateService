package monitoring

import (
	"log"
	"net/http"
	"serverStateService/entities"
	"serverStateService/models"
	"time"
)

func RunPings() {
	serverModel := models.ServerModel{}
	pingModel := models.PingResultModel{}

	servers, err := serverModel.GetServers()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, server := range servers {
		pingResultChan := make(chan entities.PingResult)

		go runPing(&server, pingResultChan)

		pingResult := <-pingResultChan

		err = pingModel.InsertResult(&pingResult)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func runPing(server *entities.Server, pingResult chan entities.PingResult) {
	response, err := http.Get(server.Ip)
	if err != nil {
		log.Fatalln(err)
		return
	}

	respond := response.StatusCode == http.StatusOK
	now := time.Now()

	pingResult <- entities.PingResult{
		IdServer:  server.Id,
		Respond:   respond,
		ScannedAt: &now,
	}
}
