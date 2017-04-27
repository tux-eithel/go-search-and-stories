package main

const (
	tplBase = `<!doctype html>

<html lang="en">
	<head>
		<meta charset="utf-8">
    	<meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width, initial-scale=1">

		<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
		 <link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.red-deep_orange.min.css" /> 
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

			.demo-card-square.mdl-card {
				width: 260px;
				height: 290px;
			}
			.demo-card-square > .mdl-card__title {
				color: #fff;
			}
			.mdl-card__action--larger {
				font-size: 20px;
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
					<div class="mdl-card__actions mdl-card--border" style="background: url('{{ .Icon }}') right no-repeat; background-size: 9% auto;">
						<a target="_blank" class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect article-button" href="{{ .URL }}">
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
						<div class="mdl-grid">
							{{ range $feed }}
								<div class="mdl-cell mdl-cell--3-col">
									<div class="demo-card-square mdl-card mdl-shadow--2dp">
										<div class="mdl-card__title mdl-card--expand" style="background: url('{{ .Image }}') top 10% left 5% no-repeat #607D8B">
											<h2 class="mdl-card__title-text">{{ .Title }}</h2>
										</div>
										<div class="mdl-card__supporting-text">
											{{ .Description }}
										</div>
										<div class="mdl-card__actions mdl-card--border mdl-card__action--larger">
											<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="source_{{ .ID }}">
												<input name="sources" value="{{ .ID }}" type="checkbox" id="source_{{ .ID }}" class="mdl-switch__input" {{ with inarray $mysources .ID }}checked{{ end }}>
												<span class="mdl-switch__label"></span>
											</label>
										</div>
									</div>
								</div>
							
							{{ end }}
						</div>	
					</div>
				{{ end }}
			</div>
			<div class="mdl-grid">
  				<div class="mdl-cell mdl-cell--4-col"></div>
  				<div class="mdl-cell mdl-cell--4-col">
					<input type="submit" value="Save Settings" class="mdl-button mdl-js-button mdl-button--raised mdl-button--accent"/>
				</div>
  				<div class="mdl-cell mdl-cell--4-col"></div>
			</div>
		</form>
	{{ end }}

	
{{ end }}`
)
