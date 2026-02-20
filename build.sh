#!/bin/bash
set -euo pipefail

# Load .env if present
if [ -f .env ]; then
    set -a
    source .env
    set +a
fi

# ----------------------------------------
#  CONFIGURATION
# ----------------------------------------

# Auto-detect Docker Hub username if not provided
DETECTED_USERNAME="$(docker info --format '{{.Username}}' 2>/dev/null || true)"
REGISTRY_USERNAME="${DOCKER_USERNAME:-$DETECTED_USERNAME}"

IMAGE_NAME="nova"
VERSION="${VERSION:-latest}"
LOCAL_TAG="${IMAGE_NAME}"
REGISTRY_TAG="${REGISTRY_USERNAME}/${IMAGE_NAME}:${VERSION}"

# ----------------------------------------
#  VALIDATION
# ----------------------------------------

# Ensure Docker is running
if ! docker info >/dev/null 2>&1; then
    echo "❌ Error: Docker is not running."
    echo "Start Docker Desktop and try again."
    exit 1
fi

# Ensure username is available
if [[ -z "$REGISTRY_USERNAME" ]]; then
    echo "❌ Error: No Docker Hub username detected."
    echo "Run: docker login"
    echo "Or export DOCKER_USERNAME=<yourname>"
    exit 1
fi

# Ensure username is lowercase (Docker requirement)
if [[ "$REGISTRY_USERNAME" =~ [A-Z] ]]; then
    echo "❌ Error: Docker Hub username must be lowercase."
    echo "Detected: $REGISTRY_USERNAME"
    exit 1
fi

# ----------------------------------------
#  ARGUMENT PARSING
# ----------------------------------------

PUSH_TO_REGISTRY=false
if [[ "${1:-}" == "--push" ]]; then
    PUSH_TO_REGISTRY=true
fi

# ----------------------------------------
#  BUILD LOGIC
# ----------------------------------------

if [[ "$PUSH_TO_REGISTRY" == true ]]; then
    echo "----------------------------------------"
    echo "🚀 Building multi-architecture image"
    echo "   Platforms: linux/amd64 + linux/arm64"
    echo "   Image:     ${REGISTRY_TAG}"
    echo "----------------------------------------"
    echo "Note: Make sure you're logged in with 'docker login'"
    echo ""

    # Create buildx builder if needed
    if ! docker buildx inspect multiarch >/dev/null 2>&1; then
        echo "🔧 Creating buildx builder..."
        docker buildx create --name multiarch --use
    else
        docker buildx use multiarch
    fi

    docker buildx build \
        --platform linux/amd64,linux/arm64 \
        -t "${REGISTRY_TAG}" \
        --push \
        .

    echo ""
    echo "✅ Build and push successful!"
    echo "   Image: ${REGISTRY_TAG}"
    echo ""
    echo "Update k8s/deployment.yaml with:"
    echo "   image: ${REGISTRY_TAG}"
    echo ""
    echo "Deploy with:"
    echo "   kubectl apply -k k8s/"
    echo ""

else
    echo "----------------------------------------"
    echo "🛠  Building local Docker image"
    echo "   Tag: ${LOCAL_TAG}"
    echo "----------------------------------------"

    docker build -t "${LOCAL_TAG}" .

    echo ""
    echo "✅ Local build successful!"
    echo "Run locally with:"
    echo "   docker run -p 8080:8080 ${LOCAL_TAG}"
    echo ""
    echo "To build and push to Docker Hub:"
    echo "   ./build.sh --push"
    echo ""
fi
