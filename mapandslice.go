package main
 
import (
    "fmt"
    "time"
)
 
type s struct {
    s []*s
}
 
type m struct {
    m map[int]*m
}
 
const (
    testLevel = 20
    testTimes = 6
)
 
func buildStruct() *s {
    root := &s{}
    cur := root
 
    for i := 0; i < testLevel; i++ {
        cur.s = make([]*s, 1)
 
        cur.s[0] = &s{}
 
        cur = cur.s[0]
    }
 
    return root
}
 
func traverseStruct(t *s) int {
    levels := 0
 
    for {
        if t.s == nil {
            return levels
        }
 
        t = t.s[0]
 
        levels++
    }
}
 
func buildMap() *m {
    root := &m{}
    cur := root
 
    for i := 0; i < testLevel; i++ {
        cur.m = make(map[int]*m, 1)
 
        cur.m[0] = &m{}
 
        cur = cur.m[0]
    }
 
    return root
}
 
func traverseMap(t *m) int {
    levels := 0
 
    for {
        current, found := t.m[0]
 
        if !found {
            return levels
        }
 
        t = current
        levels++
    }
}
 
func main() {
    fmt.Println("-----------Let's Start Building-------------")
 
    testStruct := buildStruct()
    testMap := buildMap()
 
    fmt.Println("Traversing", testStruct, testMap)
 
    for i := 0; i < testTimes; i++ {
        start := time.Now()
 
        sCount := traverseStruct(testStruct)
        sEnd := time.Since(start)
 
        start = time.Now()
        mCount := traverseMap(testMap)
        mEnd := time.Since(start)
 
        if sCount != mCount {
            panic("Count are different")
        }
 
        fmt.Printf("Traverse took: Slice %s vs Map %s\r\n", sEnd, mEnd)
    }
}