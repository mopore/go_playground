package actor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	hintText = "Press 'f' to fetch"
	processingText = "Fetching..."
	baselLat = 47.5596
	baselLon = 7.5886
	meteoUrl = "https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current_weather=true&temperature_unit=celsius"
)

type backgroundResult struct {
	Temp  string
	Error error
}

type meteoResponse struct {
	Current struct {
		Time string `json:"time"`
		Temp float64 `json:"temperature"`
	} `json:"current_weather"`
}

type FetchActor struct {
	text string
	bgResults chan backgroundResult
}

func NewFetchActor() *FetchActor {
	a := &FetchActor {
		text: hintText,
		bgResults: make(chan backgroundResult, 4),
	}
	return a
}

func (a *FetchActor) ReadInput() {
	if rl.IsKeyPressed(rl.KeyF) {
		a.text = processingText
		go fetchBaselTemp(a.bgResults)
	}
}

func (a *FetchActor) UpdateState() []ActorRequest {
	select {
	case  r := <-a.bgResults:
		if r.Error != nil {
			a.text = "Fetch failed: " + r.Error.Error()
		} else {
			a.text = "Fetched: " + r.Temp + " (current temparature ins Basel)"
		}
	default:
	//let it slip... no new result
	}
	return nil
}

func (a *FetchActor) Init(w int32, h int32) {
}

func (a *FetchActor) Render(w int32, h int32) {
	rl.DrawText(a.text, 40, 200, 20, rl.DarkBlue)
}

func fetchBaselTemp(ch chan<- backgroundResult) {
	client := http.Client{Timeout: 5 * time.Second}

	url := fmt.Sprintf(meteoUrl, baselLat, baselLon)

	resp, err := client.Get(url)
	if err != nil {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- backgroundResult{Temp: "<fetch error>", Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- backgroundResult{Temp: "<fetch error>", Error: err}
		return
	}

	var mr meteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- backgroundResult{Temp: "<fetch error>", Error: err}
		return
	}

	resText := fmt.Sprintf("%.1f C", mr.Current.Temp)

	log.Println("fetchBaselTemp: fetched result:", resText)

	ch <- backgroundResult{
		Temp: resText,
	}
}
