build:
	docker build -t roon-discovery-proxy:local .

save:
	docker save -o roon-discovery-proxy.tar roon-discovery-proxy:local

load:
	docker load -i roon-discovery-proxy.tar
