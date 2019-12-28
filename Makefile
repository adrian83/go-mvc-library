
hello:
	echo "HELLO"

docker:
	sudo systemctl start docker

deps:
	echo "starting mongodb image (version 3.4)"
	docker run -p 27017:27017 -d mongo:3.4

dart-get:
	echo "getting frontend dependencies" 
	cd library-fe && pub get 

dart-build: dart-get 
	echo "building frontend"
	pub global activate webdev
	cd library-fe && pub run build_runner build --delete-conflicting-outputs

dart-run: dart-build
	echo "running frontend"
	cd library-fe && webdev serve --hostname 0.0.0.0 --release

go-fmt:
	echo "formatting go code"
	cd library-be && gofmt -w .

go-test: go-fmt
	echo "running backend tests"
	cd library-be && go test ./... -cover

go-run: export SERVER_PORT=7070
go-run: export SERVER_HOST=0.0.0.0
go-run: export DATABASE_NAME=library
go-run: export DATABASE_HOST=localhost
go-run: export DATABASE_PORT=27017
go-run: export STATICS_PATH=../../../library-fe/build

go-run: go-test
	echo "running backend"
	cd library-be/cmd/library && go run main.go

