package bins

import (
	"encoding/json"
	"fmt"
	"struct/storage"
	"time"

	"github.com/google/uuid"
)

type BinsInterface interface {
	NewBins() *Bins
	AddAccount(name string, isPrivate bool)
	ToBytes() ([]byte, error)
	ReadBins() []Bin
	save() error
}

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

func CreateBin(name string, isPrivate bool) *Bin {

	return &Bin{
		Id:       uuid.New().String(),
		Private:  isPrivate,
		CreateAt: time.Now(),
		Name:     name,
	}
}

func (b *Bins) NewBins() *Bins {
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

func (bins *Bins) AddAccount(name string, isPrivate bool) {
	bin := CreateBin(name, isPrivate)

	bins.BinsArr = append(bins.BinsArr, *bin)
	err := bins.save()

	if err != nil {
		fmt.Println("Проблема сохранения файла")
	}
}

func (bins *Bins) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bins)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (bins *Bins) ReadBins() []Bin {
	data, err := storage.ReadFile("storage/data.json")

	if err != nil {
		fmt.Println(err)
	}

	var b Bins

	err = json.Unmarshal(data, &b)

	if err != nil {
		fmt.Println(err)
	}

	return b.BinsArr

}

func (bins *Bins) save() error {
	bins.UpdateAt = time.Now()
	data, err := bins.ToBytes()

	if err != nil {
		return err
	}

	storage.WriteFile(data, "storage/data.json")
	return nil
}
