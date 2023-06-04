package payment

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Ranco-dev/gbms/pkg/db"
	"github.com/Ranco-dev/gbms/pkg/log"

	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"

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

func PaymentReq(userId string, amount string) (Address, int, error) {
	var address Address
	var ctx = context.Background()

	//會先查詢 _address 中該 user_id 是否有地址了
	var addressCheck AddressCheck
	sqlQry := "SELECT count(id) as countt FROM _address where user_id = '" + userId + "'"
	fmt.Println(sqlQry)
	err := db.Conn.QueryRow(ctx, sqlQry).Scan(
		&addressCheck.Countt)
	if err != nil {
		log.Sugar.Errorf("PaymentReq get address err: %s", err)
		return address, http.StatusInternalServerError, err
	}
	fmt.Println(addressCheck.Countt)

	if addressCheck.Countt == "0" {
		//若無 則挑選一個無 user_id 的 _address ，並把 user_id update給該地址
		tx, err := db.Conn.Begin(ctx)
		if err != nil {
			return address, http.StatusInternalServerError, err
		}
		defer tx.Rollback(ctx)

		updateQry := fmt.Sprintf("UPDATE _address SET user_id = '%s' WHERE index = ( SELECT MIN(index) FROM _address WHERE user_id = '' );", userId)
		fmt.Println(updateQry)

		if _, err = tx.Exec(ctx, updateQry); err != nil {
			log.Sugar.Errorf("PaymentReq UPDATE address err: %s", err)
			return address, http.StatusInternalServerError, err
		}

		if err := tx.Commit(ctx); err != nil {
			return address, http.StatusInternalServerError, err
		}
	} else {
		//若有 則會把 _payment_req 資料表中 該 user_id 的 pending 資料 closed，
		tx, err := db.Conn.Begin(ctx)
		if err != nil {
			return address, http.StatusInternalServerError, err
		}
		defer tx.Rollback(ctx)

		updateQry := fmt.Sprintf("UPDATE _payment_req set status = 'closed' WHERE id in ( SELECT id FROM _payment_req WHERE user_id = '%s' and status = 'pending' );", userId)
		fmt.Println(updateQry)

		if _, err = tx.Exec(ctx, updateQry); err != nil {
			log.Sugar.Errorf("PaymentReq update paymentReq err: %s", err)
			return address, http.StatusInternalServerError, err
		}

		if err := tx.Commit(ctx); err != nil {
			return address, http.StatusInternalServerError, err
		}
	}

	//寫入一筆到 _payment_req 資料表 狀態 為 pending，寫入 初始錢包餘額，最終錢包餘額 留空
	uuid := getUuid()
	dtNow := getDtNow()
	addressString := getAddress(userId)
	initWalletBalance, err := getWalletBalance(addressString)
	if err != nil {
		return address, http.StatusInternalServerError, err
	}
	fmt.Println("initWalletBalance " + initWalletBalance)
	initWalletBalanceInt, err := strconv.Atoi(initWalletBalance)
	if err != nil {
		log.Sugar.Errorf("PaymentReq strconv.Atoi err: %s", err)
		return address, http.StatusInternalServerError, err
	}

	initWalletBalanceFloat := float64(initWalletBalanceInt) / 1000000.0
	initWalletBalanceResult := fmt.Sprintf("%.6f", initWalletBalanceFloat)

	tx, err := db.Conn.Begin(ctx)
	insertQry := "INSERT INTO _payment_req " +
		" (id, create_date, create_user, modify_date, modify_user, " +
		"	status, req_date, user_id, user_ip, amount, start_balance, end_balance, comment, address) " +
		fmt.Sprintf("values ('%s', '%s', 'program', '%s', 'program', 'pending', '%s', '%s', '', '%s', '%s', '', '', '%s')", uuid, dtNow, dtNow, dtNow, userId, amount, initWalletBalanceResult, addressString)

	fmt.Println("insertQry " + insertQry)
	if _, err = tx.Exec(ctx, insertQry); err != nil {
		fmt.Println("err " + err.Error())
		log.Sugar.Errorf("PaymentReq INSERT paymentReq err: %s", err)
		return address, http.StatusInternalServerError, err
	}

	if err := tx.Commit(ctx); err != nil {
		return address, http.StatusInternalServerError, err
	}

	address.Address = addressString
	address.DateTimeLimit = getDtLimit()
	address.TaskId = uuid

	msg := "娛樂城 " + userId + " 金額 " + amount + " 發起支付 " +
		" \n " + addressString + " \n " + uuid
	sendTelegram(msg)
	return address, http.StatusOK, nil
}

