package telegram

type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type Chat struct {
	ID                                 int64
	ChatType                           string
	Title                              string
	Username                           string
	FirstName                          string
	LastName                           string
	IsForum                            bool
	Photo                              ChatPhoto
	ActiveUsernames                    []string
	EmojiStatusCustomEmojiID           string
	EmojiStatusExpirationDate          int64
	Bio                                string
	HasPrivateForwards                 bool
	HasRestrictedVoiceAndVideoMessages bool
	JoinToSendMessages                 bool
	JoinByRequest                      bool
	Description                        string
	InviteLink                         string
	PinnedMessage                      chatMessage
	Permissions                        ChatPermissions
	SlowModeDelay                      int64
	MessageAutoDeleteTime              int64
	HasAggressiveAntiSpamEnabled       bool
	HasHiddenMembers                   bool
	HasProtectedContent                bool
	StickerSetName                     string
	CanSetStickerSet                   bool
	LinkedChatID                       int64
	Location                           ChatLocation
}

type chatMessage struct {
	MessageID                     int64
	MessageThreadID               int64
	From                          User
	Date                          int64
	ForwardFrom                   User
	ForwardFromMessageID          int64
	ForwardSignature              string
	ForwardSenderName             string
	ForwardDate                   int64
	IsTopicMessage                bool
	IsAutomaticForward            bool
	ViaBot                        User
	EditDate                      int64
	HasProtectedContent           bool
	MediaGroupID                  string
	AuthorSignature               string
	Text                          string
	Entities                      []MessageEntity
	Animation                     Animation
	Audio                         Audio
	Document                      Document
	Photo                         []PhotoSize
	Sticker                       Sticker
	Story                         Story
	Video                         Video
	VideoNote                     VideoNote
	Voice                         Voice
	Caption                       string
	CaptionEntities               []MessageEntity
	HasMediaSpoiler               bool
	Contact                       Contact
	Dice                          Dice
	Game                          Game
	Poll                          Poll
	Venue                         Venue
	Location                      Location
	NewChatMembers                []User
	LeftChatMember                User
	NewChatPhoto                  []PhotoSize
	DeleteChatPhoto               bool
	GroupChatCreated              bool
	SupergroupChatCreated         bool
	ChannelChatCreated            bool
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged
	MigrateToChatID               int64
	MigrateFromChatID             int64
	Invoice                       Invoice
	SuccessfulPayment             SuccessfulPayment
	UserShared                    UserShared
	ChatShared                    ChatShared
	ConnectedWebsite              string
	WriteAccessAllowed            WriteAccessAllowed
	PassportData                  PassportData
	ProximityAlertTriggered       ProximityAlertTriggered
	ForumTopicCreated             ForumTopicCreated
	ForumTopicEdited              ForumTopicEdited
	ForumTopicClosed              ForumTopicClosed
	ForumTopicReopened            ForumTopicReopened
	GeneralForumTopicHidden       GeneralForumTopicHidden
	GeneralForumTopicUnhidden     GeneralForumTopicUnhidden
	VideoChatScheduled            VideoChatScheduled
	VideoChatStarted              VideoChatStarted
	VideoChatEnded                VideoChatEnded
	VideoChatParticipantsInvited  VideoChatParticipantsInvited
	WebAppData                    WebAppData
	InlineKeyboardMarkup          InlineKeyboardMarkup
}

