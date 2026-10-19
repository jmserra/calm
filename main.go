//go:build !windows

// calm — a tiny terminal breathing companion. Stare and breathe.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

// mode holds a breathing pattern: inhale, hold-in, exhale, hold-out (seconds)
// plus a one-line blurb. Defaults to "coherent" — ~6 breaths/min, the
// resonance frequency associated with peak heart-rate variability.
type mode struct {
	inhale, hold1, exhale, hold2 float64
	blurb                        string
}

// modes is ordered so help output and iteration stay stable.
var modeNames = []string{"coherent", "box", "478", "exhale"}

var modes = map[string]mode{
	"coherent": {5.5, 0.0, 5.5, 0.0,
		"~6 breaths/min. Resonant, evidence-backed default for just relaxing."},
	"box": {4.0, 4.0, 4.0, 4.0,
		"Equal four counts. Grounding under stress; used by military/first responders."},
	"478": {4.0, 7.0, 8.0, 0.0,
		"Dr. Weil's sedative pattern. Good for falling asleep or coming down from anxiety."},
	"exhale": {4.0, 0.0, 6.0, 0.0,
		"Longer out than in. A gentle parasympathetic nudge, no holds to think about."},
}

const defaultMode = "coherent"

// One quiet character, filled. Nothing else.
const fill = "·"

// ANSI: hide cursor, clear, home, show cursor, reset.
const (
	hide  = "\x1b[?25l"
	show  = "\x1b[?25h"
	clear = "\x1b[2J"
	home  = "\x1b[H"
	reset = "\x1b[0m"
)

// fg returns a 256-color foreground escape.
func fg(n int) string { return fmt.Sprintf("\x1b[38;5;%dm", n) }

// phase returns the eased radius fraction 0..1 for time t (seconds).
func phase(t float64, m mode) float64 {
	cycle := m.inhale + m.hold1 + m.exhale + m.hold2
	p := math.Mod(t, cycle)
	if p < m.inhale {
		if m.inhale == 0 {
			return 1.0
		}
		return 0.5 - 0.5*math.Cos(math.Pi*(p/m.inhale))
	}
	p -= m.inhale
	if p < m.hold1 {
		return 1.0
	}
	p -= m.hold1
	if p < m.exhale {
		if m.exhale == 0 {
			return 0.0
		}
		return 0.5 + 0.5*math.Cos(math.Pi*(p/m.exhale))
	}
	return 0.0
}

// termSize returns the terminal columns and rows, falling back to 80x24.
func termSize() (int, int) {
	var ws struct{ Row, Col, Xpixel, Ypixel uint16 }
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if errno != 0 || ws.Col == 0 || ws.Row == 0 {
		return 80, 24
	}
	return int(ws.Col), int(ws.Row)
}

// draw renders one frame of the breathing circle for time t.
func draw(w *bufio.Writer, t float64, m mode) {
	cols, rows := termSize()
	cx, cy := float64(cols)/2, float64(rows)/2
	rFrac := phase(t, m)
	// Radius scales with the smaller terminal axis, leaving margin.
	rMax := math.Min(float64(cols)/2.4, float64(rows)*0.85)
	rMin := rMax * 0.45 // never collapses to a slit
	r := rMin + (rMax-rMin)*rFrac

	w.WriteString(home)
	for y := 1; y <= rows; y++ {
		for x := 1; x <= cols; x++ {
			// Terminal cells are taller than wide; squash y so it reads round.
			dx := float64(x) - cx
			dy := (float64(y) - cy) * 2.1
			if math.Hypot(dx, dy) <= r {
				w.WriteString(fill)
			} else {
				w.WriteByte(' ')
			}
		}
		if y < rows {
			w.WriteByte('\n')
		}
	}
	w.Flush()
}

// cleanup restores the terminal and exits.
func cleanup() {
	os.Stdout.WriteString(show + reset + clear + home)
	os.Exit(0)
}

func printHelp() {
	nameW := 0
	for _, n := range modeNames {
		if len(n) > nameW {
			nameW = len(n)
		}
	}
	fmt.Println("calm — a terminal breathing companion.")
	fmt.Println()
	fmt.Printf("usage: calm [mode]    (default: %s)\n\n", defaultMode)
	fmt.Println("modes:")
	for _, name := range modeNames {
		m := modes[name]
		timing := fmt.Sprintf("%s-%s-%s-%s",
			fmtNum(m.inhale), fmtNum(m.hold1), fmtNum(m.exhale), fmtNum(m.hold2))
		marker := " "
		if name == defaultMode {
			marker = "*"
		}
		fmt.Printf("  %s %s   %s  %s\n",
			marker, padRight(name, nameW), padRight(timing, 13), m.blurb)
	}
	fmt.Println()
	fmt.Println("timing is inhale-hold-exhale-hold, in seconds.  ctrl-c to leave.")
}

// fmtNum formats a float like Python's %g — no trailing zeros.
func fmtNum(x float64) string {
	return fmt.Sprintf("%g", x)
}

func padRight(s string, w int) string {
	if len(s) >= w {
		return s
	}
	return s + strings.Repeat(" ", w-len(s))
}

func main() {
	arg := defaultMode
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if arg == "-h" || arg == "--help" || arg == "help" {
		printHelp()
		return
	}
	m, ok := modes[arg]
	if !ok {
		fmt.Fprintf(os.Stderr, "calm: unknown mode '%s'\n\n", arg)
		printHelp()
		os.Exit(2)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cleanup()
	}()

	os.Stdout.WriteString(hide + clear + fg(245))
	w := bufio.NewWriter(os.Stdout)
	t0 := time.Now()
	ticker := time.NewTicker(time.Second / 30) // ~30fps, easy on eyes and CPU
	defer ticker.Stop()
	for range ticker.C {
		draw(w, time.Since(t0).Seconds(), m)
	}
}
