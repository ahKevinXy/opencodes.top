package hash

import "testing"

const secret = "i1ydX9RtHyuJTrw7frcu"
const length = 128

func TestHashidsEncode(t *testing.T) {
	str, _ := New(secret, length).HashidsEncode([]int{99, 100})
	t.Log(str)

	//GyV5pJqXvwAR  xwnoJzUm7p67
	//z5mLA21qay43JXV87W1PAgjbkYqOZoy5lwBDr44RaGqbkOdJlZoyDVWrg3xwnoJzUm7p67A8YeX02B5mKNzLj1PQ9EJNa6m0QE9Rzdx2LKGe38XVWdGD0ENB6brxoQwP
}

func TestHashidsDecode(t *testing.T) {
	ids, _ := New(secret, length).HashidsDecode("z5mLA21qay43JXV87W1PAgjbkYqOZoy5lwBDr44RaGqbkOdJlZoyDVWrg3xwnoJzUm7p67A8YeX02B5mKNzLj1PQ9EJNa6m0QE9Rzdx2LKGe38XVWdGD0ENB6brxoQwP")
	t.Log(ids)
}
