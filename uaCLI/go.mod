module github.com/tomseago/uArm/uaCLI

go 1.21.2

replace github.com/tomseago/uArm/slinger => ../slinger

replace github.com/tomseago/uArm/io => ../io

require (
	github.com/eyethereal/go-archercl v0.0.0-20190723143411-30dfa4e6a90f
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/tomseago/uArm/io v0.0.0-20231010061731-deddca7cbe2f
)

require (
	github.com/creack/goselect v0.1.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/visionmedia/go-debug v0.0.0-20180109164601-bfacf9d8a444 // indirect
	go.bug.st/serial v1.6.1 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
)
