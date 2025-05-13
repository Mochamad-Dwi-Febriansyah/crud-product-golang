package productcontrollers

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gin/database"
)

func Create(c *gin.Context) {
	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")

	// ambil file gambar dari form
	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "File wajib diunggah"})
		return
	}

	// menyimpan gambar ke direktori lokal
	uploadPath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Gagal menyimpan gambar"})
		return
	}
	//  simpan data produk ke database
	product := models.Product{ 
		NamaProduct: namaProduct,
		Deskripsi: deskripsi,
		Gambar: uploadPath,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Gagal Menyimpan Produk Cuy"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"product" : product})
}


func GetAll(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
		return
	}

	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")

	file, _ := c.FormFile("gambar")
	if file != nil {
		uploadPath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan gambar"})
			return
		}
		product.Gambar = uploadPath
	}

	product.NamaProduct = namaProduct
	product.Deskripsi = deskripsi

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui produk"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}