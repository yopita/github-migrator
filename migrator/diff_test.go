package migrator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateDiff(t *testing.T) {
	testCases := []struct {
		src, expected string
	}{
		{
			src: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
-deleted
+added
diff --git a/CHANGELOG.md b/CHANGELOG.md
index 1234567..89abcde 100644
--- a/CHANGELOG.md
+++ b/CHANGELOG.md
@@ -1,6 +1,16 @@
# CHANGELOG
-deleted
+added
`,
			expected: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
-deleted
+added
diff --git a/CHANGELOG.md b/CHANGELOG.md
index 1234567..89abcde 100644
--- a/CHANGELOG.md
+++ b/CHANGELOG.md
@@ -1,6 +1,16 @@
# CHANGELOG
-deleted
+added
`,
		},
		{
			src: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
` + strings.Repeat("\n", 20000),
			expected: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
Too large diff
`,
		},
		{
			src: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
` + strings.Repeat("\n", 20000) + `
+added
diff --git a/CHANGELOG.md b/CHANGELOG.md
index 1234567..89abcde 100644
--- a/CHANGELOG.md
+++ b/CHANGELOG.md
@@ -1,6 +1,16 @@
# CHANGELOG
-deleted
+added
`,
			expected: `diff --git a/README.md b/README.md
index 1234567..89abcde 100644
Too large diff
diff --git a/CHANGELOG.md b/CHANGELOG.md
index 1234567..89abcde 100644
--- a/CHANGELOG.md
+++ b/CHANGELOG.md
@@ -1,6 +1,16 @@
# CHANGELOG
-deleted
+added
`,
		},
		{
			src: strings.Repeat(`diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
`+strings.Repeat("\n", 5000)+`
+added
`, 20),
			expected: strings.Repeat(`diff --git a/README.md b/README.md
index 1234567..89abcde 100644
--- a/README.md
+++ b/README.md
@@ -1,6 +1,16 @@
# README
`+strings.Repeat("\n", 5000)+`
+added
`, 20)[:60000] + "\n\nToo large diff\n",
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, truncateDiff(tc.src))
	}
}
