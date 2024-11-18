REGISTRY_HOST := deejtsn/blogapplication

run:
	templ generate && go build -o ./tmp/main cmd/blog-go/main.go

build-arm:
	docker build -t ${REGISTRY_HOST}:arm --platform=linux/arm64 .

build-amd64:
	docker build -t ${REGISTRY_HOST}:amd64 --platform=linux/amd64 .
