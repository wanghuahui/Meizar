package rule

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// JandanRule struct
type JandanRule struct{}

// URLRule return the url you want
func (p *JandanRule) URLRule() (url string) {
	return "http://jandan.net/ooxx/"
}

// PageRule page num
func (p *JandanRule) PageRule(currentPage int) (page string) {
	return "page-" + strconv.Itoa(currentPage)
}

// ImageRule the rule you set
func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
