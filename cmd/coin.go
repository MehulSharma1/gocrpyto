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
		fmt.Println(idResponse)
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		start, end := 0, 0
		title := components.GetTitle(&start, &end, size, idResponse)
		current_price := components.GetCurrentPrice(&start, &end, size, idResponse)
		Low24H := components.Get24HLow(&start, &end, size, idResponse)
		High24H := components.Get24HHigh(&start, &end, size, idResponse)
		ath := components.GetATH(&start, &end, size, idResponse)
		atl := components.GetATL(&start, &end, size, idResponse)
		// p2 := widgets.NewParagraph()
		// p2.Title = "Multiline"
		// p2.Text = "Simple colored text\nwith label. It [can be](fg:red) multilined with \\n or [break automatically](fg:red,fg:bold)"
		// p2.SetRect(0, 5, 35, 10)
		// p2.BorderStyle.Fg = ui.ColorYellow

		// p3 := widgets.NewParagraph()
		// p3.Title = "Auto Trim"
		// p3.Text = "Long text with label and it is auto trimmed."
		// p3.SetRect(0, 10, 40, 15)

		// p4 := widgets.NewParagraph()
		// p4.Title = "Text Box with Wrapping"
		// p4.Text = "Press q to QUIT THE DEMO. [There](fg:blue,mod:bold) are other things [that](fg:red) are going to fit in here I think. What do you think? Now is the time for all good [men to](bg:blue) come to the aid of their country. [This is going to be one really really really long line](fg:green) that is going to go together and stuffs and things. Let's see how this thing renders out.\n    Here is a new paragraph and stuffs and things. There should be a tab indent at the beginning of the paragraph. Let's see if that worked as well."
		// p4.SetRect(40, 0, 70, 20)
		// p4.BorderStyle.Fg = ui.ColorBlue

		ui.Render(title, current_price, Low24H, High24H, ath, atl)

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
