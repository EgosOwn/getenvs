
cd src
export SET_STRING_GETENV="env-has-set"
export SET_BOOL_GETENV="true"

go mod download
go test -cover ./...