build-go:
	faas-cli build -f echo-go.yml
deploy:
	docker stack deploy -c docker-compose.yml projecte-de-xarxes
undeploy:
	docker stack rm projecte-de-xarxes