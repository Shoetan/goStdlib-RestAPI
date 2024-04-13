#use the official golang documentation
FROM golang:1.22

#set working directory
WORKDIR /app

# copy source code 
COPY . .

#Download and install dependencies
RUN go get -d -v ./...

#Build the go app
RUN go build -o ecom .

#Expose port
EXPOSE 9000

# Run the executable
CMD ["./ecom"]



