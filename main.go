// 爬取豆瓣电影 TOP250
package main

import (
	"douban-movie/model"
	"douban-movie/parse"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 新增数据
func Add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}

// 开始爬取
func Start() {
	var movies []parse.DoubanMovie

	pages := parse.GetPages(BaseUrl)

	client := &http.Client{}

	for _, page := range pages {
		request, err := parse.GetClient(strings.Join([]string{BaseUrl, page.Url}, ""))
		response, err := client.Do(request)
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Println(err)
		}

		movies = append(movies, parse.ParseMovies(doc)...)
	}

	Add(movies)
}

func main() {
	Start()

	defer model.DB.Close()
}
