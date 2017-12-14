package firewall

type Firewall struct {
	Layers map[int]*Layer
}

type Layer struct {
	Depth      int
	ScannerPos int
	Forward    bool
}

func (f *Firewall) Move() {
	for _, v := range f.Layers {
		if v.Forward {
			if v.ScannerPos == (v.Depth - 1) {
				v.Forward = false
			}
		} else {
			if v.ScannerPos == 0 {
				v.Forward = true
			}
		}
		if v.Forward {
			v.ScannerPos++
		} else {
			v.ScannerPos--
		}
	}
}
