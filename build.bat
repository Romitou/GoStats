cmd /C "set GOOS=linux&& set GOARCH=arm&& go build -o builds/panoramix-linux-arm"
git update-index --chmod=+x builds/panoramix-linux-arm
cmd /C "set GOOS=windows&& set GOARCH=amd64&& go build -o builds/panoramix-win-amd64.exe"
git update-index --chmod=+x builds/panoramix-win-amd64.exe