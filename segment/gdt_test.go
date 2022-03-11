/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  gdt_test
 * @Version: 1.0.0
 * @Date: 2022/3/11 1:48 下午
 */

package segment

import (
	"testing"
)

func Test_GDT(t *testing.T) {
	v := FromUint64(0x00c0_9a00_0000_07ff)
	//v.PrintBits()
	v.PrintJSON()
}

func Test_GDTEncode(t *testing.T) {
	p:=NewProperty(G4K,DB32Bit,L32ModeOnly,AVLReserve,PValid,DPLZero,
		SCodeData,TypeCodeData(10))
	gdt:=NewGDT(0,2047,p)
	gdt.PrintJSON()
}