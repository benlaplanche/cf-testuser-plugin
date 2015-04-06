package main

func searchIntSlice(slice []int, seek int) (result bool) {
	for _, v := range slice {
		if v == seek {
			return true
			break
		}
	}
	return false
}

func foundError(status []int) (response bool) {
	if searchIntSlice(status, 0) == true {
		return true
	}
	return false
}
