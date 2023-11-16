package main

import "fmt"

type shangping struct {
	name  string
	price float64
	stock int
}
type stockmanager interface {
	checkstock()
	UpdateStock(newStock int)
	PrintStockInfo()
}

func (s *shangping) checkstock() {
	fmt.Print("库存数量：%v\n", s.stock)
}
func (s *shangping) UpdateStock(newStock int) {
	s.stock = newStock
	fmt.Printf("库存数量已更新为：%v\n", s.stock)
}
func (s *shangping) PrintStockInfo() {
	fmt.Printf("商品名称：%v\n价格：%v\n库存数量：%v\n", s.name, s.price, s.stock)
}

type ElectronicProduct struct {
	shangping
	pingpan string
	xinhao  string
}

func (e *ElectronicProduct) ElectronicCheckStock() {
	fmt.Printf("电子产品库存数量：%v\n", e.stock)
}

func (e *ElectronicProduct) ElectronicUpdateStock(newStock int) {
	e.stock = newStock
	fmt.Printf("电子产品库存数量已更新为：%v\n", e.stock)
}
func (e *ElectronicProduct) PrintpingpanxinhaoInfo() {
	fmt.Printf("商品名称：%v\n价格：%v\n库存数量：%v\n品牌：%v\n型号：%v\n",
		e.name, e.price, e.stock, e.pingpan, e.xinhao)
}

func main() {
	book := shangping{name: "书", price: 29.99, stock: 50}
	book.PrintStockInfo()
	book.UpdateStock(45)
	laptop := ElectronicProduct{
		shangping: shangping{name: "kun", price: 999.99, stock: 20},
		pingpan:   "xiaoheizi",
		xinhao:    "ikun",
	}
	laptop.PrintpingpanxinhaoInfo()
	laptop.ElectronicUpdateStock(15)
}
