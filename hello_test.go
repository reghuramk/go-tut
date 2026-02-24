package main

import "testing"

func Test_Hello(t *testing.T) {
	// t.Run("say hello", func(t *testing.T) {
	// 	got := hello("Chris")
	// 	want := "Hello, Chris"

	// 	if got != want {
	// 		t.Errorf("we got %q, but have %q", got, want)
	// 	}
	// 	assertionhelperfunc(t, got, want)

	// })

	// t.Run("say hellp world if no name is provided", func(t *testing.T) {
	// 	got := hello("")
	// 	want := "Hello, world"
	// 	if got != want {
	// 		t.Errorf("got %q but we want %q", got, want)

	// 	}
	// 	assertionhelperfunc(t, got, want)
	// })

	t.Run("answering to a spanish request", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertionhelperfunc(t, got, want)
	})

	t.Run("should pass if french is passed", func(t *testing.T) {
		got := hello("Rachel", "French")
		want := "Bonjour, Rachel"

		assertionhelperfunc(t, got, want)
	})

}

func assertionhelperfunc(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but we want %q", got, want)
	}
}
