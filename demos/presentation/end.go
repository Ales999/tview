package main

import (
	"fmt"

	"github.com/ales999/tcell/v2"
	"github.com/ales999/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	url := "[:::https://github.com/ales999/tview]https://github.com/ales999/tview"
	fmt.Fprint(textView, url)
	return "End", Center(tview.TaggedStringWidth(url), 1, textView)
}
