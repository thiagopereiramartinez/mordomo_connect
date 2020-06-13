# Imagem para compilação
FROM golang AS go-compile

# Copiar aplicação
WORKDIR /go/src/app
COPY . .

# Baixar dependências
RUN go get -d -v ./...
RUN go install -v ./...

# Compilar aplicação
RUN go build main.go

# Imagem para execução
FROM golang AS go-run

# Copiar aplicação
WORKDIR /go/src/app
COPY --from=go-compile /go/src/app/main .

# Iniciar aplicação
CMD [ "./main" ]
