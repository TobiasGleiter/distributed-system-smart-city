[Unit]
Description=SmartCity Write Database Service
After=network.target

[Service]
User=pi
WorkingDirectory=/home/pi/studienprojekte/smartcity/server
ExecStart=./main -config=main1.json

Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
