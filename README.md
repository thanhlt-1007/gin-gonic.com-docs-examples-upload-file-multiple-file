# gin-gonic.com-docs-examples-upload-file-multiple-file

- Multiple files

- Reference: https://gin-gonic.com/docs/examples/upload-file/multiple-file/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- POST /files

```sh
curl --location 'localhost:8080/files' \
--form 'files[]=@"/Users/le.tan.thanh/Desktop/images/1.jpeg"' \
--form 'files[]=@"/Users/le.tan.thanh/Desktop/images/2.jpeg"'
```
