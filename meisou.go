package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
)

const Version string = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "meisou"
	app.Version = Version
	app.Usage = "make time of meisou"
	app.Author = "mizukmb"
	app.Email = "mizukmb6@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

type Timer struct {
	Second  int
	Minutes int
}

func (t *Timer) ToMinutes(sec int) {
	t.Minutes = sec / 60
	t.Second = sec % 60
}

var timer Timer

func tbprint(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(sec int) {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	startmsg := fmt.Sprintf("Mental concentration. Let's \"Meisou\"...")
	finishmsg := fmt.Sprintf("Finished Meisou. press Key `Esc` or `q` to exit.")

	w, h := termbox.Size()
	midx := w / 2
	midy := h / 2

	tbprint(midx-(len(startmsg)/2), midy-1, startmsg, coldef, coldef)

	for i := sec; i >= 0; i-- {
		timer.ToMinutes(i)
		timermsg := fmt.Sprintf("%.2d:%.2d\n", timer.Minutes, timer.Second)
		tbprint(midx-(len(timermsg)/2), midy, timermsg, coldef, coldef)
		termbox.Flush()
		time.Sleep(1 * time.Second)
	}
	tbprint(midx-(len(finishmsg)/2), midy+1, finishmsg, coldef, coldef)
	termbox.Flush()
}

func doMain(c *cli.Context) {
	if len(os.Args) < 2 {
		fmt.Println("Please setting time(minutes). ex. \"meisou 3\".")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if num > 60 {
		fmt.Println("Too long!! The Meisou's most suitable time is said during 30 minutes and 60 minutes.")
		return
	}

	sec := num * 60

	const coldef = termbox.ColorDefault
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw(sec)

	cmd := exec.Command("say", "\"Finished Meisou. press Key `Esc` or `q` to exit.\"")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Ch == 'q' {
				break loop
			}
		}
	}
}
