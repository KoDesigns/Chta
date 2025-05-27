# Docker Cheat Sheet 🐳

*Essential Docker commands for containerized development*

## Basic Commands

```bash
# Check Docker version
docker --version
docker version
docker info

# List running containers
docker ps
docker ps -a               # Include stopped containers

# List images
docker images
docker image ls
```

## Container Lifecycle

```bash
# Run a container
docker run <image>
docker run -d <image>      # Detached mode
docker run -it <image>     # Interactive with TTY
docker run --name <name> <image>

# Stop containers
docker stop <container>
docker stop $(docker ps -q)  # Stop all running

# Start/restart containers
docker start <container>
docker restart <container>

# Remove containers
docker rm <container>
docker rm -f <container>   # Force remove
docker container prune    # Remove all stopped
```

## Image Management

```bash
# Pull images
docker pull <image>
docker pull <image>:<tag>

# Build images
docker build .
docker build -t <name>:<tag> .
docker build -f <dockerfile> .

# Remove images
docker rmi <image>
docker image prune         # Remove unused images
docker system prune        # Remove everything unused
```

## Container Inspection

```bash
# View container logs
docker logs <container>
docker logs -f <container>  # Follow logs
docker logs --tail 50 <container>

# Execute commands in container
docker exec -it <container> bash
docker exec -it <container> sh
docker exec <container> <command>

# Inspect containers
docker inspect <container>
docker stats <container>   # Live stats
```

## Port Mapping

```bash
# Map ports
docker run -p <host>:<container> <image>
docker run -p 8080:80 nginx
docker run -p 3000:3000 node:alpine

# Expose all ports
docker run -P <image>

# Check port mappings
docker port <container>
```

## Volume Management

```bash
# Mount volumes
docker run -v <host>:<container> <image>
docker run -v $(pwd):/app <image>
docker run -v <volume>:<container> <image>

# Create volumes
docker volume create <volume>
docker volume ls
docker volume inspect <volume>

# Remove volumes
docker volume rm <volume>
docker volume prune
```

## Docker Compose

```bash
# Start services
docker-compose up
docker-compose up -d       # Detached
docker-compose up --build # Rebuild images

# Stop services
docker-compose down
docker-compose down -v     # Remove volumes

# View services
docker-compose ps
docker-compose logs
docker-compose logs -f <service>

# Execute commands
docker-compose exec <service> <command>
```

## Environment Variables

```bash
# Set environment variables
docker run -e VAR=value <image>
docker run --env-file .env <image>

# Example with multiple vars
docker run -e NODE_ENV=production \
           -e PORT=3000 \
           node:alpine
```

## Docker Registry

```bash
# Login to registry
docker login
docker login <registry>

# Tag images for push
docker tag <image> <registry>/<image>:<tag>

# Push/pull from registry
docker push <registry>/<image>:<tag>
docker pull <registry>/<image>:<tag>
```

## Networking

```bash
# List networks
docker network ls

# Create network
docker network create <network>

# Connect container to network
docker run --network <network> <image>
docker network connect <network> <container>

# Inspect network
docker network inspect <network>
```

## Cleanup Commands

```bash
# Remove stopped containers
docker container prune

# Remove unused images
docker image prune
docker image prune -a      # Remove all unused

# Remove unused volumes
docker volume prune

# Remove unused networks
docker network prune

# Clean everything
docker system prune
docker system prune -a     # Include unused images
```

## Useful Examples

```bash
# Run temporary container
docker run --rm -it ubuntu:latest bash

# Run with current directory mounted
docker run --rm -v $(pwd):/workspace ubuntu

# Run database container
docker run -d --name postgres \
  -e POSTGRES_PASSWORD=password \
  -p 5432:5432 \
  postgres:13

# Run web server with volume
docker run -d --name nginx \
  -p 8080:80 \
  -v $(pwd)/html:/usr/share/nginx/html \
  nginx:alpine
```

## Pro Tips

- Use `.dockerignore` to exclude files from build context
- Use multi-stage builds to reduce image size
- Always specify image tags in production
- Use `docker-compose` for multi-container applications
- Keep containers stateless and data in volumes
- Regularly clean up unused containers and images

---

*Happy containerizing! 🚀* 