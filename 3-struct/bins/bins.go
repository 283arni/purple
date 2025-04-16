package bins

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"struct/storage"
	"time"
)

type Bin struct {
	Id       string    `json:"id"`
	Private  bool      `json:"private"`
	CreateAt time.Time `json:"createAt"`
	Name     string    `json:"name"`
}

type Bins struct {
	BinsArr  []Bin     `json:"bins"`
	UpdateAt time.Time `json:"updateAt"`
}

func CreateBin() *Bin {
	var private, name string
	var isPrivate bool
	fmt.Println("Будет ли bin приватный Y/n?")
	fmt.Scan(&private)
	fmt.Println("Введите имя:")
	fmt.Scan(&name)

	if private == "Y" || private == "y" {
		isPrivate = true
	} else {
		isPrivate = false
	}

	return &Bin{
		Id:       name + strconv.Itoa(rand.Intn(100)),
		Private:  isPrivate,
		CreateAt: time.Now(),
		Name:     name,
	}
}

func NewBins() *Bins {
	data, err := storage.ReadFile("storage/data.json")

	if err != nil {
		return &Bins{
			BinsArr:  []Bin{},
			UpdateAt: time.Now(),
		}
	}

	var bins Bins

	err = json.Unmarshal(data, &bins)

	if err != nil {
		fmt.Println(err)
		return &Bins{
			BinsArr:  []Bin{},
			UpdateAt: time.Now(),
		}
	}

	return &bins
}

func (bins *Bins) AddAccount() {
	bin := CreateBin()

	bins.BinsArr = append(bins.BinsArr, *bin)
	bins.save()
}

func (bins *Bins) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bins)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (bins *Bins) ReadBins() {
	data, err := storage.ReadFile("storage/data.json")

	if err != nil {
		fmt.Println(err)
	}

	var b Bins

	err = json.Unmarshal(data, &b)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range b.BinsArr {
		fmt.Println(v.Name)
		fmt.Println(v.CreateAt)
	}

}

func (bins *Bins) save() {
	bins.UpdateAt = time.Now()
	data, err := bins.ToBytes()

	if err != nil {
		fmt.Println("Не удалось приобразовать данные")
	}

	storage.WriteFile(data, "storage/data.json")
}
