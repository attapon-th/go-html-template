# go-htmx-template


## Prerequisites

1. go version 1.20 or higher - [go.dev](https://go.dev)
2. task version 3.0 or higher - [taskfile.dev](https://taskfile.dev)
3. templ - for use command: `teml generate` [https://github.com/a-h/templ](https://github.com/a-h/templ)
4. (Optional) cobra-cli - [https://github.com/spf13/cobra](https://github.com/spf13/cobra)


# How to use

1. Clone this repo
2. Replace the text `go-htmx-template` with the name of `<your go module>` in all file if exists 
3. copy `example.env` to `.env`
4. run `task mod gen`
5. Start your server `task dev`
6. Open: [http://localhost:8888/go-htmx](http://localhost:8888/go-htmx) (config in `.env` file)