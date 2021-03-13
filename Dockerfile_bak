# FROM golang:alpine
# WORKDIR /app
# ENV GOPROXY=https://goproxy.io,direct
# ENV GO111MODULE on
# COPY ./ .
# RUN GOPROXY="https://goproxy.baidu.com" go mod download
# #RUN go build go.rotbotbox.com/helloweb
# #EXPOSE 8311
# # ENTRYPOINT ["go","run","helloweb.exe"]
# CMD /app

FROM golang
WORKDIR /app
ADD . /app
RUN go build go.rotbotbox.com/helloweb
EXPOSE 8311
ENTRYPOINT ["./helloweb"]