package gofortnite

import (
	"fmt"
	"log"
	"time"
)

type (
	TRNRating struct {
		Label        string  `json:"label"`
		Field        string  `json:"field"`
		Category     string  `json:"category"`
		ValueInt     int64   `json:"valueInt"`
		Value        string  `json:"value"`
		Percentile   float64 `json:"percentile"`
		DisplayValue string  `json:"displayValue"`
	}

	IntPercRankType struct {
		Label        string  `json:"label"`
		Field        string  `json:"field"`
		Category     string  `json:"category"`
		ValueInt     int64   `json:"valueInt"`
		Value        string  `json:"value"`
		Rank         int64   `json:"rank"`
		Percentile   float64 `json:"percentile"`
		DisplayValue string  `json:"displayValue"`
	}

	DecPercRankType struct {
		Label        string  `json:"label"`
		Field        string  `json:"field"`
		Category     string  `json:"category"`
		ValueDec     float64 `json:"valueDec"`
		Value        string  `json:"value"`
		Rank         int64   `json:"rank"`
		Percentile   float64 `json:"percentile"`
		DisplayValue string  `json:"displayValue"`
	}

	IntRankType struct {
		Label        string `json:"label"`
		Field        string `json:"field"`
		Category     string `json:"category"`
		ValueInt     int64  `json:"valueInt"`
		Value        string `json:"value"`
		Rank         int64  `json:"rank"`
		DisplayValue string `json:"displayValue"`
	}

	PlayList struct {
		TRNRating     TRNRating       `json:"trnRating"`
		Score         IntPercRankType `json:"score"`
		Top1          IntRankType
		Top3          IntRankType
		Top5          IntRankType
		Top6          IntRankType
		Top10         IntRankType
		Top12         IntRankType
		Top25         IntRankType
		KD            DecPercRankType `json:"kd"`
		WinRatio      DecRankType     `json:"winRatio"`
		Matches       IntPercRankType `json:"matches"`
		Kills         IntPercRankType `json:"kills"`
		KPG           DecPercRankType `json:"kpg"`
		ScorePerMatch DecPercRankType `json:"scorePerMatch"`
	}

	Stats struct {
		PlayList
	}

	LifeTimeStats struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	RecentMatches struct {
		ID            int64     `json:"id"`
		AccountID     string    `json:"accountId"`
		PlayList      string    `json:"platlist"`
		Kills         int64     `json:"kills"`
		MinutesPlayed int64     `json:"minutesPlayed"`
		Top1          int64     `json:"top1"`
		Top3          int64     `json:"top3"`
		Top5          int64     `json:"top5"`
		Top6          int64     `json:"top6"`
		Top10         int64     `json:"top10"`
		Top12         int64     `json:"top12"`
		Top25         int64     `json"top25"`
		Matches       int64     `json:"matches"`
		DateCollected Time.time `json:"dateCollected"`
		Score         int64     `json:"score"`
		Platform      int64     `json:"platform"`
	}

	FNTApi struct {
		AccountID        string          `json:"accountId"`
		PlatformID       int64           `json:"platformId"`
		PlatformName     string          `json:"platformName"`
		PlatformNameLong string          `json:"platformNameLong"`
		EpicUserHandle   string          `json:"epicUserHandle"`
		Stats            Stats           `json:"stats"`
		LifeTimeStats    []LifeTimeStats `json:"lifeTimeStats"`
		RecentMatches    []RecentMatches `json:"recentMatches"`
	}

	Fortnite struct {
		client   *http.Client
		Platform string
		Token    string
	}
)

const (
	Version   = "0.0.1"
	UserAgent = "gofornite-api-client-" + Version
	Endpoint  = "https://api.fortnitetracker.com/v1/profile/"
)

var (
	baseURL, _ = url.Parse(Endpoint)
)

func NewFortnite(client *http.Client, platform, token string) (*Fortnite, error) {
	if token == "" {
		return nil, fmt.Errorf("[goFornite %s] Please initialize an API token to continue.", Version)
	}

	if platform == "" {
		log.Printf("[goFornite %s] Platform is not initialized, defaulting to \"pc\" as platform.", Version)
		platform = "pc"
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &Fornite{client: client, Platform: platform, Token: token}, nil
}

func (fnt *Fortnite) do(resource string, opts url.Values) (*http.Response, error) {
	u, err := baseURL.Parse(resource)
	if err != nil {
		return nil, err
	}

	target := u.String()
	if opts != nil {
		target = fmt.Sprintf("%s?%s", target, opts.Encode())
	}

	req, err := http.NewRequest("GET", target, nil)
	req.Header.Set("TRN-Api-Key", fnt.Token)
	req.Header.Set("User-Agent", fnt.UserAgent)
	req.Close = true

	return fnt.client.Do(req)
}

func (fnt *Fortnite) GetDetails(epicUserName, platform string) (*FNTApi, error) {
	if platform == "" {
		platform := fnt.Platform
	}
	resource := fmt.Sprintf("%s/%s", platform, epicUserName)

	resp, err := fnt.client.do(resource, nil)
	if err != nil {
		return nil, error
	}
	defer resp.Body.Close()

	var fntAPI *FNTApi
	err = json.NewDecoder(resp.Body).Decode(&fntAPI)
	return fntAPI, err
}
