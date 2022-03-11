/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  helper
 * @Version: 1.0.0
 * @Date: 2022/3/11 4:36 下午
 */

package segment

func decodeBit(raw uint64,bitBegin uint64,bitEnd uint64) uint64{
	bitCount:=bitEnd-bitBegin+1
	if bitCount<=0{
		return 0
	}
	var flag uint64=(1<<bitCount-1)
	flag=flag<<bitBegin
	return (raw&flag)>>bitBegin
}
