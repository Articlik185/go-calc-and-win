package main

import (
    "math/rand"
    "os"
    "testing"
)

func withInput(t *testing.T, input string) {
    t.Helper()
    file, err := os.CreateTemp(t.TempDir(), "input")
    if err != nil { t.Fatal(err) }
    if _, err := file.WriteString(input); err != nil { t.Fatal(err) }
    if _, err := file.Seek(0, 0); err != nil { t.Fatal(err) }
    previous := os.Stdin
    os.Stdin = file
    t.Cleanup(func() { os.Stdin = previous; file.Close() })
}

func TestFiveMediumAttacksAccumulateDamage(t *testing.T) {
    withInput(t, "mid\nmid\nmid\nmid\nmid\n")
    previous := rnd
    rnd = rand.New(rand.NewSource(1))
    t.Cleanup(func() { rnd = previous })
    // Each medium attack deals 15–25 damage, so five deal 75–125, not 5.
    if damage := getUserAttack(); damage < 75 || damage > 125 {
        t.Fatalf("five medium attacks dealt %d; want 75–125", damage)
    }
}

func TestReplayAnswers(t *testing.T) {
    for _, tc := range []struct { answer string; want bool }{
        {"y", true}, {"Y", true}, {"n", false},
    } {
        t.Run(tc.answer, func(t *testing.T) {
            withInput(t, "mid\nmid\nmid\nmid\nmid\n"+tc.answer+"\n")
            previous := rnd
            rnd = rand.New(rand.NewSource(1))
            t.Cleanup(func() { rnd = previous })
            if got := runGame(); got != tc.want {
                t.Fatalf("replay %q = %v; want %v", tc.answer, got, tc.want)
            }
        })
    }
}
