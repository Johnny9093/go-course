package internal

type (
	MyStruct struct {
		FirstName string
		LastName string
		age int32
		internalStr internalStruct
	}

	internalStruct struct {
		Field1 int
	}
)

//func Something(){
//	myStruct:=MyStruct{"Yoni", "Lahav", 23}
//	fmt.Printf("%+v", myStruct)
//}