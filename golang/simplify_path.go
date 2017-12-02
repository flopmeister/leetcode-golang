
// string tokenization + stack
func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stk := []string{}
	for _, s := range paths {
		if s == "." || s == "" {
			continue
		}
		// pop back stk encounter ".."
		if s == ".." && len(stk) > 0 {
			stk = stk[:len(stk)-1]
		} else if s != ".." {
			stk = append(stk, s)
		}
	}
	var res string
	for _, item := range stk {
		res += "/" + item
	}
	if res == "" {
		res = "/"
	}
	return res
}
