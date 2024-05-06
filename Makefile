buildgo:
	faas-cli build -f echo-go/echo-go.yml
swarm:
	docker stack deploy -c docker-compose.yml projecte-de-xarxes
unswarm:
	docker stack rm projecte-de-xarxes
alias:
	alias ms='make swarm'
	alias mus='make unswarm'