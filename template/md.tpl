| Pkg | Version | Licenses | Direct | Advisories |
| ---- | ---- | ---- | ---- | ---- |
{{- range $key, $dep := . }}
| {{ $dep.Package }} | {{ $dep.Version }} | {{ $dep.Licenses }} | {{ $dep.Direct }} |
{{- if $dep.Advisories -}}
[Advisories({{len  $dep.Advisories }})]({{ $dep.AdvisoriesURL }})
{{- else -}}
-
{{- end }}|

{{- end }}
