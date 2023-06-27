package invade

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Ranco-dev/gbms/pkg/db"
	"github.com/Ranco-dev/gbms/pkg/log"

	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
)

type AddressCheck struct {
	Countt string `form:"Countt" json:"Countt,omitempty" swaggerignore:"true"`
}

type Address struct {
	TaskId        string `form:"TaskId" json:"TaskId,omitempty" swaggerignore:"true"`
	DateTimeLimit string `form:"DateTimeLimit" json:"DateTimeLimit,omitempty" swaggerignore:"true"`
	Address       string `form:"Address" json:"Address,omitempty" swaggerignore:"true"`
}

func InvadeLog(value string) (string, int, error) {
	var result string
	var ctx = context.Background()

	tx, err := db.Conn.Begin(ctx)
	insertQry := `INSERT INTO _invade_log 
				(id, create_date, create_user, modify_date, modify_user, value, comment) 
				values ($1, $2, 'program', $3, 'program', $4, '')`

	uuid := getUuid()
	dtNow := getDtNow()

	if _, err = tx.Exec(ctx, insertQry, uuid, dtNow, dtNow, value); err != nil {
		fmt.Println("err " + err.Error())
		log.Sugar.Errorf("PaymentReq INSERT paymentReq err: %s", err)
		result = "NO"
		return result, http.StatusInternalServerError, err
	}

	if err := tx.Commit(ctx); err != nil {
		result = "NO"
		return result, http.StatusInternalServerError, err
	}

	result = "OK"

	// msg := result + "  " + value
	// sendTelegram(msg)
	return result, http.StatusOK, nil
}

func getUuid() string {
	uuidObj := uuid.New()
	uuidString := uuidObj.String()
	uuidStringWithoutDash := regexp.MustCompile("-").ReplaceAllString(uuidString, "")
	//fmt.Println(uuidString)
	//fmt.Println(uuidStringWithoutDash)
	return uuidStringWithoutDash
}

func getDtNow() string {
	now := time.Now().UTC().Add(8 * time.Hour)
	timeString := now.Format("2006-01-02 15:04:05")
	return timeString
}

func sendTelegram(tgMsg string) {
	bot, err := tgbotapi.NewBotAPI("6005397937:AAGYlGe-J2od_DCiz5MEK47DC02nbSHb4aU")
	if err != nil {
		log.Sugar.Errorf("sendTelegram tgbotapi.NewBotAPI:", err)
	}

	bot.Debug = true
	ChatID := int64(-1001946975736)
	msg := tgbotapi.NewMessage(ChatID, tgMsg)
	if _, err := bot.Send(msg); err != nil {
		log.Sugar.Errorf("sendTelegram bot.Send:", err)
	}
}
