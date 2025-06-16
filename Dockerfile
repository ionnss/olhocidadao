# Use official Golang image
FROM golang:1.24.4

# Define work directory
WORKDIR /olhocidadao

# Copy go.mod first
COPY go.mod ./

# Copy go.sum
COPY go.sum ./

# Download dependencies and generate go.sum
RUN go mod download && go mod tidy

# Download the wait-for-it script
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh

# Make script executable
RUN chmod +x /wait-for-it.sh

# Copy code and static files
COPY . .

# Ensure static directory exists and has proper permissions
RUN mkdir -p /olhocidadao/static && chmod -R 755 /olhocidadao/static

# Compile application with optimizations for production
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app_olhocidadao -ldflags="-s -w" && \
    chmod +x /usr/local/bin/app_olhocidadao

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["/wait-for-it.sh", "db:5432", "--", "/usr/local/bin/app_olhocidadao"]