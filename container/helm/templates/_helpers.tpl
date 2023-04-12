{{/*
Expand the name of the chart.
*/}}
{{- define "golang-restapi.name" -}}
{{- default .Chart.Name .Values.golang.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "mongo.name" -}}
{{- default .Chart.Name .Values.mongo.mongonameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "redis.name" -}}
{{- default .Chart.Name .Values.redis.redisnameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "golang-restapi.fullname" -}}
{{- if .Values.golang.fullnameOverride }}
{{- .Values.golang.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.golang.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "mongo.fullname" -}}
{{- if .Values.mongo.mongofullnameOverride }}
{{- .Values.mongo.mongofullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.mongo.mongonameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "redis.fullname" -}}
{{- if .Values.redis.redisfullnameOverride }}
{{- .Values.redis.redisfullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.redis.redisnameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "golang-restapi.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "golang-restapi.labels" -}}
helm.sh/chart: {{ include "golang-restapi.chart" . }}
{{ include "golang-restapi.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "mongo.labels" -}}
helm.sh/chart: {{ include "golang-restapi.chart" . }}
{{ include "mongo.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "redis.labels" -}}
helm.sh/chart: {{ include "golang-restapi.chart" . }}
{{ include "redis.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "golang-restapi.selectorLabels" -}}
app.kubernetes.io/name: {{ include "golang-restapi.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "mongo.selectorLabels" -}}
app.kubernetes.io/name: {{ include "mongo.fullname" . }}
app.kubernetes.io/instance: {{ include "mongo.fullname" . }}
{{- end }}

{{- define "redis.selectorLabels" -}}
app.kubernetes.io/name: {{ include "redis.fullname" . }}
app.kubernetes.io/instance: {{ include "redis.fullname" . }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "golang-restapi.serviceAccountName" -}}
{{- if .Values.golang.serviceAccount.create }}
{{- default (include "golang-restapi.fullname" .) .Values.golang.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.golang.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "redis.serviceAccountName" -}}
{{- if .Values.redis.serviceAccount.create }}
{{- default (include "redis.fullname" .) .Values.redis.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.redis.serviceAccount.name }}
{{- end }}
{{- end }}
