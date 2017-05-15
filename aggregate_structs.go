package go_acceptor_structs

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type Aggregate struct {
	ProviderName         string `json:"provider_name,omitempty"`
	OperatorCode         int64  `json:"operator_code,omitempty"`
	ReportAt             int64  `json:"report_at,omitempty"`
	CampaignId           int64  `json:"id_campaign,omitempty"`
	LpHits               int64  `json:"lp_hits,omitempty"`
	LpMsisdnHits         int64  `json:"lp_msisdn_hits,omitempty"`
	MoTotal              int64  `json:"mo,omitempty"`
	MoChargeSuccess      int64  `json:"mo_charge_success,omitempty"`
	MoChargeSum          int64  `json:"mo_charge_sum,omitempty"`
	MoChargeFailed       int64  `json:"mo_charge_failed,omitempty"`
	MoRejected           int64  `json:"mo_rejected,omitempty"`
	Outflow              int64  `json:"outflow,omitempty"`
	RenewalTotal         int64  `json:"renewal,omitempty"`
	RenewalChargeSuccess int64  `json:"renewal_charge_success,omitempty"`
	RenewalChargeSum     int64  `json:"renewal_charge_sum,omitempty"`
	RenewalFailed        int64  `json:"renewal_failed,omitempty"`
	Pixels               int64  `json:"pixels,omitempty"`
}

type AggregateData struct {
	Aggregated []Aggregate `json:"aggregated,omitempty"`
}
