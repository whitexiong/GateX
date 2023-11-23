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
	"log"
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
	hostUrl := os.Getenv("AI_HOST_URL_v3")
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
	if question == "" {
		question = "你好"
	}

	chatHistory := GetChatHistoryForRoom(chatRoomID)

	currentMessage := map[string]interface{}{
		"role":    "user",
		"content": question,
	}
	chatHistory = append(chatHistory, currentMessage)
	log.Printf("上下文：%v\n", chatHistory)

	// 添加 Function Call 后期插件调用
	//functionCall := map[string]interface{}{
	//	"text": []map[string]interface{}{
	//		{
	//			"name":        "天气查询",
	//			"description": "天气插件可以提供天气相关信息。你可以提供指定的地点信息、指定的时间点或者时间段信息，来检索诗词库，精准检索到天气信息。",
	//			"parameters": map[string]interface{}{
	//				"type": "object",
	//				"properties": map[string]interface{}{
	//					"location": map[string]interface{}{
	//						"type":        "string",
	//						"description": "地点，比如北京。",
	//					},
	//					"date": map[string]interface{}{
	//						"type":        "string",
	//						"description": "日期。",
	//					},
	//				},
	//				"required": []string{"location"},
	//			},
	//		},
	//		// 这里可以继续添加其他 Function
	//	},
	//}

	return map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": os.Getenv("AI_APP_ID"),
			//"patch_id": []string{""}, //微调的模型资源ID
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain":      "generalv3",
				"temperature": 0.5,
				"max_tokens":  1024,
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": chatHistory,
			},
			//"functions": functionCall, // 添加到请求中
		},
	}
}

func StartChatWithAI(chatRoomID uint, question string) string {
	const maxRetries = 3               // 最大重试次数
	const retryDelay = 2 * time.Second // 每次失败后的等待时间

	var conn *websocket.Conn
	var resp *http.Response
	var err error
	d := websocket.Dialer{HandshakeTimeout: 5 * time.Second}

	for i := 0; i < maxRetries; i++ {
		conn, resp, err = d.Dial(AssembleAuthUrl(), nil)
		if err == nil && resp.StatusCode == 101 {
			// 连接成功
			break
		}

		// 打印错误信息
		fmt.Println("连接AI错误 (尝试 #", i+1, "):", readResp(resp), err)

		// 如果达到最大重试次数，则返回
		if i == maxRetries-1 {
			return ""
		}

		// 等待并重试
		time.Sleep(retryDelay)
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

		jsonString := string(msg)
		fmt.Println("流式数据:", jsonString)

		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			fmt.Println("解析JSON错误:", err)
			return answer
		}

		payload, exists := data["payload"].(map[string]interface{})
		if !exists {
			fmt.Println("解析payload错误")
			fmt.Printf("完整响应数据: %+v\n", data)
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

		//for _, t := range text {
		//	textMap, ok := t.(map[string]interface{})
		//	if !ok {
		//		continue
		//	}
		//
		//	// 检查并处理 function_call
		//	if functionCall, exists := textMap["function_call"]; exists {
		//		fmt.Println("Function Call 响应:", functionCall)
		//		// 根据 function_call 的内容执行相应操作
		//		if functionName, ok := functionCall.(map[string]interface{})["name"].(string); ok {
		//			switch functionName {
		//			case "天气查询":
		//				// 解析参数并执行天气查询
		//				arguments := functionCall.(map[string]interface{})["arguments"].(string)
		//				// 这里调用一个函数来处理天气查询，例如：
		//				weatherInfo := queryWeather(arguments)
		//				answer += weatherInfo
		//			}
		//		}
		//	}
		//
		//	// 累加其他文本内容
		//	if content, exists := textMap["content"].(string); exists && content != "" {
		//		answer += content
		//	}
		//}

		answer += content
		if status == 2 {
			conn.Close()
			break
		}
	}

	fmt.Println("AI回复:", answer)
	return answer
}

func queryWeather(arguments string) string {
	return "测试 func"
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
	for i := len(messages) - 1; i >= 0; i-- { // 反向遍历消息
		msg := messages[i]
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
