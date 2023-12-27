FROM node:21-alpine as frontend_builder

RUN corepack enable

WORKDIR /build

COPY frontend/package*.json ./
RUN npm install

COPY frontend ./frontend
COPY tsconfig* .
COPY svelte.config* .
COPY vite* .
COPY postcss* .
COPY tailwind* .

RUN npm run build

FROM golang:1.21.4-alpine AS backend_builder

RUN apk add --update --no-cache ca-certificates && mkdir /certguardian

WORKDIR /certguardian

COPY backend .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o certguardian

FROM alpine:latest

RUN apk add curl

ENV PATH /bin
COPY --from=frontend_builder /build/build /frontend
COPY --from=backend_builder /certguardian/certguardian /backend/certguardian

EXPOSE 7070

ENTRYPOINT ["/certguardian"]
