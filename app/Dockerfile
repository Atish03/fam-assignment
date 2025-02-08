# ========== Stage 1: Build Next.js ==========
FROM node:18-alpine AS next-builder

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./
RUN npm install --frozen-lockfile

COPY frontend ./

RUN npm run build

# ========== Stage 2: Build Go Backend ==========
FROM golang:1.23-alpine AS go-builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend ./

RUN go build -o server ./main.go

# ========== Stage 3: Final Deployment Container ==========
FROM alpine:latest

WORKDIR /app

COPY --from=go-builder /app/server ./server
COPY --from=next-builder /app/out ./frontend

EXPOSE 8080

CMD ["./server"]
