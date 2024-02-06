package systeminfo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StreamMetricWs(ch <-chan Metric, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for metric := range ch {
		data, err := json.Marshal(metric)

		if err != nil {
			fmt.Println("Error marshling", err)
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("Error writing msg", err)
			return
		}

		// move this to the channel
		time.Sleep(time.Second)
	}
}

func ServerDai() {
	metricChannel := make(chan Metric)
	defer close(metricChannel)

	// Start metric collector as a Goroutine
	// It will collect the metric and write them to the channel
	go ServeMetrics(metricChannel)

	http.Handle("/", http.FileServer(http.Dir("./systeminfo/templates")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		StreamMetricWs(metricChannel, w, r)
	})

	fmt.Println("Websocket started on 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
	}
}
