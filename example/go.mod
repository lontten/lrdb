module example

go 1.25.0

replace github.com/lontten/lrdb => ../

require (
	github.com/lontten/lcore/v2 v2.22.0
	github.com/lontten/lrdb v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.19.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96 // indirect
)
