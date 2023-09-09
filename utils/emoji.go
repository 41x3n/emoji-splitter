package utils

func IsEmoji(c rune) bool {
	return (c >= 0x1F600 && c <= 0x1F64F) || // Emoticons
		(c >= 0x1F300 && c <= 0x1F5FF) || // Misc Symbols and Pictographs
		(c >= 0x1F680 && c <= 0x1F6FF) || // Transport and Map
		(c >= 0x1F1E6 && c <= 0x1F1FF) || // Regional country flags
		(c >= 0x2600 && c <= 0x26FF) || // Misc symbols
		(c >= 0x2700 && c <= 0x27BF) || // Dingbats
		(c >= 0xFE00 && c <= 0xFE0F) || // Variation Selectors
		(c >= 0x1F900 && c <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(c >= 0x1F018 && c <= 0x1F270) || // Various asian characters
		(c >= 0x238C && c <= 0x2454) || // Misc items
		(c >= 0x20D0 && c <= 0x20FF) // Combining Diacritical Marks for Symbols
}
