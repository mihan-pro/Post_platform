CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
 

CREATE TABLE app_module (
	module_id uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4() PRIMARY KEY,
	module_name VARCHAR ( 50 ) NOT NULL,
	module_description VARCHAR ( 512 ) NOT NULL,
	module_link VARCHAR ( 512 ) NOT NULL,
	module_icon VARCHAR ( 512 ) NOT NULL,
	api_key uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4()
);


INSERT INTO app_module (module_id,module_name,module_description,module_link,module_icon,api_key) VALUES ('ac2d8b14-ce16-4306-aec6-0a65b56c95d7','Доставка цветов','Модуль позволяющий доставлять цветы на указанный адрес','http://ya.ru','http://ya.ru','72f898e6-1755-4892-8aed-8f2742a526fc');
INSERT INTO app_module (module_id,module_name,module_description,module_link,module_icon,api_key) VALUES ('18051a3a-5d89-45ec-8b38-8cc80d21959a','Модуль магазина','Модуль позволяющий заказывать товары с доставкой','http://ya.ru','http://ya.ru','a4667674-ffc8-40f6-ae9d-ac21445f7aec');
INSERT INTO app_module (module_id,module_name,module_description,module_link,module_icon,api_key) VALUES ('714592f4-7817-426f-a508-c1721df7cd41','Модуль трекера посылок','Модуль отслеживания посылок онлайн','http://ya.ru','http://ya.ru','c77db56d-be8b-40c6-ac7c-d7d840f22eb3');
;

 
CREATE TABLE app_product (
	product_id uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_name VARCHAR ( 50 ) NOT NULL,
	product_price int not null default 0,	
	delivery_time  int not null default 0,
	product_rating int not null default 0,
	product_description VARCHAR ( 512 ) NOT NULL,
	product_image VARCHAR ( 512 ) NOT NULL,
	organisation_id  uuid NOT null default null,
	module_id  uuid NOT null default null,
);
ALTER TABLE app_product ADD COLUMN module_id  uuid null default null;

CREATE TABLE app_organisation (
	organisation_id uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4() PRIMARY KEY,
	organisation_name VARCHAR ( 50 ) NOT NULL,
	organisation_description VARCHAR ( 512 ) NOT NULL,
	organisation_logo VARCHAR ( 512 ) NOT NULL,
	organisation_address VARCHAR ( 512 ) NOT NULL,
	organisation_rating int not null default 0	
);

INSERT INTO app_organisation
(organisation_name, organisation_description, organisation_logo, organisation_address, organisation_rating)
VALUES( 'ИП Иванов А.М.', 'Описание доставки ', '', 'г. Нижний Новгород ул. Минина д.1', 0);


CREATE TABLE app_comment (
	comment_id uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4() PRIMARY KEY,
	autor_id VARCHAR ( 50 ) NOT NULL,
	product_id VARCHAR ( 512 ) NOT NULL,
	comment_text VARCHAR ( 512 ) NOT NULL	
);