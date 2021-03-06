package controllers

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"wallet/ELGamal"
)
const (
	ErrorValue = "value cannot be empty"
	RejectServer = "Server Error"
)
type NewWallet struct {
	Name string `json:"name" form:"name" query:"name"`
	Id 	 string `json:"id" form:"id" query:"id"`
	Str  string `json:"str" form:"str" query:"str"`
}

func Register(c echo.Context) error {
	w := new(NewWallet)
	// 因为 echo 的 bind 无绑定检查功能
	// echo 强制要求 post 的参数写在 body 里，写在 header 里会绑定不上
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 暂时只能验证是否为空
	fmt.Println("绑定完成")
	if w.Id == "" || w.Name == "" || w.Str == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	}
	// 计算公私钥
	pub, priv, err := ELGamal.GenerateKeys(w.Str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 取哈希
	pub.G1 = new(big.Int)
	HashInfoBuf := sha256.Sum256([]byte(w.Str))
	// 向监管者提交注册请求，并返回相关信息
	// [32]byte 是一个数组，要把他转换成切片
	if resp, err := http.PostForm("http://39.106.173.191:1423/register", url.Values{"name": {w.Name}, "id": {w.Id}, "Hashky": {string(HashInfoBuf[:])}}); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		if res, err := ioutil.ReadAll(resp.Body); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, err)
		} else {
			// 判断应该返回的信息
			if bytes.Equal(res,[]byte("Successful!")) {
				return c.JSON(http.StatusOK, priv)
			} else {
				return c.JSON(http.StatusInternalServerError, RejectServer)
			}
		}
	}
}
