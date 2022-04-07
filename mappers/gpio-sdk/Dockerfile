FROM alpine:3.14
RUN mkdir -p kubeedge
COPY ./bin/gpio-sdk kubeedge/bin/
COPY ./res kubeedge/
WORKDIR kubeedge
ENTRYPOINT ["/kubeedge/gpio-sdk", "--v", "4"]