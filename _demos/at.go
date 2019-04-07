package main

import (
  "fmt"
  "math/rand"
  "os"
  "time"
  
  "github.com/gdamore/tcell"
  "github.com/gdamore/tcell"
)

func main() {
  tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
  s, e := tcell.NewScreen()
  if e != nil {
    fmt.Fprintf(os.Stderr, "%v\n", e)
    os.Exit(1)
  }
  if e = s.Init(); e != nil {
    fmt.Fprintf(os.Stderr, "%v\n", e)
    os.Exit(1)
  }

  s.SetStyle(tcell.StyleDefault.
    Foreground(tcell.ColorBlack).
    Background(tcell.ColorWhite))
  s.Clear()

  quit := make(chan struct{})
  go func() {
    for {
      ev := s.PollEvent()
      switch ev := ev.(type) {
      case *tcell.EventKey:
        switch ev.Key() {
        case tcell.KeyEscape, tcell.KeyEnter:
          close(quit)
          return
        case tcell.KeyCtrlL:
          s.Sync()
        }
      case *tcell.EventResize:
        s.Sync()
      }
    }
  }()

loop:
  for {
    select {
    case <-quit:
      break loop
    case <-time.After(time.Millisecond * 50):
    }

    // makebox(s)
  }

  s.Fini()
}