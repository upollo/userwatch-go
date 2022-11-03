# Userwatch Client Library for golang

## Usage

For a full example see [userwat.ch/docs/golang-library](https://userwat.ch/docs/golang-library)

```go
client, err := userwatchgo.NewClientBuilder(API_KEY).Build()

client.Verify(ctx, &userwatchgo.ValidationRequest{
    ValidationToken: request.UserwatchToken, // get the validation token from the javascript library
    Userinfo: &userwatchgo.UserInfo{
        UserID: request.Username,
    },
})
```

## Generating the go lib with protoc

    wget https://raw.githubusercontent.com/Userwatch/userwatch-proto/userwatch_sheperd.proto
    wget https://raw.githubusercontent.com/Userwatch/userwatch-proto/userwatch_shared.proto

    protoc  \
        --proto_path= \
        --go_out=. \
        --go-grpc_out=. \
        userwatch_shared.proto userwatch_shepherd.proto

    mv github.com/Userwatch/userwatch-go/* .
