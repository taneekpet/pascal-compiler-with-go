package main

import "github.com/taneekpet/pascal-compiler-with-go/models"

func parseProgram(slide *models.InputSlide) (
	result models.Token, err error,
) {
	result = models.Token{
		Lexemes:     nil,
		IdReference: nil,
		Type:        models.Reserved,
		ExpandedTo:  make([]models.Token, 5),
	}

	slide.PeekWithoutSpace()
	expected := []string{"P", "R", "O", "G", "R", "A", "M"}
	for i := 0; i < len(expected); i++ {
		chr := slide.Next()
		if string(chr) != expected[i] {
			return result, models.CustomError{
				ErrorIndex: slide.GetIndex(),
				Code:       models.Unrecognized,
			}
		}
	}
	result.ExpandedTo[0] = models.Token{
		Lexemes:     []rune("PROGRAM"),
		IdReference: nil,
		Type:        models.Reserved,
		ExpandedTo:  nil,
	}

	slide.PeekWithoutSpace()
	result.ExpandedTo[1], err = parseVariable(slide)
	if err != nil {
		return
	}

	chr := slide.NextWithoutSpace()
	if string(chr) != ";" {
		return result, models.CustomError{
			ErrorIndex: slide.GetIndex(),
			Code:       models.Unrecognized,
		}
	}
	result.ExpandedTo[2] = models.Token{
		Lexemes:     []rune(";"),
		IdReference: nil,
		Type:        models.Reserved,
		ExpandedTo:  nil,
	}

	slide.PeekWithoutSpace()
	result.ExpandedTo[3], err = parseBlock(slide)
	if err != nil {
		return
	}

	chr = slide.NextWithoutSpace()
	if string(chr) != "." {
		return result, models.CustomError{
			ErrorIndex: slide.GetIndex(),
			Code:       models.Unrecognized,
		}
	}
	result.ExpandedTo[4] = models.Token{
		Lexemes:     []rune("."),
		IdReference: nil,
		Type:        models.Reserved,
		ExpandedTo:  nil,
	}

	return result, nil
}

func parseVariable(slide *models.InputSlide) (
	result models.Token, err error,
) {
	nextChar := slide.NextWithoutSpace()
	variableId := ""
	for ; nextChar != models.SPACE && string(nextChar) != ";" && !slide.IsEnd(); nextChar = slide.Next() {
		variableId += string(nextChar)
	}
	return models.Token{
		Lexemes:     []rune(variableId),
		IdReference: nil, // TODO
		Type:        models.Variable,
		ExpandedTo:  nil,
	}, nil
}

func parseBlock(slide *models.InputSlide) (
	result models.Token, err error,
) {
	result = models.Token{
		Lexemes:     []rune(""),
		IdReference: nil,
		Type:        models.Block,
		ExpandedTo:  make([]models.Token, 2),
	}

	slide.PeekWithoutSpace()
	result.ExpandedTo[0], err = parseDeclarations(slide)
	if err != nil {
		return
	}

	slide.PeekWithoutSpace()
	result.ExpandedTo[1], err = parseCompoundStatement(slide)
	if err != nil {
		return
	}
	return
}

func parseDeclarations(slide *models.InputSlide) (
	result models.Token, err error,
) {
	// empty
	result = models.Token{
		Lexemes:     []rune(""),
		IdReference: nil,
		Type:        models.Declaration,
		ExpandedTo:  make([]models.Token, 0),
	}

	slide.PeekWithoutSpace()
	expected := []string{"V", "A", "R"}
	for i := 0; i < len(expected); i++ {
		chr := slide.Next()
		if string(chr) != expected[i] {
			slide.SetBack(i + 1)
			return
		}
	}
	result.ExpandedTo = append(
		result.ExpandedTo,
		models.Token{
			Lexemes:     []rune("VAR"),
			IdReference: nil,
			Type:        models.Reserved,
			ExpandedTo:  nil,
		},
	)

	for {
		slide.PeekWithoutSpace()
		tmpVarDec, parsed, err := parseVariableDeclaration(slide)
		if err != nil {
			slide.SetBack(parsed)
			break
		}

		slide.PeekWithoutSpace()
		chr := slide.Next()
		if string(chr) != ";" {
			slide.SetBack(1)
			break
		}

		// parse success, append
		result.ExpandedTo = append(
			result.ExpandedTo,
			tmpVarDec,
			models.Token{
				Lexemes:     []rune(";"),
				IdReference: nil,
				Type:        models.Reserved,
				ExpandedTo:  nil,
			},
		)
	}
	return

}

func parseCompoundStatement(slide *models.InputSlide) (
	result models.Token, err error,
) {
	result = models.Token{
		Lexemes:     []rune(""),
		IdReference: nil,
		Type:        models.CompoundStatement,
		ExpandedTo:  make([]models.Token, 0),
	}
	return
}

func parseVariableDeclaration(slide *models.InputSlide) (
	result models.Token, parsed int, err error,
) {
	result = models.Token{
		Lexemes:     []rune(""),
		IdReference: nil,
		Type:        models.VariableDeclaration,
		ExpandedTo:  make([]models.Token, 0),
	}
	return
}
