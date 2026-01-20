FROM golang:1.24-alpine


# Set the working directory (this is /backend in SkillSpark)
WORKDIR /app

# Copy only go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

WORKDIR /app/cmd 

CMD ["go", "run", "."]