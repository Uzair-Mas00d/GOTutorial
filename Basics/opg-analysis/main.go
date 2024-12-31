package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"
)

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

func Load(path string) ([]Stock, error) {
	f, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows = slices.Delete(rows, 0, 1)

	var stocks []Stock

	for _, row := range rows {
		ticker := row[0]
		gap, err := strconv.ParseFloat(row[1], 64)

		if err != nil {
			continue
		}

		openingPrice, err := strconv.ParseFloat(row[2], 64)

		if err != nil {
			continue
		}

		stocks = append(stocks, Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}

	return stocks, nil
}

var accountBalance = 10000.0

var lossTolerance = .02

var maxLossPerTrade = accountBalance * lossTolerance

var profitPercent = .8

type Position struct {
	EntryPrice      float64
	Shares          int
	TakeProfitPrice float64
	StopLossPrice   float64
	Profit          float64
}

func Calculate(gapPercent, openingPrice float64) Position {
	closingPrice := openingPrice / (1 + gapPercent)
	gapValue := closingPrice - openingPrice
	profitFromGap := profitPercent * gapValue

	stopLoss := openingPrice - profitFromGap
	takeProfit := openingPrice + profitFromGap

	shares := int(maxLossPerTrade / math.Abs(stopLoss-openingPrice))

	profit := math.Abs(openingPrice-takeProfit) * float64(shares)
	profit = math.Round(profit*100) / 100

	return Position{
		EntryPrice:      math.Round(openingPrice*100) / 100,
		Shares:          shares,
		TakeProfitPrice: math.Round(takeProfit*100) / 100,
		StopLossPrice:   math.Round(stopLoss*100) / 100,
		Profit:          math.Round(profit*100) / 100,
	}
}

type Selection struct {
	Ticker string
	Position
	Articles []Article
}

const (
	url          = "https://seeking-alpha.p.rapidapi.com/news/v2/list-by-symbol?size=5&id="
	apiKeyHeader = "x-rapidapi-key"
	apiKey       = ""
)

type attributes struct {
	PublishOn time.Time `json:"publishOn"`
	Title     string    `json:"title"`
}

type seekingAlphaNews struct {
	Attributes attributes `json:"attributes"`
}

type seekingAlphaResponse struct {
	Data []seekingAlphaNews `json:"data"`
}

type Article struct {
	PublishOn time.Time
	Headline  string
}

func FetchNews(ticker string) ([]Article, error) {
	req, err := http.NewRequest(http.MethodGet, url+ticker, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add(apiKeyHeader, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unsuccessful status code %d recvied", resp.StatusCode)
	}

	res := &seekingAlphaResponse{}

	json.NewDecoder(resp.Body).Decode(res)

	var articles []Article

	for _, item := range res.Data {
		art := Article{
			PublishOn: item.Attributes.PublishOn,
			Headline:  item.Attributes.Title,
		}

		articles = append(articles, art)
	}

	return articles, nil
}

func Deliver(filepath string, selection []Selection) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating file %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(selection)

	if err != nil {
		return fmt.Errorf("error encoding selection %w", err)
	}

	return nil
}

func main() {
	stocks, err := Load("./opg.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	stocks = slices.DeleteFunc(stocks, func(s Stock) bool {
		return math.Abs(s.Gap) < .1
	})

	var selection []Selection

	selectionChannel := make(chan Selection, len(stocks))
	for _, stock := range stocks {

		go func(s Stock, selected chan<- Selection) {

			position := Calculate(s.Gap, s.OpeningPrice)

			articles, err := FetchNews(s.Ticker)

			if err != nil {
				fmt.Printf("error loading news about %s %v", s.Ticker, err)
				return
			} else {
				fmt.Printf("Found %d article about %s", len(articles), s.Ticker)
			}

			sel := Selection{
				Ticker:   s.Ticker,
				Position: position,
				Articles: articles,
			}

			selected <- sel
		}(stock, selectionChannel)

	}

	var selections []Selection

	for sel := range selectionChannel {
		selections = append(selections, sel)
		if len(selections) == len(stocks) {
			close(selectionChannel)
		}
	}

	outputPath := "./opg.json"
	err = Deliver(outputPath, selection)
	if err != nil {
		fmt.Printf("Error writing output %v", err)
		return
	}

	fmt.Printf("Finished writing output to %s\n", outputPath)

}

