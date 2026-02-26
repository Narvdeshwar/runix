# workspace.Dockerfile
FROM ubuntu:22.04

RUN apt update && apt install -y \
    bash curl git python3 python3-pip nodejs npm ca-certificates

WORKDIR /workspace
