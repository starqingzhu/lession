package mytype

import "fmt"

/*

引用类型:切片、映射、通道、接口和函数
内置类型:字符串、bool、数值类型
结构类型:

用户自定义类型
方法
参数: 指针和值
接口+多态:指代码根据类型的具体实现采取不同行为的能力
组合+扩展类型
标识符
接口:是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现。


嵌入类型:目的代码复用
	注意:内部类型+外部类型，内部类型的方法在外部类型没有实现时，自动提升为外部类型，一旦外部类型实现了同样方法，内部类型则不提升为外部类型也不会消失。

公开和未公开的标识符:
	1.公开或者未公开的标识符，不是一个值
	2.短变量声明操作符，有能力捕获引用的类型，并创建一个未公开的类型变量

*/

type (
	//创建自定义类型
	User struct {
		name          string
		email         string
		ext           int
		intprivileged bool
	}

	//使用已有自定义类型组合新类型
	Admin struct {
		person User
		level  string
	}

	//嵌入类型
	Manager struct {
		User
		Level string
	}

	//基于已有类型重新定义类型
	Duration int64
)

func (u *User)Printp(){
	fmt.Printf("User:%+v\n",u)
}

func (ad Admin) Print(){
	fmt.Printf("Admin:---->\npserson:%+v\nlevel:%s\n",ad.person,ad.level)
}
func (ad *Admin) Printp(){
	fmt.Printf("*Admin:---->\npserson:%+v\nlevel:%s\n",ad.person,ad.level)
}

func Init() {
	fmt.Println("---------type init----------")
	var user User = User{
		name:          "Mr.sun",
		email:         "starqingzhu@163.com",
		ext:           0,
		intprivileged: true,
	}
	fmt.Printf("user实例:\n%+v\n", user)

	var user1 User
	fmt.Printf("user实例0值:\n%+v\n", user1)

	ad := Admin{
		person: user,
		level:  "super",
	}
	fmt.Printf("admin实例:\n%+v\n", ad)

	//var tm Duration
	//tm = int64(1000)
	//fmt.Printf("redefine system type int64,tm=%d\n",tm)
}

func Func(){
	var user User = User{
		name:          "Mr.sun",
		email:         "starqingzhu@163.com",
		ext:           0,
		intprivileged: true,
	}
	ad := Admin{
		person: user,
		level:  "super",
	}
	ad.Print()
	ad.Printp()
}


type Notifier interface{
	Notify()
}

type (
	IUser struct {
		name		string
		email		string
	}
	IAdmin struct {
		name		string
		email		string
	}
)

func (u *IUser) Notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}


func sendNotification(n Notifier){
	n.Notify()
}

func TestInterfaceNotify(){
	fmt.Println("---------test interface----------")
	u := IUser{
		name:  "sun",
		email: "starqingzhu@163.com",
	}
	fmt.Printf("IUser:\n%+v\n",u)
	sendNotification(&u)
}

func (a *IAdmin)Notify(){
	fmt.Printf("Sending user email to %s<%s>\n",a.name,a.email)
}

func TestMutiInterfaceNotify(){
	fmt.Println("---------test muti interface----------")
	u := IUser{
		name:  "sun",
		email: "starqingzhu@163.com",
	}
	a := IAdmin{
		name:  "a_sun",
		email: "starqingzhu@163.com",
	}
	fmt.Printf("IUser:\n%+v\nIAdmin:\n%+v\n",u,a)
	sendNotification(&u)
	sendNotification(&a)
}

func TestInsertion(){
	fmt.Println("---------test insertion----------")
	m := Manager{
		User:  User{
			name:          "m.Mr.sun",
			email:         "m.starqingzhu@163.com",
			ext:           0,
			intprivileged: true,
		},
		Level: "1",
	}
	fmt.Printf("Manager:\n%+v\n",m)
	m.User.Printp()
}

type alertCounter int
func New(value int)alertCounter{
	return alertCounter(value)
}

type (
	user1 struct {
		Name string
		email string
	}
	Admin1 struct {
		user1
		Level int
	}

	user2 struct {
		Name string
		Email string
	}
	Admin2 struct {
		user2
		Level int
	}

)


