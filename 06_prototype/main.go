package main

func main() {
	// 準備
	manager := NewManager()
	upen := NewUnderlinePen('-')
	mbox := NewMessageBox('*')
	sbox := NewMessageBox('/')
	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	// 生成テスト
	p1 := manager.Create("strong message")
	p2 := manager.Create("warning box")
	p3 := manager.Create("slash box")

	p1.Use("Hello, world")
	p2.Use("Hello, world")
	p3.Use("Hello, world")
}
