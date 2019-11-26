package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

// sliceContains returns true if `slice` contains `value`
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// pageLinks will recursively scan a `html.Node` and will return
// a list of links found, with no duplicates
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// pageTitle given a reference to a html.Node, scans it until it
// finds the title tag, and returns its value
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// parse given a string pointing to a URL will fetch and parse it
// returning an html.Node pointer
func parse(url string) (*html.Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}
	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse page")
	}
	return b, err
}


func analyze(url, baseurl string, visited *map[string]string) {
	page, err := parse(url)
	if err != nil {
		fmt.Printf("Error getting page %s %s\n", url, err)
		return
	}
	title := pageTitle(page)
	(*visited)[url] = title

	//recursively find links
	links := pageLinks(nil, page)
	for _, link := range links {
		if (*visited)[link] == "" && strings.HasPrefix(link, baseurl) {
			analyze(link, baseurl, visited)
		}
	}
}

func checkDuplicates(visited *map[string]string) {
	found := false
	uniques := map[string]string{}
	fmt.Printf("\nChecking duplicates..\n")
	for link, title := range *visited {
		if uniques[title] == "" {
			uniques[title] = link
		} else {
			found = true
			fmt.Printf("Duplicate title \"%s\" in %s but already found in %s\n", title, link, uniques[title])
		}
	}

	if !found {
		fmt.Println("No duplicates were found 😇")
	}
}

func main() {
	var url string
	var dup bool
	//flag.StringVar(&url, "url", "", "the url to parse")
	//flag.BoolVar(&dup, "dup", false, "if set, check for duplicates")
	//flag.Parse()
	url = "https://zenio.vn/"
	if url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	visited := map[string]string{}
	analyze(url, url, &visited)
	for link, title := range visited {
		fmt.Printf("%s -> %s\n", link, title)
	}

	if dup {
		checkDuplicates(&visited)
	}
}
