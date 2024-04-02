package util

import (
	"context"
	"fmt"
)

type Printer interface {
	Print(ctx context.Context, toPrint string)
}

type BasePrinter struct {
}

func NewPrinter() Printer {
	return BasePrinter{}
}

func (bp BasePrinter) Print(ctx context.Context, toPrint string) {
	fmt.Println(toPrint)
}

type VerbosePrinterDecorator struct {
	successor Printer
}

func NewVerbosePrinterDecorator(successor Printer) VerbosePrinterDecorator {
	return VerbosePrinterDecorator{successor: successor}
}

func (vp VerbosePrinterDecorator) Print(ctx context.Context, toPrint string) {
	stream := ctx.Value("stream").(string)
	switch stream {
	case "in":
		fmt.Print("< ")
	case "out":
		fmt.Print("> ")
	}
	vp.successor.Print(ctx, toPrint)
}
