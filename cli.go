// Demo code for the List primitive.
package main

import (
	// "io"
	// "log"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	// "tview"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var args []string
var buildTime string
var commsFromFile string
var textView = tview.NewTextView()
var app = tview.NewApplication()
var list = tview.NewList()
var comms []string
var commsMod []string

func linesToCommands(stri []string) []string {
	var r []string
	var lastComment = ""
	var appendd = ""
	for _, str := range stri {
		if str == "" {
			lastComment = ""
		} else if strings.Index(str, "#") == 0 {
			runes := []rune(str)
			lastComment = strings.TrimSpace(string(runes[1:]))
		} else {
			appendd = ""
			if len(lastComment) > 0 {
				appendd = lastComment + "^"
			}
			r = append(r, appendd+str)
			lastComment = ""
		}
	}
	return r
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func isRoot() bool {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		//fmt.Println(err)
		//os.Exit(1)
	}
	// fmt.Println(string(stdout))
	return strings.Replace(string(stdout), "\n", "", -1) == "root"
}
func parseCommand(command string) {

	if isRoot() {
		fmt.Println("\u001b[31mrunning as root\u001b[0m")
	}

	// check for bash if available
	useShell := "bash"
	lookPath, err := exec.LookPath(useShell)
	if err != nil {
		//fmt.Println("bash not found, using shell")
		useShell = "sh"
		lookPath, _ = exec.LookPath(useShell)
	}
	split := strings.Split(command, "^")
	if len(split) > 1 {
		split = RemoveIndex(split, 0)
		command = strings.Join(split, "^")
	}

	// find prompts
	re := regexp.MustCompile(`(<prompt:.*?>)`)
	reSplit := re.FindAllString(command, -1)
	var promptArr []string

	for i := range reSplit {
		_, found := Find(promptArr, reSplit[i])
		if !found {
			promptArr = append(promptArr, reSplit[i])
		}
	}
	//fmt.Print(promptArr)

	if len(promptArr) > 0 {
		fmt.Println("\033[33m" + command + "\033[0m")
		var strSplit []string
		var read string
		var out string
		var stri string
		for i := 0; i < len(promptArr); i++ {
			stri = strings.Trim(promptArr[i], "<>")
			strSplit = strings.Split(stri, ":")

			fmt.Print(strSplit[1])
			if len(strSplit) > 2 {
				fmt.Print(" (" + strSplit[2] + ")")
			}
			fmt.Print(": ")
			_, _ = fmt.Scanln(&read)
			if read == "" && len(strSplit) > 2 {
				out = strSplit[2]
			} else {
				out = read
			}
			command = strings.Replace(command, promptArr[i], out, -1)
			//fmt.Print(first)
		}
		//fmt.Print(promptArr)
	}

	//command = "cat"
	fmt.Println("\033[33m" + lookPath + ": " + command + "\033[0m")
	cmd := exec.Command(useShell, "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	// fmt.Print
	/* stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	// go func() {
		defer stdin.Close()
		// io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	// }()

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Print(string(stdout))
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout)) */

}
func searchRender(search string) {
	list.Clear()
	var split []string
	searchSplit := strings.Split(search, " ")
	//keywords := []string{}
	commsMod = comms
	miss := false
	var commColor = "[green]"
	if isRoot() {
		commColor = "[red]"
	}
	// fmt.Println(len(searchSplit));
	// fmt.Println(len(search));
	if len(searchSplit) > 0 && len(search) > 0 {

		commsMod = []string{}
		for n := 0; n < len(comms); n++ {
			miss = false
			for i := 0; i < len(searchSplit); i++ {
				if !strings.Contains(comms[n], searchSplit[i]) && len(searchSplit[i]) > 0 {
					miss = true
					break
				}
			}
			if !miss {
				commsMod = append(commsMod, comms[n])
			}
		}
	}

	for i := 0; i < len(commsMod); i++ {
		toParse := commsMod[i]
		//toParse = ""
		split = strings.Split(toParse, "^")
		if len(split) > 1 {
			toParse = "[white]" + split[0] + "[grey]^" + commColor
			split = RemoveIndex(split, 0)
			toParse += strings.Join(split, "^")
		} else {
			toParse = commColor + strings.Join(split, "^")
		}

		list.AddItem(toParse, "", 0, nil)
		//fmt.Println(comms[i]);
	}
}
func main() {

	removedNewLines := strings.Replace(commsFromFile, "\r\n", "\n", -1)
	replacedTicks := strings.Replace(removedNewLines, "_____single-tick_____", "'", -1)
	comms = linesToCommands(strings.Split(replacedTicks, "\n"))
	//fmt.Println(comms)
	//os.Exit(99)
	args = os.Args[1:]
	list = tview.NewList().SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		toParse := commsMod[index]
		if strings.Contains(toParse, "^") {
			breakIndex := strings.Index(toParse, "^") + 1
			toParse = string([]rune(toParse)[breakIndex:])
		}
		textView.SetText("\n[yellow]" + toParse)
	})
	inputField := tview.NewInputField()
	search := ""

	if len(args) > 0 {

		if args[0] == "-v" {
			fmt.Printf("Build Time: %s\n", buildTime)
			return
		}
		for i := 0; i < len(comms); i++ {
			if strings.Index(comms[i], args[0]+" ") == 0 || strings.Index(comms[i], args[0]+"^") == 0 {
				parseCommand(comms[i])
				break
			}
		}
		return
	}
	inputField.SetLabel("").SetPlaceholder("").SetFieldWidth(0).
		//SetAcceptanceFunc(tview.InputFieldInteger).
		SetChangedFunc(func(text string) {
			search = inputField.GetText()
			searchRender(search)
		}).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			app.Stop()
			// command, _ := list.GetItemText(list.GetCurrentItem())
			//fmt.Println(len(commsMod))
			command := commsMod[list.GetCurrentItem()]
			parseCommand(command)
		} else if key == tcell.KeyPgDn {
			list.SetCurrentItem(list.GetCurrentItem() + 20)
		} else if key == tcell.KeyDown {
			if list.GetCurrentItem() >= list.GetItemCount()-1 {
				list.SetCurrentItem(0)
			} else {
				list.SetCurrentItem(list.GetCurrentItem() + 1)
			}
		} else if key == tcell.KeyUp {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		} else {

		}
	})

	list.ShowSecondaryText(false)
	searchRender(search)
	textView.SetDynamicColors(true).
		SetBorder(false)
		//SetRegions(false).
		//SetWordWrap(true).
		/* SetChangedFunc(func() {
			app.Draw()
		}) */
	//box.SetBorder(true).SetBorderAttributes(tcell.AttrBold)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 1, 1, true).
		//AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, false).
		AddItem(textView, 5, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}

}
