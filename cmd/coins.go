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
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/olekukonko/ts"
	"github.com/spf13/cobra"
)

// coinsCmd represents the coins command
var coinsCmd = &cobra.Command{
	Use:   "coins",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := ts.GetSize()
		url := "https://api.coingecko.com/api/v3/coins/list"
		client := &http.Client{
			Timeout: time.Second * 3,
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}
		search_string, _ := cmd.Flags().GetString("filter")

		resp, errors := client.Do(req)
		if errors != nil {
			panic(errors)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		type Response struct {
			Id     string `json:"id"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
		}
		var respStruct []Response

		json.Unmarshal([]byte(body), &respStruct)
		var rows []string

		for index, w := range respStruct {
			if strings.Contains(w.Id, search_string) || strings.Contains(w.Name, search_string) || strings.Contains(w.Symbol, search_string) {
				row := fmt.Sprintf("[%d] id : %s, name: %s, symbol: %s", index, w.Id, w.Name, w.Symbol)
				rows = append(rows, row)
			}
		}
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		l := widgets.NewList()
		l.Title = "List"

		l.Rows = rows
		l.TextStyle = ui.NewStyle(ui.ColorYellow)
		l.WrapText = false
		l.SetRect(0, 0, size.Col(), size.Row())

		ui.Render(l)

		previousKey := ""
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				l.ScrollDown()
			case "k", "<Up>":
				l.ScrollUp()
			case "<C-d>":
				l.ScrollHalfPageDown()
			case "<C-u>":
				l.ScrollHalfPageUp()
			case "<C-f>":
				l.ScrollPageDown()
			case "<C-b>":
				l.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					l.ScrollTop()
				}
			case "<Home>":
				l.ScrollTop()
			case "G", "<End>":
				l.ScrollBottom()
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

			ui.Render(l)
		}

	},
}

func init() {
	getCmd.AddCommand(coinsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coinsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	coinsCmd.Flags().String("filter", "", "rew")
}
