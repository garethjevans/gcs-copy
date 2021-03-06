#!/bin/sh

# This script gets a version number from the latest tag (any branch) and appends the commit hash if not exactly on a tag.

set -e

ORIGINAL_DIR=$(pwd)

running_in_docker() {
  [ -f /.dockerenv ] 
}

is_git_repo_clean()  {
  git diff-index --quiet HEAD --
}

cleanup() {
  git checkout -q $ORIGINAL_BRANCH
  cd $ORIGINAL_DIR
  echo $VERSION
}
# Always run cleanup...
trap cleanup INT TERM EXIT

# Store the original branch name
ORIGINAL_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Get the commit hash
COMMIT=$(git rev-parse --verify HEAD)

if running_in_docker; then
  #echo "Build inside docker container detected. Checking out commit..."
  git checkout -q $COMMIT
elif ! is_git_repo_clean; then
  #echo "You have uncommitted work in this project. Please commit your changes to git before making a build. Aborted."
  exit 1
fi

# Get the nearest tag for the release (will append hash to release version if not at exact release"
TAG=$(git describe --tags HEAD --always || true)

# Set version based on tag or nearest tag and hash
if [ -z "$TAG" ]; then
  #echo "HEAD not at tag. Appending git commit to last tag found..."
  VERSION="$(git describe --tags `git rev-list --tags --max-count=1`)-$COMMIT"
else
  #echo "Using tag $TAG"
  VERSION="$TAG"
fi
