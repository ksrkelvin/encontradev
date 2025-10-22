package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    c.HTML(http.StatusOK, "home.html", gin.H{
        "title": "Página Inicial",
    })
}

func PartialExample(c *gin.Context) {
    c.HTML(http.StatusOK, "partial.html", gin.H{
        "msg": "Conteúdo carregado dinamicamente via HTMX 🚀",
    })
}