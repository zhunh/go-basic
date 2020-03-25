package demo_struct

import "fmt"

type Animal struct{
	Color string
}

type Dog struct {
	Animal
	ID int
	Name string
	Age int
}

func (d *Dog)Eat() string {
	fmt.Println("dog --> yummy yummy")
	return "dog --> yummy yummy!"
}

func (d *Dog)Run() string {
	fmt.Println("ID:",d.ID,"Dog is running")
	return "Dog is running"
}

type Cat struct {
	Animal
	ID int
	Name string
	Age int
}

func (c *Cat)Eat() string {
	fmt.Println("cat --> yummy yummy")
	return "cat -> yummy yummy!"
}

func (c *Cat)Run() string {
	fmt.Println("ID:",c.ID,"Cat is running")
	return "Cat is running"
}

func Test()  {
	//方式一
	//var dog Dog
	//dog.ID = 101
	//dog.Name = "wangcai"
	//dog.Age = 2
	//方式二
	//dog := Dog{ID:1,Name:"ui",Age:3}
	//方式三
	//dog := new(Dog)
	//dog.ID = 101
	//dog.Name = "wangcai"
	//dog.Age = 2
	//fmt.Println("dog", dog)

	//测试接口定义变量
	dog := new(Dog)
	cat := new(Cat)
	action(dog)
	action(cat)

}

func action(b Behavior) string {
	b.Run()
	b.Eat()
	return ""
}
