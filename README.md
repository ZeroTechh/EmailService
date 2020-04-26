
# EmailService

A velocity service which deals with sending emails

## Functions

SendSimpleEmail -- sends a simple text only email

```go
_, err := client.SendSimpleEmail(ctx, &proto.EmailData{
    Text: "YOUR TEXT",
    Email: "Email to send",
})
```
