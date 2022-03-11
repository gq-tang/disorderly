/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  var
 * @Version: 1.0.0
 * @Date: 2022/3/11 4:28 下午
 */

package segment

const (
	GByte PropertyG=iota
	G4K
)

const (
	DB16Bit PropertyDB=iota
	DB32Bit
)

const (
	L32ModeOnly PropertyL=iota
)

const (
	AVLReserve PropertyAVL=iota
)

const (
	PUnValid PropertyP=iota
	PValid
)

const (
	DPLZero PropertyDPL=iota
	DPLOne
	DPLTwo
	DPLThree
)

const (
	SSystem PropertyS=iota
	SCodeData
)