package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func SendMail(receiver string, ProjectName string, SiteName string, SiteURL string, Trigger string) {
	EMAIL_SENDER := os.Getenv("EMAIL_SENDER")
	EMAIL_PASSWORD := os.Getenv("EMAIL_PASSWORD")

	if EMAIL_SENDER == "" || EMAIL_PASSWORD == "" {
		color.Red("Unable to configure email credentials")
		return
	}

	from := EMAIL_SENDER
	password := EMAIL_PASSWORD

	to := []string{receiver}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	mimeHeaders := textproto.MIMEHeader{}
	mimeHeaders.Add("From", from)
	mimeHeaders.Add("To", receiver)
	mimeHeaders.Add("Subject", "This is a notification from Stratus")
	mimeHeaders.Add("MIME-Version", "1.0")
	mimeHeaders.Add("Content-Type", `multipart/related; boundary="`+writer.Boundary()+`"`)

	for k, v := range mimeHeaders {
		body.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v[0])))
	}
	body.Write([]byte("\r\n"))

	htmlPart, _ := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type": {"text/html; charset=UTF-8"},
	})
	t, _ := template.ParseFiles("./templates/template.html")
	htmlBuffer := new(bytes.Buffer)
	t.Execute(htmlBuffer, struct {
		Name        string
		ProjectName string
		SiteName    string
		SiteURL     string
		Trigger     string
	}{
		Name:        receiver,
		ProjectName: ProjectName,
		SiteName:    SiteName,
		SiteURL:     SiteURL,
		Trigger:     Trigger,
	})
	htmlPart.Write(htmlBuffer.Bytes())

	imagePath := "./templates/icon.jpeg"
	imagePart, _ := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":              {"image/jpeg"},
		"Content-Disposition":       {fmt.Sprintf(`inline; filename="%s"`, filepath.Base(imagePath))},
		"Content-Transfer-Encoding": {"base64"},
		"Content-ID":                {"<icon>"},
	})

	imageFile, err := os.Open(imagePath)
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer imageFile.Close()

	encoder := base64.NewEncoder(base64.StdEncoding, imagePart)
	_, err = io.Copy(encoder, imageFile)
	if err != nil {
		color.Red(err.Error())
		return
	}
	encoder.Close()

	writer.Close()

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Email sent successfully")
}
