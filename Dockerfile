FROM golang:1.19-bullseye as build



# Prepare /app
RUN mkdir -p /app/
WORKDIR /app/

# Packages
COPY go.mod go.sum /app/
RUN go mod download

# App files
COPY . /app/
RUN find /app

# Build
RUN go build .




FROM debian:bullseye

# Prepare /app
RUN mkdir -p /app/
WORKDIR /app/

# App files
COPY --from=build /app/go-telegram-statusbot /app/


# Finalize
USER nobody
EXPOSE 8080
LABEL org.opencontainers.image.source https://github.com/kolypto/go-telegram-statusbot
ENV SESSION_FILE=/app/data/session.json
ENV LISTEN=:8080
CMD /app/go-telegram-statusbot
