package functions

import (
	"encoding/json"
	"html/template"
	"io"
	"movieapp/schema"
	"net/http"
	"strings"
)


func fetchRequest(key, host, url string) ([]schema.Show, error) {
	req, _:= http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiResp schema.APIResponse
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return nil, err
	}

	return apiResp.Shows, nil
}

func HomePage(w http.ResponseWriter, r *http.Request, key, host string) {
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "us"
	}

	shows, err := fetchRequest(key, host, "https://streaming-availability.p.rapidapi.com/shows/search/filters?country="+country+"&series_granularity=episode&order_direction=asc&order_by=original_title&genres_relation=and&output_language=en")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
		"eq": func(a, b string) bool { return a == b },
	})
	tmpl, err = tmpl.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := schema.HomePageData{
		PageTitle: "Movie Explorer",
		Shows: shows,
		CurrentCountry: country,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}