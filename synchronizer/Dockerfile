FROM golang as build

COPY . /code
WORKDIR /code
RUN go build -o app -a .

FROM jrottenberg/ffmpeg
COPY --from=build /code/app /usr/local/bin/
RUN chmod 333 /usr/local/bin/app

ENTRYPOINT [ "/usr/local/bin/app" ]

