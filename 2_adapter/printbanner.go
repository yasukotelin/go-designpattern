package main

// PrintBanner がBannerとPrinterとのアダプターとなる構造体
type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(s string) *PrintBanner {
	return &PrintBanner{
		banner: NewBanner(s),
	}
}

func (pb *PrintBanner) PrintWeek() {
	pb.banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
