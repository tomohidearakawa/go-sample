#!/bin/bash

# Dockerイメージのビルド
docker build -t my-go-app .

# Dockerコンテナの実行
docker run -p 8080:8080 my-go-app
