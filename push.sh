#!/bin/bash
set -e

echo $GCLOUD_SERVICE_KEY | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

docker tag ${PROJECT_NAME}/${DOCKER_IMAGE_NAME}:latest gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}:$TRAVIS_COMMIT
gcloud docker push gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}

yes | gcloud beta container images add-tag gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}:$TRAVIS_COMMIT gcr.io/${PROJECT_NAME}/${DOCKER_IMAGE_NAME}:latest
