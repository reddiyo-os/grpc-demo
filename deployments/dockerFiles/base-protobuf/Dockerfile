FROM golang:alpine
ENV PROTOBUF_TAG=v3.1.0
ADD build.sh /tmp/build.sh
RUN /tmp/build.sh
#Timezone information - needed because scratch won't have it
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .