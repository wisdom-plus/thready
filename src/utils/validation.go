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

func ValidateUsername(username string) (string, string) {
	if len(username) < 3 {
		return username, "ユーザー名は3文字以上でなければなりません。"
	}
	if len(username) > 20 {
		return username, "ユーザー名は20文字以下でなければなりません。"
	}
	if strings.TrimSpace(username) == "" {
		return username, "ユーザー名が空白です。"
	}
	return username, ""
}

func ValidatePassword(password string) (string, string) {
	if len(password) < 6 {
		return password, "パスワードは6文字以上でなければなりません。"
	}
	if len(password) > 100 {
		return password, "パスワードは100文字以下でなければなりません。"
	}
	if strings.TrimSpace(password) == "" {
		return password, "パスワードが空白です。"
	}
	return password, ""
}
