package main

import (
	"regexp"
	"strings"
)

var (
	ReIgnoreBlock = map[string]*regexp.Regexp{
		"doctype":  regexp.MustCompile(`(?ims)<!DOCTYPE.*?>`),               // raw doctype
		"comment":  regexp.MustCompile(`(?ims)<!--.*?-->`),                  // raw comment
		"script":   regexp.MustCompile(`(?ims)<script.*?>.*?</script>`),     // javascript
		"noscript": regexp.MustCompile(`(?ims)<noscript.*?>.*?</noscript>`), // javascript
		"style":    regexp.MustCompile(`(?ims)<style.*?>.*?</style>`),       // css
		"link":     regexp.MustCompile(`(?ims)<link.*?>`),                   // css
	}
	ReNewLineBlock = map[string]*regexp.Regexp{
		"<div>": regexp.MustCompile(`(?is)<div.*?>`),
		"<p>":   regexp.MustCompile(`(?is)<p.*?>`),
		"<br>":  regexp.MustCompile(`(?is)<br.*?>`),
		"<hr>":  regexp.MustCompile(`(?is)<hr.*?>`),
		"<li>":  regexp.MustCompile(`(?is)<li.*?>`),
	}
	ReMultiNewLine = regexp.MustCompile(`(?m)\n+`)
	ReTag          = regexp.MustCompile(`(?ims)<.*?>`)
)

func clean(rawhtml string) string {
	lines := strings.Split(rawhtml, "\n")
	for i, _ := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	rawhtml = strings.Join(lines, "\n")
	for _, v := range ReIgnoreBlock {
		rawhtml = v.ReplaceAllString(rawhtml, "")
	}
	for k, v := range ReNewLineBlock {
		rawhtml = v.ReplaceAllString(rawhtml, "\n"+k)
	}
	rawhtml = ReMultiNewLine.ReplaceAllString(rawhtml, "\n")
	text := strings.TrimSpace(ReTag.ReplaceAllString(rawhtml, ""))
	lines = strings.Split(text, "\n")
	for i, _ := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return strings.Join(lines, "\n")
}
