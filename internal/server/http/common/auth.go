package common

// const (
// 	RWD  = 0 // 1 = 1<<RWD read write delete permission
// 	READ = 1 // 2 = 1<<READ access permission
// 	USE  = 2 // 4 = 1<<USE use permission
// )

// /*
// 1 归属权
// 2 访问权
// 4 使用权
// 7 归属权-访问权-使用权
// 3 归属权-访问权
// 5 归属权-使用权
// 6 访问权-使用权
// */

// /*
// 目前可给用户授予的是：
// 访问权
// 使用权
// 访问权+使用权
// */

// const (
// 	RES_OWNER_AUTH   = (1 << RWD) | (1 << READ) | (1 << USE) //资源归属权
// 	ACCRESS_AUTH     = 1 << READ                             //auth=2      //访问权
// 	USE_AUTH         = 1 << USE                              //auth=4   //使用权
// 	ACCRESS_USE_AUTH = 1<<READ | 1<<USE                      //auth=6 //访问权+使用权
// )

// var AuthMap = map[uint16]string{
// 	ACCRESS_AUTH:     "访问权",
// 	USE_AUTH:         "使用权",
// 	ACCRESS_USE_AUTH: "访问权+使用权",
// }

// func CheckAuthValid(auth uint16) bool {
// 	if _, ok := AuthMap[auth]; ok {
// 		return true
// 	}
// 	return false
// }
