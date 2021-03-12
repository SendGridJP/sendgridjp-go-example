package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"io/ioutil"
	"encoding/base64"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func main() {

	err_read := godotenv.Load()
	if err_read != nil {
	  log.Fatalf("error: %v", err_read)
	}

	// .envから環境変数読み込み
    API_KEY := os.Getenv("API_KEY")
    TOS := strings.Split(os.Getenv("TOS"), ",")
    FROM := os.Getenv("FROM")

	message := mail.NewV3Mail()
	// 送信元
	from := mail.NewEmail("", FROM)
	message.SetFrom(from)

	// 宛先
	p := mail.NewPersonalization()
	to := mail.NewEmail("", TOS[0])
	p.AddTos(to)
	p.SetSubstitution("%fullname%", "田中 太郎")
	p.SetSubstitution("%familyname%", "田中")
	p.SetSubstitution("%place%", "中野")
	message.AddPersonalizations(p)

	p2 := mail.NewPersonalization()
	to2 := mail.NewEmail("", TOS[1])
	p2.AddTos(to2)
	p2.SetSubstitution("%fullname%", "佐藤 次郎")
	p2.SetSubstitution("%familyname%", "佐藤")
	p2.SetSubstitution("%place%", "目黒")
	message.AddPersonalizations(p2)

	p3 := mail.NewPersonalization()
	to3 := mail.NewEmail("", TOS[2])
	p3.AddTos(to3)
	p3.SetSubstitution("%fullname%", "鈴木 三郎")
	p3.SetSubstitution("%familyname%", "鈴木")
	p3.SetSubstitution("%place%", "中野")
	message.AddPersonalizations(p3)
	// 件名
	message.Subject = "[sendgrid-go-example] フクロウのお名前は%fullname%さん"
	// テキストパート
	c := mail.NewContent("text/plain", "%familyname% さんは何をしていますか？\r\n 彼は%place%にいます。")
	message.AddContent(c)
	// HTMLパート
	c = mail.NewContent("text/html", "<strong> %familyname% さんは何をしていますか？</strong><br>彼は%place%にいます。")	
	message.AddContent(c)
	// カテゴリ
	message.AddCategories("category1")
	// カスタムヘッダ
	message.SetHeader("X-Sent-Using", "SendGrid-API")
	// 添付ファイル
	a := mail.NewAttachment()
	file, _ := os.OpenFile("./gif.gif", os.O_RDONLY, 0600)
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	data_enc := base64.StdEncoding.EncodeToString(data)
	a.SetContent(data_enc)
	a.SetType("image/gif")
	a.SetFilename("owl.gif")
	a.SetDisposition("attachment")
	message.AddAttachment(a)

	client := sendgrid.NewSendClient(API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