func PaymentCheck(taskId string) (string, int, error) {
	var ctx = context.Background()

	//會先查詢 _address 中該 user_id 是否有地址了
	var taskStatus string
	sqlQry := "SELECT status FROM _payment_req where id = '" + taskId + "'"
	//fmt.Println(sqlQry)
	err := db.Conn.QueryRow(ctx, sqlQry).Scan(
		&taskStatus)
	if err != nil {
		log.Sugar.Errorf("PaymentCheck db.Conn.QueryRow err: %s", err)
		return taskStatus, http.StatusInternalServerError, err
	}
	//fmt.Println(taskStatus)

	return taskStatus, http.StatusOK, nil
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
	now := time.Now()
	timeString := now.Format("2006-01-02 15:04:05")
	return timeString
}

func getDtLimit() string {
	now := time.Now()
	afterHalfHour := now.Add(30 * time.Minute)
	timeString := afterHalfHour.Format("2006-01-02 15:04:05")
	return timeString
}

func getAddress(userId string) string {
	address := ""
	var ctx = context.Background()
	sqlQry := "SELECT address as countt FROM _address where user_id = '" + userId + "'"
	//fmt.Println(sqlQry)
	err := db.Conn.QueryRow(ctx, sqlQry).Scan(
		&address)
	if err != nil {
		log.Sugar.Errorf("getAddress db.Conn.QueryRow err: %s", err)
		return address
	}
	//fmt.Println(address)
	return address
}

func getWalletBalance(address string) (string, error) {
	//url := "https://apilist.tronscanapi.com/api/accountv2?address=THL1P3gA9xq8AZ6maufU68K16Rs1xbPyqt"
	url := "https://apilist.tronscanapi.com/api/accountv2?address=" + address
	//TODO 輪流API KEY
	web, err := fetchClient(url, "afdabcc8-9bda-4526-9d24-4573a57ec164")
	//fmt.Println("web " + web)
	if err != nil {
		return "", err
	}
	reStr, err := fetch(web)
	if err != nil {
		return "", err
	}
	//fmt.Println("reStr " + reStr)
	return reStr, nil
}

func fetchClient(url string, apiKey string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Sugar.Errorf("fetchClient http.NewRequest err: %s", err)
		return "", err
	}

	req.Header.Add("TRON-PRO-API-KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Sugar.Errorf("fetchClient client.Do err: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("fetchClient Unexpected status code %d", resp.StatusCode)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Sugar.Errorf("fetchClient ioutil.ReadAll err: %s", err)
		return "", err
	}

	return string(body), nil
}

type Token struct {
	TokenId   string `json:"tokenId"`
	Balance   string `json:"balance"`
	TokenName string `json:"tokenName"`
}

type JSONData struct {
	WithPriceTokens []Token `json:"withPriceTokens"`
}

func fetch(web string) (string, error) {
	reStr := "0"

	var data JSONData
	err := json.Unmarshal([]byte(web), &data)
	if err != nil {
		log.Sugar.Errorf("fetch Error:", err)
		//fmt.Println("Error:", err)
		return "0", err
	}

	for _, token := range data.WithPriceTokens {
		if token.TokenId == "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" {
			reStr = token.Balance
			break
		}
	}

	return reStr, nil
}

func sendTelegram(tgMsg string) {
	bot, err := tgbotapi.NewBotAPI("5729751481:AAE3RpJzg4vr3xAQ85zU4G4pz_IaHGUhqqY")
	if err != nil {
		log.Sugar.Errorf("sendTelegram tgbotapi.NewBotAPI:", err)
	}

	bot.Debug = true
	ChatID := int64(-1001873872614)
	msg := tgbotapi.NewMessage(ChatID, tgMsg)
	if _, err := bot.Send(msg); err != nil {
		log.Sugar.Errorf("sendTelegram bot.Send:", err)
	}
}
