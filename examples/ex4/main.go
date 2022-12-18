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
