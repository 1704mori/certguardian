FROM golang:1.21.5-alpine AS backend_builder

RUN apk add --update --no-cache ca-certificates && mkdir /certguardian

WORKDIR /certguardian

COPY backend .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o certguardian

FROM node:21-alpine as frontend_builder

RUN corepack enable

WORKDIR /build

RUN mkdir frontend

RUN echo "PUBLIC_NODE_ENV=ENV_BUILD" > ./frontend/.env

COPY frontend/package*.json ./frontend
RUN cd frontend && npm install

COPY frontend ./frontend
RUN cd frontend && npm run build

FROM alpine:latest

ENV PATH /bin
ENV GIN_MODE=release

COPY --from=frontend_builder /build/frontend/build /app/frontend/build
COPY --from=backend_builder /certguardian/certguardian /app/bin/certguardian

EXPOSE 7070

ENTRYPOINT ["/app/bin/certguardian"]
