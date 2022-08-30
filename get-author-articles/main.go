package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getArticleTitles(author string) []string {
	type Article struct {
		Title      string `json:"title"`
		Author     string `json:"author"`
		StoryTitle string `json:"story_title"`
	}

	type ArticleResponse struct {
		Page       int       `json:"page"`
		PerPage    int       `json:"per_page"`
		Total      int       `json:"total"`
		TotalPages int       `json:"total_pages"`
		Data       []Article `json:"data"`
	}

	limit := 10
	page := 0
	var articles []string
	for page < limit {
		url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?author=%s&page=%d", author, page)
		fmt.Println(url)

		response := new(ArticleResponse)
		getJson(url, response)
		limit = response.TotalPages
		fmt.Println(fmt.Sprintf("%v", response))
		for _, d := range response.Data {
			if d.Author == author {
				if d.Title != "" {
					articles = append(articles, d.Title)
				} else if d.StoryTitle != "" {
					articles = append(articles, d.StoryTitle)
				}
			}
		}
		page++
	}

	return articles
}

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return json.Unmarshal(body, target)
}

func main() {
	getArticleTitles("epaga")
}
