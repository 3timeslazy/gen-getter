test:
	@echo "Executing tests"
	@go build
	@./gen-getter structs.go getters.go
	@go build
	@go clean -i
	@rm -f getters.go
	@echo "Done"