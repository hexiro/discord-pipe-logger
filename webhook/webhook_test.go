package webhook

import (
	"testing"
)

const testURL = "https://discord.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"
const testCanaryURL = "https://canary.discord.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"
const testOldURL = "https://canary.discord.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"
const testPtbURL = "https://canary.discord.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"

const id = "671422873239289884"
const token = "G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"

var webhook = &Webhook{
	ID:    id,
	Token: token,
}

func TestFromURL(t *testing.T) {
	webhook1, err := FromURL(testURL)
	if err != nil {
		t.Error("Got unexpected error ", err)
	}
	webhook2, _ := FromURL(testCanaryURL)
	webhook3, _ := FromURL(testOldURL)
	webhook4, _ := FromURL(testPtbURL)
	if webhook1.URL() != testURL {
		t.Fail()
	}

	for index, hook := range []*Webhook{webhook1, webhook2, webhook3, webhook4} {
		if *webhook != *hook {
			t.Errorf("Expected %+v, got %+v (#%d)", webhook, webhook1, index+1)
		}
	}
}

func TestFromIDAndToken(t *testing.T) {
	webhook1, err := FromIDAndToken(id, token)
	if err != nil {
		t.Error("Got unexpected error ", err)
	}
	if *webhook != *webhook1 {
		t.Fail()
	}
}

func TestWebhookFromURLStringWithWrongURL(t *testing.T) {
	testData := map[string]string{
		"Another website":    "https://google.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
		"Missing channel id": "https://discord.com/api/webhooks/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
		"Invalid route":      "https://discord.com/api//671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
	}
	for test, link := range testData {
		t.Run(test, func(t *testing.T) {
			webhook, err := FromURL(link)
			if webhook != nil || err == nil {
				t.Error("Wanted nothing, got webhook")
			}
		})
	}
}