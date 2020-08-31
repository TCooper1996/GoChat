module github.com/GoChat

go 1.14

require (
	github.com/models v0.0.0
	github.com/views v0.0.0
	github.com/jroimartin/gocui v0.4.0
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
)

replace (
	github.com/models => ./
)