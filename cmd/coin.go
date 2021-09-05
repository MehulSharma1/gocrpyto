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

	"github.com/MehulSharma1/gocrpyto/internal"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Args:  cobra.MaximumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("coin called")
		id := args[0]
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

		header := widgets.NewParagraph()
		header.Text = "Press q to quit, Press h or l to switch tabs"
		header.SetRect(0, 0, 50, 1)
		header.Border = false
		header.TextStyle.Bg = ui.ColorBlue
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
