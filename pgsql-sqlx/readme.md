
i use some method to response result of the api ^^
* the first one `1`
```go
return ctx.Status(fiber.StatusOK).JSON(&book.WebResponse{
    Code:   ctx.Response().StatusCode(),
    Status: "OK",
    Data:   newBook,
})
```
* the second one `2`
```go
ctx.Status(fiber.StatusNotFound)
return ctx.JSON(helpers.RespApi(ctx.Response().StatusCode(),"Not Found", result))
```
* the third one `3`
```go
return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
    "code": ctx.Response().StatusCode(),
    "status": "OK",
    "data": result,
})
```

* delete go.mod and go.sum
* delete go.mod and go.sum

* `run go mod init again`
```bash
go mod init ur-package
go mod tidy
```