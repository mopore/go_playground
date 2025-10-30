package main

import (
	"io"
	"net/http"
	"runtime"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define what the worker sends back
type FetchResult struct {
	URL   string
	Body  string
	Error error
}

func main() {
	// Important for raylib (esp. on macOS)
	runtime.LockOSThread()

	rl.InitWindow(800, 450, "Async Event Example (press F to fetch)")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Channel for communicating async results
	results := make(chan FetchResult, 4)

	statusMsg := "Press F to fetch from example.com"
	lastFetch := time.Time{}

	for !rl.WindowShouldClose() {
		// Handle input (raylib input API must stay on main thread)
		if rl.IsKeyPressed(rl.KeyF) {
			if time.Since(lastFetch) > time.Second {
				lastFetch = time.Now()
				statusMsg = "Fetching..."
				go fetchURL("https://example.com", results)
			}
		}

		// Non-blocking read from results channel
		select {
		case r := <-results:
			if r.Error != nil {
				statusMsg = "Fetch failed: " + r.Error.Error()
			} else {
				statusMsg = "Fetched " + r.URL + " (" + shortBody(r.Body) + ")"
			}
		default:
			// no new result — do nothing
		}

		// --- RENDER ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(statusMsg, 40, 200, 20, rl.DarkGray)
		rl.DrawText("Press 'F' to start fetch", 40, 240, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

// Runs in a background goroutine
func fetchURL(url string, ch chan<- FetchResult) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- FetchResult{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- FetchResult{URL: url, Error: err}
		return
	}

	ch <- FetchResult{
		URL:  url,
		Body: string(body),
	}
}

// Helper to limit long response
func shortBody(s string) string {
	if len(s) > 50 {
		return s[:50] + "..."
	}
	return s
}

