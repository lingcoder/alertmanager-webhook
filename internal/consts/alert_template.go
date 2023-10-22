package consts

const DefaultFeishuTemplate = `
**{#alert-level-status}**: {{ .Severity }} {{ if .IsRecovered }}{#alert-status-recovered}{{else }}{#alert-status-triggered}{{ end }}
**{#alert-rule-name}**: {{ .RuleName }}
{{ if .RuleNotes }}**{#alert-rule-notes}**: {{ .RuleNotes }}{{end }}
**{#alert-monitoring-metrics}**: {{.TagsJSON }}
{{ if .IsRecovered }}**{#alert-recovery-time}**: {{ .LastEvalTime }}{{ else }}**{#alert-trigger-time}**: {{  .TriggerTime }}
**{#alert-trigger-value}**: {{ .TriggerValue }}{{ end }}
**{#alert-sending-time}**: {{ .SendingTime }}`
