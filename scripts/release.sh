#!/bin/bash
set -e

# Release script for go-athenahealth
# This script ensures releases are always tagged from the 'internal' branch

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if version is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Version number required${NC}"
    echo "Usage: $0 <version>"
    echo "Example: $0 v0.0.4"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-.*)?$ ]]; then
    echo -e "${RED}Error: Invalid version format${NC}"
    echo "Version must be in format: vX.Y.Z or vX.Y.Z-suffix"
    echo "Example: v0.0.4 or v0.0.4-beta"
    exit 1
fi

# Check if we're on the internal branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ "$CURRENT_BRANCH" != "internal" ]; then
    echo -e "${RED}Error: Releases must be tagged from the 'internal' branch${NC}"
    echo "Current branch: $CURRENT_BRANCH"
    echo ""
    echo "To fix this:"
    echo "  git checkout internal"
    echo "  git pull origin internal"
    echo "  ./scripts/release.sh $VERSION"
    exit 1
fi

# Ensure we're up to date
echo -e "${YELLOW}Fetching latest changes...${NC}"
git fetch origin

# Check if branch is up to date with remote
LOCAL=$(git rev-parse @)
REMOTE=$(git rev-parse @{u})
BASE=$(git merge-base @ @{u})

if [ $LOCAL != $REMOTE ]; then
    if [ $LOCAL = $BASE ]; then
        echo -e "${RED}Error: Local branch is behind remote${NC}"
        echo "Run: git pull origin internal"
        exit 1
    elif [ $REMOTE = $BASE ]; then
        echo -e "${RED}Error: Local branch has unpushed commits${NC}"
        echo "Run: git push origin internal"
        exit 1
    else
        echo -e "${RED}Error: Local and remote branches have diverged${NC}"
        echo "Run: git pull origin internal"
        exit 1
    fi
fi

# Check if tag already exists
if git rev-parse "$VERSION" >/dev/null 2>&1; then
    echo -e "${RED}Error: Tag $VERSION already exists${NC}"
    exit 1
fi

# Confirm release
echo -e "${YELLOW}Ready to create release:${NC}"
echo "  Branch: internal"
echo "  Tag: $VERSION"
echo "  Commit: $(git rev-parse --short HEAD)"
echo ""
read -p "Continue? (y/N) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Release cancelled"
    exit 1
fi

# Create and push tag
echo -e "${GREEN}Creating tag $VERSION...${NC}"
git tag -a "$VERSION" -m "Release $VERSION"

echo -e "${GREEN}Pushing tag to origin...${NC}"
git push origin "$VERSION"

echo -e "${GREEN}✓ Release $VERSION created successfully!${NC}"
echo ""
echo "Next steps:"
echo "  1. Create release notes on GitHub"
echo "  2. Monitor CI/CD pipeline"
echo "  3. Verify the release appears in pkg.go.dev"
