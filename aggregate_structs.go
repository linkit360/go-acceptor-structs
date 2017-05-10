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
