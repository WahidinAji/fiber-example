## Base API
@base
```bash
https://127.0.0.1:8080/api
```
* header
```json
{
    "Content-Type": "application/json"
}
```
## GET
* url
```bash
{@base}/api/books
```
* responses
```json
{
    "code": 200,
    "status" : "Success get all books",
    "data": {
        "books": {
            "id" : "1",
            "author:" : "John",
            "title": "John Smith Title",
            "description": "John Smith Description"
        }
    }
}
```

## POST
* url
```bash
{@base}/api/books
```
* fomr-data
```json
{
    "author:" : "required 'string' ",
    "title": "required 'string' ",
    "description": "required 'string' "
}
```
* responses
```json
{
    "code": 200,
    "status" : "Success create book",
    "data": {
        "id" : "1",
        "author:" : "John",
        "title": "John Smith Title",
        "description": "John Smith Description"
    }
}
```

## GET BY ID
* url must have parameter => id
```bash
{@base}/api/books/{bookId}
```
* responses
```json
{
    "code": 200,
    "status" : "Success get book",
    "data": {
        "id" : "1",
        "author:" : "John",
        "title": "John Smith Title",
        "description": "John Smith Description"
    }
}
```

## UPDATE BY ID
* url must have parameter => id
```bash
{@base}/api/books/{bookId}
```
* fomr-data
```json
{
    "author:" : "required 'string' ",
    "title": "required 'string' ",
    "description": "required 'string' "
}
* responses
```json
{
    "code": 200,
    "status" : "Success update book",
    "data": {
        "id" : "1",
        "author:" : "John ",
        "title": "John Smith Title Update",
        "description": "John Smith Description Update"
    }
}
```

## DELETE BY ID
* url must have parameter => id
```bash
{@base}/api/books/{bookId}
```
* responses
```json
{
    "code": 200,
    "status" : "Success delete book"
}
```

### ID NOT FOUND REPONSE
```json
{
    "code": 404,
    "status" : "Data not found"
}
```