LC_GCR=gcr.io/learncheck-258301

gen: design/design.go
	./goa gen goauth/design

build: gen cmd/*
#	go build ./cmd/oauth-cli
	go build ./cmd/oauth

docker: build Dockerfile
	docker build -t  $(LC_GCR)/goauth:latest .

run: docker local_cleanup
	docker run --name g2 $(LC_GCR)/goauth:latest

local_cleanup:
	docker rm g2

push: docker
	docker push  $(LC_GCR)/goauth:latest

deploy: push
	gcloud run deploy goauth --image  $(LC_GCR)/goauth:latest  --region us-east1 --platform managed

exmp:
	./goa example goauth/design

reset:
	rm -rf oauth_secured.go cmd



