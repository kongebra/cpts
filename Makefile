cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

dpl ?= deploy.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

build:
    docker build -t $(APP_NAME) .

build-nc:
    docker build -t --no-cache $(APP_NAME) .

up:
    build run

stop:
    docker stop $(APP_NAME); docker rm $(APP_NAME)


run:
    docker run -i -t --rm --env-file=./config.env -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)