all:
	@go run main.go
fmt:
	go fmt wego/wxapi
	go fmt main.go

add:
	git add *
