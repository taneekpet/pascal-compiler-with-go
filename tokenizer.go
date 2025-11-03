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
	result.ExpandedTo[2] = models.Token{
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
	return models.Token{}, nil
}

func parseBlock(slide *models.InputSlide) (
	result models.Token, err error,
) {
	return models.Token{}, nil
}
