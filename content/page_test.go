package content

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewPage(t *testing.T) {
	tempDir := t.TempDir()

	err := NewPage(tempDir, "blogs/test-blog.md")
	assert.NoError(t, err, "Should not return an error if input is valid")

	expectedFile := filepath.Join(tempDir, "blogs/test-blog.md")
	data, err := os.ReadFile(expectedFile)
	assert.NoError(t, err, "Should be able to read generated file")

	dataStr := string(data)
	assert.Contains(t, dataStr, `title = "test-blog"`, "Front matter should contain correct title")
	assert.Contains(t, dataStr, `url = "/blogs/test-blog"`, "Front matter should contain correct url")
	assert.Contains(t, dataStr, `date = "`, "Front matter should contain a date field")
	assert.Contains(t, dataStr, `layout = "single.html"`, "Front matter should contain correct layout")
}
