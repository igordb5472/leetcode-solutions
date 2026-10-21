package main

type H2O struct {
	new, oxygen, hydrogen chan struct{}
}

func NewH2O() *H2O {
	new, oxygen, hydrogen := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	new <- struct{}{}
	h := &H2O{new, oxygen, hydrogen}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	select {
	case <-h.new:
		releaseHydrogen()
		h.oxygen <- struct{}{}
	case <-h.hydrogen:
		releaseHydrogen()
		h.new <- struct{}{}
	}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oxygen
	releaseOxygen()
	h.hydrogen <- struct{}{}
}
