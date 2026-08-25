# Stage 1: Build React UI
FROM node:18-alpine AS ui-builder
WORKDIR /app/ui
COPY ui/package*.json ./
RUN npm install
COPY ui/ ./
RUN npm run build

# Stage 2: Build Go Server & Agent
FROM golang:1.21-alpine AS go-builder
WORKDIR /app
COPY server/ ./server
COPY agent/ ./agent
RUN cd server && go build -o /meshscope-server main.go
RUN cd agent && go build -o /meshscope-agent main.go

# Stage 3: Final Lightweight Runtime
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=ui-builder /app/ui/dist ./ui/dist
COPY --from=go-builder /meshscope-server ./meshscope-server
COPY --from=go-builder /meshscope-agent ./meshscope-agent

EXPOSE 8080

CMD ["/bin/sh", "-c", "/app/meshscope-server & /app/meshscope-agent"]