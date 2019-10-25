Сборщик данных счетчиков воды
=============================
Сервис сбора данных счетчиков по http(get)

Установка
---------
1. Скопировать фаилы 
water_meter.exe (Windows) water_meter (Linux) и .env
в любую директорию
2. Поправить настройки подключения в БД (MySQL) в файле .env
3. Запустить (желательно как сервис/демон)

Как оно работает
----------------
Оно висит на порту 9001 локалхоста и ждет данные GET формата 

    ?id=K3n210Pm7l&cval1=720&cval2=0&hval1=483&hval2=486&vp=77&dt=7770304567060
    
Где 
    
    id      номер счетчика (string)
    cval1   Показания счетчика №1 холодной воды (int)
    cval2   Показания счетчика №2 холодной воды (int)
    hval1   Показания счетчика №1 горяче воды (int)
    hval2   Показания счетчика №2 горяче воды (int)
    vp      Уровень заряда аккумулятора (int(0-255))
    dt      Дата в unix time format (int)
    
Скрипт для БД (в следующей версии сделаю встроенным):

    create table if not exists w_meter.meters_data
    (
    	id int auto_increment primary key ,
    	MeterId varchar(16) null,
    	WCold1 int null,
    	WCold2 int null,
    	WHot1 int null,
    	WHot2 int null,
    	Power int null,
    	Date timestamp default CURRENT_TIMESTAMP not null,
    	constraint meters_data_id_uindex
    		unique (id)
    );
    
Данные полученые со счетчика проверяются(пока не полностью) пишутся в БД

Что дальше
----------
Надо дописать логер и автоматическое развертывание таблиц в БД


[Gosha.say](mailto:gosha.say@gmail.com) https://aboutblank.pro