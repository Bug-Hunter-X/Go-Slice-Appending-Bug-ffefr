func main() {

        x := make([]int, 0, 10) // Set initial capacity to 10

        for i := 0; i < 10; i++ {
                x = append(x, i)
        }
        fmt.Println(x)

}     