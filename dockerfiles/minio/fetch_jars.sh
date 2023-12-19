#!/bin/bash
rm -rf dockerfiles/minio/jar/*
rsync --exclude algo/common/ algo/*/target/*-latest.jar dockerfiles/minio/jar/