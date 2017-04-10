package go_acceptor_structs

type Response struct{}

type Aggregate struct {
	ReportAt     int64  `json:"report_at,omitempty"`
	CampaignId   int64  `json:"id_campaign,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	OperatorCode int64  `json:"operator_code,omitempty"`
	LpHits       int64  `json:"lp_hits,omitempty"`
	LpMsisdnHits int64  `json:"lp_msisdn_hits,omitempty"`
	Mo           int64  `json:"mo,omitempty"`
	MoUniq       int64  `json:"mo_uniq,omitempty"`
	MoSuccess    int64  `json:"mo_success,omitempty"`
	RetrySuccess int64  `json:"retry_success,omitempty"`
	Pixels       int64  `json:"pixels,omitempty"`
}

type AggregateData struct {
	Aggregated []Aggregate `json:"aggregated,omitempty"`
}

/*
	Example:

	ProviderName: cheese
	Time: time.Now().Format("2006-01-02")
*/
type BlackListGetParams struct {
	ProviderName string `json:"provider_name,omitempty"`
	Time         string `json:"time,omitempty"`
}

type BlackListAddParams struct {
	ProviderName string   `json:"provider_name,omitempty"`
	Msisdns      []string `json:"msisdns,omitempty"`
}

type BlackListResponse struct {
	Msisdns []string `json:"msisdns,omitempty"`
}

type CampaignsGetParams struct {
	Country uint `json:"id_country,omitempty"`
}

type CampaignsResponse struct {
	Campaigns []CampaignsCampaign `json:"campaigns,omitempty"`
}

type CampaignsCampaign struct {
	Id    string `json:"id,omitempty"` // UUID
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}
