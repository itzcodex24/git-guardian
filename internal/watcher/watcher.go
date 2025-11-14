package watcher

import (
    "fmt"
    "path/filepath"
    "time"

    "github.com/fsnotify/fsnotify"
)

func WatchAndDebounce(folder string, debounceStr string, fn func()) error {
    debounce, err := time.ParseDuration(debounceStr)
    if err != nil || debounce <= 0 {
        debounce = 30 * time.Second
    }

    w, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    defer w.Close()

    if err := w.Add(folder); err != nil {
        return err
    }

    entries, _ := filepath.Glob(filepath.Join(folder, "*"))
    for _, e := range entries {
        _ = w.Add(e)
    }

    var timer *time.Timer
    events := w.Events
    errs := w.Errors

    for {
        select {
        case ev := <-events:
            fmt.Println("[watcher] event:", ev)
            if timer != nil {
                timer.Stop()
            }
            timer = time.AfterFunc(debounce, fn)
        case er := <-errs:
            if er != nil {
                fmt.Println("[watcher] error:", er)
            }
        }
    }
}
