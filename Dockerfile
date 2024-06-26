FROM ubi8/go-toolset
COPY * .

RUN go build .
CMD ./GO-CI-CD

