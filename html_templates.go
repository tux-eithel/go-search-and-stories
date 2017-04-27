package main

const (
	tplBase = `<!doctype html>

<html lang="en">
	<head>
		<meta charset="utf-8">
    	<meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width, initial-scale=1">

		<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
		<link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.red-light_blue.min.css" /> 
		<script defer src="https://code.getmdl.io/1.3.0/material.min.js"></script>

		<title>DuckDuckGo - Search and Stories</title>
		<meta name="description" content="Desktop Web App for DuckDuckGo Search and Stories">
		
		<style>
			.demo-card-wide.mdl-card {
			width: 512px;
			}
			.demo-card-wide > .mdl-card__title {
				color: #fff;
				height: 176px;
			}
			.demo-card-wide > .mdl-card__menu {
			color: #fff;
			}
		</style>

	</head>
	<body>
		<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
			<header class="mdl-layout__header">
				<div class="mdl-layout__header-row">
				<!-- Title -->
				<span class="mdl-layout-title">Desktop DuckDuckGo - Search and Stories</span>
				</div>
				<!-- Tabs -->
				<div class="mdl-layout__tab-bar mdl-js-ripple-effect">
				{{ block "menu" . }}{{ end }}
				</div>
			</header>
			
			<main class="mdl-layout__content">
				{{ block "page" . }}{{end}}
			</main>
		</div>
	</body>
</html>`

	tplIndex = `
{{ define "menu" }}
	<a href="/" class="mdl-layout__tab is-active">News</a>
	<a href="/settings" class="mdl-layout__tab">Settings</a>
{{ end }}
	
{{ define "page" }}
	{{ with .error }}
		<h3>Some error occurs:</h3>
<pre>
{{ . }}
</pre>
	{{ end }}

	{{ with .news }}
		<div class="mdl-grid">
		{{ range . }}
  			<div class="mdl-cell mdl-cell--4-col">
				<div class="demo-card-wide mdl-card mdl-shadow--2dp">
					<div class="mdl-card__title" style="background: url('{{ .Image }}') center / cover;">
						<h2 class="mdl-card__title-text">{{ .Category }}</h2>
					</div>
					<div class="mdl-card__supporting-text">
						{{ .Title }}
					</div>
					<div class="mdl-card__actions mdl-card--border">
						<a target="_blank" class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" href="{{ .URL }}">
						Read Article
						</a>
					</div>
				</div>
			</div>
		{{ end }}
		</div>
	{{ end }}

	
{{ end }}`

	tplSettings = `

{{ define "menu" }}
	<a href="/" class="mdl-layout__tab">News</a>
	<a href="/settings" class="mdl-layout__tab is-active">Settings</a>
{{ end }}

{{ define "page" }}
	{{ with .error }}
		<h3>Some error occurs:</h3>
<pre>
{{ . }}
</pre>
	{{ end }}

	{{ $mysources := .mysources }}
	{{ with .sources }}
		<form action="/settings" method="POST" />
			<div class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">
				<div class="mdl-tabs__tab-bar">
					{{ range $cat, $feed := . }}
					<a href="#{{ validid $cat }}" class="mdl-tabs__tab {{ if eq $cat "Answers" }}is-active{{ end }}">{{ $cat }}</a>
					{{ end }}
				</div>
				{{ range  $cat, $feed := . }}
					<div class="mdl-tabs__panel {{ if eq $cat "Answers" }}is-active{{ end }}" id="{{ validid $cat }}">
						{{ range $feed }}
						<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="source_{{ .ID }}">
							<input name="sources" value="{{ .ID }}" type="checkbox" id="source_{{ .ID }}" class="mdl-switch__input" {{ with inarray $mysources .ID }}checked{{ end }}>
							<span class="mdl-switch__label">{{ .Title }}</span>
						</label>
					{{ end }}
					</div>
				{{ end }}
			</div>
		<input type="submit" value="Save Settings" class="mdl-button mdl-js-button mdl-button--raised mdl-button--accent"/>
		</form>
	{{ end }}

	
{{ end }}`
)
