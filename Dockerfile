FROM node:21-alpine as frontend_builder

RUN apk add --update --no-cache ca-certificates curl wget
RUN corepack enable

WORKDIR /build

ADD backend ./backend

RUN export GOVERSION=$(curl -s "https://go.dev/VERSION?m=text" | head -1) \
  && echo $GOVERSION \
  && wget https://go.dev/dl/${GOVERSION}.linux-amd64.tar.gz \
  && rm -rf /usr/local/go \
  && tar -C /usr/local -xzf ${GOVERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

RUN cd backend && go mod download
RUN cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o certguardian

# frontend
RUN mkdir frontend
COPY frontend ./frontend
RUN echo "PUBLIC_NODE_ENV=ENV_BUILD" > ./frontend/.env
RUN cd frontend \
  && npm install \
  && npm run build

#run
ENV GIN_MODE=release

# COPY --from=frontend_builder /build/frontend/build /app/frontend/build
# COPY --from=backend_builder /certguardian/certguardian /app/bin/certguardian

EXPOSE 7070

ENTRYPOINT ["./backend/certguardian"]
