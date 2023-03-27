.PHONY: proto

# generate proto
proto:
	kratos proto client pkg
	kratos proto client internal
	kratos proto client --proto_path=$(shell pwd)/api api