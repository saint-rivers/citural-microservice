##
## STAGE 1 - BUILD
##
FROM golang:1.20 AS build
# RUN useradd -u 1001 -m iamuser

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./

# the scratch wasn't able to find the binary in the deployment stage
# unless I specified CGO, GOOS and GOARCH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/root /app/cmd/tinker/main.go


##
## STAGE 2 - DEPLOY
##
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/root /app/

RUN apk add --update docker openrc
RUN rc-update add docker boot

# USER 1001
EXPOSE 8010

# runs the imported binary
ENTRYPOINT ["/app/root"]