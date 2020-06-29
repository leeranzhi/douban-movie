package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	html := `
	<div class="paginator">
        <span class="prev">
            <link rel="prev" href="?start=75&amp;filter=">
            <a href="?start=75&amp;filter=">&lt;前页</a>
        </span>
        
        

                
            <a href="?start=0&amp;filter=">1</a>
        
                
            <a href="?start=25&amp;filter=">2</a>
        
                
            <a href="?start=50&amp;filter=">3</a>
        
                
            <a href="?start=75&amp;filter=">4</a>
        
                <span class="thispage">5</span>
                
            <a href="?start=125&amp;filter=">6</a>
        
                
            <a href="?start=150&amp;filter=">7</a>
        
                
            <a href="?start=175&amp;filter=">8</a>
        
                
            <a href="?start=200&amp;filter=">9</a>
        
                
            <a href="?start=225&amp;filter=">10</a>
        
        <span class="next">
            <link rel="next" href="?start=125&amp;filter=">
            <a href="?start=125&amp;filter=">后页&gt;</a>
        </span>

            <span class="count">(共250条)</span>
        </div>
`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".paginator>a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Attr("href"))
	})
}
