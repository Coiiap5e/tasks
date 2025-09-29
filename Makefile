BINARY_NAME = algorithms
TARGET_DIR =  bin # имя папки для бинарников
SRC_MAIN = main.go # имя файла с главной функцией
.DEFAULT_GOAL := run

.PHONY: build build-all run clean test bench info

build:
	go build -o $(TARGET_DIR)/$(BINARY_NAME) $(SRC_MAIN)
build-all:
	GOARCH=amd64 GOOS=windows go build -o $(TARGET_DIR)/$(BINARY_NAME)-windows.exe $(SRC_MAIN)
	GOARCH=amd64 GOOS=darwin go build -o $(TARGET_DIR)/$(BINARY_NAME)-darwin $(SRC_MAIN)
	GOARCH=amd64 GOOS=linux go build -o $(TARGET_DIR)/$(BINARY_NAME)-linux $(SRC_MAIN)


run: build
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(TARGET_DIR)

test:
	go test ./... -v

bench:
	go test ./... -bench . -benchmem -v

info:
	@echo "== Tasks Golang Solution =="
	@echo ""
	@echo "Go Version:"
	@go version | cut -d' ' -f3
	@echo ""
	@echo "Module Info:"
	@echo " Name: $(shell go list -m)"
	@echo " Version: $(shell go list -m -f {{.GoVersion}})"
	@echo ""
	@echo "Dependencies:"
	@if "$(shell go mod tidy -v 2>nul)" == "" (\
		echo "Up to date"\
	) else (\
  		echo "Needs update"\
  		echo "Run: go mod tidy && go list -u -m all"\
	)
	@echo ""
	@echo "Module validation:"
	@if "$(shell go mod verify 2>nul)" == "all modules verified" (\
    		echo "Valid"\
    	) else (\
      		echo "Invalid"\
    	)
	@echo ""
	@echo "Repository status"
	@echo "Branch: $(shell git branch --show-current 2>nul)"
	@echo "Changes: "
	@if "$(shell git status --short 2>nul)" == "" (\
    		echo "no changes"\
    	) else (\
      		git status --short 2>nul\
    	)
	@echo "Legend:"
	@powershell -Command "Write-Host 'R' -ForegroundColor Green -NoNewline; Write-Host 'M' -ForegroundColor Red -NoNewline; Write-Host ' - renamed'"
	@powershell -Command "Write-Host 'M' -ForegroundColor Red -NoNewline; Write-Host ' - modified'"
	@powershell -Command "Write-Host '??' -ForegroundColor Red -NoNewline; Write-Host ' - untracked'"
	@powershell -Command "Write-Host 'D' -ForegroundColor Red -NoNewline; Write-Host ' - deleted'"
	@echo ""
	@echo "Structure:"
	@echo "  algorithms/ - algorithms and tasks"
	@echo "  structures/ - data structures"
	@echo ""
	@echo " Commands:"
	@echo " make build - build project for current platform"
	@echo " make build-all - Cross-compilation"
	@echo " make run - build for current platform and run"
	@echo " make test - run tests"
	@echo " make bench - run benchmarks"
	@echo " make clean - clean build"
	@echo " make info - this info"
	@echo " make - default command: run"
	@echo ""
