package AMJ

func WhatToPrint(i int, word, filename, res string) (string, error) {
	for _, letter := range word {
		line, err := GetLine(1+(int(letter)-32)*9+i, filename)
		res += line
		if err != nil {
			return "", err
		}
	}
	return res, nil
}
