#!/usr/bin/env bash

mkdir -p ./release/
sudo dpkg-deb --build --debug --verbose . ./release/gpt-beta.deb
