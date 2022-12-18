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
