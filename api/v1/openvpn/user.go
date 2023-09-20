package openvpn

import (
	"bufio"
	"fmt"
	"gateway/api/v1/setting"
	"gateway/apierrors"
	"gateway/ssh"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type User struct {
	ID             string `json:"id"`
	CommonName     string `json:"common_name"`
	RealAddress    string `json:"real_address"`
	BytesReceived  string `json:"bytes_received"`
	BytesSent      string `json:"bytes_sent"`
	ConnectedSince string `json:"connected_since"`
}

func GetUserList(c *gin.Context) {
	output, err := ssh.ExecuteCustomSSHCommand("echo 'status' | nc 127.0.0.1 7505")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to run command via SSH: %v", err),
		})
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	var userList []User
	insideClientList := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "ROUTING TABLE") {
			break
		}

		if strings.Contains(line, "OpenVPN CLIENT LIST") {
			insideClientList = true
			continue
		}

		if insideClientList && !strings.Contains(line, "Common Name,") {
			parts := strings.Split(line, ",")
			if len(parts) >= 5 {
				user := User{
					ID:             parts[0],
					CommonName:     parts[0],
					RealAddress:    parts[1],
					BytesReceived:  parts[2],
					BytesSent:      parts[3],
					ConnectedSince: parts[4],
				}
				userList = append(userList, user)
			}
		}
	}

	if len(userList) == 0 {
		setting.SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	setting.SendResponse(c, http.StatusOK, apierrors.Success, userList)
}

// ... 其他代码

func CreateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func GetUserLogs(c *gin.Context) {

}
