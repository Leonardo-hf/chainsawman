#!/bin/bash
rm -rf dockerfiles/minio/jar/*
rsync --exclude common-latest.jar algo/*/target/*-latest.jar dockerfiles/minio/jar/