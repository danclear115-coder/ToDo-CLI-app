package taskFunc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var stdin = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	return readLineInternal(prompt, "", false)
}

func ReadLineWithDefault(prompt, initial string) string {
	return readLineInternal(prompt, initial, false)
}

func ReadPassword(prompt string) string {
	return readLineInternal(prompt, "", true)
}

func readLineInternal(prompt, initial string, mask bool) string {
	fd := int(os.Stdin.Fd())

	if !term.IsTerminal(fd) {
		fmt.Print(prompt + initial)
		s, _ := stdin.ReadString('\n')
		if trimmed := strings.TrimSpace(s); trimmed != "" {
			return trimmed
		}
		return initial
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Print(prompt + initial)
		s, _ := stdin.ReadString('\n')
		if trimmed := strings.TrimSpace(s); trimmed != "" {
			return trimmed
		}
		return initial
	}
	defer term.Restore(fd, oldState)

	buf := []rune(initial)
	pos := len(buf)

	redraw := func() {
		display := string(buf)
		if mask {
			display = strings.Repeat("*", len(buf))
		}
		fmt.Printf("\r\033[K%s%s", prompt, display)
		if back := len(buf) - pos; back > 0 {
			fmt.Printf("\033[%dD", back)
		}
	}
	redraw()

	for {
		r, _, err := stdin.ReadRune()
		if err != nil {
			break
		}

		switch {
		case r == '\r' || r == '\n':
			fmt.Print("\r\n")
			return strings.TrimSpace(string(buf))

		case r == 3:
			term.Restore(fd, oldState)
			fmt.Printf("\n%s[!] Прервано пользователем.%s\n", Red, Reset)
			os.Exit(0)

		case r == 127 || r == 8:
			if pos > 0 {
				buf = append(buf[:pos-1], buf[pos:]...)
				pos--
				redraw()
			}

		case r == 1:
			pos = 0
			redraw()

		case r == 5:
			pos = len(buf)
			redraw()

		case r == 21:
			buf = buf[:0]
			pos = 0
			redraw()

		case r == 27:
			next, _, err := stdin.ReadRune()
			if err != nil || (next != '[' && next != 'O') {
				continue
			}
			seq, _, err := stdin.ReadRune()
			if err != nil {
				continue
			}
			switch seq {
			case 'D':
				if pos > 0 {
					pos--
					redraw()
				}
			case 'C':
				if pos < len(buf) {
					pos++
					redraw()
				}
			case 'H':
				pos = 0
				redraw()
			case 'F':
				pos = len(buf)
				redraw()
			case '1', '3', '4', '7', '8':
				tail, _, _ := stdin.ReadRune()
				if tail != '~' {
					continue
				}
				switch seq {
				case '3':
					if pos < len(buf) {
						buf = append(buf[:pos], buf[pos+1:]...)
						redraw()
					}
				case '1', '7':
					pos = 0
					redraw()
				case '4', '8':
					pos = len(buf)
					redraw()
				}
			}

		case r >= 32:
			buf = append(buf, 0)
			copy(buf[pos+1:], buf[pos:])
			buf[pos] = r
			pos++
			redraw()
		}
	}

	return strings.TrimSpace(string(buf))
}
