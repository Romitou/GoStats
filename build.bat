cmd /C "set GOOS=linux&& set GOARCH=arm&& go build -o builds/gostats-linux-arm"
git update-index --chmod=+x builds/gostats-linux-arm
cmd /C "set GOOS=windows&& set GOARCH=amd64&& go build -o builds/gostats-win-amd64.exe"
git update-index --chmod=+x builds/gostats-win-amd64.exe