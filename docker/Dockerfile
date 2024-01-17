FROM golang:latest AS build-stage

WORKDIR /app

COPY . ./

EXPOSE 8080

RUN apt-get -y update
RUN apt-get install -y software-properties-common
RUN add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe restricted multiverse"
RUN apt-get -y update

CMD ["tail", "-f", "/dev/null"]