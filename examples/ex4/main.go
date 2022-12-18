package main

import (
	"fmt"
	"strings"

	parser "github.com/JaydenTeoh/html-parser/pkg/parse"
	utils "github.com/JaydenTeoh/html-parser/pkg/utils"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>

`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := parser.Parse(r)
	utils.CheckErrors(err)
	fmt.Printf("%+v\n", links)
}
