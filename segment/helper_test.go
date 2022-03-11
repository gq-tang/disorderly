/**
 * @Author: tangguangqiang
 * @Description:
 * @File:  helper_test
 * @Version: 1.0.0
 * @Date: 2022/3/11 4:38 下午
 */

package segment

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func Test_DecodeBit(t *testing.T) {
	type param struct {
		raw   uint64
		begin uint64
		end   uint64
	}
	tests := []struct {
		give param
		want uint64
	}{
		{
			give: param{
				raw:   0b0000_1100,
				begin: 2,
				end:   3,
			},
			want: 0b11,
		},
		{
			give: param{
				raw:   0b0101_1100,
				begin: 3,
				end:   3,
			},
			want: 0b1,
		},
	}
	for i, test := range tests {
		name:=fmt.Sprintf("test:%d",i)
		t.Run(name,func(t *testing.T){
			p:=test.give
			res:=decodeBit(p.raw,p.begin,p.end)
			assert.Equal(t,test.want,res)
		})
	}
}
