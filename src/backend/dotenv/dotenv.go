package dotenv

import (
	"os"
	"regexp"
)

var variables = make(map[string]string)

func GetVar(name string) string {
	if len(variables) == 0 {
		setup()
	}

	value, exists := variables[name]

	if !exists {
		panic("Enviroment variable not found: " + name)
	}

	return value
}

func setup() {
	file, err := os.ReadFile("./../.env")
	if err != nil {
		panic(err.Error())
	}

	reg := regexp.MustCompile("(?P<key>.*)=(?P<value>.*)")

	matches := reg.FindAllStringSubmatch(string(file), -1)

	for _, line := range matches {
		variables[line[1]] = line[2]
	}
}
