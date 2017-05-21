package models

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteInput(t *testing.T) {
	work := &Work{
		Path: "/tmp",
	}
	expected := "lorem\nipsum\n"

	err := work.WriteInput(expected)
	assert.NoError(t, err)

	var actual string
	actual, err = work.ReadInput()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestReadInput(t *testing.T) {
	work := &Work{
		Path: "/tmp",
	}

	expected := "1e5\n23e-1\n"
	filepath := fmt.Sprintf("%s/input.txt", work.Path)

	f, err := os.Create(filepath)
	assert.NoError(t, err)

	defer f.Close()
	defer os.Remove(filepath)

	w := bufio.NewWriter(f)
	_, err = w.WriteString(expected)
	assert.NoError(t, err)

	w.Flush()

	actual, err := work.ReadInput()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestReadOutput(t *testing.T) {
	work := &Work{
		Path: "/tmp",
	}

	expected := "1e5\n23e-1\n"
	filepath := fmt.Sprintf("%s/output.txt", work.Path)

	f, err := os.Create(filepath)
	assert.NoError(t, err)

	defer f.Close()
	defer os.Remove(filepath)

	w := bufio.NewWriter(f)
	_, err = w.WriteString(expected)
	assert.NoError(t, err)

	w.Flush()

	actual, err := work.ReadOutput()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
