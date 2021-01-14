package controllers

import (
	"bytes"
	"crypto/sha256"
	"github.com/labstack/echo"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"wallet/ELGamal"
)
const (
	ErrorValue = "value cannot be blank"
	RejectServer = "Server Error"
)
type newWallet struct {
	name string `json:"name" form:"name"`
	id 	 string `json:"id" form:"id"`
	str  string `json:"str" form:"str"`
}

func Register(c echo.Context) error {
	w := new(newWallet)
	// 因为 echo 的 bind 无绑定检查功能
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 暂时只能验证是否为空
	if w.id == "" || w.name == "" || w.str == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	}
	// 计算公私钥
	pub, priv, err := ELGamal.GenerateKeys(w.str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 取哈希
	pub.G1 = new(big.Int)
	HashInfoBuf := sha256.Sum256([]byte(w.str))
	// 向监管者提交注册请求，并返回相关信息
	// [32]byte 是一个数组，要把他转换成切片
	if resp, err := http.PostForm("http://39.106.173.191：1423/register", url.Values{"name": {w.name}, "id": {w.id}, "Hashky": {string(HashInfoBuf[:])}}); err != nil {
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
