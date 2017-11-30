GOPATH := ${PWD}
export GOPATH

build:
		cd src/app && go fmt
		go build -v app
		sudo systemctl restart cco.golang.service

ensure:
		cd src/app && dep ensure -vendor-only

update:
		cd src/app && dep ensure

test:
		go test -v app

vet:
		go vet ./src/app/...
