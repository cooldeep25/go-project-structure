# Using wheezy from the official golang docker repo
# alternatively the FROM golang:latest  can be used to fetch latest image.
FROM google/golang
MAINTAINER Kuldeep <cooldeep25@hotmail.com>

# Setting up working directory
# WORKDIR /app 
#Add . /go/src/github.com/cooldeep25/milkbills/
#https://git.heroku.com/kopenvclear

# build the milkbill
# RUN go build -o milkbill .

# Get godeps from main repo
#RUN go get github.com/tools/godep
#Install
#RUN godep go install

# running the script file to execute 
#RUN ./scripts/run_stg.sh

# Setting up environment variables
#ENV ENV STG
#ENV DATABASE_URL=postgres://vjcwapvkggaszc:de60399c466f50b35c4667dd43a98eabcbdbb1e89df8db9d222801108ce5b59f@ec2-23-23-248-162.compute-1.amazonaws.com:5432/drht35sqmocpt
#ENV PORT 80

#EXPOSE 80

# CMD ["/app/main"]
#ENTRYPOINT ["/go/bin/kopenvclear"]
