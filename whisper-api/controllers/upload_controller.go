package controllers

import (
    "fmt"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/AHKAYY007/Whisper-backend/models"
    "github.com/AHKAYY007/Whisper-backend/config"
    "os"
)

func UploadBusinessImage(c *gin.Context) {
    businessID := c.Param("id")

    var business models.Business
    if err := config.DB.First(&business, "id = ?", businessID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
        return
    }

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
        return
    }

    uploadPath := "uploads/businesses/"
    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
        return
    }

    filename := fmt.Sprintf("%s%s", businessID, filepath.Ext(file.Filename))
    filePath := filepath.Join(uploadPath, filename)

    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    // Update business record
    imageURL := fmt.Sprintf("/uploads/businesses/%s", filename)
    business.ImageURL = imageURL
    config.DB.Save(&business)

    c.JSON(http.StatusOK, gin.H{
        "message":   "Image uploaded successfully",
        "image_url": imageURL,
    })

}