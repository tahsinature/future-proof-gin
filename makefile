prepare:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
run:
	go run pkg/main.go -r
syncdb:
	go run pkg/main.go --syncf
seed:
	go run pkg/main.go --syncf --seed
test:
	go test -v ./pkg/tests