type Message struct {
	MessageID                     int64                         `json:"message_id"`
	MessageThreadID               int64                         `json:"message_thread_id"`
	From                          User                          `json:"from"`
	SenderChat                    messageChat                   `json:"sender_chat"`
	Date                          int64                         `json:"date"`
	Chat                          messageChat                   `json:"chat"`
	ForwardFrom                   User                          `json:"forward_from"`
	ForwardFromChat               messageChat                   `json:"forward_from_chat"`
	ForwardFromMessageID          int64                         `json:"forward_from_message_id"`
	ForwardSignature              string                        `json:"forward_signature"`
	ForwardSenderName             string                        `json:"forward_sender_name"`
	ForwardDate                   int64                         `json:"forward_date"`
	IsTopicMessage                bool                          `json:"is_topic_message"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`
	ReplyToMessage                chatMessage                   `json:"reply_to_message"`
	ViaBot                        User                          `json:"via_bot"`
	EditDate                      int64                         `json:"edit_date"`
	HasProtectedContent           bool                          `json:"has_protected_content"`
	MediaGroupID                  string                        `json:"media_group_id"`
	AuthorSignature               string                        `json:"author_signature"`
	Text                          string                        `json:"text"`
	Entities                      []MessageEntity               `json:"entities"`
	Animation                     Animation                     `json:"animation"`
	Audio                         Audio                         `json:"audio"`
	Document                      Document                      `json:"document"`
	Photo                         []PhotoSize                   `json:"photo"`
	Sticker                       Sticker                       `json:"sticker"`
	Story                         Story                         `json:"story"`
	Video                         Video                         `json:"video"`
	VideoNote                     VideoNote                     `json:"video_note"`
	Voice                         Voice                         `json:"voice"`
	Caption                       string                        `json:"caption"`
	CaptionEntities               []MessageEntity               `json:"caption_entities"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler"`
	Contact                       Contact                       `json:"contact"`
	Dice                          Dice                          `json:"dice"`
	Game                          Game                          `json:"game"`
	Poll                          Poll                          `json:"poll"`
	Venue                         Venue                         `json:"venue"`
	Location                      Location                      `json:"location"`
	NewChatMembers                []User                        `json:"new_chat_members"`
	LeftChatMember                User                          `json:"left_chat_member"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo"`
	GroupChatCreated              bool                          `json:"group_chat_created"`
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int64                         `json:"migrate_to_chat_id"`
	MigrateFromChatID             int64                         `json:"migrate_from_chat_id"`
	PinnedMessage                 chatMessage                   `json:"pinned_message"`
	Invoice                       Invoice                       `json:"invoice"`
	SuccessfulPayment             SuccessfulPayment             `json:"successful_payment"`
	UserShared                    UserShared                    `json:"user_shared"`
	ChatShared                    ChatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                        `json:"connected_website"`
	WriteAccessAllowed            WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	ForumTopicCreated             ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    WebAppData                    `json:"web_app_data"`
	InlineKeyboardMarkup          InlineKeyboardMarkup          `json:"inline_keyboard_markup"`
}

type messageChat struct {
	ID                                 int64        `json:"id"`
	ChatType                           string       `json:"type"`
	Title                              string       `json:"title"`
	Username                           string       `json:"username"`
	FirstName                          string       `json:"first_name"`
	LastName                           string       `json:"last_name"`
	IsForum                            bool         `json:"is_forum"`
	Photo                              ChatPhoto    `json:"photo"`
	ActiveUsernames                    []string     `json:"active_usernames"`
	EmojiStatusCustomEmojiID           string       `json:"emoji_status_custom_emoji_id"`
	EmojiStatusExpirationDate          int64        `json:"emoji_status_expiration_date"`
	Bio                                string       `json:"bio"`
	HasPrivateForwards                 bool         `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool         `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool         `json:"join_to_send_messages"`
	JoinByRequest                      bool         `json:"join_by_request"`
	Description                        string       `json:"description"`
	InviteLink                         string       `json:"invite_link"`
	SlowModeDelay                      int64        `json:"slow_mode_delay"`
	MessageAutoDeleteTime              int64        `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       bool         `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   bool         `json:"has_hidden_members"`
	HasProtectedContent                bool         `json:"has_protected_content"`
	StickerSetName                     string       `json:"sticker_set_name"`
	CanSetStickerSet                   bool         `json:"can_set_sticker_set"`
	LinkedChatID                       int64        `json:"linked_chat_id"`
	Location                           ChatLocation `json:"location"`
}

type ChatPhoto struct {
	SmallFileID       string
	SmallFileUniqueID string
	BigFileID         string
	BigFileUniqueID   string
}

type ChatInviteLink struct {
	InviteLink              string
	Creator                 User
	CreatesJoinRequest      bool
	IsPrimary               bool
	IsRevoked               bool
	Name                    string
	ExpireDate              int64
	MemberLimit             int64
	PendingJoinRequestCount int64
}

