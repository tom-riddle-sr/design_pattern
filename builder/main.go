package main

import "fmt"

// Builder：定義「用什麼材料」
// Director：定義「建什麼組合」

type Computer struct {
	CPU         string
	RAM         string
	Storage     string
	GPU         string
	PowerSupply string
}

type IComputerBuilder interface {
	BuildCPU()
	BuildRAM()
	BuildStorage()
	BuildGPU()
	BuildPowerSupply()
	GetComputer() *Computer
}

type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{computer: &Computer{}}
}

func (b *GamingComputerBuilder) BuildCPU() {
	b.computer.CPU = "高端 CPU"
}

func (b *GamingComputerBuilder) BuildRAM() {
	b.computer.RAM = "32GB RAM"
}

func (b *GamingComputerBuilder) BuildStorage() {
	b.computer.Storage = "高速 SSD"
}

func (b *GamingComputerBuilder) BuildGPU() {
	b.computer.GPU = "RTX 4090"
}

func (b *GamingComputerBuilder) BuildPowerSupply() {
	b.computer.PowerSupply = "850W Gold PSU"
}

func (b *GamingComputerBuilder) GetComputer() *Computer {
	computer := b.computer
	b.computer = &Computer{} // 重置 builder 狀態
	return computer
}

type OfficeComputerBuilder struct {
	computer *Computer
}

func NewOfficeComputerBuilder() *OfficeComputerBuilder {
	return &OfficeComputerBuilder{computer: &Computer{}}
}

func (b *OfficeComputerBuilder) BuildCPU() {
	b.computer.CPU = "普通 CPU"
}

func (b *OfficeComputerBuilder) BuildRAM() {
	b.computer.RAM = "16GB RAM"
}

func (b *OfficeComputerBuilder) BuildStorage() {
	b.computer.Storage = "標準 SSD"
}

func (b *OfficeComputerBuilder) BuildGPU() {
	b.computer.GPU = "集成顯卡"
}

func (b *OfficeComputerBuilder) BuildPowerSupply() {
	b.computer.PowerSupply = "500W Bronze PSU"
}

func (b *OfficeComputerBuilder) GetComputer() *Computer {
	computer := b.computer
	b.computer = &Computer{} // 重置 builder 狀態
	return computer
}

type Director struct {
	builder IComputerBuilder
}

func NewDirector(builder IComputerBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildBudgetComputer() {
	d.builder.BuildCPU()
	d.builder.BuildRAM()
	d.builder.BuildStorage()
}
func (d *Director) BuildMidRangeComputer() {
	d.builder.BuildCPU()
	d.builder.BuildRAM()
	d.builder.BuildStorage()
	d.builder.BuildGPU()
}
func (d *Director) BuildHighEndComputer() {
	d.builder.BuildCPU()
	d.builder.BuildRAM()
	d.builder.BuildStorage()
	d.builder.BuildGPU()
	d.builder.BuildPowerSupply()
}

// Show: 輸出電腦配置
func (c *Computer) Show() {
	fmt.Println("CPU:", c.CPU)
	fmt.Println("RAM:", c.RAM)
	fmt.Println("Storage:", c.Storage)
	if c.GPU != "" {
		fmt.Println("GPU:", c.GPU)
	}
	if c.PowerSupply != "" {
		fmt.Println("Power:", c.PowerSupply)
	}
}

func main() {
	// 建造一台高端遊戲電腦
	gamingBuilder := NewGamingComputerBuilder()
	director := NewDirector(gamingBuilder)
	director.BuildHighEndComputer()
	gamingComputer := gamingBuilder.GetComputer()
	fmt.Println("遊戲電腦配置：")
	gamingComputer.Show()

	// 建造一台中端辦公電腦
	officeBuilder := NewOfficeComputerBuilder()
	director = NewDirector(officeBuilder)
	director.BuildMidRangeComputer()
	officeComputer := officeBuilder.GetComputer()
	fmt.Println("辦公電腦配置：")
	officeComputer.Show()
}
