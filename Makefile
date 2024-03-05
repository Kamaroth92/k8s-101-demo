build:
	go build -o k8s-101-demo

run: build
	./k8s-101-demo -i --all

deps:
	git clone https://github.com/GoogleCloudPlatform/microservices-demo