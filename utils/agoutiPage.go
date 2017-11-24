// +build !linux

package utils

import (
	"fmt"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

type AgoutiPage struct {
	*agouti.Page
}
// SwitchToRootFrame focuses on the original, default page frame before any calls
// to Selection.Frame were made. After switching, all new and existing selections
// will refer to the root frame. All further Page methods will apply to this frame
// as well.
func (p *AgoutiPage) SwitchToRootFrameByName(frame *api.Element) error {
	if err := p.Session().Frame(frame); err != nil {
		return fmt.Errorf("failed to switch to original page frame: %s", err)
	}
	return nil
}