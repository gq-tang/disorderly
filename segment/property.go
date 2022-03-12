/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  property
 * @Version: 1.0.0
 * @Date: 2022/3/11 4:16 下午
 */

package segment

import (
	"fmt"
)

type (
	GDTProperty struct {
		G    PropertyG
		DB   PropertyDB
		L    PropertyL
		AVL  PropertyAVL
		P    PropertyP
		DPL  PropertyDPL
		S    PropertyS
		Type GDTType
	}

	PropertyG    byte // Granularity 粒度
	PropertyDB   byte // Default operation size(0=16-bit segment;1=32-bit segment)
	PropertyL    byte // 64-bit code segment(IA-32e mode only)
	PropertyAVL  byte // Available for use by system software
	PropertyP    byte // Segment present
	PropertyDPL  byte // Descriptor privilege level
	PropertyS    byte // Descriptor type(0=system,1=code or data)
	TypeCodeData byte
	TypeSystem   byte
)

type GDTType interface {
	MarshalText() ([]byte, error)
	Value() byte
}

func NewProperty(g PropertyG, db PropertyDB, l PropertyL, avl PropertyAVL, p PropertyP, dpl PropertyDPL,
	s PropertyS, typ GDTType) GDTProperty {
	return GDTProperty{
		G:    g & 1,
		DB:   db & 1,
		L:    l & 1,
		AVL:  avl & 1,
		P:    p & 1,
		DPL:  dpl & 3,
		S:    s & 1,
		Type: typ,
	}
}

func (pr *GDTProperty) DecodeHigh(high uint64) {
	g := decodeBit(high, 23, 23)
	pr.G = PropertyG(g)

	db := decodeBit(high, 22, 22)
	pr.DB = PropertyDB(db)

	l := decodeBit(high, 21, 21)
	pr.L = PropertyL(l)

	avl := decodeBit(high, 20, 20)
	pr.AVL = PropertyAVL(avl)

	p := decodeBit(high, 15, 15)
	pr.P = PropertyP(p)

	dpl := decodeBit(high, 13, 14)
	pr.DPL = PropertyDPL(dpl)

	s := decodeBit(high, 12, 12)
	pr.S = PropertyS(s)

	typ := decodeBit(high, 8, 11)
	if s == 0 {
		pr.Type = TypeSystem(typ)
	} else {
		pr.Type = TypeCodeData(typ)
	}
}

var format = "value: %d,desc: %s"

func (p PropertyG) MarshalText() ([]byte, error) {
	var info string
	switch p {
	case 0:
		info = fmt.Sprintf(format, p, "1 byte")
	case 1:
		info = fmt.Sprintf(format, p, "4K byte")
	default:
		info = fmt.Sprintf(format, p, "invalid value")
	}
	return []byte(info), nil
}

func (p PropertyDB) MarshalText() ([]byte, error) {
	var info string
	switch p {
	case 0:
		info = fmt.Sprintf(format, p, "16-bit")
	case 1:
		info = fmt.Sprintf(format, p, "32-bit")
	default:
		info = fmt.Sprintf(format, p, "invalid value")
	}
	return []byte(info), nil
}

func (p PropertyL) MarshalText() ([]byte, error) {
	var info string
	switch p {
	case 0:
		info = fmt.Sprintf(format, p, "IA-32e mode only")

	default:
		info = fmt.Sprintf(format, p, "invalid value")
	}
	return []byte(info), nil
}

func (p PropertyP) MarshalText() ([]byte, error) {
	var info string
	switch p {
	case 0:
		info = fmt.Sprintf(format, p, "无效描述符")
	case 1:
		info = fmt.Sprintf(format, p, "有效")
	default:
		info = fmt.Sprintf(format, p, "invalid value")
	}
	return []byte(info), nil
}

func (p PropertyS) MarshalText() ([]byte, error) {
	var info string
	switch p {
	case 0:
		info = fmt.Sprintf(format, p, "system")
	case 1:
		info = fmt.Sprintf(format, p, "code/data")
	default:
		info = fmt.Sprintf(format, p, "invalid value")
	}
	return []byte(info), nil
}

func (t TypeCodeData) Value() byte {
	return byte(t)
}

func (t TypeCodeData) MarshalText() ([]byte, error) {
	var info string
	switch t {
	case 0:
		info = fmt.Sprintf(format, t, "数据_只读")
	case 1:
		info = fmt.Sprintf(format, t, "数据_只读,已访问")
	case 2:
		info = fmt.Sprintf(format, t, "数据_可读/写")
	case 3:
		info = fmt.Sprintf(format, t, "数据_可读/写,已访问")
	case 4:
		info = fmt.Sprintf(format, t, "数据_向下扩展,只读")
	case 5:
		info = fmt.Sprintf(format, t, "数据_向下扩展,只读,已访问")
	case 6:
		info = fmt.Sprintf(format, t, "数据_向下扩展,可读/写")
	case 7:
		info = fmt.Sprintf(format, t, "数据_向下扩展,可读/写,已访问")

	case 8:
		info = fmt.Sprintf(format, t, "代码_仅执行")
	case 9:
		info = fmt.Sprintf(format, t, "代码_仅执行,已访问")
	case 10:
		info = fmt.Sprintf(format, t, "代码_执行/可读")
	case 11:
		info = fmt.Sprintf(format, t, "代码_执行/可读,已访问")
	case 12:
		info = fmt.Sprintf(format, t, "代码_一致性段,仅执行")
	case 13:
		info = fmt.Sprintf(format, t, "代码_一致性段,仅执行,已访问")
	case 14:
		info = fmt.Sprintf(format, t, "代码_一致性段,执行/可读")
	case 15:
		info = fmt.Sprintf(format, t, "代码_一致性段,执行/可读,已访问")
	default:
		info = fmt.Sprintf(format, t, "invalid value")
	}
	return []byte(info), nil
}

func (t TypeSystem) Value() byte {
	return byte(t)
}

func (t TypeSystem) MarshalText() ([]byte, error) {
	var info string
	switch t {
	case 0:
		info = fmt.Sprintf(format, t, "系统_保留")
	case 1:
		info = fmt.Sprintf(format, t, "系统_16-Bit TSS(Available)")
	case 2:
		info = fmt.Sprintf(format, t, "系统_LDT")
	case 3:
		info = fmt.Sprintf(format, t, "系统_16-Bit TSS(Busy)")
	case 4:
		info = fmt.Sprintf(format, t, "系统_16-Bit Call Gate")
	case 5:
		info = fmt.Sprintf(format, t, "系统_Task Gate")
	case 6:
		info = fmt.Sprintf(format, t, "系统_16-Bit Interrupt Gate")
	case 7:
		info = fmt.Sprintf(format, t, "系统_16-Bit Trap Gate")

	case 8:
		info = fmt.Sprintf(format, t, "系统_保留")
	case 9:
		info = fmt.Sprintf(format, t, "系统_32-Bit TSS(Available)")
	case 10:
		info = fmt.Sprintf(format, t, "系统_保留")
	case 11:
		info = fmt.Sprintf(format, t, "系统_32-Bit TSS(Busy)")
	case 12:
		info = fmt.Sprintf(format, t, "系统_32-Bit Call Gate")
	case 13:
		info = fmt.Sprintf(format, t, "系统_保留")
	case 14:
		info = fmt.Sprintf(format, t, "系统_32-Bit Interrupt Gate")
	case 15:
		info = fmt.Sprintf(format, t, "系统_32-Bit Trap Gate")
	default:
		info = fmt.Sprintf(format, t, "invalid value")
	}
	return []byte(info), nil
}
