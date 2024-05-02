# Projecte de xarxes - High Availability

## Open-Faas
- Crear plantilla Go: `faas-cli new --lang golang-middleware --prefix u1062049 projecte-de-xarxes`
- Crear plantilla Go: `faas-cli new --lang golang-http-template --prefix u1062049 projecte-de-xarxes-2`
- Construir imatge: `faas-cli build -f projecte-de-xarxes.yml`

# Docker Swarm
- `docker swarm init`
- `docker stack deploy -c docker-compose.yml projecte-de-xarxes`
- `docker service ls`
- `docker stack rm projecte-de-xarxes`