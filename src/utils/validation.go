package utils

import "strings"

func ValidateThreadTitle(title string) (string, string) {
	if len(title) < 5 {
		return title, "スレッドのタイトルは5文字以上でなければなりません。"
	}
	if len(title) > 100 {
		return title, "スレッドのタイトルは100文字以下でなければなりません。"
	}
	if strings.TrimSpace(title) == "" {
		return title, "スレッドのタイトルが空白です。"
	}
	return title, ""
}

func ValidateMessageContent(content string) (string, string) {
	if len(content) < 1 {
		return content, "メッセージの内容は1文字以上でなければなりません。"
	}
	if len(content) > 1000 {
		return content, "メッセージの内容は1000文字以下でなければなりません。"
	}
	if strings.TrimSpace(content) == "" {
		return content, "メッセージの内容が空白です。"
	}
	return content, ""
}
