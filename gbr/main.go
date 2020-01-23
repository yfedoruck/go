package main

import (
	"os/exec"

	ui "github.com/airking05/termui"
)

func main() {
	//****  git for-each-ref --format '%(refname:short)' refs/heads/

	res1, _ := exec.Command("git", "for-each-ref", "--format", "%(refname:short)", "refs/heads/").Output()
	//fmt.Printf("%s", res1)
	//fmt.Print(res1)
	//fmt.Print(string(res1))
	//return

	//if err != nil {
	//	fmt.Printf("An error has happened %v", err)
	//	return
	//}

	//for _, item := range items {
	//	fmt.Println(item["message"])
	//}

	//fmt.Println(items[0]["message"])

	err1 := ui.Init()
	if err1 != nil {
		panic(err1)
	}
	defer ui.Close()

	//for _, item := range items {
	p := ui.NewPar(string(res1))
	p.Height = 30
	p.Width = 100
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Beautiful It"
	p.BorderFg = ui.ColorCyan
	//}

	draw := func(t int) {
		ui.Render(p)
	}
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})
	ui.Loop()
}
