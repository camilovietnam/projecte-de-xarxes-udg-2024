buildgo:
	faas-cli build -f echo-go.yml
swarm:
	docker stack deploy -c docker-compose.yml projecte-de-xarxes
unswarm:
	docker stack rm projecte-de-xarxes