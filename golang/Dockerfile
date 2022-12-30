FROM golang:1.13

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY ./.bashrc /root

#RUN \
#    go get -u github.com/99designs/gqlgen@v0.10.2
#    go install github.com/golang/mock/mockgen@v1.5.0 && \
#    go get -u github.com/spf13/cobra@latest && \
#    go get -u github.com/BurntSushi/toml && \
#    go get -u github.com/xuri/excelize/v2


RUN usermod -u 1000 www-data

#USER www-data

USER root

CMD ["tail", "-f", "/dev/null"]