
cd src
export SET_STRING_GETENV="env-has-set"
export SET_BOOL_GETENV="true"
export SET_INT_GETENV=20

go mod download
go test -cover ./...