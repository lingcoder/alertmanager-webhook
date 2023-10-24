package model

type WebhookAlert struct {
	Severity     string
	IsRecovered  bool
	RuleName     string
	RuleNotes    string
	Description  string
	TagsJSON     string
	LastEvalTime string
	TriggerTime  string
	TriggerValue string
	SendingTime  string
}
