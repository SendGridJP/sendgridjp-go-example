package main

import (
    "os"
    "fmt"
    "strings"
    "log"
    "github.com/sendgrid/sendgrid-go"
    "github.com/joho/godotenv"
)

func main() {

    err_read := godotenv.Load()
    if err_read != nil {
        log.Fatalf("error: %v", err_read)
    }

    SENDGRID_USERNAME := os.Getenv("SENDGRID_USERNAME")
    SENDGRID_PASSWORD := os.Getenv("SENDGRID_PASSWORD")
    TOS := strings.Split(os.Getenv("TOS"), ",")
    FROM := os.Getenv("FROM")

    email := sendgrid.NewMail()
    email.SMTPAPIHeader.SetTos(TOS) // 宛先はこちらで指定したものが使用される
    email.AddTo(FROM)               // 実際には使用されない。エラー回避のため記載。
    email.SetFrom(FROM)
    email.SetFromName("送信者名")
    email.SetSubject("[sendgrid-go-example] フクロウのお名前はfullnameさん")
    email.SetText("familyname さんは何をしていますか？\r\n 彼はplaceにいます。")
    email.SetHTML("<strong> familyname さんは何をしていますか？</strong><br />彼はplaceにいます。")
    sub := make(map[string][]string)
    sub["fullname"] = []string{"田中 太郎", "佐藤 次郎", "鈴木 三郎"}
    sub["familyname"] = []string{"田中", "佐藤", "鈴木"}
    sub["place"] = []string{"office", "home", "office" }
    email.SetSubstitutions(sub)
    email.AddSection("office", "中野")
    email.AddSection("home", "目黒")
    email.AddCategory("category1")
    file, _ := os.OpenFile("./gif.gif", os.O_RDONLY, 0600)
    email.AddAttachment("gif.gif", file)
    defer file.Close()

    sg := sendgrid.NewSendGridClient(SENDGRID_USERNAME, SENDGRID_PASSWORD)
    if r := sg.Send(email); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}
