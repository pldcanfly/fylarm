package components

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//	HourlyUnits struct {
//			Time string `json:"time"`
//			Temp string `json:"temperature_2m"`
//			Change string
//
//		} `json:"hourly_units"`
type ApiResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Current   struct {
		Temp float32 `json:"temperature_2m"`
	} `json:"current"`
	Hourly struct {
		Time        [72]string  `json:"time"`
		Temperature [72]float32 `json:"temperature_2m"`
		RainChance  [72]float32 `json:"percipitation_propability"`
		RainAmount  [72]float32 `json:"precipitation"`
	} `json:"hourly"`
}

type TimePoint struct {
	Time       time.Time
	Temp       float32
	RainChance float32
}

func GetWeather() *widget.Label {
	data, err := fetchWeather()
	if err != nil {
		log.Println("get weather:", err)
	}

	if err != nil {
		log.Println("parsing error:", err)
	}

	// Get 8:00 14:00 20:00
	fmt.Println(data.Current.Temp)
	ws := fmt.Sprintf("Aktuell: %3.0f °C\n", data.Current.Temp)
	now := time.Now()

	cnt := 0
	for i, v := range data.Hourly.Time {
		time, _ := time.Parse("2006-01-02T15:04", v)
		h := time.Hour()

		if h == 8 || h == 14 || h == 20 {
			if now.Before(time) {

				ws += "Heute "
			} else {
				ws += "Morgen "
			}

			ws += fmt.Sprintf("%02d:00 Uhr   %3.0f °C", h, data.Hourly.Temperature[i])
			if data.Hourly.RainChance[i] > 0 {
				ws += fmt.Sprintf(" - %.0f%% %.0fmm", data.Hourly.RainChance[i], data.Hourly.RainAmount[i])
			}
			ws += "\n"

			cnt++
			if cnt >= 3 {
				break
			}

		}
	}

	w := widget.NewLabel(ws)
	w.Alignment = fyne.TextAlignCenter

	return w
}

func fetchWeather() (*ApiResponse, error) {
	// https://open-meteo.com/en/docs/#latitude=47.0667&longitude=15.45&hourly=temperature_2m,precipitation_probability,precipitation&timezone=Europe%2FBerlin

	url := "https://api.open-meteo.com/v1/forecast?latitude=47.0667&longitude=15.45&current=temperature_2m&hourly=temperature_2m,precipitation_probability,precipitation&timezone=Europe%2FBerlin&forecast_days=3"

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch api: %v", err)
	}
	defer res.Body.Close()
	var v *ApiResponse

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading bytes: %v", err)
	}
	err = json.Unmarshal(r, &v)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling:: %v", err)
	}

	return v, nil
}
