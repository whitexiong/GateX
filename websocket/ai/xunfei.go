package ai

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gateway/models"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// 构造讯飞AI认证URL
func AssembleAuthUrl() string {
	apiSecret := os.Getenv("AI_API_SECRET")
	apiKey := os.Getenv("AI_API_KEY")
	hostUrl := os.Getenv("AI_HOST_URL")
	ul, err := url.Parse(hostUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	date := time.Now().UTC().Format(time.RFC1123)
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	signature := HmacWithShaTobase64(signString, apiSecret)

	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"hmac-sha256\", headers=\"host date request-line\", signature=\"%s\"", apiKey, signature)
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)

	return hostUrl + "?" + v.Encode()
}

// 生成与讯飞AI交互的参数
func GenParams(chatRoomID uint, question string) map[string]interface{} {
	// 如果问题为空，则使用默认问题
	if question == "" {
		question = "你好"
	}

	// 获取历史聊天记录
	chatHistory := GetChatHistoryForRoom(chatRoomID)

	// 将当前问题添加到历史聊天记录中
	currentMessage := map[string]interface{}{
		"role":    "user",
		"content": question,
	}
	chatHistory = append(chatHistory, currentMessage)
	return map[string]interface{}{
		"header": map[string]interface{}{"app_id": os.Getenv("AI_APP_ID")},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain":      "generalv2",
				"temperature": 0.8,
				"top_k":       6,
				"max_tokens":  2048,
				"auditing":    "default",
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{"text": chatHistory},
		},
	}
}

func StartChatWithAI(chatRoomID uint, question string) string {
	d := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
	conn, resp, err := d.Dial(AssembleAuthUrl(), nil)

	if err != nil || resp.StatusCode != 101 {
		fmt.Println("连接AI错误:", readResp(resp), err)
		return ""
	}

	go func() {
		// 在这里，我们传递chatRoomID以获取上下文
		if err := conn.WriteJSON(GenParams(chatRoomID, question)); err != nil {
			fmt.Println("发送问题错误:", err)
		}
	}()

	var answer string
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息错误:", err)
			break
		}

		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			fmt.Println("解析JSON错误:", err)
			return answer
		}

		payload, exists := data["payload"].(map[string]interface{})
		if !exists {
			fmt.Println("解析payload错误")
			return answer
		}

		choices, exists := payload["choices"].(map[string]interface{})
		if !exists {
			fmt.Println("解析choices错误")
			return answer
		}

		header, exists := data["header"].(map[string]interface{})
		if !exists {
			fmt.Println("解析header错误")
			return answer
		}

		code, _ := header["code"].(float64)
		if code != 0 {
			fmt.Println("AI返回错误:", data["payload"])
			return answer
		}

		status, _ := choices["status"].(float64)
		text, exists := choices["text"].([]interface{})
		if !exists || len(text) == 0 {
			fmt.Println("解析text错误")
			return answer
		}

		content, exists := text[0].(map[string]interface{})["content"].(string)
		if !exists {
			fmt.Println("解析content错误")
			return answer
		}

		answer += content
		if status == 2 {
			conn.Close()
			break
		}
	}

	fmt.Println("AI回复:", answer)
	return answer
}

func HmacWithShaTobase64(signString []string, key string) string {
	data := strings.Join(signString, "\n")
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func GetChatHistoryForRoom(chatRoomID uint) []map[string]interface{} {
	var messages []models.Message
	models.DB.Where("chat_room_id = ?", chatRoomID).Order("created_at desc").Limit(3).Find(&messages)

	var chatHistory []map[string]interface{}
	for _, msg := range messages {
		role := "user"
		if msg.AIProvider != models.None {
			role = "assistant"
		}
		chatHistory = append(chatHistory, map[string]interface{}{
			"role":    role,
			"content": msg.Content,
		})
	}
	return chatHistory
}
