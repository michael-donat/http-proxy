all:
	docker build -t michaeldonat/http-proxy .
	docker push michaeldonat/http-proxy