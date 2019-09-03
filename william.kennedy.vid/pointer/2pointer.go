// 2.3 Pointer - Part 3 (Escape Analysis), William Kennedy vid

package main

type user struct {
	name  string
	email string
}

func main() {
	u1 := CreateUserV1()
	u2 := CreateUserV2()
	u3 := CreateUserV3()
	u4 := CreateUserV4()

	println("u1", &u1)
	println("u2", &u2)
	println("u3", &u3)
	println("u4", &u4)
}

// value symantics because the caller gets its own copy after the function returns
func CreateUserV1() user {
	// literal construction
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// print fuction is in it's own memory frame. It can not access u directly from the V1 frame.
	println("V1", &u)

	// return a copy of u to the main memory frame and them m becomes the AF again
	return u
}

// pointer symantics because the caller gets the memory address of u
func CreateUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	// return the memory address of u, up the stack to main
	// the compiler will consturct u on the heap because when control is returned to main, the V2 memory frame will be destroyed
	return &u
}

// don't do this. Poor readability
func CreateUserV3() *user {
	// u is a pointer here. Don't use pointer symantics during construction
	u := &user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V3", &u)

	// pointer symantics
	return u
}

// don't do this. Poor readability
// pointer symantics during construction, then value symantics for the return.
func CreateUserV4() user {
	u := &user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V4", &u)

	// value symantics
	return *u
}
