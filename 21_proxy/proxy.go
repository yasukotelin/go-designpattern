package main

type PrinterProxy struct {
	name string
	real Printable
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{
		name: name,
	}
}

func (pp *PrinterProxy) SetPrinterName(name string) {
	pp.name = (name)

	if pp.real != nil {
		pp.real.SetPrinterName(name)
	}
}

func (pp *PrinterProxy) GetPrinterName() string {
	return pp.name
}

func (pp *PrinterProxy) Print(s string) {
	if pp.real == nil {
		pp.real = NewPrinter(pp.name)
	}
	pp.real.Print(s)
}
