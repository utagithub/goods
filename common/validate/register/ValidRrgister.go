/*
 * @Author: 尤_Ta
 * @Date:  22:12
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:12
 */

package register

/*
 * ValidRegister 收集所有自定义验证器-方便在engine中注册
 * 新增的自定义验证器,直接在 ValidRegister方法中调用即可
 */

func ValidRegister() {
	GoodValidRrgister{}.goodValidRrgister()
}
