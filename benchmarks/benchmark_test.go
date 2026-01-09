package benchmarks

import (
	"testing"

	"github.com/yourusername/gosensitive"
	"github.com/yourusername/gosensitive/dict"
)

// BenchmarkDetector_Contains benchmarks the Contains method
func BenchmarkDetector_Contains(b *testing.B) {
	words := generateWords(1000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词和测试内容的文本，用于性能基准测试。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Contains(text)
	}
}

// BenchmarkDetector_Find benchmarks the Find method
func BenchmarkDetector_Find(b *testing.B) {
	words := generateWords(1000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词和测试内容的文本，用于性能基准测试。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkDetector_Replace benchmarks the Replace method
func BenchmarkDetector_Replace(b *testing.B) {
	words := generateWords(1000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词和测试内容的文本，用于性能基准测试。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Filter(text)
	}
}

// BenchmarkDetector_SmallDict benchmarks with small dictionary
func BenchmarkDetector_SmallDict(b *testing.B) {
	words := generateWords(100)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词100的测试文本。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkDetector_MediumDict benchmarks with medium dictionary
func BenchmarkDetector_MediumDict(b *testing.B) {
	words := generateWords(10000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词5000的测试文本。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkDetector_LargeDict benchmarks with large dictionary
func BenchmarkDetector_LargeDict(b *testing.B) {
	words := generateWords(100000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	text := "这是一段包含敏感词50000的测试文本。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkDetector_LongText benchmarks with long text
func BenchmarkDetector_LongText(b *testing.B) {
	words := generateWords(1000)
	detector, _ := gosensitive.New().LoadWords(words).Build()
	
	// Generate long text
	text := ""
	for i := 0; i < 100; i++ {
		text += "这是一段很长的文本，包含了很多内容。"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkBuild benchmarks the Build operation
func BenchmarkBuild_1K(b *testing.B) {
	words := generateWords(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosensitive.New().LoadWords(words).Build()
	}
}

func BenchmarkBuild_10K(b *testing.B) {
	words := generateWords(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosensitive.New().LoadWords(words).Build()
	}
}

// Helper function to generate test words
func generateWords(count int) []dict.Word {
	words := make([]dict.Word, count)
	for i := 0; i < count; i++ {
		words[i] = dict.Word{
			Text:     "测试词" + string(rune(i%10000)),
			Category: dict.CategoryOther,
			Level:    dict.LevelMedium,
		}
	}
	return words
}

