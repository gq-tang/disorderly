/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  var
 * @Version: 1.0.0
 * @Date: 2022/3/11 4:28 下午
 */

package segment

const (
	GByte PropertyG = iota
	G4K
)

const (
	DB16Bit PropertyDB = iota
	DB32Bit
)

const (
	L32ModeOnly PropertyL = iota
)

const (
	AVLReserve PropertyAVL = iota
)

const (
	PUnValid PropertyP = iota
	PValid
)

const (
	DPLZero PropertyDPL = iota
	DPLOne
	DPLTwo
	DPLThree
)

const (
	SSystem PropertyS = iota
	SCodeData
)

const (
	DataRO   TypeCodeData = iota // Read-Only
	DataROA                      // Read-Only,accessed
	DataRW                       // Read/Write
	DataRWA                      // Read/Write,accessed
	DataROE                      // Read-Only,expand-down
	DataROEA                     // Read-Only,expand-down,accessed
	DataRWE                      // Read/Write,expand-down
	DataRWEA                     // Read/Write,expand-down,accessed

	CodeE   // Execute-Only
	CodeEA  // Execute-Only,accessed
	CodeER  // Execute/Read
	CodeERA // Execute/Read,accessed
	CodeEC  // Execute-Only,conforming
	CodeECA // Execute-Only,conforming,accessed
	CodeRC  // Execute/Read-Only,conforming
	CodeRCA // Execute/Read-Only,conforming,accessed
)
