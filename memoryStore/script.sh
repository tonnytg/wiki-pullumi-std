#!/bin/bash

gcloud memcache instances create my-memcache-instance \
    --node-count=2 \
    --node-cpu=1 \
    --node-memory=3GB \
    --region=us-central1 \
    --zone=us-central1-a \
    --network=default


gcloud redis instances describe my-redis-instance --region=us-central1


gcloud memcache instances describe my-memcache-instance --region=us-central1


gcloud redis instances update my-redis-instance --size=2


gcloud redis instances delete my-redis-instance --region=us-central1
gcloud memcache instances delete my-memcache-instance --region=us-central1



