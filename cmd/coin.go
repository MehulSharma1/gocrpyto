/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MehulSharma1/gocrpyto/components"
	"github.com/MehulSharma1/gocrpyto/internal"

	ui "github.com/gizak/termui/v3"
	"github.com/olekukonko/ts"
	"github.com/spf13/cobra"
)

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("coin called")
		id := args[0]
		size, _ := ts.GetSize()
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s?localization=false", id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(fmt.Sprintf("Failed to create request for %s", url))
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Failed to make request for %s", url))
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("Failed to read response body from %s", resp.Body))
		}
		var idResponse internal.CoinResponse
		json.Unmarshal(body, &idResponse)
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=usd&days=1", id)
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			panic(fmt.Sprintf("Failed to create request for %s", url))
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Failed to make request for %s", url))
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("Failed to read response body from %s", resp.Body))
		}
		var history24H internal.HistoricalPrice
		json.Unmarshal(body, &history24H)

		start, end := 0, 0
		title := components.GetTitle(&start, &end, size, idResponse)
		current_price := components.GetCurrentPrice(&start, &end, size, idResponse)
		Low24H := components.Get24HLow(&start, &end, size, idResponse)
		High24H := components.Get24HHigh(&start, &end, size, idResponse)
		ath := components.GetATH(&start, &end, size, idResponse)
		atl := components.GetATL(&start, &end, size, idResponse)
		volume := components.GetVolume(&start, &end, size, idResponse)
		CirculatingSupply := components.GetCirculatingSupply(&start, &end, size, idResponse)
		PriceChange := components.GetPriceChange(&start, &end, size, idResponse)
		Plot24H := components.Get24HPlot(&start, &end, size, history24H)
		ui.Render(title, current_price, Low24H, High24H, ath, atl, volume, CirculatingSupply, PriceChange, Plot24H)

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	},
}

func init() {
	describeCmd.AddCommand(coinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	coinCmd.Flags().String("id", "0", "Id of the coin")
}
