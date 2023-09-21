package openvpn

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"gateway/api/v1/setting"
	"gateway/apierrors"
	"gateway/ssh"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type Certificate struct {
	Filename     string    `json:"filename"`
	CommonName   string    `json:"common_name"`
	Issuer       string    `json:"issuer"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidTo      time.Time `json:"valid_to"`
	SerialNumber string    `json:"serial_number"`
}

func GetCertificateList(c *gin.Context) {
	// 获取.crt文件列表
	output, err := ssh.ExecuteCustomSSHCommand("ls /etc/openvpn/easy-rsa/pki/issued/*.crt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to run command via SSH: %v", err),
		})
		return
	}

	certFiles := strings.Split(output, "\n")
	var certList []Certificate

	for _, certFile := range certFiles {
		if certFile == "" {
			continue
		}

		// 通过SSH获取证书内容
		certContent, err := ssh.ExecuteCustomSSHCommand("cat " + certFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to get content of certificate %s: %v", certFile, err),
			})
			return
		}

		block, _ := pem.Decode([]byte(certContent))
		if block == nil {
			// Handle error
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			// Handle error
			continue
		}

		certInfo := Certificate{
			Filename:     filepath.Base(certFile),
			CommonName:   cert.Subject.CommonName,
			Issuer:       cert.Issuer.CommonName,
			ValidFrom:    cert.NotBefore,
			ValidTo:      cert.NotAfter,
			SerialNumber: cert.SerialNumber.String(),
		}
		certList = append(certList, certInfo)
	}

	setting.SendResponse(c, http.StatusOK, apierrors.Success, certList)

}

type CreateCertificateRequest struct {
	ClientName string `json:"clientName"`
	Password   string `json:"password"`
}

func CreateCertificate(c *gin.Context) {
	var request CreateCertificateRequest

	// 解析请求
	if err := c.BindJSON(&request); err != nil {
		log.Println("Failed to parse request:", err) // 增加日志输出
		setting.SendResponse(c, http.StatusOK, apierrors.InternalServerError, nil)
		return
	}

	// 构建命令
	command := fmt.Sprintf("./scripts/generate_ovpn.sh %s", request.ClientName)

	// 执行命令
	output, err := ssh.ExecuteCustomSSHCommand(command)

	log.Println("Command Output:", output) // 增加日志输出

	// 判断结果
	if err != nil {
		log.Println("Error executing command:", err) // 增加日志输出
		setting.SendResponse(c, http.StatusOK, apierrors.InternalServerError, gin.H{"output": output, "error": err.Error()})
		return
	}

	if strings.Contains(output, ".ovpn 文件已创建成功!") {
		setting.SendResponse(c, http.StatusOK, apierrors.Success, nil)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to create certificate. Output: %s", output),
		})
	}
}

// RevokeCertificate 吊销证书
func RevokeCertificate(c *gin.Context) {
	// 从URL获取证书ID
	certificateID := c.Param("id")

	// 伪代码: 根据证书ID从数据库或其他存储中删除证书
	// ...

	c.JSON(200, gin.H{
		"status": "success",
		"data":   "Certificate " + certificateID + " revoked.",
	})
}
