# telegram-bot
  Telegram-бот предлагает следующие пункты меню:\
   &ensp;1 до валидации \
       &emsp;* добавить компанию\
       &emsp;* войти в компанию\
   &ensp;2 после успешной валидации ID компании:\
       &emsp;* создать счёт (последовательно запрашивать действия: ввести сумму, ввести описание счёта, ввести email для отправки счёта)\
       &emsp;* посмотреть последние 10 операций\
       &emsp;* получить информацию по общей сумме платежей за день
       
## Пример работы
<img width="300" alt="image" src="https://user-images.githubusercontent.com/102811233/224660092-bf29244b-ce45-40e4-8698-1eddf0981d0b.png">
## Описание кода __ 
Для хранения данных используется БД PostgreSQL, для взаимодействия с ней библеотека github.com/jmoiron/sqlx. Часть кода для открытия БД и CRUD-операции находятся в директории pkg/repository. С помощью интерфейса хранилища, в котором перечислены обязательные функции, хранилище можно заменить на другую БД без изменений в других директориях.
Для взаимодействия с telegram github.com/go-telegram-bot-api/telegram-bot-api. Связь с ботом и обработка запросов располагается в pkg/telegram.
В папке configs хранятся данные для подключения к БД и токен для бота, которые получаются с помощью библиотек github.com/spf13/viper и github.com/joho/godotenv.
