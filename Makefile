.PHONY: check mocks build

check:
	@go fmt $$(go list ./... | grep -v /vendor/) && go vet $$(go list ./... | grep -v /vendor/)
	@go test -cover -race $$(go list ./... | grep -v /vendor/)

mocks: ##	generate mocks for unit tests
	rm -f mocks/*_gen.go
	genmocks 'mockOpenWeatherClient *MockOpenWeatherClient' weather.OpenWeatherClient mocks > mocks/datastore_mock_gen.go
	goimports -w mocks/*
	gofmt -w mocks/*

build:
	@go build -o weather cmd/weather/*