type ChatPermissions struct {
	CanSendMessages       bool
	CanSendAudios         bool
	CanSendDocuments      bool
	CanSendPhotos         bool
	CanSendVideos         bool
	CanSendVideoNotes     bool
	CanSendVoiceNotes     bool
	CanSendPolls          bool
	CanSendOtherMessages  bool
	CanAddWebPagePreviews bool
	CanChangeInfo         bool
	CanInviteUsers        bool
	CanPinMessages        bool
	CanManageTopics       bool
}

type ChatLocation struct {
	Location Location
	Address  string
}

type MessageEntity struct {
	MessageEntityType string
	Offset            int64
	Length            int64
	URL               string
	User              User
	Language          string
	CustomEmojiID     string
}

type Animation struct {
	FileID       string
	FileUniqueID string
	Width        int64
	Height       int64
	Duration     int64
	Thumbnail    PhotoSize
	FileName     string
	MimeType     string
	FileSize     int64
}

type Audio struct {
	FileID       string
	FileUniqueID string
	Duration     int64
	Performer    string
	Title        string
	FileName     string
	MimeType     string
	FileSize     int64
	Thumbnail    PhotoSize
}

type Document struct {
	FileID       string
	FileUniqueID string
	Thumbnail    PhotoSize
	FileName     string
	MimeType     string
	FileSize     int64
}

type PhotoSize struct {
	FileID       string
	FileUniqueID string
	Width        int64
	Height       int64
	FileSize     int64
}

type Sticker struct {
}

type Story struct {
}

type Video struct {
	FileID       string
	FileUniqueID string
	Width        int64
	Height       int64
	Duration     int64
	Thumbnail    PhotoSize
	FileName     string
	MimeType     string
	FileSize     int64
}

type VideoNote struct {
	FileID       string
	FileUniqueID string
	Length       int64
	Duration     int64
	Thumbnail    PhotoSize
	FileSize     int64
}

type Voice struct {
	FileID       string
	FileUniqueID string
	Duration     int64
	MimeType     string
	FileSize     int64
}

type Contact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserIDD     int64
	Vcard       string
}

type Dice struct {
	Emoji string
	Value int64
}

type Game struct {
}

type Poll struct {
	ID                    string
	Question              string
	Options               []PollOption
	TotalVoterCount       int64
	IsClosed              bool
	IsAnonymous           bool
	PollType              PollType
	AllowsMultipleAnswers bool
	CorrectOptionID       int64
	Explanation           string
	ExplanationEntities   []MessageEntity
	OpenPeriod            int64
	CloseDate             int64
}

type PollOption struct {
	Text       string
	VoterCount int64
}

type PollAnswer struct {
	PollID    string
	VoterChat Chat
	User      User
	OptionIDs []int64
}

type Venue struct {
	Location        Location
	Title           string
	Address         string
	FoursquareID    string
	FoursquareType  string
	GooglePlaceID   string
	GooglePlaceType string
}

type Location struct {
	Longitude            float64
	Latitude             float64
	HorizontalAccuracy   float64
	LivePeriod           int64
	Heading              int64
	ProximityAlertRadius int64
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64
}

type Invoice struct {
}

type SuccessfulPayment struct {
}

type UserShared struct {
	RequestID int64
	UserID    int64
}

type ChatShared struct {
	RequestID int64
	ChatID    int64
}

type WriteAccessAllowed struct {
	FromRequest        bool
	WebAppName         string
	FromAttachmentMenu bool
}

type PassportData struct {
}

type BotCommand struct {
	Command     string
	Description string
}

type BotCommandScopeDefault struct {
	BotCommandScopeDefaultType string
}

type BotCommandScopeAllPrivateChats struct {
	BotCommandScopeAllPrivateChatsType string
}

type BotCommandScopeAllGroupChats struct {
	BotCommandScopeAllGroupChatsType string
}

type BotCommandScopeAllChatAdministrators struct {
	BotCommandScopeAllChatAdministratorsType string
}

type BotCommandScopeChat struct {
	BotCommandScopeChatType string
	ChatID                  string
}

type BotCommandScopeChatAdministrators struct {
	BotCommandScopeChatAdministratorsType string
	ChatID                                string
}

