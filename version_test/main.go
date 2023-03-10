package main

/*	优点：
	1、扩展性强
	2、对已有方法修改幅度小
	3、可同时比较大或小
	解决问题：
	1、原有代码存在逻辑漏洞
	2、返回值的二义性
	3、代码冗余
*/
func main() {
	caseAllClient()
	caseBeiKeApp()
	caseBeiKeAppError()
	caseExtend()
}