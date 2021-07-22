compose = docker-compose -f

alpine = docker-compose.alpine.yml
buster = docker-compose.buster.yml
distroless = docker-compose.distroless.yml
scratch = docker-compose.scratch.yml

build-alpine:
	${compose} ${alpine} build

start-alpine:
	${compose} ${alpine} up

stop-alpine:
	${compose} ${alpine} down

build-buster:
	${compose} ${buster} build

start-buster:
	${compose} ${buster} up

stop-buster:
	${compose} ${buster} down

build-distroless:
	${compose} ${distroless} build

start-distroless:
	${compose} ${distroless} up

stop-distroless:
	${compose} ${distroless} down

build-scratch:
	${compose} ${scratch} build

start-scratch:
	${compose} ${scratch} up

stop-scratch:
	${compose} ${scratch} down

build-all: build-scratch build-distroless build-alpine build-buster