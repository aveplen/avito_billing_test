# Тестовое задание на позицию стажера-бекендера

## Задача:

Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.

&nbsp;

[x] - выполнено

&nbsp;

[x] - метод начисления средств

[x] - метод списания средств

[x] - метод перевода средств

[x] - метод получения текущего баланса

[x] - доп1 возмножность вывода баланса в указанной валюте

[ ] - доп2 возмножность получить историю зачислений и списаний


## Стек:

**Go** - тут понятно

**gin-gonic/gin** - фреймворк

**jackc/pgx/v4** - драйвер для работы с postgres

**ilyakaznacheev/cleanenv** - библиотека для чтения файлов конфигурации

**uber-go/zap** - логгер

**PostgreSQL** - СУБД

## Что реализовано:

1. POST /billing/deposit - зачисление на баланс пользователя с данным user_id средств в размере money (принимается только в виде string)
2. POST /billing/withdraw - списание с баланса пользователя с данным user_id средств в размере money
3. POST /billing/transact - перевод средств в размере money от пользователя с from_id к пользователю to_id
4. POST /billing/balance - получение текущего баланса пользователя с данным user_id в рублях.
5. POST /billing/balance?currency=USD - получение текущего баланса пользователя в желаемой валюте (www.currencyconverterapi.com)

По deposit, withdraw transact в случае успешного выполнения возвращается 204 - No Content

## Тонкости реализации:

Все запросы имеют тип POST, потому что принимают json, даже если там всего одно поле, как у /billing/balance.

"Работаем" с реальными деньгами реальных людей, поэтому хранить размер транзакции в типе float нельзя, потому что из-за точности могут полететь копейки. Принято решение представлять деньги string'ом, в бд они хранятся как money (8-байтовый decimal). Поскольку деньги это string, то обработчик запроса с конвертацией из рублей в другую валюту так же должен работать со строками, для этого можно было бы взять какую-нибудь библиотеку для работы со строками, как с числами, с хорошим тестовым покрытием, но мне было интересно самому написать строкове умножение.

Все расчёты по зачислениям и списаниям производятся базой, баланс не опускается ниже нуля, потому что на поле наложено ограничение типа CHECK.

## Запуск:

```docker-compose up --build```

## Примеры запросов:

```
curl http://localhost:8080/billing/deposit --request POST --header "Content-Type: application/json" --data '{"user_id": 1, "money": "123.45"}'

curl http://localhost:8080/billing/withdraw --request POST --header "Content-Type: application/json" --data '{"user_id": 1, "money": "123.45"}'

curl http://localhost:8080/billing/transact --request POST --header "Content-Type: application/json" --data '{"from_id": 1, "to_id": 2,  "money": "123.45"}'

curl http://localhost:8080/billing/balance --request POST --header "Content-Type: application/json" --data '{"user_id": 1}'

curl http://localhost:8080/billing/balance?currency=USD --request POST --header "Content-Type: application/json" --data '{"user_id": 1}'

```