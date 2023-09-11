FROM golang:1.20 AS backend-builder
WORKDIR /app
COPY api/v1/handlers/auth /app/auth
COPY ./config /app/config
COPY ./gateway-ui/*.go /app/gateway-ui/
COPY api/v1/routes /app/routes
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway-system .

FROM node:14 AS frontend-builder
WORKDIR /app
COPY ./gateway-ui/package*.json ./gateway-ui/
WORKDIR /app/gateway-ui
RUN npm install
COPY ./gateway-ui ./
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=backend-builder /app/gateway-system /app/gateway-system
COPY --from=frontend-builder /app/gateway-ui/dist /app/frontend

EXPOSE 8050 8051

CMD ["/app/gateway-system"]
