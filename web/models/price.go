package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Price struct in sql
// ['id', 'court', 'court_number', 'auction', 'city_area', 'address', 'total_auction_price',
// 'pointtopay', 'empty', 'memory', 'remark', 'look_fig', 'communication', 'land']
type Price struct {
	ID            string `orm:"column(id)"`
	Court         string `orm:"column(court)"`
	CourtNumber   string `orm:"column(court_number)"`
	Auction       string `orm:"column(auction)"`
	CityArea      string
	Address       string
	TotalAuction  string
	Pointtopay    string
	Empty         string
	Memory        string
	Remark        string
	LookFig       string
	Communication string
	Land          string
}

/*筆次                      | text | YES  |     | NULL    |       |
| 法院名稱                  | text | YES  |     | NULL    |       |
| 字號股別                  | text | YES  |     | NULL    |       |
| 拍賣日期拍賣次數          | text | YES  |     | NULL    |       |
| 縣市                      | text | YES  |     | NULL    |       |
| 房屋地址/樓層面積         | text | YES  |     | NULL    |       |
| 總拍賣底價(元)            | text | YES  |     | NULL    |       |
| 點交                      | text | YES  |     | NULL    |       |
| 空屋                      | text | YES  |     | NULL    |       |
| 標別                      | text | YES  |     | NULL    |       |
| 備註                      | text | YES  |     | NULL    |       |
| 看圖                      | text | YES  |     | NULL    |       |
| 採通訊投標                | text | YES  |     | NULL    |       |
| 土地有無遭受污染*/

// Qunia is a function to Qurey unique area for sql select distict
//SELECT   distict RIGHT(city_area, 3)   FROM house.house_price;

// QureyPrice is a function form sql select * form house_price
func QureyPrice() []Price {
	o := orm.NewOrm()
	var Prices []Price
	//beego.Debug(Prices)
	_, err := o.Raw("SELECT *  FROM house_price").QueryRows(&Prices)
	if err != nil {
		fmt.Println(err)
	}
	//beego.Debug(Prices)
	return Prices
}
