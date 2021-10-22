#!/bin/bash

go get -v -u \
    github.com/gin-contrib/cors \
    github.com/go-playground/assert/v2 \
    github.com/go-playground/validator/v10 \
    github.com/golang/protobuf \
    github.com/joho/godotenv \
    github.com/json-iterator/go \
    github.com/leodido/go-urn \
    github.com/modern-go/concurrent \
    github.com/modern-go/reflect2 \
    github.com/ugorji/go \
    golang.org/x/crypto \
    golang.org/x/sys \
    golang.org/x/text \
    gopkg.in/yaml.v2 \
    gorm.io/driver/mysql \
    gorm.io/driver/sqlite \
    gorm.io/gorm \
    github.com/appleboy/gin-jwt \
    github.com/appleboy/gin-jwt/v2
