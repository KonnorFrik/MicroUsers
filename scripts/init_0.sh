set -e
go mod tidy
go build -o ../bin/ .
useradd -r user -p user
