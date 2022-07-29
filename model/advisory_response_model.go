package model

type AdvisoryResponseModel struct {
	AdvisoryQueryModel
	RuleSet []*AdvisoryLayer `json:"rule_set,omitempty"`
}

type AdvisoryLayer struct {
	Name          string                  `json:"name,omitempty"`
	Details       *map[string]interface{} `json:"details,omitempty"`
	QueryDuration int                     `json:"query_duration,omitempty"`
}
