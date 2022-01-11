run-local:  ## Build and run the go application outside of Docker
	go run main.go

size:  ## Show the current size of the example app container
	docker images|grep exampleapp

console:  ## Run shell inside the container.
	docker run -ti --rm --entrypoint=busybox exampleapp sh

diff2:  ## Show difference between 1 and 2
	colordiff -u Dockerfile1 Dockerfile2

diff3:  ## Show difference between 2 and 3
	colordiff -u Dockerfile2 Dockerfile3

diff4:  ## Show difference between 3 and 4
	colordiff -u Dockerfile3 Dockerfile4

diff5:  ## Show difference between 4 and 5
	colordiff -u Dockerfile4 Dockerfile5

run1:  ## Build and run the simple demo app
	docker build -f Dockerfile1 -t exampleapp .
	docker run -ti --rm exampleapp

run2:  ## Build and run the app after adding an unprivileged user
	docker build -f Dockerfile2 -t exampleapp .
	docker run -ti --rm exampleapp

run3:  ## Build and run the app after adding dumb-init as entrypoint
	docker build -f Dockerfile3 -t exampleapp .
	docker run -ti --rm exampleapp

run4:  ## Build and run the app after avoiding temporary files
	docker build -f Dockerfile4 -t exampleapp .
	docker run -ti --rm exampleapp

run5:  ## Build and run the app after using multi-stage build
	docker build -f Dockerfile5 -t exampleapp .
	docker run -ti --rm exampleapp


.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
