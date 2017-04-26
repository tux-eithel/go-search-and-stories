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
{{ end }}`

	tplSettings = `{{ define "page" }}
	{{ with .error }}
		<h3>Some error occurs:</h3>
<pre>
{{ . }}
</pre>
	{{ end }}

	{{ with .sources }}
		<h2>List of Sources:</h2>
		{{ range . }}
			<input type="checkbox" value="{{ .ID }}" id="source_{{ .ID }}"/>
			<label for="source_{{ .ID }}">{{ .Category }} - {{ .Title }}</label>
			<br />
		{{ end }}
	{{ end }}

	
{{ end }}`
)
