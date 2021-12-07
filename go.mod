module github.com/ozonmp/trv-airport-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.1
	github.com/aphistic/golf v0.0.0-20180712155816-02c07f170c5a // indirect
	github.com/aphistic/sweet v0.3.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/gammazero/workerpool v1.1.2
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/halink0803/zerolog-graylog-hook v0.0.0-20180726102656-e138ca268b12
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozonmp/trv-airport-api/pkg/trv-airport-api v0.0.0-00010101000000-000000000000
	github.com/pressly/goose/v3 v3.1.0
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.24.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/yoheimuta/protolint v0.35.2 // indirect
	golang.org/x/net v0.0.0-20211205041911-012df41ee64c // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	google.golang.org/genproto v0.0.0-20211203200212-54befc351ae9 // indirect
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozonmp/trv-airport-api/pkg/trv-airport-api => ./pkg/trv-airport-api
