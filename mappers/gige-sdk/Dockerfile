FROM golang
WORKDIR /
COPY ./res ./res
COPY ./bin ./bin
COPY ./baumer/Ubuntu-16.04/arm64 ./bin/cti
ENV GENICAM_GENTL64_PATH /bin/cti
RUN sed -i '1a /bin/genicam/Linux64_ARM' /etc/ld.so.conf && ldconfig
ENTRYPOINT ["/bin/main"]
CMD ["--v=4"]
EXPOSE 80/tcp
EXPOSE 80/udp

