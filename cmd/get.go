package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
    "io"
	"os"
	"text/tabwriter"
	"time"
    "strings"
)

type canteen struct {
	WeekNumber int   `json:"weekNumber"`
	Days       []day `json:"days"`
}

type day struct {
	DayOfWeek string `json:"dayOfWeek"`
	Date      string `json:"date"`
	Menu      []menu `json:"menus"`
}

type menu struct {
	Type string `json:"type"`
	Dish string `json:"menu"`
}

func GetMenu(offset int) {
    bytes := getMenuFromSource(offset)
    actualMenu := canteen{}

    weekDays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

    err := json.Unmarshal(bytes, &actualMenu)

    if err != nil {
        fmt.Println(err)
        return
    }

    w := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)
    fmt.Fprintln(w, "\033[01mDay\tDate\tDish\t\033[0m")

    for i := range len(actualMenu.Days) {
        for j := range len(actualMenu.Days[i].Menu) {
            if actualMenu.Days[i].Menu[j].Type == "Dagens vegetar ret" {
                continue
            }

            theTime, err := time.Parse("2006-01-02T15:04:05", actualMenu.Days[i].Date)

            if err != nil {
                fmt.Println(actualMenu.Days[i].Date)
                fmt.Println(err)
                return
            }

            formattedDate := theTime.Format("02.01.2006")
            formattedNow := time.Now().Format("02.01.2006")

            if formattedNow == formattedDate {
                fmt.Fprintln(w, "\033[32m"+weekDays[i]+"\t"+formattedDate+"\t"+actualMenu.Days[i].Menu[j].Dish+"\t\033[0m")
                break
            }

            fmt.Fprintln(w, "\033[90m"+weekDays[i]+"\t"+formattedDate+"\t"+actualMenu.Days[i].Menu[j].Dish+"\t\033[0m")
        }
    }

    w.Flush()
}

func getMenuFromSource(offset int) []byte {

    menuTime := time.Now().Add(time.Hour * 24 * 7 * time.Duration(offset))
    formattedDate := menuTime.Format("Mon Jan _2 2006")
    formattedDate = strings.ReplaceAll(formattedDate, " ", "%20")
	url := "https://shop.foodandco.dk/api/WeeklyMenu?restaurantId=1234&languageCode=da-DK&date="+formattedDate

	foodClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

    res, getErr := foodClient.Do(req)

	if getErr != nil {
		panic(getErr)
	}

    if res.Body != nil {
        defer res.Body.Close()
    }

    body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		panic(readErr)
	}

	return body
}
