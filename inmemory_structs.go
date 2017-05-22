package go_acceptor_structs

type HandShake struct {
	InstanceId string `json:"insance_id,omitempty"` // provider name by instance id
}

type InitResponse struct {
	BlackList []string            `json:"blacklist,omitempty"`
	Campaigns map[string]Campaign `json:"campaigns,omitempty"`
	Services  map[string]Service  `json:"services,omitempty"`
}

type Service struct {
	Id                  string  `json:"id,omitempty"`            // unique id
	Code                string  `json:"code,omitempty"`          // previous service id
	Price               int     `json:"price,omitempty"`         // в целых рублях
	RetryDays           int     `json:"retry_days,omitempty"`    // for retries - days to keep retries, for periodic - subscription is alive
	InactiveDays        int     `json:"inactive_days,omitempty"` // days of unsuccessful charge turns subscription into inactive state
	GraceDays           int     `json:"grace_days,omitempty"`    // days in end of subscription period where always must be charged OK
	PaidHours           int     `json:"paid_hours,omitempty"`    // rejected rule
	DelayHours          int     `json:"delay_hours,omitempty"`   // repeat charge delay
	SMSOnSubscribe      string  `json:"sms_on_unsubscribe,omitempty"`
	SMSOnCharged        string  `json:"sms_on_charged,omitempty"`
	SMSOnUnsubscribe    string  `json:"sms_on_subscribe,omitempty"` // if empty, do not send
	SMSOnContent        string  `json:"sms_on_content,omitempty"`
	SMSOnRejected       string  `json:"sms_on_rejected,omitempty"`
	SMSOnBlackListed    string  `json:"sms_on_blacklisted,omitempty"`
	SMSOnPostPaid       string  `json:"sms_on_postpaid,omitempty"`
	PeriodicAllowedFrom int     `json:"periodic_allowed_from,omitempty"` // send content in sms allowed from and to times.
	PeriodicAllowedTo   int     `json:"periodic_allowed_to,omitempty"`
	PeriodicDays        string  `json:"periodic_days,omitempty"` // days of week to charge subscriber
	MinimalTouchTimes   int     `json:"minimal_touch_times,omitempty"`
	ContentIds          []int64 `json:"content_ids,omitempty"`
}

type Campaign struct {
	Id               string `json:"id,omitempty"`   // UUID
	Code             string `json:"code,omitempty"` // previous id
	Title            string `json:"title,omitempty"`
	Link             string `json:"link,omitempty"`
	Lp               string `json:"lp,omitempty"` // UUID
	Hash             string `json:"hash,omitempty"`
	ServiceCode      string `json:"service_code,omitempty"` // previous service code
	AutoClickRatio   int64  `json:"auto_click_ratio,omitempty"`
	AutoClickEnabled bool   `json:"auto_click_enabled,omitempty"`
	PageSuccess      string `json:"page_success,omitempty"`
	PageError        string `json:"page_error,omitempty"`
	PageThankYou     string `json:"page_thank_you,omitempty"`
	PageWelcome      string `json:"page_welcome,omitempty"`
}

type Content struct {
	Id   int64
	Name string
}
