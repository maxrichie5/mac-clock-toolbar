#!/bin/sh

mkdir zones
cd zones

cp "/opt/homebrew/Cellar/go/1.19.1/libexec/lib/time/zoneinfo.zip" "."
unzip zoneinfo.zip
rm zoneinfo.zip

find . -type f -not -name "." -not -name "zones.txt" | awk '{sub("\./", "")}1' > zones.txt
mv zones.txt ..

cd ..
rm -rf zones
