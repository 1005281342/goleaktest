build:
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/leak
	chmod +x ./bin/leak


dockerbuild:
	docker build  -t ccr.ccs.tencentyun.com/oyjx/leak:v1.0.0 .

dockerrun:
	docker  run  -p  6060:6060 ccr.ccs.tencentyun.com/oyjx/leak:v1.0.0