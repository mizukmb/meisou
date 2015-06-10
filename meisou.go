package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

const Version string = "0.1.0"

type Timer struct {
	Second  int
	Minutes int
}

func (t *Timer) ToMinutes(sec int) {
	t.Minutes = sec / 60
	t.Second = sec % 60
}

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
	var timer Timer
	timer.ToMinutes(sec)

	fmt.Println("Mental concentration. Let's \"Meisou\"...")
	for i := sec; i > 0; i-- {
		timer.ToMinutes(i)
		fmt.Printf("%.2d:%.2d\n", timer.Minutes, timer.Second)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("finished")

}
