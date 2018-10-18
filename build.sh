#!/bin/sh

DIR_OUTPUT=build/
FILENAME=reminder

rm -rf ${DIR_OUTPUT}
git clone https://github.com/phanmduong/go-reminder-build
mv go-reminder-build ${DIR_OUTPUT}
GOOS=linux GOARCH=amd64 go build -o ${DIR_OUTPUT}${FILENAME} -v main.go
cp env.example ${DIR_OUTPUT}
cp -r ./public/. ${DIR_OUTPUT}public
cd ${DIR_OUTPUT}
git add .
git commit -a -m "ADD build at `date`"
git push -u origin master
cd ..
