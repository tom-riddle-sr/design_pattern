package main

// 實例

type Storage interface{}

type LocalStorage struct{}
type CloudStorage struct{}

// 抽象

type Note interface{}

type SimpleNote struct {
	storage Storage
}

type EncryptedNote struct {
	storage Storage
}

func NewSimpleNote(storage Storage) *SimpleNote {
	return &SimpleNote{storage: storage}
}

func NewEncryptedNote(storage Storage) *EncryptedNote {
	return &EncryptedNote{storage: storage}
}

func choseConfig(storageType string, noteType string) interface{} {
	var storage Storage

	// 選擇存儲方式
	if storageType == "local" {
		storage = &LocalStorage{}
	} else if storageType == "cloud" {
		storage = &CloudStorage{}
	}

	// 選擇筆記類型
	if noteType == "simple" {
		return NewSimpleNote(storage)
	} else if noteType == "encrypted" {
		return NewEncryptedNote(storage)
	}

	return nil
}

func main() {
	note1 := choseConfig("local", "simple")
	_ = note1

	note2 := choseConfig("cloud", "encrypted")
	_ = note2
}
