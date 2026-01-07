package main

// ========== 產品介面 ==========
type IButton interface {
	Render() string
	GetStyle() string
}

type ITextField interface {
	Render() string
	GetBorder() string
}

type ICheckBox interface {
	Render() string
	GetShape() string
}

// ========== 工廠介面 ==========
type IGUIFactory interface {
	GetButton() IButton
	GetTextField() ITextField
	GetCheckBox() ICheckBox
}

// ========== Windows 產品族 ==========
type WindowsButton struct{}

func (b *WindowsButton) Render() string {
	return "Rendering Windows Button"
}

func (b *WindowsButton) GetStyle() string {
	return "Square Style"
}

type WindowsTextField struct{}

func (t *WindowsTextField) Render() string {
	return "Rendering Windows TextField"
}

func (t *WindowsTextField) GetBorder() string {
	return "Single Line Border"
}

type WindowsCheckBox struct{}

func (c *WindowsCheckBox) Render() string {
	return "Rendering Windows CheckBox"
}

func (c *WindowsCheckBox) GetShape() string {
	return "Square Shape"
}

// ========== Windows 工廠 ==========
type WindowsGUI struct{}

func (f *WindowsGUI) GetButton() IButton {
	return &WindowsButton{}
}

func (f *WindowsGUI) GetTextField() ITextField {
	return &WindowsTextField{}
}

func (f *WindowsGUI) GetCheckBox() ICheckBox {
	return &WindowsCheckBox{}
}

// ========== MacOS 產品族 ==========
type MacOSButton struct{}

func (b *MacOSButton) Render() string {
	return "Rendering MacOS Button"
}

func (b *MacOSButton) GetStyle() string {
	return "Rounded Style"
}

type MacOSTextField struct{}

func (t *MacOSTextField) Render() string {
	return "Rendering MacOS TextField"
}

func (t *MacOSTextField) GetBorder() string {
	return "Rounded Border"
}

type MacOSCheckBox struct{}

func (c *MacOSCheckBox) Render() string {
	return "Rendering MacOS CheckBox"
}

func (c *MacOSCheckBox) GetShape() string {
	return "Circle Shape"
}

// ========== MacOS 工廠 ==========
type MacOSGUI struct{}

func (f *MacOSGUI) GetButton() IButton {
	return &MacOSButton{}
}

func (f *MacOSGUI) GetTextField() ITextField {
	return &MacOSTextField{}
}

func (f *MacOSGUI) GetCheckBox() ICheckBox {
	return &MacOSCheckBox{}
}

// ========== 工廠選擇器 ==========
func FactoryOS(os string) IGUIFactory {
	switch os {
	case "Windows":
		return &WindowsGUI{}
	case "MacOS":
		return &MacOSGUI{}
	default:
		return nil
	}
}

// ========== 客戶端使用 ==========
func main() {
	// Windows 風格
	windowsFactory := FactoryOS("Windows")
	println("=== Windows UI ===")
	btn := windowsFactory.GetButton()
	println(btn.Render(), "-", btn.GetStyle())

	txt := windowsFactory.GetTextField()
	println(txt.Render(), "-", txt.GetBorder())

	chk := windowsFactory.GetCheckBox()
	println(chk.Render(), "-", chk.GetShape())

	// MacOS 風格
	macFactory := FactoryOS("MacOS")
	println("\n=== MacOS UI ===")
	btn2 := macFactory.GetButton()
	println(btn2.Render(), "-", btn2.GetStyle())

	txt2 := macFactory.GetTextField()
	println(txt2.Render(), "-", txt2.GetBorder())

	chk2 := macFactory.GetCheckBox()
	println(chk2.Render(), "-", chk2.GetShape())
}
