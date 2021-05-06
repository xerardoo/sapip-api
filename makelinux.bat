@echo on

echo MAKE SAPIP

env GOOS=linux GOARCH=amd64 go build -v -o sapip main.go
MOVE sapip dist/sapip

CP .env.example dist/.env.example

pause
exit
