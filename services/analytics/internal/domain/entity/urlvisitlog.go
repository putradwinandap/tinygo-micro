package entity

type URLVisitLog struct {
	ID          int    `json:"id"`
	UrlID       string `json:"url_id"`
	IPAdress    string `json:"ip_address"`
	UserAgent   string `json:"user_agent"`
	Referer     string `json:"referer"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}
