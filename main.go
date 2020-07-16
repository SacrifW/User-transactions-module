package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)
//Написать бэкенд сервис, по обработке транзакций игрока.
//Каждый запрос имеет token, без которого запрос не действительный.
//Пользователи должны храниться в кэше с помощью map[uint64]*User//Сколько храниться? Поставить по умолчанию хранение КЭШа час???
//Пользователи, которые подверглись изменению, должны сохраняться в базу с периодичностью в 10 секунд.
//Статистика(depositCount, depositSum, ....) должны считаться отдельно(не в структуре пользователя) и в реалтайме, и тоже держаться в кэше.
//Метод GetUser не должен взаимодействовать с базой, а только с кэшом.

type Cache struct {
	mux sync.Mutex
	Users map[uint64]*User
}

type User struct {
	UserId int `json:"user_id"`
	Balance float64 `json:"balance"`
	//Token string `json:"token"`//???
	Account Account `json:"cash"`

}

type Account struct { //перепроверить поля структуры!!!!
	DepositId int     `json:"deposit_id"`
	Deposit   Deposit `json:"deposit"`
	Bet       Bet     `json:"bet"`
	Win       Win `json:"win"`
}
type Deposit struct {

	DepositCount int `json:"deposit_count"`
	DepositSum int `json:"deposit_sum"`
}

type Bet struct {
	BetCount int `json:"bet_count"`
	BetSum int `json:"bet_sum"`
}

type Win struct {
	WinCount int `json:"win_count"`
	WinSum int `json:"win_sum"`
}


func DepSum(deposit, entrance int) interface{}{
	depositSum := Deposit{
		DepositCount: deposit,
		DepositSum:   deposit+entrance,
	}

return depositSum
}
func BetSum(bet, entrance int) interface{} {
	betSum := Bet{
		BetCount: bet,
		BetSum:   bet+entrance,
	}

	return betSum
}
func WinSum(win, entrance int) interface{} {
	winSum := Win{
		WinCount: win,
		WinSum:   win+entrance,
	}

	return winSum
}

func New () *Cache{//инициализация нового контейнера для хранения данных юзеров
	Users := make(map[uint64]*User)

	cache := Cache{
		Users:             Users,
	}

return &cache//возвращаем кэш
}

func (c *Cache) AddUser(key uint64, UserId int, Balance float64, Token string){//Set добавляет новый элемент в кэш или заменяет существующий

	c.mux.Lock()
	defer c.mux.Unlock()

	c.Users[key] = &User{
		UserId:     UserId,
	//	Token:      Token,
		Balance: Balance,
	}
}

func (c *Cache) GetUser(key uint64) (interface{}, bool) {

	c.mux.Lock()
	defer c.mux.Unlock()

	user, found := c.Users[key]
	// ключ не найден
	if !found {
		return nil, false
	}

	return user, true
}


func main(){
	router := gin.Default()

	New()//инициализировали кеш

	router.POST("/user/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"user created"})
	})

	router.GET("user/get", func(c *gin.Context) {//в инструкции тут ПОСТ, я не понимаю, почему тут пост, а не гет.
		var user User
		c.JSON(http.StatusOK, user)
	})

	router.POST("/user/deposit", )
//	router.POST("/transaction", )


	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":8080")
}

