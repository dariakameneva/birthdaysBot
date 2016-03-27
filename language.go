package main

const (
	unknownCommand = "I don't know what you're saying"

	passwOkReply   = "You're goddamnt right!"
	passwFailReply = "I don't know... It just doesn't seem right. Why don't you give it another try? '/passw <answer>'"
	passwAddition  = " Reply with '/passw <answer>'"

	setBDOkFmtReply      = "Got it: %d of %s!"
	setBDFailReply       = "Ima stupid bot, I don't understand this. Send me your birthday like this: '/set Feb 29'"
	setBDOutOfRangeReply = "No can do. It's way out of range!"

	helpReply = `I am a bot that stores birthdays.
/start 	- Start a conversation
/passw - send me a password so I know I can trust you
/help - I guess you know what that is, you've used it already
/set Jan 31 - set your birthday as January 31st
`
)
