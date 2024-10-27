@REM SET CGO_ENABLED=0
@REM SET GOOS=darwin
@REM SET GOARCH=arm64
@REM go build -ldflags="-s -w" -o ./build/yuki-image-drawin-arm64 main.go
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags="-s -w" -o ./build/yuki-image-linux-arm64 main.go
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=arm64
go build -ldflags="-s -w" -o ./build/yuki-image-windows-arm64.exe main.go
@REM SET CGO_ENABLED=0
@REM SET GOOS=darwin
@REM SET GOARCH=amd64
@REM go build -ldflags="-s -w" -o ./build/yuki-image-drawin-amd64 main.go
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./build/yuki-image-linux-amd64 main.go
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./build/yuki-image-windows-amd64.exe main.go

pause>nul