# Desktop version of DuckDuckGo - Search and Stories



## Intro
Why this web app?

I really like the app from DuckDuckGo (https://duckduckgo.com/app), especially the part **Top Stories**, but I don't like to read them from my phone/tablet.

So I've create this simple web app (still beta) version to read the feed from my pc.



## Screenshot

#### Stories
![List Stories](/screenshot/stories.png?raw=true)

#### Settings
![Settings](/screenshot/settings.png?raw=true)



## Install

If you've go toolchain installed do:
```
go get -u github.com/tux-eithel/go-search-and-stories
```



### CLI options
```
-p int
    	listen server port (default 8080)
-t
    use html files instead of compiled templates (default false)
```



## Usage

Run the server `go-search-and-stories` and visit http://localhost:8080



## Settings

The app save the settings in your `$HOME` in a file called `.ddg_settings` in plain text.



## Theme

The html and css are inside the file `html_templates.go`... so I guess you are thinking *Why the fuck did he put the template thing inside a go file?!?*.

Well... so when you do `go get ...` you will have an executable with all the stuff inside. In the future maybe I'm going to read external html file.



## Theme Files

Inside `templates` directory the are all the file used to render pages.

Names **must** be keep it unchanged



## Why did you use Go to make this instead of X ?

Well, it's the language I'm learning off-work, so that is!  
I'm learning also Dart and looking a bit of Rust, but it's too early to develop with those languages.



## In the next episodes....

- [ ] Save the last visited items
- [ ] CLI settings: configuration file
- [ ] Import/Export settings
- [ ] Sync settings using someone else's computer (a.k.a **cloud** services)
- [ ] Encrypt setting file
- [ ] Add tests
