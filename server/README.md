`env GOOS=linux GOARCH=arm go build -o main main.go`

`scp main3.json pi@192.168.0.23:studienprojekte/smartcity/server`

`scp `

`ssh pi@192.168.0.21`

`top` <- für Monitoring

Programm als Service mit autostart

Alle Services anzeigen:
`systemctl list-units --type=service`

--- Remote

`ssh pi@jukebox.dynalias.org -p 12221`

`ssh pi@192.168.180.66`
`ssh pi@192.168.180.67`

PW: `Vs24!DhWb20!`
192.169.180.65

`scp -P 12221 main pi@jukebox.dynalias.org:studienprojekte/smartcity/server`

`scp -P 12221 main.py pi@jukebox.dynalias.org:studienprojekte/smartcity/server`
`scp -P 12221 sensor_config.json pi@jukebox.dynalias.org:studienprojekte/smartcity/server`
`scp -P 12221 sensors.py pi@jukebox.dynalias.org:studienprojekte/smartcity/server`

`scp -P 12221 main1.json pi@jukebox.dynalias.org:studienprojekte/smartcity/server`
`scp -P 12221 main2.json pi@jukebox.dynalias.org:studienprojekte/smartcity/server`
`scp -P 12221 main3.json pi@jukebox.dynalias.org:studienprojekte/smartcity/server`

`scp main pi@192.168.180.66:studienprojekte/smartcity/server`
`scp main pi@192.168.180.67:studienprojekte/smartcity/server`

`scp main2.json pi@192.168.180.66:studienprojekte/smartcity/server`
`scp main3.json pi@192.168.180.67:studienprojekte/smartcity/server`

---

`sudo systemctl daemon-reload`
`sudo systemctl start server-smartcity.service`

`cd /home/pi/studienprojekte/smartcity/server`
`cd /etc/systemd/system`

`./main -config=main2.json`

---

`curl -X GET 192.168.180.66:23311/bully/health`
`curl -X GET 10.8.0.1:23312/bully/health`

pw: abc14az

# Locust

`scp -P 12221 simulation.py pi@jukebox.dynalias.org:studienprojekte/smartcity/server`

Run the Locust script on the .65 Pi:

`locust -P 9025 -f simulation.py`

Mit Locust Dashboard verbinden:

`http://10.8.0.1:9025`

Applikationen auf Pis starten

Locust test starten:

- New:
- Host: `http://192.168.180.67:23312`
- Anzahl User auswählen und starten.
