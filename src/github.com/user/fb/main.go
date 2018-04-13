package main

import (
	"fmt"

	"os/exec"

	ui "github.com/airking05/termui"
	fb "github.com/huandu/facebook"
)

func main() {
	//tok := "EAACEdEose0cBAIoqCOBM00MdZAu2xZB3VZAD5qkLz1YjASIIg8O20ZBkjmtELG4k8KTQd0MGy5EFBWdTS69buVC8fZCvsJ6oyCsxpdM32u3JDsDvOTgDHj5L5eZCnndKutjYOeFD13gxzVky2M1vTyWiHoc0T6ix0pUT9UEkRDq5z4dtplcLzHHJNATcqcB78ZD"
	//res, _ := fb.Get("/831976246906620/feed", fb.Params{
	//	"access_token": tok,
	//})

	//git for-each-ref --format '%(refname:short)' refs/heads/

	res1, _ := exec.Command("git", "branch").Output()
	//fmt.Printf("%s", res1)
	fmt.Print(string(res1))
	//fmt.Print(res1)
	return

	var items []fb.Result
	//err := res.DecodeField("data", &items)

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
	p := ui.NewPar(items[0]["message"].(string))
	p.Height = 30
	p.Width = 100
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Beautiful It"
	p.BorderFg = ui.ColorCyan
	//}

	//strs := []string{"[0] gizak/termui", "[1] editbox.go", "[2] interrupt.go", "[3] keyboard.go", "[4] output.go", "[5] random_out.go", "[6] dashboard.go", "[7] nsf/termbox-go"}
	//list := ui.NewList()
	//list.Items = strs
	//list.ItemFgColor = ui.ColorYellow
	//list.BorderLabel = "List"
	//list.Height = 7
	//list.Width = 25
	//list.Y = 4
	//
	//g := ui.NewGauge()
	//g.Percent = 50
	//g.Width = 50
	//g.Height = 3
	//g.Y = 11
	//g.BorderLabel = "Gauge"
	//g.BarColor = ui.ColorRed
	//g.BorderFg = ui.ColorWhite
	//g.BorderLabelFg = ui.ColorCyan

	//spark := ui.Sparkline{}
	//spark.Height = 1
	//spark.Title = "srv 0:"
	//spdata := []int{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	//spark.Data = spdata
	//spark.LineColor = ui.ColorCyan
	//spark.TitleColor = ui.ColorWhite
	//
	//spark1 := ui.Sparkline{}
	//spark1.Height = 1
	//spark1.Title = "srv 1:"
	//spark1.Data = spdata
	//spark1.TitleColor = ui.ColorWhite
	//spark1.LineColor = ui.ColorRed
	//
	//sp := ui.NewSparklines(spark, spark1)
	//sp.Width = 25
	//sp.Height = 7
	//sp.BorderLabel = "Sparkline"
	//sp.Y = 4
	//sp.X = 25

	//sinps := (func() []float64 {
	//	n := 220
	//	ps := make([]float64, n)
	//	for i := range ps {
	//		ps[i] = 1 + math.Sin(float64(i)/5)
	//	}
	//	return ps
	//})()
	//
	//lc := ui.NewLineChart()
	//lc.BorderLabel = "dot-mode Line Chart"
	//lc.Data = sinps
	//lc.Width = 50
	//lc.Height = 11
	//lc.X = 0
	//lc.Y = 14
	//lc.AxesColor = ui.ColorWhite
	//lc.LineColor = ui.ColorRed | ui.AttrBold
	//lc.Mode = "dot"
	//
	//bc := ui.NewBarChart()
	//bcdata := []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	//bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	//bc.BorderLabel = "Bar Chart"
	//bc.Width = 26
	//bc.Height = 10
	//bc.X = 51
	//bc.Y = 0
	//bc.DataLabels = bclabels
	//bc.BarColor = ui.ColorGreen
	//bc.NumColor = ui.ColorBlack
	//
	//lc1 := ui.NewLineChart()
	//lc1.BorderLabel = "braille-mode Line Chart"
	//lc1.Data = sinps
	//lc1.Width = 26
	//lc1.Height = 11
	//lc1.X = 51
	//lc1.Y = 14
	//lc1.AxesColor = ui.ColorWhite
	//lc1.LineColor = ui.ColorYellow | ui.AttrBold
	//
	//p1 := ui.NewPar("Hey!\nI am a borderless block!")
	//p1.Border = false
	//p1.Width = 26
	//p1.Height = 2
	//p1.TextFgColor = ui.ColorMagenta
	//p1.X = 52
	//p1.Y = 11

	draw := func(t int) {
		//g.Percent = t % 101
		//list.Items = strs[t%9:]
		//sp.Lines[0].Data = spdata[:30+t%50]
		//sp.Lines[1].Data = spdata[:35+t%50]
		//lc.Data = sinps[t/2%220:]
		//lc1.Data = sinps[2*t%220:]
		//bc.Data = bcdata[t/2%10:]
		//ui.Render(p, list, g, sp, lc, bc, lc1, p1)
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
