package supervisor

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "sync"
    "time"

    "github.com/itzcodex24/git-guardian/internal/git"
    "github.com/itzcodex24/git-guardian/internal/state"
    "github.com/itzcodex24/git-guardian/internal/watcher"
)

var (
    running = map[string]chan struct{}{}
    mu      sync.Mutex
)

func StartAll() {
    list, _ := state.Load()
    for _, w := range list {
        if !w.Paused && (w.Mode == "watch" || w.Mode == "interval") {
            startWatcher(w)
        }
    }
}


func StartAllBlocking() {
    StartAll()
    select {}
}

func startWatcher(st state.WatcherState) {
    mu.Lock()
    if _, ok := running[st.ID]; ok {
        mu.Unlock()
        return 
    }
    stop := make(chan struct{})
    running[st.ID] = stop
    mu.Unlock()

    go func(s state.WatcherState, stopCh chan struct{}) {
        defer func() {
            mu.Lock()
            delete(running, s.ID)
            mu.Unlock()
        }()

        for {
            if _, err := os.Stat(s.Folder); os.IsNotExist(err) {
                fmt.Println("[supervisor] folder removed, cancelling:", s.Folder)
                Remove(s.ID)
                return
            }

            switch s.Mode {
            case "interval":
                d, err := time.ParseDuration(s.Interval)
                if err != nil || d <= 0 {
                    d = 5 * time.Minute
                }
                if err := git.CommitAndPush(s.Folder); err == nil {
                    updateLastRun(s.ID)
                } else {
                    fmt.Println("[supervisor] interval commit error:", err)
                }
                select {
                case <-time.After(d):
                    
                case <-stopCh:
                    return
                }

            case "watch":
                done := make(chan struct{})
                go func() {
                    _ = watcher.WatchAndDebounce(s.Folder, s.Debounce, func() {
                        if err := git.CommitAndPush(s.Folder); err == nil {
                            updateLastRun(s.ID)
                        } else {
                            fmt.Println("[supervisor] watch commit error:", err)
                        }
                    })
                    close(done)
                }()

                select {
                case <-stopCh:
                    return
                case <-done:
                }
            default:
                return
            }
        }
    }(st, stop)
}

func Pause(id string) error {
    mu.Lock()
    if ch, ok := running[id]; ok {
        close(ch)
        delete(running, id)
    }
    mu.Unlock()

    list := state.Get()
    for i := range list {
        if list[i].ID == id {
            list[i].Paused = true
        }
    }
    state.Update(list)
    return nil
}

func Resume(id string) error {
    list := state.Get()
    for i := range list {
        if list[i].ID == id {
            if !list[i].Paused {
                return fmt.Errorf("watcher already running: %s", id)
            }
            list[i].Paused = false
            state.Update(list)
            startWatcher(list[i])
            return nil
        }
    }
    return fmt.Errorf("watcher not found: %s", id)
}

func Remove(id string) error {
    mu.Lock()
    if ch, ok := running[id]; ok {
        close(ch)
        delete(running, id)
    }
    mu.Unlock()

    list := state.Get()
    newList := []state.WatcherState{}
    for _, w := range list {
        if w.ID != id {
            newList = append(newList, w)
        }
    }
    state.Update(newList)
    return nil
}

func List() []state.WatcherState {
    return state.Get()
}

func updateLastRun(id string) {
    list := state.Get()
    for i := range list {
        if list[i].ID == id {
            list[i].LastRun = time.Now().Format(time.RFC3339)
            break
        }
    }
    state.Update(list)
}

func GenerateLaunchdPlist() string {
    bin, _ := exec.LookPath("guardian")
    if bin == "" {
        bin = "/usr/local/bin/guardian"
    }
    home, _ := os.UserHomeDir()
    log := filepath.Join(home, "Library", "Logs", "gitguardian.log")

    return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.gitguardian.agent</string>

    <key>ProgramArguments</key>
    <array>
        <string>%s</string>
        <string>daemon</string>
    </array>

    <key>RunAtLoad</key><true/>
    <key>KeepAlive</key><true/>

    <key>StandardOutPath</key><string>%s</string>
    <key>StandardErrorPath</key><string>%s</string>
</dict>
</plist>`, bin, log, log)
}
