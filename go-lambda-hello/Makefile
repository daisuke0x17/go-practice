.PHONY: generate
generate:
	GOOS=linux go build -o handler handler.go

.PHONY: zip
zip:
	zip function.zip handler
