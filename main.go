package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func postFilesHandler(context *gin.Context) {
    form, _ := context.MultipartForm()
    files :=  form.File["files[]"]
    for _, file := range files {
        context.SaveUploadedFile(file, "tmp/" + file.Filename)
    }

    context.JSON(
        http.StatusOK,
        gin.H {
            "len": len(files),
        },
    )
}

func main() {
    engine := gin.Default()
    engine.MaxMultipartMemory = 8 << 20
    engine.POST("/files", postFilesHandler)
    engine.Run()
}
