#!/bin/sh

# Configurable port with default
PORT=${1:-8080}
IMAGE_NAME="rework-lbpk"
CONTAINER_NAME="reworklbpk"

# Build image if not present
if ! docker image inspect $IMAGE_NAME >/dev/null 2>&1; then
  echo "Docker image '$IMAGE_NAME' not found. Building..."
  docker build -t $IMAGE_NAME .
else
  echo "Docker image '$IMAGE_NAME' already exists."
fi

# Clean up any running container with the same name
if docker ps -a --format '{{.Names}}' | grep -Eq "^${CONTAINER_NAME}\$"; then
  echo "Stopping and removing previous container '$CONTAINER_NAME'..."
  docker stop $CONTAINER_NAME >/dev/null
  docker rm $CONTAINER_NAME >/dev/null
fi

# Run container in detached mode, auto-restart on failure
echo "Starting container '$CONTAINER_NAME' on port $PORT..."
docker run -d --name $CONTAINER_NAME -p $PORT:8080 --restart unless-stopped $IMAGE_NAME

# Wait a moment, then show logs
sleep 2
echo "Showing container logs (press Ctrl+C to exit logs, container keeps running)..."
docker logs -f $CONTAINER_NAME

# Optional: Prompt for cleanup when script ends
echo
read -p "Do you want to stop and remove the container '$CONTAINER_NAME'? [y/N] " CLEANUP
if [ "$CLEANUP" = "y" ] || [ "$CLEANUP" = "Y" ]; then
  docker stop $CONTAINER_NAME
  docker rm $CONTAINER_NAME
  echo "Container cleaned up."
else
  echo "Container left running."
fi
