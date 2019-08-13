#!/bin/bash

#Get dependencies
go get github.com/fatih/color

#install packetos
cd ../src/github.com/packetos ; go install 

#install ip
cd ../ip ; go install

#install facility
cd ../facility ; go install

#install plan
cd ../plan ; go install

#install device
cd ../device ; go install

#install main
cd ../main ; go install