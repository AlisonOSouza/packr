package fix

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlisonOSouza/packr/v2/jam/store"
)

//YesToAll will be used by the command to skip the questions
var YesToAll bool

var replace = map[string]string{
	"github.com/AlisonOSouza/packr": "github.com/AlisonOSouza/packr/v2",
}

var ic = ImportConverter{
	Data: replace,
}

var checks = []Check{
	// packrClean,
	ic.Process,
}

func packrClean(r *Runner) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	store.Clean(pwd)
	return nil
}

func ask(q string) bool {
	if YesToAll {
		return true
	}

	fmt.Printf("? %s [y/n]\n", q)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ToLower(strings.TrimSpace(text))
	return text == "y" || text == "yes"
}
