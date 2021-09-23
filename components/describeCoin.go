package components

import (
	"fmt"

	"github.com/MehulSharma1/gocrpyto/internal"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/olekukonko/ts"
)

func GetTitle(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {
	p1 := widgets.NewParagraph()
	p1.Text = fmt.Sprintf("[Details of %s](fg:red)", idResponse.Name)
	*end += 3
	p1.SetRect(0, *start, size.Col(), *end)
	p1.Border = false
	return p1
}

func GetCurrentPrice(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {
	p2 := widgets.NewParagraph()
	p2.Title = "Current Price"
	p2.Text = fmt.Sprintf("USD %f",
		idResponse.MarketData.CurrentPrice.Usd)
	*start = *end
	*end += 5
	p2.SetRect(0, *start, size.Col()/3, *end)
	return p2
}
func Get24HLow(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p3 := widgets.NewParagraph()
	p3.Title = "24 Hour Low"
	p3.Text = fmt.Sprintf("USD %f",
		idResponse.MarketData.Low24H.Usd)
	p3.SetRect(size.Col()/3+1, *start, size.Col()/3*2, *end)
	return p3
}
func Get24HHigh(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p4 := widgets.NewParagraph()
	p4.Title = "24 Hour Low"
	p4.Text = fmt.Sprintf("USD %f",
		idResponse.MarketData.High24H.Usd)
	p4.SetRect(size.Col()/3*2+1, *start, size.Col()/3*3, *end)
	return p4
}

func GetATH(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p3 := widgets.NewParagraph()
	p3.Title = "All Time High"
	p3.Text = fmt.Sprintf("USD %f on %s up %.2f%% from current price",
		idResponse.MarketData.Ath.Usd, idResponse.MarketData.AthDate["usd"], idResponse.MarketData.AthChangePercentage.Usd)
	*start = *end
	*end += 5
	p3.SetRect(0, *start, size.Col()/2, *end)
	return p3
}
func GetATL(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p4 := widgets.NewParagraph()
	p4.Title = "7 Day High"
	p4.Text = fmt.Sprintf("USD %f on %s down %.2f%% from current price",
		idResponse.MarketData.Atl.Usd, idResponse.MarketData.AtlDate["usd"], idResponse.MarketData.AtlChangePercentage.Usd)
	p4.SetRect(size.Col()/2+1, *start, size.Col(), *end)
	return p4
}

func GetVolume(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p4 := widgets.NewParagraph()
	p4.Title = "Total Volume"
	*start = *end
	*end += 5
	p4.Text = fmt.Sprintf("USD %f", idResponse.MarketData.TotalVolume.Usd)
	p4.SetRect(0, *start, size.Col()/3, *end)
	return p4
}

func GetCirculatingSupply(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p4 := widgets.NewParagraph()
	p4.Title = "Circulating Supply"
	p4.Text = fmt.Sprintf("%f %s",
		idResponse.MarketData.CirculatingSupply, idResponse.Symbol)
	p4.SetRect(size.Col()/3+1, *start, size.Col()/3*2, *end)
	return p4
}
func GetPriceChange(start, end *int, size ts.Size, idResponse internal.CoinResponse) *widgets.Paragraph {

	p4 := widgets.NewParagraph()
	p4.Title = "24 Hour Price Change"
	p4.Text = fmt.Sprintf("%f", idResponse.MarketData.PriceChangePercentage24H)
	p4.SetRect(size.Col()/3*2+1, *start, size.Col(), *end)
	return p4
}

func Get24HPlot(start, end *int, size ts.Size, history24H internal.HistoricalPrice) *widgets.Plot {

	plot24H := func() [][]float64 {
		data := make([][]float64, 1)
		// var x_axis []float64
		var y_axis []float64
		for _, val := range history24H.Prices {
			// x_axis = append(x_axis, val[0])
			y_axis = append(y_axis, val[0])
		}
		// data[0] = x_axis
		fmt.Println(len(data[0]))

		data[0] = y_axis
		return data
	}()
	*start = *end
	*end += 25
	p0 := widgets.NewPlot()
	p0.Title = "24H Prices"
	p0.HorizontalScale = 1
	p0.Data = plot24H
	p0.SetRect(0, *start, size.Col(), *end)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen
	return p0
}
