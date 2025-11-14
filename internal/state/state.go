package state

import (
    "encoding/json"
    "os"
    "path/filepath"
    "sync"
)

type WatcherState struct {
    ID       string `json:"id"`
    Folder   string `json:"folder"`
    Mode     string `json:"mode"` 
    Interval string `json:"interval,omitempty"`
    Debounce string `json:"debounce,omitempty"`
    Paused   bool   `json:"paused"`
    LastRun  string `json:"last_run,omitempty"`
}

var (
    mu        sync.Mutex
    watchers  []WatcherState
    statePath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "gitguardian", "watchers.json")
)

func loadFromDisk() {
    mu.Lock()
    defer mu.Unlock()
    if _, err := os.Stat(statePath); os.IsNotExist(err) {
        watchers = []WatcherState{}
        return
    }
    data, err := os.ReadFile(statePath)
    if err != nil {
        watchers = []WatcherState{}
        return
    }
    _ = json.Unmarshal(data, &watchers)
}

func saveToDisk() {
    mu.Lock()
    defer mu.Unlock()
    _ = os.MkdirAll(filepath.Dir(statePath), 0755)
    data, _ := json.MarshalIndent(watchers, "", "  ")
    _ = os.WriteFile(statePath, data, 0644)
}

func Load() ([]WatcherState, error) {
    loadFromDisk()
    return watchers, nil
}

func Get() []WatcherState {
    loadFromDisk()
    return watchers
}

func Update(list []WatcherState) {
    mu.Lock()
    watchers = list
    mu.Unlock()
    saveToDisk()
}

func Append(w WatcherState) {
    mu.Lock()
    watchers = append(watchers, w)
    mu.Unlock()
    saveToDisk()
}
