package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type TemplateAttr struct {
	Title     string
	Source    string
	OrderInfo OrderInfo
}

type OrderInfo struct {
	OrderNumber    string
	OrderDateTime  string
	ShippingMethod string
	PaymentMethod  string
	SourceOrder    string
}

func main() {
	buf := &bytes.Buffer{}

	directoryInternal := "template/telegram/"
	file, err := os.Open(directoryInternal + "EDCTransactionMsgTemplate.txt")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()

	byteString, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
		return
	}

	telegramMsg := string(byteString)
	template, err := template.New("telegramMsg").Parse(telegramMsg)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// process order
	orderInfo := OrderInfo{
		OrderNumber:    "1000005767",
		OrderDateTime:  "18-02-2022 19:28:56",
		ShippingMethod: "JNE Trucking",
		PaymentMethod:  "Virtual Account BCA",
		SourceOrder:    "eraspace",
	}

	// combine final struct
	templatAttr := TemplateAttr{
		Title:     "WE GOT NEW ORDER !!! ",
		Source:    "ERASPACE",
		OrderInfo: orderInfo,
	}

	err = template.Execute(buf, templatAttr)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(buf.String())
}