package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/util"
	"gateway/websocket/ai"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadDocument(c *gin.Context) {
	// 从请求中获取文件
	//file, header, err := c.Request.FormFile("file")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
	//	return
	//}
	//defer file.Close()

	// 准备表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//part, err := writer.CreateFormFile("file", filepath.Base(header.Filename))
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in uploading file"})
	//	return
	//}
	//_, err = io.Copy(part, file)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in uploading file"})
	//	return
	//}
	writer.WriteField("fileType", "wiki")
	//writer.WriteField("fileName", "测试pdf")
	writer.WriteField("url", "https://pic1.591adb.cn/data/b2bs/yueyuan/banner/05/96/51199605/1700445837655abe8d033d9.pdf")
	writer.Close()

	// 创建请求
	signature := ai.AssembleSignature() // 获取签名
	req, err := http.NewRequest("POST", "https://chatdoc.xfyun.cn/openapi/fileUpload", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating request"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("appId", os.Getenv("AI_APP_ID"))
	req.Header.Set("signature", signature)
	req.Header.Set("timestamp", fmt.Sprint(time.Now().Unix()))

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in sending request"})
		return
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(responseBody)})
}

func Chatdoc(c *gin.Context) {
	content := c.PostForm("content")
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No content provided"})
		return
	}
	sseWriter := util.NewSSEWriter(c.Writer)

	signature := ai.AssembleSignature()
	wsURL := fmt.Sprintf("wss://chatdoc.xfyun.cn/openapi/chat?appId=%s&timestamp=%d&signature=%s",
		os.Getenv("AI_APP_ID"), time.Now().Unix(), signature)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to WebSocket"})
		return
	}
	defer conn.Close()

	// 发送问答消息
	message := map[string]interface{}{
		"fileIds": []string{"4016a6a0757f4508b3f6fc9e95e7e8d4", "fa3c620f4afa4873a2237119f764a3fd"}, // 文档 ID 列表
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": content, // 用户的提问
			},
		},
		"chatExtends": map[string]interface{}{
			"wikiPromptTpl":             "请将以下内容作为已知信息：\n<wikicontent>\n请根据以上内容回答用户的问题。\n问题:<wikiquestion>\n回答:",
			"wikiFilterScore":           0.82, //取值范围为(0,1] 参考值为：0.80非常宽松 0.82宽松 0.83标准0.84严格 0.86非常严格
			"sparkWhenWithoutEmbedding": true, //大模型兜底
			"temperature":               0.5,  //大模型问答时的温度，取值 0-1，temperature 越大，大模型回答随机度越高
		},
	}
	err = conn.WriteJSON(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			// 解析返回的消息
			var respData map[string]interface{}
			if err := json.Unmarshal(msg, &respData); err != nil {
				break
			}

			// 发送数据到前端
			if err := sseWriter.SendEvent(string(msg)); err != nil {
				break
			}

			// 检查状态码，如果为 2 则结束会话
			if status, ok := respData["status"].(float64); ok && status == 2 {
				break
			}
		}
	}()

	select {
	case <-c.Request.Context().Done():
		return
	}
}
