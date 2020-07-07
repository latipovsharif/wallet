1) Установить переменные среды окружения
* `WALLET_DB_ADDR` - адрес и порт базы данных // 127.0.0.1:5432
* `WALLET_DB_USER` - пользователь БД // postgres
* `WALLET_DB_PASSWORD` - пароль пользователя БД 
* `WALLET_HOST_ADDR` - Хост на котором будет запущен сервис

2) Выполнить скрипт `sql/initial.sql` в БД
3) Запустить сервис `go run main.go`

<b>Доступные endpoint - ы</b>

a) Создание кошелька: <b>POST</b> `WALLET_HOST_ADDR/v1/wallet/`

    Тело запроса: 
    {
        "name": "My super wallet"
    }
    
b) Перевод средств между кошельками: <b>POST</b> `WALLET_HOST_ADDR/v1/wallet/transfer/`

    Тело запроса: 
    {
     	"src_wallet": 1,
	    "dst_wallet": 2,
	    "amount": 200,
	    "transaction_id": "ad338340-ec3e-4ff2-a289-f8cb5dfd9fd3"
    }

c) Выписка по кошельку <b>GET</b> `WALLET_HOST_ADDR/v1/wallet/excerpt/{wallet_id}/`

    Query params: date=(asc|desc), operation=(asc|desc)


d) Депозит <b>POST</b> `WALLET_HOST_ADDR/v1/wallet/excerpt/`

    Тело запроса

    {
        "wallet": 1,
        "amount": 200,
        "transaction_id": "ad338340-ec3e-4ff2-a289-f8cb5dfd9fd3"
    }