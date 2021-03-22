  
FROM golang:latest as builder

WORKDIR /build

COPY . .

RUN go build -ldflags="-s -w" -o alist

# upx 压缩
FROM hairyhenderson/upx:latest as upx
# 文件拷贝
COPY --from=builder /build/alist /usr/bin/alist

RUN upx --best --lzma /usr/bin/alist

FROM scratch
WORKDIR /

COPY --from=upx /usr/bin/alist alist
# 目录复制
COPY --from=builder /build/dist ./dist
COPY --from=builder /build/env.yml ./env.yml

EXPOSE 3000

ENTRYPOINT [ "/alist" ]

# # build stage
# FROM alpine:3.7 as build

# ARG WRK_VERSION=4.0.2

# RUN apk add --no-cache --virtual .build-deps \
#     build-base \
#     perl \
#  && wget -O /tmp/wrk.tar.gz https://github.com/wg/wrk/archive/${WRK_VERSION}.tar.gz \
#  && tar -xzf /tmp/wrk.tar.gz -C /tmp/ \
#  && make -C /tmp/wrk-${WRK_VERSION} \
#  && mv /tmp/wrk-${WRK_VERSION}/wrk /wrk

# # upx stage
# FROM gruebel/upx:latest as upx

# COPY --from=build /wrk /wrk.org

# RUN upx --best --lzma -o /wrk /wrk.org

# # final stage
# FROM alpine:3.7

# ARG BUILD_DATE

# LABEL org.label-schema.build-date=${BUILD_DATE} \
#       org.label-schema.schema-version="1.0"

# RUN apk add --no-cache \
#     libgcc

# COPY --from=upx /wrk /usr/bin/wrk

# ENTRYPOINT ["wrk"]

# docker run --rm gruebel/wrk:latest -t2 -c200 -d10s https://www.example.com

# -c, --connections: total number of HTTP connections to keep open with
#                    each thread handling N = connections/threads

# -d, --duration:    duration of the test, e.g. 2s, 2m, 2h

# -t, --threads:     total number of threads to use

# -s, --script:      LuaJIT script, see SCRIPTING

# -H, --header:      HTTP header to add to request, e.g. "User-Agent: wrk"

#     --latency:     print detailed latency statistics

#     --timeout:     record a timeout if a response is not received within
#                    this amount of time.