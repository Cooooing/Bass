package model

import "errors"

var (
	ErrCharacterNotFound        = errors.New("game idle character not found")
	ErrCharacterInvalid         = errors.New("game idle character invalid")
	ErrCharacterLimitExceeded   = errors.New("game idle character limit exceeded")
	ErrCharacterNameDuplicate   = errors.New("game idle character name duplicate")
	ErrBackpackChangeInvalid    = errors.New("game idle backpack change invalid")
	ErrBackpackInsufficient     = errors.New("game idle backpack insufficient")
	ErrItemInvalid              = errors.New("game idle item invalid")
	ErrRecipeInvalid            = errors.New("game idle recipe invalid")
	ErrRecipeOutputEmpty        = errors.New("game idle recipe output empty")
	ErrActionInvalid            = errors.New("game idle action invalid")
	ErrActionQueueFull          = errors.New("game idle action queue full")
	ErrActionQueueStateConflict = errors.New("game idle action queue state conflict")
	ErrChatMessageInvalid       = errors.New("game idle chat message invalid")
)
