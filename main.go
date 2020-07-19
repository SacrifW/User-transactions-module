package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type Cache struct {
	mux sync.Mutex
	Users map[uint64]*User
}

type User struct {
	UserId int
	Balance float64 `json:"balance"`
	DepositCount int `json:"deposit_count"`
	DepositSum float64 `json:"deposit_sum"`
	BetCount int `json:"bet_count"`
	BetSum float64 `json:"bet_sum"`
	WinCount int `json:"win_count"`
	WinSum float64 `json:"win_sum"`
}

type Deposit struct{
	ForUser uint64 `json:"for_user"`
	DepositId int `json:"deposit_id"`
	Amount float64 `json:"deposit_amount"`//сумма попоплнения
	Time time.Time `json:"deposit_time"`//время пополнения
}

type Transaction struct {
	TransactionOfUser uint64 `json:"transaction_of_user"`
	TransactionId int `json:"transaction_id"`
	Type bool `json:"type"`
	Amount float64 `json:"amount_transaction"`
	Time time.Time `json:"time_transaction"`
}

var token = "testtask"

func New () *Cache{//инициализация нового контейнера для хранения данных юзеров
	Users := make(map[uint64]*User)
	cache := Cache{
		Users:             Users,
	}

	return &cache//возвращаем кэш
}
var Users = make(map[uint64]*User)


func AddUser (balance float64, token string) (Users map[uint64]*User, err error){
	//time.Sleep(15 * time.Second)//допилить обновление юзера? как реализовать? или дать таймер функции Баланса?
var Id uint64
var userId int
userId = int(Id)

	if token == "" {
		fmt.Println("Нет токена")
	}

	user := User{
		Balance: balance,
	}

	userId = 0
	if userId == len(Users){
		userId += 1
		Id = uint64(userId)
	}

	Users[Id] = &user

return Users, err
}

func AddDeposit (amount, balance float64) (Balance float64, err error){
	if token == "" {
		fmt.Println("Нет токена")
	}

	var depositId int
	var userId uint64


	user := User{
		Balance: balance,
	}

	deposit := Deposit{
		ForUser: userId,
		DepositId: depositId,
		Amount:    amount,
		Time:      time.Now(),
	}

	balanceDepositSum(deposit.Amount, user.Balance)

return user.Balance, err
}

func balanceDepositSum (amount float64, balance float64) float64{
	balance += amount
	return balance
}

func AddTransaction (TransactionOfUser uint64, amount float64, balance float64, token string,userId int, transactionId int) (Balance float64, err error){

	if token == "" {
		fmt.Println("Нет токена")
	}

	user := User{
		Balance: balance,
	}

	var	transType bool
	switch transType {
	case true:
		balanceWinSum(amount, user.Balance)
	case false:
		balanceBet(amount, user.Balance)
	}

	transaction := Transaction{
		TransactionOfUser: TransactionOfUser,
		TransactionId:     transactionId,
		Type:              transType,
		Amount:            amount,
		Time:              time.Now(),
	}

	Transactions := make(map[int]Transaction)

	key := 0
	if key == len(Transactions){
		key += 1
	}
	Transactions[key] = transaction

	return user.Balance, err
}
func balanceWinSum (amount float64, balance float64) float64 {
	balance += amount
	return balance
}
func balanceBet (amount float64, balance float64) float64{
	balance -= amount
	return balance
}

func GetUser(userId uint64, token string, Users map[uint64]User) (user User, err error) {
	time.Sleep(15*time.Second)

	if token == "" {
		fmt.Println("Нет токена")
	}

	for userId, _ = range Users{
		user := Users[userId]
		return user, err
	}

	return user, err
}
func main(){
	router := gin.Default()
	
	New()//инициализировали кеш

	router.POST("/user/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, AddUser)
	})

	router.GET("user/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetUser)
	})//ДОДЕЛАТЬ

	router.POST("/user/deposit", func(c *gin.Context) {
		c.JSON(http.StatusOK, AddDeposit)
	})

	router.POST("/transaction", func(c *gin.Context) {
		c.JSON(http.StatusOK, AddTransaction)
	})


	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})


	router.Run(":8080")
}
