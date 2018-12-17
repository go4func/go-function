package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

const audioURL = "http://www.internetradiouk.com/"

// const audioURL = "http://www.google.com/"

func main() {
	TestHTML()
}

func TestHTML() {
	resp, err := http.Get(audioURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for _, l := range visit(resp, nil, doc) {
		log.Println(l)
	}
}

func visit(resp *http.Response, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// parse url
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					break
				}
				links = append(links, link.String())
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(resp, links, c)
	}
	return links

}

func TestOutline() {
	resp, err := http.Get(audioURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		log.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
