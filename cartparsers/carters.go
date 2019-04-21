package cartparsers

import (
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const (
	cartersUrl = "https://carters.com"
	userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"

	cartersCartName  = "#cart-items-form"
	cartersAttrsName = ".mini-cart-attributes"

	cartersProductName   = ".product-name"
	cartersProductStyle  = ".product-number"
	cartersProductSize1  = ".Size .label"
	cartersProductSize2  = ".Size .value"
	cartersProductColor1 = ".Color .label"
	cartersProductColor2 = ".Color .value"
	cartersProductQty    = ".qtyCount .value"
	cartersProductPrice  = ".price .alt-price"
)

// cartersParser implements CartParser for Carters shop
type cartersParser struct {
	result []CartItem
}

// NewCartersParser creates new Carter's parser
func NewCartersParser() CartParser {
	return &cartersParser{}
}

// Parse parses Carters shopping cart
func (p *cartersParser) Parse(reader io.Reader) ([]CartItem, error) {
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create document")
	}

	sel := document.Find(cartersCartName)
	if sel == nil {
		return nil, errors.New("invalid HTML - no cart object")
	}

	attr := sel.Find(cartersAttrsName).Each(func(index int, product *goquery.Selection) {
		p.parseProduct(product)
	})
	if attr == nil {
		return nil, errors.New("invalid HTML - no attributes object")
	}

	return p.result, nil
}

// parseProduct parse one product
func (p *cartersParser) parseProduct(product *goquery.Selection) {
	cartItem := CartItem{
		ProductUrl:   p.getUrl(product, cartersProductName),
		ProductName:  p.getSimpleText(product, cartersProductName),
		ProductStyle: p.getSimpleText(product, cartersProductStyle),
		ProductSize:  p.getMergeText(product, cartersProductSize1, cartersProductSize2),
		ProductColor: p.getMergeText(product, cartersProductColor1, cartersProductColor2),
		ProductQty:   p.getSimpleInt(product, cartersProductQty),
		ProductPrice: p.getPrice(product, cartersProductPrice),
	}
	cartItem.ProductPrice /= float64(cartItem.ProductQty)

	if len(cartItem.ProductName) > 0 {
		if p.result == nil {
			p.result = make([]CartItem, 0)
		}
		p.result = append(p.result, cartItem)
	}
}

// getUrl retrieve product's URL
func (p *cartersParser) getUrl(product *goquery.Selection, selector string) string {
	node := product.Find(selector)
	if node != nil {
		href := node.Find("a")
		if href != nil {
			url, exists := href.Attr("href")
			if exists {
				return p.actualUrl(cartersUrl + url)
			}
		}
	}
	return ""
}

// actualUrl retrieve actual product's URL
func (p *cartersParser) actualUrl(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}

	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	return resp.Request.URL.Scheme + "://" + resp.Request.URL.Host + resp.Request.URL.Path
}

// getSimpleText retrieve simple text case
func (p *cartersParser) getSimpleText(product *goquery.Selection, selector string) string {
	node := product.Find(selector)
	if node != nil {
		return p.innerText(node)
	}
	return ""
}

// getSimpleInt retrieve simple integer case
func (p *cartersParser) getSimpleInt(product *goquery.Selection, selector string) int {
	value, err := strconv.Atoi(p.getSimpleText(product, selector))
	if err != nil {
		return 0
	}
	return value
}

// getPrice retrieve price
func (p *cartersParser) getPrice(product *goquery.Selection, selector string) float64 {
	text := p.getSimpleText(product, selector)
	text = strings.ReplaceAll(text, "$", "")
	price, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0
	}
	return price
}

// getMergeText retrieve text case with merging
func (p *cartersParser) getMergeText(product *goquery.Selection, selector1 string, selector2 string) string {
	return p.getSimpleText(product, selector1) + " " + p.getSimpleText(product, selector2)
}

// innerText retrieve inner text for HTML node
func (p *cartersParser) innerText(node *goquery.Selection) string {
	if node != nil {
		text := node.Text()
		text = strings.ReplaceAll(text, "\n", " ")
		text = strings.ReplaceAll(text, "\t", " ")

		space := regexp.MustCompile(`\s+`)
		text = space.ReplaceAllString(text, " ")
		return strings.Trim(text, " \n\t")
	}
	return ""
}
