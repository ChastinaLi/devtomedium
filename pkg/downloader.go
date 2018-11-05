package pkg

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	resty "gopkg.in/resty.v1"
)

func Download_markdown(url string) string {
	html_source := download_html(url)
	return extract_markdown(html_source)
}

func download_html(url string) string {
	resty.SetCookie(&http.Cookie{
		Name:     "remember_user_token",
		Value:    "W1s3OTQ0Nl0sIjEzclp2TlV6bXVFRWEzVmJZSzVzIiwiMTU0MDUxMTk3Ny4xODIyNSJd--c642e37ccfe6136cf8fb11915ca2137cb3d10ac2",
		Path:     "/",
		Domain:   "dev.to",
		HttpOnly: true,
		Secure:   false,
	})

	resp, err := resty.R().
		SetHeader("Accept", "text/html").
		SetHeader("X-DevTools-Emulate-Network-Conditions-Client-Id", "B8971F415BB09694F5CC55B5284DCABA").
		Get(url + "/edit")
	if err != nil {
		return ""
	}
	return resp.String()
}

func extract_markdown(html_source string) string {
	doc, err := html.Parse(strings.NewReader(html_source))
	if err != nil {
		log.Fatal(err)
	}
	var node *html.Node
	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "textarea" {
			node = n.FirstChild
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)
	if node == nil {
		return ""
	}
	markdown := node.Data
	return markdown
}
