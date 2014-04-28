package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "github.com/sendgrid/sendgrid-go"
    "encoding/json"
)

type T struct {
    SENDGRID_USERNAME string
    SENDGRID_PASSWORD string
    TOS []string
    FROM string
}

func main() {

    b, err_read := ioutil.ReadFile("config.json")
    if err_read != nil { panic(err_read) }

    t := T{}
    err := json.Unmarshal(b, &t)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    email := sendgrid.NewMail()
    for _, to := range t.TOS {
        email.AddTo(to)
    }
    email.AddFrom(t.FROM)
    email.AddFromName("送信者名")
    email.AddSubject("[sendgrid-go-example] フクロウのお名前はfullnameさん")
    email.AddText("familyname さんは何をしていますか？\r\n 彼はplaceにいます。")
    email.AddHTML("<strong> familyname さんは何をしていますか？</strong><br />彼はplaceにいます。")
    sub := make(map[string][]string)
    sub["fullname"] = []string{"田中 太郎", "佐藤 次郎", "鈴木 三郎"}
    sub["familyname"] = []string{"田中", "佐藤", "鈴木"}
    sub["place"] = []string{"office", "home", "office" }
    email.SetSubstitutions(sub)
    email.AddSection("office", "中野")
    email.AddSection("home", "目黒")
    email.AddCategory("カテゴリ1") 
    email.AddAttachment("./gif.gif")

    sg := sendgrid.NewSendGridClient(t.SENDGRID_USERNAME, t.SENDGRID_PASSWORD)
    if r := sg.Send(email); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}
