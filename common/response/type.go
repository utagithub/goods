/*
 * @Author: 尤_Ta
 * @Date:  23:13
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:13
 */

package response

type Responses interface {
	SetCode(int32)
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses
}
