.PHONY: proto

# generate proto
proto:
	kratos proto client pkg
	kratos proto client internal
	kratos proto client api
	kratos proto client app