package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

//id product
var (
	id    int
	mutex sync.Mutex
)

//Product struct -
type Product struct {
	Code        string
	Name        string
	Description string
	Image       string
	Quantity    int
	Price       float64
}

//MarshalBinary -
func (p *Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

//UnmarshalBinary -
func (p *Product) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	return nil
}

var (
	client, errconn = NewClient()
)

//InsertProduct -
func InsertProduct(p Product) string {
	if errconn != nil {
		return "can't connect to database server"
	}

	cacheKey := "product"
	cacheId := strconv.Itoa(id)

	mutex.Lock()
	id++
	mutex.Unlock()

	val, _ := p.MarshalBinary()

	if err := client.HSet(cacheKey, cacheId, val).Err(); err != nil {
		return "err"
	}

	return "insert success"
}

//GetProduct -
func GetProduct(cacheKey string, id int) (Product, error) {
	var result Product

	if errconn != nil {
		return result, errconn
	}

	data, err := client.HGet(cacheKey, strconv.Itoa(id)).Result()

	if err != nil {
		panic(err)
	}

	if err := result.UnmarshalBinary([]byte(data)); err != nil {
		return result, err
	}

	return result, nil
}

//GetAllProduct -
func GetAllProduct() ([]Product, error) {
	var result []Product

	data, err := client.Do("product", "*").Result()

	if err != nil {
		return result, err
	}

	fmt.Println("DATA: ", data)

	// for key, value := range data {
	// 	if err := result.UnmarshalBinary([]byte(data)); err != nil {

	// 	}
	// }

	return result, nil
}
