prepare:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
run:
	go run pkg/main.go -r
seed:
	go run pkg/main.go --seed --syncf
