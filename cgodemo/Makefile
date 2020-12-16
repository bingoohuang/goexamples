.PHONY: all c go run env

TESTLIBPATH=ctestlib
all: run

env:
	docker build --tag cgo .
	docker run --rm -ti -v $(shell pwd):/src cgo

c:
	gcc -c -Wall -Werror -fpic -o ${TESTLIBPATH}/test.o ${TESTLIBPATH}/test.c
	ar cru ${TESTLIBPATH}/libtest.a  ${TESTLIBPATH}/test.o
	# gcc -shared -o ${TESTLIBPATH}/libtest.so ${TESTLIBPATH}/test.o

go:c
	go build

run:go
	./cgodemo

