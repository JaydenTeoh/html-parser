package link

import (
	"io"
	"strings"

	html "golang.org/x/net/html"
)

// Link represents a link (<a href="..."></a>) in a HTML
type Link struct {
	Href string
	Text string
}

// Parse will take in a reader for HTML and return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	// 1. Find <a> nodes in document
	// 2. for each link node:
	// 2a. build a Link
	// 3. return the Links
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break // only get the parent <a> href key; ignore nested links. append the href key to the Href value of the return Link
		}
	}
	ret.Text = getText(n) // get all the text data from all the TextNodes in each parent <a> node and append it to the return Link
	return ret
}

// get all the text data from all the TextNodes in a parent node
func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c)
	}
	return strings.Join(strings.Fields(ret), " ") //Fields func splits string sperated by consec whitespace/newline (remove extra spaces) and return an array of strings
	//Join connects these array of strings with the separator string (in our cases => " ") and return a single string
}

// get an array of all the <a> nodes from the root node
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
