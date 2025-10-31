package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	baselLat = 47.5596
	baselLon = 7.5886
	meteoUrl = "https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&current_weather=true&temperature_unit=celsius"
)

type BackgroundResult struct {
	Temp  string
	Error error
}

type MeteoResponse struct {
	Current struct {
		Time string `json:"time"`
		Temp float64 `json:"temperature"`
	} `json:"current_weather"`
}

func main() {
	// Important for raylib (esp. on macOS)
	runtime.LockOSThread()

	rl.InitWindow(800, 450, "Async Event Example (press F to fetch)")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	bgResults := make(chan BackgroundResult, 4)

	statusMsg := "Press F to fetch current temperature in basel"
	lastFetch := time.Time{}

	for !rl.WindowShouldClose() {
		// Handle input (raylib input API must stay on main thread)
		if rl.IsKeyPressed(rl.KeyF) {
			if time.Since(lastFetch) > time.Second {
				lastFetch = time.Now()
				statusMsg = "Fetching..."
				go fetchBaselTemp(bgResults)
			}
		}

		// Non-blocking read from results channel
		select {
		case r := <-bgResults:
			if r.Error != nil {
				statusMsg = "Fetch failed: " + r.Error.Error()
			} else {
				statusMsg = "Fetched: " + r.Temp + " (current temperature in Basel)"
			}
		default:
			// no new result — do nothing
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(statusMsg, 40, 200, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}

func fetchBaselTemp(ch chan<- BackgroundResult) {
	client := http.Client{Timeout: 5 * time.Second}

	url := fmt.Sprintf(meteoUrl, baselLat, baselLon)

	resp, err := client.Get(url)
	if err != nil {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- BackgroundResult{Temp: "<fetch error>", Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- BackgroundResult{Temp: "<fetch error>", Error: err}
		return
	}

	var mr MeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
		log.Println("fetchBaselTemp: error fetching:", err)
		ch <- BackgroundResult{Temp: "<fetch error>", Error: err}
		return
	}

	resText := fmt.Sprintf("%.1f C", mr.Current.Temp)

	log.Println("fetchBaselTemp: fetched result:", resText)

	ch <- BackgroundResult{
		Temp: resText,
	}
}


