# SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>
#
# SPDX-License-Identifier: CC0-1.0

.PHONY: default clean all aix dragonfly darwin freebsd illumos netbsd linux openbsd plan9 solaris windows

default:
	go build -o out/go-webring .

clean:
	rm -rf out

# aix, dragonfly, freebsd, netbsd, openbsd, plan9, and windows arm64 are disabled due to library bugs
all: aix darwin dragonfly freebsd  illumos netbsd openbsd plan9 linux solaris windows

aix:
	CGO_ENABLED=0 GOOS=aix GOARCH=ppc64 CGO_ENABLED=0 go build -o out/go-webring-aix-ppc64 .

darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o out/go-webring-darwin-arm64 .

dragonfly:
	GOOS=dragonfly GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-dragonfly-amd64 .

freebsd:
	GOOS=freebsd GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-freebsd-i386 .
	GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-freebsd-amd64 .
	GOOS=freebsd GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-freebsd-arm .

illumos:
	GOOS=illumos GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-illumos-amd64 .

linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-linux-i386 .
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-linux-amd64 .
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-linux-arm .
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o out/go-webring-linux-arm64 .
	GOOS=linux GOARCH=ppc64 CGO_ENABLED=0 go build -o out/go-webring-linux-ppc64 .
	GOOS=linux GOARCH=ppc64le CGO_ENABLED=0 go build -o out/go-webring-linux-ppc64le .
	GOOS=linux GOARCH=mips CGO_ENABLED=0 go build -o out/go-webring-linux-mips .
	GOOS=linux GOARCH=mipsle CGO_ENABLED=0 go build -o out/go-webring-linux-mipsle .
	GOOS=linux GOARCH=mips64 CGO_ENABLED=0 go build -o out/go-webring-linux-mips64 .
	GOOS=linux GOARCH=mips64le CGO_ENABLED=0 go build -o out/go-webring-linux-mips64le .
	GOOS=linux GOARCH=riscv64 CGO_ENABLED=0 go build -o out/go-webring-linux-riscv64 .
	GOOS=linux GOARCH=s390x CGO_ENABLED=0 go build -o out/go-webring-linux-s390x .

netbsd:
	GOOS=netbsd GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-netbsd-i386 .
	GOOS=netbsd GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-netbsd-amd64 .
	GOOS=netbsd GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-netbsd-arm .

openbsd:
	GOOS=openbsd GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-openbsd-i386 .
	GOOS=openbsd GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-openbsd-amd64 .
	GOOS=openbsd GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-openbsd-arm .
	GOOS=openbsd GOARCH=arm64 CGO_ENABLED=0 go build -o out/go-webring-openbsd-arm64 .

plan9:
	GOOS=plan9 GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-plan9-i386 .
	GOOS=plan9 GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-plan9-amd64 .
	GOOS=plan9 GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-plan9-arm .

solaris:
	GOOS=solaris GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-solaris-amd64 .

windows:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o out/go-webring-windows-i386.exe .
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o out/go-webring-windows-amd64.exe .
	GOOS=windows GOARCH=arm CGO_ENABLED=0 go build -o out/go-webring-windows-arm.exe .
	GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -o out/go-webring-windows-arm64.exe .
