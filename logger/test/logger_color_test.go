package test

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/logger"
	"testing"
)

func TestOriginal(t *testing.T) {
	//简单快速的使用，跟 fmt.Print* 类似
	color.Redp("Simple to use color")
	color.Redln("Simple to use color")
	color.Greenp("Simple to use color\n")
	color.Cyanln("Simple to use color")
	color.Yellowln("Simple to use color")

	// 简单快速的使用，跟 fmt.Print* 类似
	color.Red.Println("Simple to use color")
	color.Green.Print("Simple to use color\n")
	color.Cyan.Printf("Simple to use %s\n", "color")
	color.Yellow.Printf("Simple to use %s\n", "color")

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// 自定义颜色
	color.New(color.FgWhite, color.BgBlack).Println("custom color style")

	// 也可以:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

	// internal style:
	color.Info.Println("message")
	color.Warn.Println("message")
	color.Error.Println("message")

	// 使用内置颜色标签
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")
	// 自定义标签: 支持使用16色彩名称，256色彩值，rgb色彩值以及hex色彩值
	color.Println("<fg=11aa23>he</><bg=120,35,156>llo</>, <fg=167;bg=232>wel</><fg=red>come</>")

	// apply a style tag
	color.Tag("info").Println("info style text")

	// prompt message
	color.Info.Prompt("prompt style message")
	color.Warn.Prompt("prompt style message")

	// tips message
	color.Info.Tips("tips style message")
	color.Warn.Tips("tips style message")
}

func TestColor(t *testing.T) {
	black := color.FgBlack.Render
	red := color.FgRed.Render
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render
	blue := color.FgBlue.Render
	magenta := color.FgMagenta.Render
	cyan := color.FgCyan.Render
	white := color.FgWhite.Render
	defaultC := color.FgDefault.Render
	fmt.Println(black("black "), red("red "), green("green "), yellow("yellow "), blue("blue "), magenta("magenta "),
		cyan("cyan "), white("white "), defaultC(" defaultC "))

	black1 := color.FgDarkGray.Render
	red1 := color.FgLightRed.Render
	green1 := color.FgLightGreen.Render
	yellow1 := color.FgLightYellow.Render
	blue1 := color.FgLightBlue.Render
	magenta1 := color.FgLightMagenta.Render
	cyan1 := color.FgLightCyan.Render
	white1 := color.FgLightWhite.Render
	gray1 := color.FgGray.Render
	fmt.Println(black1("black "), red1("red "), green1("green "), yellow1("yellow "), blue1("blue "), magenta1("magenta "),
		cyan1("cyan "), white1("white "), gray1(" gray "))

	black2 := color.BgBlack.Render
	red2 := color.BgRed.Render
	green2 := color.BgGreen.Render
	yellow2 := color.BgYellow.Render
	blue2 := color.BgBlue.Render
	magenta2 := color.BgMagenta.Render
	cyan2 := color.BgCyan.Render
	white2 := color.BgWhite.Render
	gray2 := color.BgDefault.Render
	fmt.Println(black2("black "), red2("red "), green2("green "), yellow2("yellow "), blue2("blue "), magenta2("magenta "),
		cyan2("cyan "), white2("white "), gray2(" gray "))

	black3 := color.BgDarkGray.Render
	red3 := color.BgLightRed.Render
	green3 := color.BgLightGreen.Render
	yellow3 := color.BgLightYellow.Render
	blue3 := color.BgLightBlue.Render
	magenta3 := color.BgLightMagenta.Render
	cyan3 := color.BgLightCyan.Render
	white3 := color.BgLightWhite.Render
	gray3 := color.BgGray.Render
	fmt.Println(black3("black "), red3("red "), green3("green "), yellow3("yellow "), blue3("blue "), magenta3("magenta "),
		cyan3("cyan "), white3("white "), gray3(" gray "))

	reset := color.OpReset.Render
	bold := color.OpBold.Render
	fuzzy := color.OpFuzzy.Render
	italic := color.OpItalic.Render
	underscore := color.OpUnderscore.Render
	blink := color.OpBlink.Render
	fastBlink := color.OpFastBlink.Render
	reverse := color.OpReverse.Render
	concealed := color.OpConcealed.Render
	strikeThrough := color.OpStrikethrough.Render
	fmt.Println(reset("reset "), bold("bold "), fuzzy("fuzzy "), italic("italic "), underscore("underscore "), blink("blink "),
		fastBlink("fastBlink "), reverse("reverse "), concealed("concealed "), strikeThrough("strikeThrough "))
}

func TestLoggerColor(t *testing.T) {
	config.LoadYamlFile("./application-color.yaml")
	logger.InitLog()

	logger.Debug("debug data")
	logger.Info("info data")
	logger.Warn("warn data")
	logger.Error("error data")
	logger.Fatal("fatal data")
	//logger.Panic("panic data")
}
