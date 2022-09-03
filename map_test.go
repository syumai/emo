package emo

import "testing"

func TestGetFromEmojiString(t *testing.T) {
	tests := map[string]struct {
		emoji string
		want  *Emoji
	}{
		"valid jp": {
			emoji: "🇯🇵",
			want:  Get("jp"),
		},
		"invalid non-emoji": {
			emoji: "X",
			want:  nil,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := GetFromEmojiString(tc.emoji)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}
