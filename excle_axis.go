package gutils

import (
	"strconv"
)

//excel col cycleNum
const (
	cycleNum int = 26
	maxCol   int = 702
)

//col index
const (
	a = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
)

//col chars
const (
	ac = "A"
	bc = "B"
	cc = "C"
	dc = "D"
	ec = "E"
	fc = "F"
	gc = "G"
	hc = "H"
	ic = "I"
	jc = "J"
	kc = "K"
	lc = "L"
	mc = "M"
	nc = "N"
	oc = "O"
	pc = "P"
	qc = "Q"
	rc = "R"
	sc = "S"
	tc = "T"
	uc = "U"
	vc = "V"
	wc = "W"
	xc = "X"
	yc = "Y"
	zc = "Z"
)

func getColumnKey(colNo int) string {
	switch colNo {
	case a:
		return ac
	case b:
		return bc
	case c:
		return cc
	case d:
		return dc
	case e:
		return ec
	case f:
		return fc
	case g:
		return gc
	case h:
		return hc
	case i:
		return ic
	case j:
		return jc
	case k:
		return kc
	case l:
		return lc
	case m:
		return mc
	case n:
		return nc
	case o:
		return oc
	case p:
		return pc
	case q:
		return qc
	case r:
		return rc
	case s:
		return sc
	case t:
		return tc
	case u:
		return uc
	case v:
		return vc
	case w:
		return wc
	case x:
		return xc
	case y:
		return yc
	case z:
		return zc
	}

	return ""
}

//只考虑26*26+26=676列，一般报表不会超过这个列数
func calcColChar(colNo int) string {

	if colNo < cycleNum {
		return getColumnKey(colNo)
	}

	if colNo+1 >= maxCol {
		return zc + zc
	}

	ss := make([]byte, 2)

	cs := colNo + 1
	//余数
	ys := cs % cycleNum

	ci := 0
	if ys == 0 {
		ci = z
	} else {
		ci = ys - 1
	}

	ss[1] = getColumnKey(ci)[0]

	hi := (cs - cycleNum - ys) / cycleNum
	if ys == 0 {
		hi--
	}

	ss[0] = getColumnKey(hi)[0]

	return string(ss)
}

func CalcAxis(ri, ci int) string {

	colKey := calcColChar(ci)
	rowKey := strconv.Itoa(ri + 1)

	return colKey + rowKey
}

const EG = "$"

func CalcAbsAxis(ri, ci int) string {

	colKey := calcColChar(ci)
	rowKey := strconv.Itoa(ri + 1)

	return EG + colKey + EG + rowKey
}

//excel sheet axis computer
type SheetAxiser struct {
	span map[string]bool
}

func (this *SheetAxiser) Span(ri, ci, rs, cs int) {
	for r := ri; r <= ri+rs; r++ {
		for c := ci; c <= ci+cs; c++ {
			if r == ri && c == ci {
				continue
			} else {
				//非起始位置
				key := strconv.Itoa(r) + "_" + strconv.Itoa(c)
				this.span[key] = true
			}

		}
	}
}

func (this *SheetAxiser) IsSpan(ri, ci int) bool {
	key := strconv.Itoa(ri) + "_" + strconv.Itoa(ci)
	if v, ok := this.span[key]; ok {
		return v
	}
	return false
}

//ri is row index
//ci is col index
//isAbs if true return $
func (this *SheetAxiser) Axis(ri, ci int, isAbs bool) string {
	if isAbs {
		return CalcAbsAxis(ri, ci)
	} else {
		return CalcAxis(ri, ci)
	}
}
