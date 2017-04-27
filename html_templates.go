package main

const (
	tplBase = `<!doctype html>

<html lang="en">
	<head>
		<meta charset="utf-8">
    	<meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width, initial-scale=1">

		<title>DuckDuckGo - Search and Stories</title>
		<meta name="description" content="Desktop Web App for DuckDuckGo Search and Stories">
		
		<style>
			
		</style>

	</head>
	<body>
		{{ block "page" . }}asd{{end}}
	</body>
</html>`

	tplIndex = `{{ define "page" }}
	<h1>The index page</h1>
	{{ with .error }}
		<h3>Some error occurs:</h3>
<pre>
{{ . }}
</pre>
	{{ end }}

	{{ with .news }}
		<h2>List of News:</h2>
		{{ range . }}
		<div>
		<h4>{{ .Title }}</h4>
		<img src="{{ .Image }}" height="300px" width="300px" />
		<a href="{{ .URL }}" title="see article" target="_blank"/>Read Article</a>
		</div>
		{{ end }}
	{{ end }}

	
{{ end }}`

	tplSettings = `{{ define "page" }}
	{{ with .error }}
		<h3>Some error occurs:</h3>
<pre>
{{ . }}
</pre>
	{{ end }}

	{{ $mysources := .mysources }}
	{{ with .sources }}
		<h2>List of Sources:</h2>
		<form action="/settings" method="POST" />
		{{ range . }}
			{{ $id := .ID }}
			<input name="sources" type="checkbox" {{ with inarray $mysources $id }}checked="checked"{{ end }} value="{{ .ID }}" id="source_{{ .ID }}"/>
			<label for="source_{{ .ID }}">{{ .Category }} - {{ .Title }}</label>
			<br />
		{{ end }}
		<input type="submit" value="Save Settings"/>
		</form>
	{{ end }}

	
{{ end }}`
)
