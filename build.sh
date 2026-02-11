#!/bin/bash

# Configuration
REGISTRY_USERNAME="${DOCKER_USERNAME:-YOUR_DOCKERHUB_USERNAME}"
IMAGE_NAME="nova"
VERSION="${VERSION:-latest}"
LOCAL_TAG="nova"
REGISTRY_TAG="${REGISTRY_USERNAME}/${IMAGE_NAME}:${VERSION}"

# Parse arguments
PUSH_TO_REGISTRY=false
if [[ "$1" == "--push" ]]; then
    PUSH_TO_REGISTRY=true
fi

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running."
    echo "Please start Docker and try again."
    exit 1
fi

if [ "$PUSH_TO_REGISTRY" = true ]; then
    # Build multi-architecture image for k8s (ARM64 for Raspberry Pi + AMD64)
    echo "Building multi-architecture image for AMD64 and ARM64..."
    echo "Image: ${REGISTRY_TAG}"

    # Check if logged in to Docker Hub
    if ! docker info | grep -q "Username"; then
        echo "Error: Not logged in to Docker Hub."
        echo "Run 'docker login' to authenticate."
        exit 1
    fi

    # Create buildx builder if it doesn't exist
    if ! docker buildx inspect multiarch > /dev/null 2>&1; then
        echo "Creating buildx builder..."
        docker buildx create --name multiarch --use
    else
        docker buildx use multiarch
    fi

    docker buildx build \
        --platform linux/amd64,linux/arm64 \
        -t "${REGISTRY_TAG}" \
        --push \
        .

    if [ $? -eq 0 ]; then
        echo ""
        echo "✓ Build and push successful!"
        echo "Image: ${REGISTRY_TAG}"
        echo ""
        echo "Update k8s/deployment.yaml with: image: ${REGISTRY_TAG}"
        echo "Then deploy: kubectl apply -k k8s/"
    else
        echo "Build failed!"
        exit 1
    fi
else
    # Local build only
    echo "Building Docker image locally..."
    docker build -t "${LOCAL_TAG}" .

    if [ $? -eq 0 ]; then
        echo ""
        echo "✓ Build successful!"
        echo "To run: docker run -p 8080:8080 ${LOCAL_TAG}"
        echo ""
        echo "To build and push to registry: ./build.sh --push"
    else
        echo "Build failed!"
        exit 1
    fi
fi