FROM golang
COPY . /server
WORKDIR /server
EXPOSE 80
CMD go run . config.yml