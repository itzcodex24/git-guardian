package watcher

import (
    "fmt"
    "path/filepath"
    "time"

    "github.com/fsnotify/fsnotify"
)

// WatchAndDebounce watches a folder (non-recursive) and debounces events.
// If you need recursive watching, consider walking the tree and adding subdirs.
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

    // add the folder itself
    if err := w.Add(folder); err != nil {
        return err
    }

    // optional: also watch immediate subdirectories (common case)
    // NOTE: large projects with many subdirs may need a recursive strategy.
    entries, _ := filepath.Glob(filepath.Join(folder, "*"))
    for _, e := range entries {
        // watch only directories
        // ignore errors — best-effort
        // if fi, err := os.Stat(e); err == nil && fi.IsDir() { w.Add(e) }
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
