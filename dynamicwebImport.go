package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type DynamicwebPageResponse struct {
	Data []DynamicwebPage `json:"data"`
}
type DynamicwebPage struct {
	ID   int    `json:"Id"`
	Slug string `json:"Slug"`
	Name string `json:"Name"`
}
type HugoPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (dest *HugoPage) mapDynamicwebPage(
	src DynamicwebPage,
	dynamicwebHostURL string,
	client *http.Client) {
	dest.ID = src.ID
	dest.Title = src.Name
	dest.URL = src.Slug
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dynamicwebHostURL, ok := os.LookupEnv("DYNAMICWEB_HOST_URL")
	if !ok || dynamicwebHostURL == "" {

	}
	var productsEndpoint = dynamicwebHostURL + "/api/getallpages"

	// Fetch pages
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // this line should be removed in production
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest(http.MethodGet, pagesEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Parse response JSON
	var pages DynamicwebPageResponse
	if err = json.Unmarshal(body, &pages); err != nil {
		log.Fatal(err)
	}
	// Clear content/product directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.RemoveAll(dir + "/content/page"); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(dir+"/content/page", 0777); err != nil {
		log.Fatal(err)
	}
	// Write product content files
	for _, sourcePage := range pages.Data {
		var destPage = HugoPage{}
		destPage.mapDynamicwebPage(sourcePage, dynamicwebHostURL, client)
		entryJSON, err := json.MarshalIndent(destPage, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(dir + "/content/page/" + destPage.URL + ".md")
		if err != nil {
			log.Fatal(err)
		}
		writer := bufio.NewWriter(file)
		writer.WriteString(string(entryJSON) + "\n")
		//writer.WriteString("\n")
		//writer.WriteString(destPage.Title)
		writer.Flush()
		file.Close()
	}
}
