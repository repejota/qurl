FROM golang:latest
WORKDIR /go/src/github.com/repejota/qurl
COPY . .
RUN make deps
RUN make install
CMD ["qurl"]