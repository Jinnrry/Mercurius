# linux x86_64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/release/linux_x86_64/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/release/linux_x86_64/server ./server/main.go
zip "bin/linux_x86_64.zip" bin/release/linux_x86_64/*

# linux x86
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/release/linux_x86/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/release/linux_x86/server ./server/main.go
zip "bin/linux_x86.zip" bin/release/linux_x86/*

# linux arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/release/linux_arm/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/release/linux_arm/server ./server/main.go
zip "bin/linux_arm.zip" bin/release/linux_arm/*

# linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/release/linux_arm64/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/release/linux_arm64/server ./server/main.go
zip "bin/linux_arm64.zip" bin/release/linux_arm64/*

# linux ppc64
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -o bin/release/linux_ppc64/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -o bin/release/linux_ppc64/server ./server/main.go
zip "bin/linux_ppc64.zip" bin/release/linux_ppc64/*

# linux ppc64le
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -o bin/release/linux_ppc64le/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -o bin/release/linux_ppc64le/server ./server/main.go
zip "bin/linux_ppc64le.zip" bin/release/linux_ppc64le/*

# linux mips
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o bin/release/linux_mips/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o bin/release/linux_mips/server ./server/main.go
zip "bin/linux_mips.zip" bin/release/linux_mips/*

# linux mipsle
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o bin/release/linux_mipsle/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o bin/release/linux_mipsle/server ./server/main.go
zip "bin/linux_mipsle.zip" bin/release/linux_mipsle/*

# linux mips64
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o bin/release/linux_mips64/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o bin/release/linux_mips64/server ./server/main.go
zip "bin/linux_mips64.zip" bin/release/linux_mips64/*

# linux mips64le
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o bin/release/linux_mips64le/client ./client/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o bin/release/linux_mips64le/server ./server/main.go
zip "bin/linux_mips64le.zip" bin/release/linux_mips64le/*

# windows x86_64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/release/windows_x86_64/client.exe ./client/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/release/windows_x86_64/server.exe ./server/main.go
zip "bin/windows_x86_64.zip" bin/release/windows_x86_64/*

# windows x86
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/release/windows_x86/client.exe ./client/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/release/windows_x86/server.exe ./server/main.go
zip "bin/windows_x86.zip" bin/release/windows_x86/*

#mac os x86_64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/release/macos/client ./client/main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/release/macos/server ./server/main.go
zip "bin/macos.zip" bin/release/macos/*