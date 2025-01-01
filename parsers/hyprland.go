package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bind struct {
	mods   []string
	keys   []string
	action string
	desc   string
}

type Parser struct {
	files        []string
	pattern      func(string, map[string]string) (Bind, bool)
	translations map[string]string
	output       string
}

var Parsers = map[string]Parser{
	"hyprland": {
		files: []string{"~/dotfiles/hypr/.config/hypr/hyprland/keybinds.conf"},
		pattern: func(line string, translations map[string]string) (Bind, bool) {
			split := strings.Split(line, "=")
			if strings.Trim(split[0], " ") != "bind" {
				return Bind{}, false
			}

			contents := translate(split[1], translations)

			args := strings.Split(contents, ",")

			if len(args) <= 3 {
				return Bind{
					mods:   strings.Split(strings.Trim(args[0], " "), "_"),
					keys:   strings.Split(strings.Trim(args[1], " "), "_"),
					action: strings.Trim(args[2], " "),
				}, true
			} else {
				return Bind{
					mods:   strings.Split(strings.Trim(args[0], " "), "_"),
					keys:   strings.Split(strings.Trim(args[1], " "), "_"),
					action: strings.Trim(args[2]+": "+args[3], " "),
				}, true
			}
		},
		translations: map[string]string{
			"$mods": "Mod_Shift",
			"$modc": "Mod_Ctrl",
			"$moda": "Mod_Alt",
			"$mod":  "Mod",
		},
	},
	"zsh": {
		files: []string{"~/.config/zsh/.zshrc", "~/.zshrc"},
		translations: map[string]string{
			"^M": "",
		},
	},
}

func fixString(str string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	str = strings.Replace(str, "~", homeDir, 1)
	str = strings.Trim(str, " ")
	return str
}

func (p Parser) Compile() {

	binds := []Bind{}

	for _, filepath := range p.files {
		file, err := os.Open(fixString(filepath))
		defer file.Close()
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			bind, isValid := p.pattern(line, p.translations)
			if isValid {
				binds = append(binds, bind)
			}
		}
	}
	fmt.Println(binds)
}

func (p Parser) Render() {
}

func translate(line string, words map[string]string) string {
	for key, val := range words {
		line = strings.ReplaceAll(line, key, val)
	}
	return line
}
