package monitoring

import (
	"log"
	"net/http"
	"serverStateService/entities"
	"serverStateService/models"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func RunPings() {
	servers, err := models.GetServers()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, server := range servers {
		pingResultChan := make(chan entities.PingResult)

		go runPing(&server, pingResultChan)

		pingResult := <-pingResultChan

		err = models.InsertResult(&pingResult)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func runPing(server *entities.Server, pingResult chan entities.PingResult) {
	response, err := http.Get(server.Ip)
	now := time.Now()
	if err != nil {
		log.Println(err)
		pingResult <- entities.PingResult{
			IdServer:  server.Id,
			Respond:   false,
			ScannedAt: &now,
		}
		return
	}

	respond := response.StatusCode == http.StatusOK

	pingResult <- entities.PingResult{
		IdServer:  server.Id,
		Respond:   respond,
		ScannedAt: &now,
	}
}
