
cd src
export SET_STRING_GETENV="env-has-set"
go mod download
go test -cover ./...