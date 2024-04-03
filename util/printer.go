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

var basePrinter *BasePrinter

func NewBasePrinter() BasePrinter {
	if basePrinter != nil {
		return *basePrinter
	}
	basePrinter = &BasePrinter{}
	return *basePrinter
}

func (bp BasePrinter) Print(ctx context.Context, toPrint string) {
	fmt.Println(toPrint)
}

type VerbosePrinterDecorator struct {
	successor Printer
}

var verbosePrinterDecorator *VerbosePrinterDecorator

func NewVerbosePrinterDecorator(successor Printer) VerbosePrinterDecorator {
	if verbosePrinterDecorator != nil {
		return *verbosePrinterDecorator
	}
	verbosePrinterDecorator := &VerbosePrinterDecorator{successor: successor}
	return *verbosePrinterDecorator
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
