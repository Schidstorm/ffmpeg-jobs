FROM golang as build

COPY . /code
WORKDIR /code
RUN GOOS=linux GOARCH=amd64 go build  -o app -a -ldflags '-w -extldflags "-static"' .

FROM jrottenberg/ffmpeg
COPY --from=build /code/app /usr/local/bin/
RUN chmod 111 /usr/local/bin/app

ENTRYPOINT [ "/usr/local/bin/app" ]

