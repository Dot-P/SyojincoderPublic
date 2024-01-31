package main

import (
	"fmt"
	"encoding/csv"
    "log"
    "os"
	"time"
	"github.com/gocolly/colly"
	"regexp"
)

func main() {

	fmt.Println("Start scraping...");

	// create struct for inputing sone information
	type Data struct {
		Contest, Content string
	}

	var dataes []Data
	
	// create the object
	c := colly.NewCollector()

	// add user agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// access to the url
	// url is like this (From 20, the last letter is a)

	// https://atcoder.jp/contests/abc001/tasks/abc001_1
	// https://atcoder.jp/contests/abc020/tasks/abc020_a
	// https://atcoder.jp/contests/abc328/tasks/abc328_a

	c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

	// execute
	// ABC001 A; section(id:tast-statement) -> section -> no rule:(
	// ABC328 A; div(class:part) -> section -> p(s) (ABC10 ~)
	for j := 0; j < 4; j++ {

		diff := ""
		if j == 0 {
			diff = "a"
		} else if j == 1 { 
			diff = "b"
		} else if j == 2 { 
			diff = "c"
		} else if j == 3 { 
			diff = "d"
		}

		for i := 20; i < 329; i++ {

			if i == 316{
				continue
			}

			if i >= 42 && i <= 60 && (j == 2 || j == 3) {
				continue
			}

			content := ""
			c.OnHTML("div.part section", func(e *colly.HTMLElement) {
				e.ForEach("h3, pre, div, p, ul > li", func(_ int, el *colly.HTMLElement) { 
					text := el.Text
					content += text + " "
				})
			})

			result := fmt.Sprintf("%03d%s", i, diff)

			err := c.Visit("https://atcoder.jp/contests/abc" + fmt.Sprintf("%03d", i) + "/tasks/abc"+ fmt.Sprintf("%03d", i) +"_" + diff)
			if err != nil {
				fmt.Println(result + " is Fatal")
				continue
			}

			c.Wait()

			space := regexp.MustCompile(`\s+`)
			content = space.ReplaceAllString(content, " ")

			fmt.Println(result + " is success")

			dataes = append(dataes, Data{Contest: diff + fmt.Sprintf("%03d", i), Content: content})
			
			time.Sleep(3 * time.Second) 

		}
	}

	// open the output CSV file
	file, err := os.Create("problem_d.csv")

	if err != nil {
	log.Fatalln("Failed to create the output CSV file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// add the header row to the CSV
	headers := []string{
		"Name",
		"Content",
	}
	writer.Write(headers)

	for _, data := range dataes {

	record := []string{
		data.Contest,
		data.Content,
	}
	
	writer.Write(record)
	}

	fmt.Println("Done!");

}