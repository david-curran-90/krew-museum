{{/*
Set the name of the deployment
*/}}
{{- define "krewmuseum.name" -}}
{{- $name := default .Release.Name .Values.nameOverride }}
{{- printf "%s" $name -}}
{{- end -}}

{{/*
Set pod annotations
*/}}
{{- define "krewmuseum.annotations" -}}
{{- if .Values.annotations -}}
{{- toYaml .Values.annotations -}}
{{- end -}}
{{- end -}}

{{/*
Set pod labels
*/}}
{{- define "krewmuseum.labels" -}}
app: {{ include "krewmuseum.name" . | quote }}
chart: {{ .Chart.Name | quote }}
heritage: {{ .Release.Service | quote }}
release: {{ .Release.Name | quote }}
{{- if .Values.labels -}}
{{- toYaml .Values.labels -}}
{{- end -}}
{{- end -}}

{{/*
Set the volumes for persistent storage
*/}}
{{- define "krewmuseum.volumes" -}}
{{- end -}}

{{/*
Stateful set persistence
*/}}
{{- define "krewmuseum.volumeClaimTemplate" -}}
- metadata:
    name: "plugin-dir"
  spec:
    accessModes: 
    - {{ .Values.persistence.accessMode | quote }}
    {{- if .Values.persistence.storageClass }}
    {{- if (eq "-" .Values.persistence.storageClass) }}
    storageClassName: ""
    {{- else }}
    storageClassName: {{ .Values.persistence.storageClass | quote }}
    {{- end }}
    {{- end }}
    resources:
      requests:
        storage: {{ .Values.persistence.size | quote }}
{{- end -}}
