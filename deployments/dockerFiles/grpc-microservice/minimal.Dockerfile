#NOTE:  This is the "official" build for webservice_redfalcon in docker.   it will reach out to github and download the source from there.  
# It will also create a minimal image that has 0 extra shit in it.   This should be used in production
FROM mornindew/base-demo-protobuf:tagname as builder

RUN apk add git
RUN apk --update add ca-certificates
WORKDIR /go/src


RUN git clone https://github.com/reddiyo-os/grpc-demo.git


## TODO: Paramaterize the root folder name
# Gets the necessary dependencies that are called by the code
RUN cd /go/src/grpc-demo/ && go get ./...
#COMPILE PARAMETERS TO TELL THE COMPLIER TO STATICALLY LINK THE RUNTIME LIBRARIES INTO THE BINARY
RUN cd /go/src/grpc-demo/pkg/grpc-service/main && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app/server


# Buidling the final container from the minimilist image
FROM scratch

COPY --from=builder /app/ /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/email-secret/ /etc/email-secret/

WORKDIR /app
#CMD ["./server"]
EXPOSE 50051
ENTRYPOINT ["./server"]