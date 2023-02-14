package controller

type Test struct {
	id   int8
	name string
}

//func GetHappinessList(context *gin.Context) {
//	result := 123
//	fmt.Println(result)
//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": result})
//}

//func GetClubs(context *gin.Context) {
//	var request model.Persons
//	test := data.DB.Model(&model.Persons{}).First(&request)
//	if test.Error != nil {
//		panic(test.Error)
//	}
//	fmt.Println(test)
//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": test.Value})
//}
//func PostTest(context *gin.Context) {
//	var request model.Persons
//	test := data.DB.Model(&model.Persons{}).First(&request)
//	if test.Error != nil {
//		panic(test.Error)
//	}
//	fmt.Println(test)
//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": test.Value})
//}
