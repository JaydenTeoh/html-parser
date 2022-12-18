package main

import (
	"fmt"
	parser "parse_html/pkg/parse"
	utils "parse_html/pkg/utils"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := parser.Parse(r)
	utils.CheckErrors(err)
	fmt.Printf("%+v\n", links)
}
