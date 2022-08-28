{{- define "prudent.labels" -}}
app: {{ .Release.Name }}
deploymentTime: {{ now | unixEpoch | quote }}
{{- end -}}
