# kickstart-generator
Kickstart File Generator in Go

### Compile and Run - Single golang binary with kickstart template

```
$go version
go version go1.8.1 linux/amd64
$go build kickstart_generator.go
$./kickstart_generator &
[1] 26065
$curl 'localhost:8080/ks_generate/?os=centos7&version=7.3.1611&fqdn=myhost.example.com'
```
or just go to the browser and type the same url

### Add your own parameters

Modify the ks.tmpl file by replacing the value you want to parameterized with {{.your_variable}} and then just use it in the url 
```
curl 'localhost:8080/ks_generate/?os=centos7&version=7.3.1611&fqdn=myhost.example.com&your_variable=foo'
```

### Stop the running process
```
sudo kill -9 26065
[1]+  Killed                  ./kickstart_generator
```

### Run the Kickstart Generator in Docker

1. Compile the code 
```
$CGO_ENABLED=0 go build -a -installsuffix cgo -o kickstart_generator kickstart_generator.go
```

2. Create Dockerfile
```
FROM alpine:latest

RUN apk update && apk add curl && rm -rf /var/cache/apk/*

ADD kickstart_generator /kickstart_generator
ADD ks.tmpl /ks.tmpl

EXPOSE 8080

HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
      CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT ["/kickstart_generator"]

```
NOTE: You could reduce the image size by not installing curl and remove the healthcheck


3. Build the image (Size is 13.6 MB )
```
$sudo docker build -t kickstart_generator -f Dockerfile .
Sending build context to Docker daemon 8.283 MB
Step 1/7 : FROM alpine:latest
 ---> a41a7446062d
Step 2/7 : RUN apk update && apk add curl && rm -rf /var/cache/apk/*
 ---> Running in 56700eb6536f
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/community/x86_64/APKINDEX.tar.gz
v3.6.0-36-gc60c2243c4 [http://dl-cdn.alpinelinux.org/alpine/v3.6/main]
v3.6.0-28-gc11d515dc8 [http://dl-cdn.alpinelinux.org/alpine/v3.6/community]
OK: 8429 distinct packages available
(1/4) Installing ca-certificates (20161130-r2)
(2/4) Installing libssh2 (1.8.0-r1)
(3/4) Installing libcurl (7.54.0-r0)
(4/4) Installing curl (7.54.0-r0)
Executing busybox-1.26.2-r4.trigger
Executing ca-certificates-20161130-r2.trigger
OK: 5 MiB in 15 packages
 ---> 3d9b76b5d956
Removing intermediate container 56700eb6536f
Step 3/7 : ADD kickstart_generator /kickstart_generator
 ---> 8746a2af96aa
Removing intermediate container dd26359b2e7e
Step 4/7 : ADD ks.tmpl /ks.tmpl
 ---> 189f918bd546
Removing intermediate container f27766339507
Step 5/7 : EXPOSE 8080
 ---> Running in 9278aeb7d30f
 ---> 3747089ebb18
Removing intermediate container 9278aeb7d30f
Step 6/7 : HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD curl -f http://localhost:8080 || exit 1
 ---> Running in 6dbbda98b45d
 ---> 07a38a057ef5
Removing intermediate container 6dbbda98b45d
Step 7/7 : ENTRYPOINT /kickstart_generator
 ---> Running in cb2a0f6851c1
 ---> 78eca4c4da37
Removing intermediate container cb2a0f6851c1
Successfully built 78eca4c4da37

```

4. Docker run

```
sudo docker run --rm -td -p 8080:8080 kickstart_generator
c7d534ab18386f418f6b2965b3231b5c0d1b9a1d7fafad2eb119242981652172
$curl http://your_container:8080/status
API is up and running
```

And Enjoy! :+1:
