package main

import (
  "fmt"
  "os"
  "time"

  "github.com/gdamore/tcell"
  "github.com/mattn/go-runewidth"
)

// holds coordinates of where we want player to be
type player struct {
	x int
	y int
}
// use above struct to tell emitStr() where we are, allows us to move on the screen

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
  for _, c := range str {
    var comb []rune
    w := runewidth.RuneWidth(c)
    if w == 0 {
      comb = []rune{c}
      c = ' '
      w = 1
    }
    s.SetContent(x, y, c, comb, style)
    x += w
  }
}

func main() {
// will be printing out location to the screen, don't need to use debug
  debug:=false
// initialize player instance at 0,0
  player:= player{
	  x:0,
	  y:0,
  }
 
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

// pass white directly to our printing function
  white := tcell.StyleDefault.
    Foreground(tcell.ColorWhite).
    Background(tcell.ColorBlack)

// swap forground and background colors
  s.SetStyle(tcell.StyleDefault.
    Foreground(tcell.ColorWhite).
    Background(tcell.ColorBlack))
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
	// clear our screen inside our loop
    s.Clear()
		emitStr(s, 0, 0, white, "@")
	// call show to actually display it on screen
	s.Show()
  }

  s.Fini()
}