Когда завершите задачу, в этом README опишите свой ход мыслей: как вы пришли к решению, какие были варианты и почему выбрали именно этот. 

# Что нужно сделать

Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение. 

# Необязательно, но было бы круто

Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.

# Способ запуска

Записываем необходимые параметры в файле config/conf.yml 

TimeLimit - промежуток времени на котором может быть максимум N запросов в наносекундах

MaxQuantityQuery - максимальное количество запросов

Запускаем Redis:

cd ./deployments

docker compose up

Запускаем приложение:

cd ./cmd

go run main.go

У программы есть flag port, если он не установлен за дефолтное значение применяется 8080

# Способ взаимодействия

Отправить запрос на url localhost:8080\query с параметром id
