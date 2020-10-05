package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func clean(program []byte) []byte {
	re := regexp.MustCompile(`[^\+-\.,\[\]><]`)
	return re.ReplaceAll(program, []byte{})
}

// findMatchingOpen will find a matching open bracket given a program and a position of the closing bracket.
func findMatchingOpen(program []byte, pos int) int {
	var count int
	index := pos

	if program[pos] != ']' {
		panic("wrong call to findMatchingOpen, expecting close bracket")
	}

	for !(count == 0 && index != pos) {
		curr := program[index]

		switch curr {
		case ']':
			count++
		case '[':
			count--
		}

		index--
	}

	return index + 1
}

// findMatchingClose will find a matching closing bracket given a program and a position of the opening bracket.
func findMatchingClose(program []byte, pos int) int {
	var count int
	index := pos

	if program[pos] != '[' {
		panic("wrong call to findMatchingClose, expecting close bracket")
	}

	for !(count == 0 && index != pos) {
		curr := program[index]

		switch curr {
		case '[':
			count++
		case ']':
			count--
		}

		index++
	}

	return index - 1
}

// instruction represents one Brainf*** operation.
// As well as the character, it stores two additional things: the location and the position of the matching bracket, which is only set
// if this instruction is also a bracket ('[' or ']'). It's not very memory efficient.
type instruction struct {
	ch       byte
	pos      int
	matchPos int
}

func createInstructions(program []byte) []instruction {
	instructions := []instruction{}

	for index := 0; index < len(program); index++ {
		curr := program[index]

		switch curr {
		case '+', '-', '.', ',', '>', '<':
			instructions = append(instructions, instruction{curr, index, 0})
		case '[':
			instructions = append(
				instructions,
				instruction{curr, index, findMatchingClose(program, index)},
			)
		case ']':
			instructions = append(
				instructions,
				instruction{curr, index, findMatchingOpen(program, index)},
			)
		}
	}

	return instructions
}

func assemble(instructions []instruction) string {
	var out bytes.Buffer

	out.WriteString(`defineRegisters:
	mov r0,#0
	mov r1,#body` + "\n")
	out.WriteString("run:\n")

	var additions int
	var subtractions int

	var pointerAdditions int
	var pointerSubtractions int

	for _, instruction := range instructions {
		if additions != 0 && instruction.ch != '+' {
			out.WriteString("	add r0,r0,#")
			out.WriteString(fmt.Sprint(additions))
			out.WriteString("\n")
			out.WriteString("	str r0,[r1]\n")
			additions = 0
		}

		if subtractions != 0 && instruction.ch != '-' {
			out.WriteString("	sub r0,r0,#")
			out.WriteString(fmt.Sprint(subtractions))
			out.WriteString("\n")
			subtractions = 0
		}

		if pointerAdditions != 0 && instruction.ch != '>' {
			out.WriteString("	str r0,[r1]\n")
			out.WriteString("	add r1,r1,#")
			out.WriteString(fmt.Sprint(pointerAdditions))
			out.WriteString("\n")
			out.WriteString("	ldr r0,[r1]\n")
			pointerAdditions = 0
		}

		if pointerSubtractions != 0 && instruction.ch != '<' {
			out.WriteString("	str r0,[r1]\n")
			out.WriteString("	sub r1,r1,#")
			out.WriteString(fmt.Sprint(pointerSubtractions))
			out.WriteString("\n")
			out.WriteString("	ldr r0,[r1]\n")
			pointerSubtractions = 0
		}

		switch instruction.ch {
		case '+':
			additions++
		case '-':
			subtractions++
		case '>':
			pointerAdditions++
		case '<':
			pointerSubtractions++
		case '.':
			out.WriteString("	out r0,7\n")
		case ',':
			out.WriteString("	inp [r1],7\n")
		case '[':
			loopLabelStart := fmt.Sprintf("LS%d", instruction.pos)
			loopLabelEnd := fmt.Sprintf("LE%d", instruction.matchPos)

			// Writes "LSd:"
			out.WriteString(loopLabelStart)
			out.WriteString(":\n")
			out.WriteString("	cmp r0,#0\n")

			// Writes "beq LEd"
			out.WriteString("	beq ")
			out.WriteString(loopLabelEnd)
			out.WriteString("\n")
		case ']':
			loopLabelStart := fmt.Sprintf("LS%d", instruction.matchPos)
			loopLabelEnd := fmt.Sprintf("LE%d", instruction.pos)

			// Writes "b LSd"
			out.WriteString("	b ")
			out.WriteString(loopLabelStart)
			out.WriteString("\n")

			// Writes "LEd:"
			out.WriteString(loopLabelEnd)
			out.WriteString(":\n")
		}
	}

	out.WriteString("	HALT\n")
	out.WriteString("body: dat 0\n")

	return out.String()
}

func simplify(program string) string {
	var out bytes.Buffer
	lines := strings.Split(program, "\n")

	prevLine := ""

	for _, line := range lines {
		switch {
		case line == "	ldr r0,[r1]" && prevLine == "	str r0,[r1]":
			continue
		case line == "	str r0,[r1]" && prevLine == "	str r0,[r1]":
			continue
		}

		out.WriteString(line + "\n")
		prevLine = line
	}

	return out.String()
}
