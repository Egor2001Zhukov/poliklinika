module poliklinika_gateway

go 1.22rc2

replace (
	github.com/zhumorist/common_go latest => ../common_go latest
)

require github.com/zhumorist/common_go v0.0.0-20240405140207-6245885cc1bb // indirect
