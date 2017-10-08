package htmllinkextractor

import (
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func parseLinksFromHTMLDocument(n *html.Node) []Link {
	var links []Link
	findLinksInHTMLDocumentAndAppendToList(n, &links)
	return links
}

func findLinksInHTMLDocumentAndAppendToList(htmlNode *html.Node, links *[]Link) {
	if isAnchorElement(htmlNode) {
		var link Link
		link.Href = getHrefAttributeFromHTMLElement(htmlNode)
		link.Text = getInnerTextFromHtmlNode(htmlNode.FirstChild)
		*links = append(*links, link)
	}
	for c := htmlNode.FirstChild; c != nil; c = c.NextSibling {
		findLinksInHTMLDocumentAndAppendToList(c, links)
	}
}

func getHrefAttributeFromHTMLElement(htmlNode *html.Node) string {
	var href string
	for _, attribute := range htmlNode.Attr {
		if isHrefAttribute(attribute) {
			href = attribute.Val
			break
		}
	}
	return href
}

func getInnerTextFromHtmlNode(innerTextNode *html.Node) string {
	var innerText string
	if isTextNode(innerTextNode) {
		innerText = innerTextNode.Data + getTextWithinTagsInsideTextNode(innerTextNode)
	}
	return trim(innerText)
}

func getTextWithinTagsInsideTextNode(textNode *html.Node) string {
	textWithinHtmlTag := getTextWrappedInHtmlTagsInsideTextNode(textNode)
	textAfterHtmlTag := getTextAfterTagsInsideTextNode(textNode)
	return textWithinHtmlTag + textAfterHtmlTag
}

func getTextAfterTagsInsideTextNode(textNode *html.Node) string {
	var innerText string
	if hasNextSibling(textNode) && hasNextSibling(textNode.NextSibling) && isTextNode(textNode.NextSibling.NextSibling) {
		innerText = innerText + textNode.NextSibling.NextSibling.Data
	}
	return innerText
}

func hasNextSibling(node *html.Node) bool {
	return node.NextSibling != nil
}

func isTextNode(htmlNode *html.Node) bool {
	return htmlNode.Type == html.TextNode
}

func getTextWrappedInHtmlTagsInsideTextNode(textNode *html.Node) string {
	var textEnclosedByHtmlTags string
	if hasNextSibling(textNode) && isHtmlElementNode(textNode.NextSibling) && hasFirstChild(textNode.NextSibling) {
		textEnclosedByHtmlTags = textNode.NextSibling.FirstChild.Data
	}
	return textEnclosedByHtmlTags
}

func hasFirstChild(htmlNode *html.Node) bool {
	return htmlNode.FirstChild != nil
}

func trim(innerText string) string {
	return strings.Trim(innerText, "\n ")
}

func isHrefAttribute(attr html.Attribute) bool {
	return attr.Key == "href"
}

func isAnchorElement(n *html.Node) bool {
	return isHtmlElementNode(n) && n.Data == "a"
}

func isHtmlElementNode(n *html.Node) bool {
	return n.Type == html.ElementNode
}

func ExtractLinks(htmlText string) []Link {
	reader := strings.NewReader(htmlText)
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	return parseLinksFromHTMLDocument(doc)
}
