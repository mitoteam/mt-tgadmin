# Set up variables
EXECUTABLE_NAME := mt-tgadmin

#do the job
include internal/goapp/Makefile.inc.mk


# additional targets
.PHONY: run
run:
	clear
	go build
	${EXECUTABLE_NAME} run