type BotCommandScopeChatMember struct {
	BotCommandScopeChatMemberType string
	ChatID                        string
	UserID                        int64
}

type BotName struct {
	Name string
}

type BotDescription struct {
	Description string
}

type BotShortDescription struct {
	ShortDescription string
}

type MenuButtonCommands struct {
	MenuButtonCommandsType string
}

type MenuButtonDefault struct {
	MenuButtonDefaultType string
}

type MenuButtonWebApp struct {
	MenuButtonWebAppType string
	Text                 string
	WebApp               WebAppInfo
}

type ResponseParameters struct {
	MigrateToChatID int64
	RetryAfter      int64
}

type ProximityAlertTriggered struct {
	Traveler User
	Watcher  User
	Distance int64
}

type ForumTopic struct {
	MessageThreadID   int64
	Name              string
	IconColor         int64
	IconCustomEmojiID string
}

type ForumTopicCreated struct {
	Name              string
	IconColor         int64
	IconCustomEmojiID string
}

type ForumTopicEdited struct {
	Name              string
	IconCustomEmojiID string
}

type ForumTopicClosed struct {
}

type ForumTopicReopened struct {
}

type GeneralForumTopicHidden struct {
}

type GeneralForumTopicUnhidden struct {
}

type VideoChatScheduled struct {
}

type VideoChatStarted struct {
}

type VideoChatEnded struct {
	Duration int64
}

type VideoChatParticipantsInvited struct {
	Users []User
}

