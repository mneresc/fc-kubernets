docker build -t mneresc/go-test-http .
docker run --rm -p 8080:8080 neres/go-test

docker login
docker push mneresc/go-test-http