package str

import "testing"

func TestIsPhone(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		// 中国号码测试
		{"CN-Valid-1", "+86 13812345678", true},
		{"CN-Valid-2", "13812345678", true},
		{"CN-Valid-3", "0086-139-1234-5678", true},
		{"CN-Valid-4", "+86 15912345678", true},
		{"CN-Invalid-1", "+86 12345678", false},
		{"CN-Invalid-2", "+86 123456789012", false},
		{"CN-Invalid-3", "+86 2234567890", false},

		// 香港号码测试
		{"HK-Valid-1", "+852 1234 5678", true},
		{"HK-Valid-2", "00852-12345678", true},
		{"HK-Invalid", "+852 123456789", false},

		// 北美号码测试（美国/加拿大）
		{"US-Valid-1", "+1 234 567 8900", true},
		{"US-Valid-2", "+1 2345678900", true},
		{"US-Valid-3", "+1-321-555-8888", true},
		{"US-Valid-4", "001 345 666 7777", true},
		{"US-Invalid-1", "+1 123 456 789", false},   // 太短
		{"US-Invalid-2", "+1 123 456 78901", false}, // 太长

		// 日本号码测试
		{"JP-Valid-1", "+81 90 1234 5678", true},
		{"JP-Valid-2", "+81-80-1234-5678", true},
		{"JP-Invalid", "+81 123 456", false},

		// 韩国号码测试
		{"KR-Valid-1", "+82 10 1234 5678", true},
		{"KR-Valid-2", "+82-10-8888-9999", true},
		{"KR-Invalid", "+82 1234", false},

		// 英国号码测试
		{"GB-Valid-1", "+44 7911 123456", true},
		{"GB-Valid-2", "+44-7700-900000", true},
		{"GB-Invalid", "+44 123", false},

		// 德国号码测试
		{"DE-Valid-1", "+49 151 12345678", true},
		{"DE-Valid-2", "+49-170-1234567", true},
		{"DE-Invalid", "+49 12", false},

		// 法国号码测试
		{"FR-Valid-1", "+33 6 12 34 56 78", true},
		{"FR-Valid-2", "+33-7-12345678", true},
		{"FR-Invalid", "+33 123", false},

		// 新加坡号码测试
		{"SG-Valid-1", "+65 8123 4567", true},
		{"SG-Valid-2", "+65-9123-4567", true},
		{"SG-Invalid", "+65 1234", false},

		// 澳大利亚号码测试
		{"AU-Valid-1", "+61 4 1234 5678", true},
		{"AU-Valid-2", "+61-4-1234-5678", true},
		{"AU-Invalid", "+61 123", false},

		// 特殊情况测试
		{"Empty", "", false},
		{"OnlyPlus", "+", false},
		{"OnlyCountryCode", "+86", false},
		{"InvalidCountryCode", "+999 12345678", false},
		{"LettersInNumber", "+86 1234a5678", false},
		{"SpecialChars", "+86 123#4567", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPhone(tt.phone)
			if result != tt.expected {
				t.Errorf("IsPhone(%s) = %v; want %v", tt.phone, result, tt.expected)
			}
		})
	}
}
