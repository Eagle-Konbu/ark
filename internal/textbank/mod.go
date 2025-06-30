package textbank

import _ "embed"

//go:embed description_template/description.md
var DescriptionTemplateMarkdown string

//go:embed description_template/description.txt
var DescriptionTemplateText string

//go:embed description_template/description.xml
var DescriptionTemplateXML string

//go:embed description_template/description.arklite.txt
var DescriptionTemplateArklite string

//go:embed description_template/arklite_compless_header.arklite.txt
var ArkliteComplessHeaderTemplate string

const (
	EmojiSuccess     = "✅"
	EmojiInterrupted = "🛑"
	EmojiAlmost      = "🔒"
	EmojiDone        = "🎉"
	EmojiArchive     = "📜"
	EmojiArk         = "🪨"
	EmojiBoard       = "🪧"
	EmojiStar        = "🌟"
	EmojiHourglass   = "⏳"
)
