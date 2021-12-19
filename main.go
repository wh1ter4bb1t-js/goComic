package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mxschmitt/playwright-go"
	"github.com/spf13/viper"
)

type comic struct {
	Url   string `mapstructure:"url"`
	Title string `mapstructure:"title"`
	Date  string `mapstructure:"date"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.config/goComic/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	downloadFolder := viper.GetString("folder")
	configDate := viper.GetString("start_date")

	// set up comics that I want
	var comics []comic
	viper.UnmarshalKey("comics", &comics)

	//get date difference
	today := time.Now()
	startDate, err := time.Parse("2006/01/02", fmt.Sprintf("%s", configDate))
	checkError(err, "failed to parse date")
	difference := today.Sub(startDate)

	//start playwright
	pw, err := playwright.Run()
	checkError(err, "failed to start playwright")

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	checkError(err, "failed to launch browser")

	page, err := browser.NewPage()
	checkError(err, "failed to open a new page in browser")

	//going through the comic and downloading
	for _, comic := range comics {
		comicStartDate, err := time.Parse("2006/01/02", comic.Date)
		checkError(err, "failed to parse date")

		todayComicDate := comicStartDate.Add(difference).Format("2006/01/02")

		todayComicUrl := fmt.Sprintf("%s/%s", comic.Url, todayComicDate)
		if _, err = page.Goto(todayComicUrl); err != nil {
			log.Fatal("failed to open comics website")
		}

		image, err := page.QuerySelector(".item-comic-image img")
		checkError(err, "failed to get image")

		if image != nil {
			src, err := image.GetAttribute("src")
			checkError(err, "failed to get comic image source")

			downloadComic(src, comic.Title, downloadFolder)
		}
	}

}

func checkError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func downloadComic(src string, title string, downloadFolder string) error {
	resp, err := http.Get(src)
	checkError(err, "failed to http call comic image source")
	defer resp.Body.Close()

	os.MkdirAll(downloadFolder, os.ModePerm)
	out, err := os.Create(fmt.Sprintf("%s%s", downloadFolder, title))
	checkError(err, "failed to create file")
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	checkError(err, "failed to copy image to file")

	return nil
}
