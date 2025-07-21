# Go Captcha Solver

Bu paket, Go dili ile yazılmış basit bir captcha çözücü modülüdür. 2Captcha servisini kullanarak, captcha resimlerini otomatik olarak çözmenizi sağlar.

## Başlamadan Önce

### 2Captcha API Anahtarı Nasıl Alınır?

1. [2Captcha](https://2captcha.com/) sitesine gidin ve ücretsiz/kullanıcı hesabı oluşturun.
2. Hesabınıza giriş yaptıktan sonra, **Dashboard** (Kontrol Paneli) sayfasında API anahtarınızı (API Key) bulun.
3. Bu API anahtarını `.env` dosyanıza yazacağız.


## Kurulum

```bash
go get github.com/fizikciyim/go-captcha-solver
```

## Kullanım

.env.example dosyasını .env olarak kopyalayıp içine kendi 2Captcha API anahtarınızı yazın:

```bash
CAPTCHA_API_KEY="api_keyinizi_buraya_yazin"
```

Ardınan projenizde aşağıdaki gibi kullanabilirsiniz:

```
package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/kullaniciAdi/go-captcha-solver/captcha"
)

func main() {
    // .env dosyasını yükle
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("CAPTCHA_API_KEY")
    if apiKey == "" {
        log.Fatal("CAPTCHA_API_KEY is not set")
    }

    answer, err := captcha.CaptchaSolver("debug/kare.png", apiKey)
    if err != nil {
        log.Fatal(err)
    }

    println("Captcha sonucu:", answer)
}
```

## Özellikler
* Base64 formatında captcha resmi gönderir.

* 2Captcha servisi üzerinden çözüm sonucunu alır.

* Basit ve kullanımı kolay.

## Dikkat Edilmesi Gerekenler

* 2Captcha servisi ücretli olabilir, API anahtarınızı güvenli tutun.

* Gönderilen resmin dosya yolu ve formatı doğru olmalıdır.

* Sunucuya istekler belirli aralıklarla yapılır, bu yüzden çözüm süresi biraz zaman (genellikle 5-10 saniye) alabilir.
