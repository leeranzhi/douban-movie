package parse

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}

var htmlContent = `<div class="paginator">
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
        </div>`

// 获取分页
func GetPages(url string) []Page {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	request, err := GetClient(url)
	res, err := client.Do(request)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return ParsePages(doc)
}

var Headers = map[string]string{
	"Host":                      "movie.douban.com",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer":                   "https://movie.douban.com/top250",
}

func GetClient(baseUrl string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	for index, header := range Headers {
		req.Header.Add(index, header)
	}
	return req, nil
}

// 分析分页
func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	//doc, err := goquery.NewDocumentFromReader(doc.)
	//if err != nil {
	//	log.Fatal(err)
	//}

	doc.Find(".paginator>a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		fmt.Println(s.Text())
		url, _ := s.Attr("href")
		fmt.Println("url", url)

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})
	fmt.Println("pages:", len(pages))
	return pages
}

// 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []DoubanMovie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := DoubanMovie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		log.Printf("i: %d, movie: %v", i, movie)

		movies = append(movies, movie)
	})

	return movies
}
