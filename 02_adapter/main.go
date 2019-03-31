package main

func main() {
	// interfaceを用いたアダプターパターン実装のサンプル
	var printer Printer = NewPrintBanner("hello")
	printer.PrintWeek()
	printer.PrintStrong()

	// 委譲を用いたアダプターパターン実装のサンプル
}
