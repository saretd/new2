package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/trv-airport-api/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/trv-airport-api/internal/app/sender EventSender
//go:generate mockgen -destination=./mocks/airport_repo_mock.go -package=mocks github.com/ozonmp/trv-airport-api/internal/repo  AirportRepo
