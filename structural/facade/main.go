package main

import "fmt"

// ===== Subsystems =====

// 照明系統
type LightingSystem struct{}

func (l *LightingSystem) TurnOn() {
	fmt.Println("💡 Lights turned on")
}

func (l *LightingSystem) TurnOff() {
	fmt.Println("💡 Lights turned off")
}

func (l *LightingSystem) SetBrightness(level int) {
	fmt.Printf("💡 Lights brightness set to %d%%\n", level)
}

// 空調系統
type AirConditioner struct{}

func (a *AirConditioner) TurnOn() {
	fmt.Println("❄️  Air conditioner turned on")
}

func (a *AirConditioner) TurnOff() {
	fmt.Println("❄️  Air conditioner turned off")
}

func (a *AirConditioner) SetTemperature(temp int) {
	fmt.Printf("❄️  Temperature set to %d°C\n", temp)
}

// 音響系統
type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("🔊 Sound system turned on")
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("🔊 Sound system turned off")
}

func (s *SoundSystem) SetVolume(volume int) {
	fmt.Printf("🔊 Volume set to %d%%\n", volume)
}

func (s *SoundSystem) PlayMusic(song string) {
	fmt.Printf("🎵 Playing: %s\n", song)
}

// 窗簾系統
type CurtainSystem struct{}

func (c *CurtainSystem) Open() {
	fmt.Println("🪟 Curtains opened")
}

func (c *CurtainSystem) Close() {
	fmt.Println("🪟 Curtains closed")
}

// ===== Facade =====

// 智慧家居外觀類
type SmartHomeFacade struct {
	lights  *LightingSystem
	ac      *AirConditioner
	sound   *SoundSystem
	curtain *CurtainSystem
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lights:  &LightingSystem{},
		ac:      &AirConditioner{},
		sound:   &SoundSystem{},
		curtain: &CurtainSystem{},
	}
}

// 離家模式
func (s *SmartHomeFacade) LeavingHome() {
	fmt.Println("\n🚪 Activating Leaving Home Mode...")
	s.lights.TurnOff()
	s.ac.TurnOff()
	s.sound.TurnOff()
	s.curtain.Close()
	fmt.Println("✅ All systems secured\n")
}

// 回家模式
func (s *SmartHomeFacade) ComingHome() {
	fmt.Println("\n🏠 Activating Coming Home Mode...")
	s.lights.TurnOn()
	s.lights.SetBrightness(80)
	s.ac.TurnOn()
	s.ac.SetTemperature(24)
	s.curtain.Open()
	fmt.Println("✅ Welcome home!\n")
}

// 電影模式
func (s *SmartHomeFacade) MovieMode() {
	fmt.Println("\n🎬 Activating Movie Mode...")
	s.lights.TurnOn()
	s.lights.SetBrightness(20)
	s.curtain.Close()
	s.sound.TurnOn()
	s.sound.SetVolume(60)
	s.sound.PlayMusic("Movie Soundtrack")
	fmt.Println("✅ Enjoy your movie!\n")
}

// 睡眠模式
func (s *SmartHomeFacade) SleepMode() {
	fmt.Println("\n😴 Activating Sleep Mode...")
	s.lights.TurnOff()
	s.sound.TurnOff()
	s.ac.TurnOff()
	s.curtain.Close()
	fmt.Println("✅ Good night!\n")
}

// 派對模式
func (s *SmartHomeFacade) PartyMode() {
	fmt.Println("\n🎉 Activating Party Mode...")
	s.lights.TurnOn()
	s.lights.SetBrightness(100)
	s.sound.TurnOn()
	s.sound.SetVolume(85)
	s.sound.PlayMusic("Party Playlist")
	s.ac.TurnOn()
	s.ac.SetTemperature(22)
	fmt.Println("✅ Let's party!\n")
}

func main() {
	fmt.Println("===== Smart Home Facade Pattern Example =====")

	// 創建智慧家居 Facade
	smartHome := NewSmartHomeFacade()

	// 回家場景
	smartHome.ComingHome()

	// 電影場景
	smartHome.MovieMode()

	// 派對場景
	smartHome.PartyMode()

	// 睡眠場景
	smartHome.SleepMode()

	// 離家場景
	smartHome.LeavingHome()

	fmt.Println("===== Facade Pattern Benefits =====")
	fmt.Println("✓ 簡化了複雜子系統的操作")
	fmt.Println("✓ 客戶端不需要知道各個子系統的細節")
	fmt.Println("✓ 一個方法就能控制多個子系統")
	fmt.Println("✓ 降低了客戶端與子系統之間的耦合度")
}
