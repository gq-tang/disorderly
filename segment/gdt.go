/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  gdt
 * @Version: 1.0.0
 * @Date: 2022/3/11 11:02 上午
 */

package segment

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unsafe"
)

type (
	GDT struct{
	Raw RawData `json:"raw"`
	high uint64
	Base uint64 `json:"base"`
	Limit uint64 `json:"limit"`
	Property GDTProperty `json:"property"`
}
	RawData uint64
)

func FromUint64(v uint64) *GDT{
	gdt:=&GDT{
		Raw: RawData(v),
		high: v>>32,
		Property: GDTProperty{},
	}
	gdt.decodeBase()
	gdt.decodeLimit()
	gdt.Property.DecodeHigh(gdt.high)
	return gdt
}

func NewGDT(base uint32,limit uint32,p GDTProperty) (*GDT){
	g:=GDT{
		Base:     uint64(base),
		Limit:    uint64(limit),
		Property: p,
	}

	g.encode()
	return &g
}

func (gdt *GDT) encode(){
	var high uint64
	high|=decodeBit(gdt.Base,24,31)<<24
	high|=uint64(gdt.Property.G)<<23
	high|=uint64(gdt.Property.DB)<<22
	high|=uint64(gdt.Property.L)<<21
	high|=uint64(gdt.Property.AVL)<<20
	high|= decodeBit(gdt.Limit,16,19)<<16
	high|=uint64(gdt.Property.P)<<15
	high|=uint64(gdt.Property.DPL)<<13
	high|=uint64(gdt.Property.S)<<12
	high|=uint64(gdt.Property.Type.Value())<<8
	high|=decodeBit(gdt.Base,16,23)

	var low uint64
	low|=decodeBit(gdt.Base,0,15)<<16
	low|=decodeBit(gdt.Limit,0,15)

	gdt.Raw=RawData(high<<32+low)
	gdt.high=high
}

func (gdt *GDT) decodeBase(){
	b31_24:=decodeBit(gdt.high,24,31)<<24
	b23_0:=decodeBit(uint64(gdt.Raw),16,38)>>16
	gdt.Base=b31_24+b23_0
}

func (gdt *GDT) decodeLimit(){
	l19_16:=decodeBit(gdt.high,16,19)<<16
	l15_0:=decodeBit(uint64(gdt.Raw),0,15)
	gdt.Limit=l19_16+l15_0
}

func (gdt *GDT) PrintBits() {
	gdt.printHigh()
	gdt.printLow()
}

func (gdt *GDT) PrintJSON() {
	b,err:=json.MarshalIndent(gdt,""," ")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b))
}

func (gdt *GDT) printHigh(){
	fmt.Println("|                      |  |D |  |A |   Seg.    |  |  D  |  |         |              |")
	fmt.Println("|   Base 31:24         |G |/ |L |V |   Limit   |P |  P  |S |   Type  |  Base 23:16  |")
	fmt.Println("|                      |  |B |  |L |   19:16   |  |  L  |  |         |              |")
	fmt.Println("31 30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0")
	var i int64
	for i=31;i>=0;i--{
		bit:=decodeBit(gdt.high,uint64(i),uint64(i))
		if i>9{
			fmt.Printf("%-2d ",bit)
		}else {
			fmt.Printf("%-1d ",bit)
		}

	}
	fmt.Println()
}

func (gdt *GDT) printLow(){
	fmt.Println("|              Base Address 15:00              |           Segment Limit 15:00      |")
	fmt.Println("31 30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0")
	var i int64
	for i=31;i>=0;i--{
		bit:=decodeBit(uint64(gdt.Raw),uint64(i),uint64(i))
		if i>9{
			fmt.Printf("%-2d ",bit)
		}else {
			fmt.Printf("%-1d ",bit)
		}

	}
	fmt.Println()
}

func (r RawData) MarshalText()([]byte,error){
	rp:=(*[8]byte) (unsafe.Pointer(&r))
	hexS:=hex.EncodeToString((*rp)[:])
	return []byte(hexS),nil
}