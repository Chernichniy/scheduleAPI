package resources

type user struct {
	id         int    `json: id`
	is_bot     bool   `json: is_bot`
	first_name string `json: first_name`
}

type chat struct {
	id        int    `json: id`
	chat_type string `json: chat_type`
}

type linkPreviewOption struct {
	is_disabled bool `json: is_disabled` //true
}

type messageWithoutPhoto struct {
	chat_id              int               `json: chat_id`
	text                 string            `json: text`
	disable_notification bool              `json: disable_notification` //true
	parse_mode           string            `json parse_mode`            //MarkdownV2
	link_preview_options linkPreviewOption `json: link_preview_options`
}

type messgaeWithPhoto struct {
	chat_id              int               `json: chat_id`
	photo                string            `json: photo`
	caption              string            `json: caption`
	disable_notification bool              `json: disable_notification` //true
	parse_mode           string            `json parse_mode`            //MarkdownV2
	link_preview_options linkPreviewOption `json: link_preview_options`
}

//should a struct of message for the telegram
