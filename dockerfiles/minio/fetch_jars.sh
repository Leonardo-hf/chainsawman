#!/bin/bash
rsync --exclude algo/common/ algo/*/target/*-latest.jar dockerfiles/minio/jar/