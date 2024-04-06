`env GOOS=linux GOARCH=arm go build -o main main.go`

`scp main3.json pi@192.168.0.23:studienprojekte/smartcity/server`

`ssh pi@192.168.0.21`

`top` <- für Monitoring

Programm als Service mit autostart

Alle Services anzeigen:
`systemctl list-units --type=service`
