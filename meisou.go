package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
)

const Version string = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "meisou"
	app.Version = Version
	app.Usage = "make time of meisou"
	app.Author = "mizukmb"
	app.Email = "mizukmb6@gmail.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "say, s", Usage: "shaberu in English."},
	}
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

	timer.ToMinutes(sec)
	timermsg := fmt.Sprintf("%.2d:%.2d\n", timer.Minutes, timer.Second)
	tbprint(midx-(len(timermsg)/2), midy, timermsg, coldef, coldef)
	if sec <= 0 {
		tbprint(midx-(len(finishmsg)/2), midy+1, finishmsg, coldef, coldef)
	}
	termbox.Flush()
}

func canUseTimer(num int) (bool, error) {
	if num <= 0 {
		return false, errors.New(fmt.Sprintf("Cannot use this time:%d.\n You should set time higher than 0.", num))
	}
	if num > 60 {
		return false, errors.New(fmt.Sprintf("Too long this time:%d.\n!! The Meisou's most suitable time is said during 30 minutes and 60 minutes.", num))
	}
	return true, nil
}

func termboxEvent(ev chan termbox.Event) {
	for {
		ev <- termbox.PollEvent()
	}
}

func doMain(c *cli.Context) {
	if len(os.Args) < 2 {
		fmt.Println("Please setting time(minutes). ex. \"meisou [global options] 3\".")
		return
	}

	var arg int
	if c.Bool("say") {
		arg = 2
	} else {
		arg = 1
	}

	num, err := strconv.Atoi(os.Args[arg])
	if err != nil {
		fmt.Println(err)
		return
	}

	ok, err := canUseTimer(num)
	if !ok {
		fmt.Println(err)
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

	du := 1
	tick := time.Tick(1 * time.Second)
	ev := make(chan termbox.Event)

	go termboxEvent(ev)

loop:
	for {
		select {
		case <-tick:
			if sec-du >= 0 {
				draw(sec - du)
				du++
			} else {
				break loop
			}
		case event := <-ev:
			if event.Key == termbox.KeyEsc {
				return
			}
			if event.Ch == 'q' {
				return
			}
		}
	}

	if c.Bool("say") {
		cmd := exec.Command("say", "\"Finished Meisou. press Key `Esc` or `q` to exit.\"")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	for {
		event := termbox.PollEvent()
		if event.Key == termbox.KeyEsc {
			return
		}
		if event.Ch == 'q' {
			return
		}
	}
}
