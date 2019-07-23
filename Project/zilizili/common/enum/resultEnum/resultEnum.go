package resultEnum


type Status int

// iota 初始化后会自动递增
const (
    Success Status = iota + 10000// value --> 0
    Failed              // value --> 1
    Error            // value --> 2
)
