package models

type Tiss []struct {
	Brand            string  `json:"brand"`
	Brand_alt        string  `json:"brand_alt"`
	Article          string  `'json:"article"`
	Min_price        float32 `json:"min_price"`
	Warehouse_offers []struct {
		Id    string  `json:"id"`
		Price float32 `json:"price"`
	} `json:"warehouse_offers"`
}
type TissBrand struct {
	Article   string `json:"Article"`
	BrandList []struct {
		BrandID   int    `json:"BrandID"`
		BrandName string `json:"BrandName"`
	} `json:"BrandList"`
}
type Froza struct {
	Min_price map[string]struct {
		Price_full float32 `json:"price_full"`
	} `json:"min_price"`
}
type Froza2 struct {
	Min_price map[string]struct {
		Price_full string `json:"price_full"`
		Detail_num string `json:"detail_num"`
	} `json:"min_price"`
}
type Mikado struct {
	List []struct {
		Code_List_Row []struct {
			Supplier     string `xml:"Supplier"`
			ProducerCode string `xml:"ProducerCode"`
			Brand        string `xml:"Brand"`
			PriceRUR     string `xml:"PriceRUR"`
		} `xml:"Code_List_Row"`
	} `xml:"List"`
}
type LoginType struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	Errors  string `json:"errors"`
	Result  struct {
		Content string `json:"content"`
		Params  struct {
			Success bool   `json:"success"`
			Token   string `json:"token"`
		} `json:"params"`
	} `json:"result"`
}
type Forum struct {
	Gid   string  `json:"gid"`
	Brand string  `json:"brand"`
	Art   string  `json:"art"`
	Price float32 `json:"price"`
}

type Logout struct {
	Success bool `json:"success"`
}

var Login = "29771" //Это для ivers
var Pass = "23y1sm"

type Ivers struct {
	Success bool `json:"success"`
	Result  struct {
		Params struct {
			Success bool `json:"success"`
			Objects struct {
				Original []struct {
					Brand      string `json:"brand"`
					BaseNumber string `json:"baseNumber"`
					Price      string `json:"price"`
					Prices     []struct {
						Price float32 `json:"price"`
						Title string  `json:"title"`
					} `json:"prices"`
					WarehouseCitys []string `json:"warehouseCitys"`
				} `json:"original"`
			} `json:"objects"`
		} `json:"params"`
	} `json:"result"`
}

type Str struct {
	Firm  string
	Price float32
}
type Str2 struct {
	Firm   string
	Prices []Str
}

var Rez struct {
	Article string
	//Keys    []string
	Firms map[string][]Str
	Viev  []Str2
}

type Rezult struct {
	Firm  string
	Price map[string]float32
}

/*
type Orders struct {
	Number      string `json:"Number"`
	Datatime    string `json:"Datatime"`
	State       string `json:"State"`
	Branch_Name string `json:"Branch_Name"`
	Branch_Code string `json:"Branch_Code"`
	Price       string `json:"Price"`
	ID          string `json:"ID"`
	Arhive      int    `json:"Arhive`
}
type Item struct {
	ID           string `json:"ID"`
	Brand        string `json:"Brand"`
	Article      string `json:"Article"`
	Article_Name string `json:"Article_Name"`
	Quantity     string `json:"Quantity"`    //Количество позиций
	Price        string `json:"Price"`       //Стоимость за единицу
	Total_Price  string `json:"Total_Price"` //Общая цена
	State        string `json:"State"`       //Состояние конкретной позиции
	InProcess    string `json:"InProcess"`   //Количество позиций в обработке
	Withdraw     string `json:"Withdraw"`    //Количество позиций в отказе
	Shipped      string `json:"Shipped"`     //Количество позиций отгружено
	//Comment      string `json:"Comment"`
}
type OrderItems struct {
	ID    string `json:"ID"`
	State string `json:"State"`
	Items []Item `json:"Items"`
}
type ArticleBrandList struct {
	Article   string      `json:"Article"`
	BrandList []BrandList `json:"BrandList"`
}
type BrandArticle struct {
	Brand   string `json:"Brand"`
	Article string `json:"Article"`
}
type BrandList struct {
	BrandID   int    `json:"BrandID"`
	BrandName string `json:"BrandName"`
}
type OurData struct {
	//Data        string  //дата записи этого заказа в файл
	ACTT        string  //Номер позиции по АСТТ(номеру блока)
	Id          string  //Id заказа
	Brand       string  //фирма детали
	Article     string  //партнамбер
	ArticleName string  //название товара
	Quantity    int     //количество
	Price       float32 //цена
	State       bool    //состояние позиции. если true - отгружено
	Remark      string  //наша ремарка
}

var TimeOffset = -1 //период обработки в месяцах
//Структура таблицы japautozap/tiss
// id string длиной 36(37 с запасом)
// ACTT string(11)
//
// number string (30) номер детали
// name string(58) название детали
// firm string(15) фирма
// price string (10) цена
// quantity string (10) количество
// remark string (20) примечание
// status bool   отгружено/нет
// date DATE время записи заказа

// создаём таблицу
// mysql -u vazman -p
// create database japautozap;  - если ещё не создана
// use japautozap
/* create table tiss (id varchar (37),
ACTT varchar(11),
number varchar(30),
name varchar(80),
firm varchar(15),
price varchar(10),
quantity varchar(10),
remark varchar(20),
status bool,
date DATE);*/
//ALTER TABLE tiss modify column name varchar(50);   изменяем столбец
//ALTER TABLE tiss modify column number varchar(30);
//mysqldump -u vazman -p japautozap tiss > japautozap_tiss.sql
//mysql -uvazman -prbhgbxb1 japautozap < japautozap_tiss.sql
// update tiss set status=false;
//ALTER TABLE tiss ADD COLUMN quanity2  VARCHAR(10) AFTER status;
//ALTER TABLE tiss ADD COLUMN status2  bool  AFTER quanity2;
//ALTER TABLE tiss DROP COLUMN status2;
//ALTER TABLE tiss ADD COLUMN subid  TINYINT UNSIGNED AFTER id;
//update tiss set status=false where number="C-110";
