package iteration

func main() {

}

const count = 5

func Repeat(arg string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += arg
	}
	return repeated
}
