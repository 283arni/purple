package bins

import (
	"encoding/json"
	"fmt"
	"struct/storage"
	"time"
)

type BinsInterface interface {
	AddAccount(name string, file string)
	DeleteBin(binRes string)
	ToBytes() ([]byte, error)
	ReadBins() []Bin
	save() error
}

type Metadata struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Private   bool      `json:"private"`
}

type MetadataDelete struct {
	Id              string `json:"id"`
	VersionsDeleted int    `json:"versionsDeleted"`
}

type Record struct {
	Name string `json:"name"`
}

type Request struct {
	Record   Record   `json:"record"`
	Metadata Metadata `json:"metadata"`
}

type RequestDelete struct {
	Metadata MetadataDelete `json:"metadata"`
	Message  string         `json:"message"`
}

type Bin struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
}

type Bins struct {
	BinsArr  []Bin     `json:"bins"`
	UpdateAt time.Time `json:"updateAt"`
}

func CreateBin(binRes string) *Bin {
	var bin Request
	json.Unmarshal([]byte(binRes), &bin)

	return &Bin{
		Id:        bin.Metadata.Id,
		Name:      bin.Record.Name,
		Private:   bin.Metadata.Private,
		CreatedAt: bin.Metadata.CreatedAt,
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

func (bins *Bins) AddAccount(binRes string, file string) {
	bin := CreateBin(binRes)

	bArr := bins.ReadBins()
	bins.BinsArr = append(bArr, *bin)
	err := bins.save()

	if err != nil {
		fmt.Println("Проблема сохранения файла")
	}
}

func (bins *Bins) DeleteBin(binRes string) {
	var bin RequestDelete
	json.Unmarshal([]byte(binRes), &bin)

	bArr := bins.ReadBins()
	bCopy := []Bin{}

	for _, v := range bArr {
		if v.Id != bin.Metadata.Id {
			bCopy = append(bCopy, v)
		}
	}

	bins.BinsArr = bCopy

	err := bins.save()

	if err != nil {
		fmt.Println("Проблема сохранения файла")
		return
	}

	fmt.Println(bin.Message)
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
