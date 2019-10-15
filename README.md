[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/noting59/Golang-Solid-Implementation/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/noting59/Golang-Solid-Implementation/?branch=master)

[![Build Status](https://scrutinizer-ci.com/g/noting59/Golang-Solid-Implementation/badges/build.png?b=master)](https://scrutinizer-ci.com/g/noting59/Golang-Solid-Implementation/build-status/master)

# Golang-Solid-Implementation
Fake online shop page with pay, one-click pay and refund using SOLID

----------
INSTALL
-------

Clone the source

    git clone https://github.com/noting59/Golang-Solid-Implementation && cd ./Golang-Solid-Implementation

Copy file config.yaml.dist to config.yaml
  
    postresql_conn: "user=user dbname=test sslmode=disable"
    
Setup dependencies

    go get github.com/lib/pq
    go get github.com/thedevsaddam/renderer
    go get -u github.com/go-chi/chi
    go get github.com/afex/hystrix-go/hystrix
    go get github.com/noting59/Golang-Solid-Implementation

Setup sqlite data structure

     psql -d {database} -a -f setup.sql 

Run the app

    go build -o index && ./index

And visit

    http://localhost:5000/

----------

По функціоналу:

1. Сторінка / - сторінка товару
2. Сторінка /cart - корзина товару
3. Сторінка /order/new - створити ордер - (якщо кнопка Pay - немає в юзера card_token, якщо One click Pay - то є)
4. Якщо немає one click pay - то переходимо на сторінку введення карти - iframe
5. Якщо є one click pay - то переходимо на нову сторінку
6. Сторінка /order/list - сторінка всіх ордерів юзера і статуси цих ордерів (якщо ордер статус "approved" - то є кнопка "refund")

7. Також в фоні працює таска, що чекає статуси ордерів "processing" - раз в 10 секунд

Зауваження:
1. всі корнер кейси не розглядались, а тому перевіряти слід лише карту номер 1 з https://solidgate.atlassian.net/wiki/spaces/API/pages/9207830/Cards+for+test+payments
2. Вебхуки для відловлення зміни статусів не робив, тому що локально в них сенсу немає
3. Якщо будуть питання, звертайтесь
