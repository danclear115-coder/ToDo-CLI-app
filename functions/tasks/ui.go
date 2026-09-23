package taskFunc

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

const (

	Green  = "\033[32m"
	Cyan   = "\033[36m"
	Red    = "\033[31m"
	Gray   = "\033[90m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	
)

const (
	minTermWidth = 50
	maxTermWidth = 140
	defTermWidth = 80
)

func TermWidth() int {

	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		w = defTermWidth
	}
	if w < minTermWidth {
		w = minTermWidth
	}
	if w > maxTermWidth {
		w = maxTermWidth
	}
	return w
}

func FormWidth() int {
	
	w := TermWidth() - 4
	if w > 92 {
		w = 92
	}
	if w < 36 {
		w = 36
	}
	return w
}

func WrapText(s string, width int) []string {

	if width < 1 {
		width = 1
	}
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	cur := ""

	flush := func() {
		if cur != "" {
			lines = append(lines, cur)
			cur = ""
		}
	}

	for _, word := range words {
		w := []rune(word)
		for len(w) > width {
			flush()
			lines = append(lines, string(w[:width]))
			w = w[width:]
		}
		word = string(w)

		switch {
		case cur == "":
			cur = word
		case len([]rune(cur))+1+len([]rune(word)) <= width:
			cur += " " + word
		default:
			flush()
			cur = word
		}
	}
	flush()

	if len(lines) == 0 {
		lines = []string{""}
	}
	return lines
}

func Pad(s string, w int) string {
	r := []rune(s)
	if len(r) >= w {
		return s
	}
	return s + strings.Repeat(" ", w-len(r))
}

type Column struct {
	Header string

	Fixed int

	Weight int

	MinWidth int
}

type Cell struct {
	Text  string
	Color string
}

func columnWidths(cols []Column) []int {
	total := TermWidth()
	overhead := 1 + len(cols)*3
	avail := total - overhead
	if avail < len(cols) {
		avail = len(cols)
	}

	widths := make([]int, len(cols))
	fixedSum, weightSum := 0, 0
	for i, c := range cols {
		if c.Fixed > 0 {
			widths[i] = c.Fixed
			fixedSum += c.Fixed
		} else {
			weightSum += c.Weight
		}
	}
	if weightSum == 0 {
		weightSum = 1
	}

	flexAvail := avail - fixedSum
	if flexAvail < 0 {
		flexAvail = 0
	}

	for i, c := range cols {
		if c.Fixed > 0 {
			continue
		}
		w := flexAvail * c.Weight / weightSum
		if w < c.MinWidth {
			w = c.MinWidth
		}
		widths[i] = w
	}

	for i, c := range cols {
		if hw := len([]rune(c.Header)); widths[i] < hw {
			widths[i] = hw
		}
	}

	return widths
}

func gridLine(left, mid, right string, widths []int) string {
	parts := make([]string, len(widths))
	for i, w := range widths {
		parts[i] = strings.Repeat("─", w+2)
	}
	return left + strings.Join(parts, mid) + right
}

func printGridRow(borderColor string, cells []Cell, widths []int) {
	wrapped := make([][]string, len(cells))
	maxLines := 1
	for i, cell := range cells {
		wrapped[i] = WrapText(cell.Text, widths[i])
		if len(wrapped[i]) > maxLines {
			maxLines = len(wrapped[i])
		}
	}

	for line := 0; line < maxLines; line++ {
		var b strings.Builder
		b.WriteString(borderColor + "│" + Reset)
		for i := range cells {
			text := ""
			if line < len(wrapped[i]) {
				text = wrapped[i][line]
			}
			b.WriteString(" " + cells[i].Color + Pad(text, widths[i]) + Reset + " " + borderColor + "│" + Reset)
		}
		fmt.Println(b.String())
	}
}

func RenderTable(borderColor string, cols []Column, rows [][]Cell) {
	widths := columnWidths(cols)

	fmt.Println()
	fmt.Println(borderColor + gridLine("┌", "┬", "┐", widths) + Reset)

	headerCells := make([]Cell, len(cols))
	for i, c := range cols {
		headerCells[i] = Cell{Text: c.Header, Color: Bold + borderColor}
	}
	printGridRow(borderColor, headerCells, widths)
	fmt.Println(borderColor + gridLine("├", "┼", "┤", widths) + Reset)

	for i, row := range rows {
		printGridRow(borderColor, row, widths)
		if i != len(rows)-1 {
			fmt.Println(borderColor + gridLine("├", "┼", "┤", widths) + Reset)
		}
	}

	fmt.Println(borderColor + gridLine("└", "┴", "┘", widths) + Reset)
	fmt.Println()
}

func BoxTop(title, color string) int {
	w := FormWidth()
	t := "─ [ " + title + " ] "
	fill := w - len([]rune(t))
	if fill < 0 {
		fill = 0
	}
	fmt.Printf("\n%s┌%s%s┐%s\n", color, t, strings.Repeat("─", fill), Reset)
	return w
}

func BoxBottom(width int, color string) {
	fmt.Printf("%s└%s┘%s\n\n", color, strings.Repeat("─", width), Reset)
}

func BoxLine(width int, borderColor, textColor, text string) {
	inner := width - 2
	for _, line := range WrapText(text, inner) {
		fmt.Printf("%s│%s %s%s%s %s│%s\n",
			borderColor, Reset,
			textColor, Pad(line, inner), Reset,
			borderColor, Reset)
	}
}

func BoxPrompt(borderColor, label string) string {
	return fmt.Sprintf("%s│%s %s", borderColor, Reset, label)
}

func AskID(width int, color string) (uint, bool) {
	idStr := ReadLine(BoxPrompt(color, "ID задачи: "))

	val64, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		BoxLine(width, color, Red, "ID должен быть целым положительным числом.")
		BoxBottom(width, color)
		return 0, false
	}

	return uint(val64), true
}
