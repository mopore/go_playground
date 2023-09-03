package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MutexScoreBoardManager struct {
    l sync.RWMutex
    scoreboard map[string]int
}

func newMutexScoreBoardManager() *MutexScoreBoardManager {
    return &MutexScoreBoardManager{
        scoreboard: make(map[string]int),
    }
}

func (m *MutexScoreBoardManager) UpdateScore(name string, score int) {
    m.l.Lock()
    defer m.l.Unlock()
    m.scoreboard[name] = score
}

func (m *MutexScoreBoardManager) GetScore(name string) (int, bool) {
    m.l.RLock()
    defer m.l.RUnlock()
    val, ok := m.scoreboard[name]
    return val, ok
}

func main() {
    wg := sync.WaitGroup{}
    wg.Add(2)
    manager := newMutexScoreBoardManager()
    go func() {
        for i := 0; i < 10; i++ {
            score := rand.Intn(100)
            name := fmt.Sprintf("Player%d", i)
            manager.UpdateScore(name, score)
            time.Sleep(time.Millisecond * 500)
        }
        wg.Done()
    }()
    go func() {
        for i := 0; i <= 10; i++ {
            fmt.Println("Scoreboard")
            for j := 0; j <= 10; j++ {
                name := fmt.Sprintf("Player%d", j)
                score, ok := manager.GetScore(name)
                if ok {
                    fmt.Printf("%s: %d\n", name, score)
                }
            }
            fmt.Println()
            time.Sleep(time.Millisecond * 550)
        }
        wg.Done()
    }()
    wg.Wait()
}
