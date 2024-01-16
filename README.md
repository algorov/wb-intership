# WB-Intership

## Тестовое задание
Необходимо разработать демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе. Модель данных в формате JSON прилагается к заданию.	
				
## Что нужно сделать:
* **Развернуть локально PostgreSQL:**
1. Создать свою БД
2. Настроить своего пользователя
3. Создать таблицы для хранения полученных данных

* **Разработать сервис:**
1. Реализовать подключение и подписку на канал в nats-streaming 
2. Полученные данные записывать в БД
3. Реализовать кэширование полученных данных в сервисе (сохранять in memory)
4. В случае падения сервиса необходимо восстанавливать кэш из БД
5. Запустить http-сервер и выдавать данные по id из кэша 

* **Разработать простейший интерфейс отображения полученных данных по id заказа**

## Бонус-задание						
* Покройте сервис автотестами — будет плюсик вам в карму. 
* Устройте вашему сервису стресс-тест: выясните на что он способен.
