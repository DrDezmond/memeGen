observe: *.go
ignore: /vendor
formation: web=2
build-server: make build
web: restart=fail waitfor=localhost:8888 ./apiserver serve