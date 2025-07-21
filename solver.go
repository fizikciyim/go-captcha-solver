package captcha

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func CaptchaSolver(filePath, apiKey string) (string, error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(imageData)

	// 2Captcha'ya gönder
	resp, err := http.PostForm("https://2captcha.com/in.php", url.Values{
		"key":    {apiKey},
		"method": {"base64"},
		"body":   {encoded},
		"json":   {"1"},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Status  int    `json:"status"`
		Request string `json:"request"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Status != 1 {
		return "", fmt.Errorf("gönderim hatası: %s", result.Request)
	}

	// Sonucu bekle
	for {
		time.Sleep(5 * time.Second)
		checkURL := fmt.Sprintf("https://2captcha.com/res.php?key=%s&action=get&id=%s&json=1", apiKey, result.Request)
		resp, err := http.Get(checkURL)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		var checkResult struct {
			Status  int    `json:"status"`
			Request string `json:"request"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&checkResult); err != nil {
			return "", err
		}

		if checkResult.Status == 1 {
			return checkResult.Request, nil
		} else if checkResult.Request != "CAPCHA_NOT_READY" {
			return "", fmt.Errorf("hata: %s", checkResult.Request)
		}
	}
}
