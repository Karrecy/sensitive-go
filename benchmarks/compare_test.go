package benchmarks

import (
	"testing"

	"github.com/yourusername/gosensitive"
)

// BenchmarkDFA_Match benchmarks DFA algorithm
func BenchmarkDFA_Match(b *testing.B) {
	words := generateWords(5000)
	detector, _ := gosensitive.New().
		UseAlgorithm(gosensitive.AlgorithmDFA).
		LoadWords(words).
		Build()
	
	text := "这是一段包含测试词100和其他内容的文本。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkAC_Match benchmarks AC algorithm
func BenchmarkAC_Match(b *testing.B) {
	words := generateWords(5000)
	detector, _ := gosensitive.New().
		UseAlgorithm(gosensitive.AlgorithmAC).
		LoadWords(words).
		Build()
	
	text := "这是一段包含测试词100和其他内容的文本。"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

// BenchmarkDFA_Build benchmarks DFA build time
func BenchmarkDFA_Build(b *testing.B) {
	words := generateWords(5000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosensitive.New().
			UseAlgorithm(gosensitive.AlgorithmDFA).
			LoadWords(words).
			Build()
	}
}

// BenchmarkAC_Build benchmarks AC build time
func BenchmarkAC_Build(b *testing.B) {
	words := generateWords(5000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosensitive.New().
			UseAlgorithm(gosensitive.AlgorithmAC).
			LoadWords(words).
			Build()
	}
}

// Comparison test with different text lengths
func BenchmarkCompare_ShortText_DFA(b *testing.B) {
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmDFA, "短文本敏感词")
}

func BenchmarkCompare_ShortText_AC(b *testing.B) {
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmAC, "短文本敏感词")
}

func BenchmarkCompare_MediumText_DFA(b *testing.B) {
	text := "这是一段中等长度的文本，包含了一些敏感词和测试内容。" +
		"我们需要检测这段文本中的所有敏感词汇。"
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmDFA, text)
}

func BenchmarkCompare_MediumText_AC(b *testing.B) {
	text := "这是一段中等长度的文本，包含了一些敏感词和测试内容。" +
		"我们需要检测这段文本中的所有敏感词汇。"
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmAC, text)
}

func BenchmarkCompare_LongText_DFA(b *testing.B) {
	text := ""
	for i := 0; i < 50; i++ {
		text += "这是一段很长的文本，包含了大量的内容和信息。"
	}
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmDFA, text)
}

func BenchmarkCompare_LongText_AC(b *testing.B) {
	text := ""
	for i := 0; i < 50; i++ {
		text += "这是一段很长的文本，包含了大量的内容和信息。"
	}
	benchmarkWithAlgorithm(b, gosensitive.AlgorithmAC, text)
}

// Helper function for comparison benchmarks
func benchmarkWithAlgorithm(b *testing.B, algo gosensitive.AlgorithmType, text string) {
	words := generateWords(5000)
	detector, _ := gosensitive.New().
		UseAlgorithm(algo).
		LoadWords(words).
		Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detector.Find(text)
	}
}