type WebAppData struct {
	Data       string
	ButtonText string
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type LoginUrl struct {
	URL                string
	ForwardText        string
	BotUsername        string
	RequestWriteAccess bool
}

type SwitchInlineQueryChosenChat struct {
	Query             string
	AllowUserChats    bool
	AllowBotChats     bool
	AllowGroupChats   bool
	AllowChannelChats bool
}

type CallbackQuery struct {
	ID              string
	From            User
	Message         Message
	InlineMessageID string
	ChatInstance    string
	Data            string
	GameShortName   string
}

type ForceReply struct {
	ForceReply            bool
	InputFieldPlaceholder string
	Selective             bool
}

type CallbackGame struct {
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	URL                          string                       `json:"url"`
	CallbackData                 string                       `json:"callback_data"`
	WebApp                       *WebAppInfo                  `json:"web_app"`
	SwitchInlineQuery            string                       `json:"switch_inline_query"`
	LoginURL                     *LoginUrl                    `json:"login_url"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
	CallbackGame                 *CallbackGame                `json:"callback_game"`
	Pay                          bool                         `json:"pay"`
}

type UserProfilePhotos struct {
	TotalCount int64
	Photos     [][]PhotoSize
}

type File struct {
	FileID       string
	FileUniqueID string
	FileSize     int64
	FilePath     string
}

type WebAppInfo struct {
	URL string
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton
	IsPersistent          bool
	ResizeKeyboard        bool
	OneTimeKeyboard       bool
	InputFieldPlaceholder string
	Selective             bool
}

type KeyboardButton struct {
	Text            string
	RequestUser     KeyboardButtonRequestUser
	RequestChat     KeyboardButtonRequestChat
	RequestContact  bool
	RequestLocation bool
	RequestPoll     KeyboardButtonPollType
	WebApp          WebAppInfo
}

type KeyboardButtonRequestUser struct {
	RequestID     int64
	UserIsBot     bool
	UserIsPremium bool
}

type KeyboardButtonRequestChat struct {
	RequestID               int64
	ChatIsChannel           bool
	ChatIsForum             bool
	ChatHasUsername         bool
	ChatIsCreated           bool
	UserAdministratorRights ChatAdministratorRights
	BotAdministratorRights  ChatAdministratorRights
	BotIsMember             bool
}

type KeyboardButtonPollType struct {
	KeyboardButtonPollType string
}

type ChatAdministratorRights struct {
	IsAnonymous         bool
	CanManageChat       bool
	CanDeleteMessages   bool
	CanManageVideoChats bool
	CanRestrictMembers  bool
	CanPromoteMembers   bool
	CanChangeInfo       bool
	CanInviteUsers      bool
	CanPostMessages     bool
	CanEditMessages     bool
	CanPinMessages      bool
	CanPostStories      bool
	CanEditStories      bool
	CanDeleteStories    bool
	CanManageTopics     bool
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool
	Selective      bool
}

type ChatMemberOwner struct {
	Status      string
	User        User
	IsAnonymous bool
	CustomTitle string
}

type ChatMemberAdministrator struct {
	Status              string
	User                User
	CanBeEdited         bool
	IsAnonymous         bool
	CanManageChat       bool
	CanDeleteMessages   bool
	CanManageVideoChats bool
	CanRestrictMembers  bool
	CanPromoteMembers   bool
	CanChangeInfo       bool
	CanInviteUsers      bool
	CanPostMessages     bool
	CanEditMessages     bool
	CanPinMessages      bool
	CanPostStories      bool
	CanEditStories      bool
	CanDeleteStories    bool
	CanManageTopics     bool
	CustomTitle         string
}

type ChatMemberRestricted struct {
	Status                string
	User                  User
	IsMember              bool
	CanSendMessages       bool
	CanSendAudios         bool
	CanSendDocuments      bool
	CanSendPhotos         bool
	CanSendVideos         bool
	CanSendVideoNotes     bool
	CanSendVoiceNotes     bool
	CanSendPolls          bool
	CanSendOtherMessages  bool
	CanAddWebPagePreviews bool
	CanChangeInfo         bool
	CanInviteUsers        bool
	CanPinMessages        bool
	CanManageTopics       bool
	UntilDate             int64
}

type ChatMemberMember struct {
	Status string
	User   User
}

type ChatMemberLeft struct {
	Status string
	User   User
}

type ChatMemberBanned struct {
	Status    string
	User      User
	UntilDate int64
}

type ChatMemberUpdated struct {
	Chat                    Chat
	From                    User
	Date                    int64
	OldChatMember           ChatMember
	NewChatMember           ChatMember
	InviteLink              ChatInviteLink
	ViaChatFolderInviteLink bool
}

type ChatJoinRequest struct {
	Chat       Chat
	From       User
	UserChatID int64
	Date       int64
	Bio        string
	InviteLink ChatInviteLink
}

type ChatMember struct {
}

type InputMediaPhoto struct {
	InputMediaPhotoType string
	Media               string
	Caption             string
	ParseMode           string
	CaptionEntities     []MessageEntity
	HasSpoiler          bool
}

type InputMediaVideo struct {
	InputMediaVideoType string
	Media               string
	Thumbnail           InputFile
	Caption             string
	ParseMode           string
	CaptionEntities     []MessageEntity
	Width               int64
	Height              int64
	Duration            int64
	SupportsStreaming   bool
	HasSpoiler          bool
}

type InputMediaAnimation struct {
	InputMediaVideoType string
	Media               string
	Thumbnail           InputFile
	Caption             string
	ParseMode           string
	CaptionEntities     []MessageEntity
	Width               int64
	Height              int64
	Duration            int64
	HasSpoiler          bool
}

type InputMediaAudio struct {
	InputMediaVideoType string
	Media               string
	Thumbnail           InputFile
	Caption             string
	ParseMode           string
	CaptionEntities     []MessageEntity
	Width               int64
	Height              int64
	Duration            int64
	Performer           string
	HasSpoiler          bool
}

type InputMediaDocument struct {
	InputMediaVideoType         string
	Media                       string
	Thumbnail                   InputFile
	Caption                     string
	ParseMode                   string
	CaptionEntities             []MessageEntity
	DisableContentTypeDetection bool
}

type InputFile struct {
}

type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType string   `json:"chat_type"`
	Location Location `json:"location"`
}

type ChosenInlineResult struct {
	ResultID        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageID string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}

type PreCheckoutQuery struct {
	ID               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int64     `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionID string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info"`
}

type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IpAddress                    string   `json:"ip_address"`
	LastErrorDate                int      `json:"last_error_date"`
	LastErrorMessage             string   `json:"last_error_message"`
	LastSynchronizationErrorDate int      `json:"last_synchronization_error_date"`
	MaxConnections               int      `json:"max_connections"`
	AllowedUpdates               []string `json:"allowed_updates"`
}
