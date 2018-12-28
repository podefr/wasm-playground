FROM golang:1.11.3

RUN mkdir /test

WORKDIR /test

COPY index.html main.wasm ./

RUN cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./

EXPOSE 8080

RUN go get -u github.com/shurcooL/goexec

ENTRYPOINT ["goexec", "http.ListenAndServe(\":8080\", http.FileServer(http.Dir(\".\")))"]

