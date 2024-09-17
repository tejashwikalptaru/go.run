package utils

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputState represents the current state of user inputs.
type InputState struct {
	Jump      bool
	MoveLeft  bool
	MoveRight bool
	Sprint    bool // Applicable for desktop only
}

// InputHandler manages input detection and state.
type InputHandler struct {
	isMobile     bool
	screenWidth  int
	screenHeight int

	// Movement touch tracking for mobile
	movementTouchID   ebiten.TouchID
	movementOriginX   int
	movementOriginY   int
	movementCurrentX  int
	movementCurrentY  int
	movementThreshold int // Minimum movement in pixels to trigger direction
}

// NewInputHandler initializes and returns a new InputHandler.
// screenWidth and screenHeight should match your game's resolution.
func NewInputHandler(screenWidth, screenHeight int) *InputHandler {
	return &InputHandler{
		isMobile:          IsMobile(),
		screenWidth:       screenWidth,
		screenHeight:      screenHeight,
		movementTouchID:   -1, // -1 indicates no active movement touch
		movementThreshold: 10, // Adjust as needed for sensitivity
	}
}

// Update processes the current input and returns the InputState.
// It should be called once per frame, typically within the Game's Update method.
func (ih *InputHandler) Update() InputState {
	var state InputState

	if !ih.isMobile {
		// Handle Desktop Inputs
		state.Jump = inpututil.IsKeyJustPressed(ebiten.KeySpace)
		state.MoveLeft = ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
		state.MoveRight = ebiten.IsKeyPressed(ebiten.KeyArrowRight)
		state.Sprint = ebiten.IsKeyPressed(ebiten.KeyShift)
	} else {
		// Handle Mobile Inputs

		// 1. Detect Touch Releases for Jumping
		releasedTouchIDs := inpututil.AppendJustReleasedTouchIDs(nil)
		for _, id := range releasedTouchIDs {
			x, y := ebiten.TouchPosition(id)
			if ih.isWithinLeftHalf(x, y) {
				state.Jump = true
			}
			// If the released touch was the movement touch, reset it
			if id == ih.movementTouchID {
				ih.movementTouchID = -1
			}
		}

		// 2. Detect Touch Presses for Movement
		pressedTouchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		for _, id := range pressedTouchIDs {
			x, y := ebiten.TouchPosition(id)
			if ih.isWithinRightHalf(x, y) && ih.movementTouchID == -1 {
				// Assign this touch as the movement touch
				ih.movementTouchID = id
				ih.movementOriginX = x
				ih.movementOriginY = y
				ih.movementCurrentX = x
				ih.movementCurrentY = y
			}
		}

		// 3. Track Movement Touches
		if ih.movementTouchID != -1 {
			// Check if the movement touch is still active
			isActive := false
			activeTouchIDs := ebiten.AppendTouchIDs(nil)
			for _, id := range activeTouchIDs {
				if id == ih.movementTouchID {
					isActive = true
					x, y := ebiten.TouchPosition(id)
					ih.movementCurrentX = x
					ih.movementCurrentY = y
					break
				}
			}

			if isActive {
				// Calculate horizontal movement delta
				deltaX := ih.movementCurrentX - ih.movementOriginX

				if deltaX < -ih.movementThreshold {
					state.MoveLeft = true
				} else if deltaX > ih.movementThreshold {
					state.MoveRight = true
				}
				// If movement is within threshold, do not set movement
			} else {
				// Movement touch was released without significant movement
				ih.movementTouchID = -1
			}
		}

		// 4. Detect Touch Presses for Jumping (Immediate Jump)
		// Optional: Implement immediate jump on touch press if desired
		// For now, jump is triggered on touch release
	}

	return state
}

// isWithinLeftHalf checks if the touch is within the left half of the screen.
func (ih *InputHandler) isWithinLeftHalf(x, y int) bool {
	return x < ih.screenWidth/2
}

// isWithinRightHalf checks if the touch is within the right half of the screen.
func (ih *InputHandler) isWithinRightHalf(x, y int) bool {
	return x >= ih.screenWidth/2
}

// LogTouches Optional: LogTouches is a debugging method to log current touch events.
// Uncomment and call this method within Update if you need to debug touch inputs.
func (ih *InputHandler) LogTouches() {
	if ih.isMobile {
		// Log active touches
		activeTouchIDs := ebiten.AppendTouchIDs(nil)
		for _, id := range activeTouchIDs {
			x, y := ebiten.TouchPosition(id)
			log.Printf("Active Touch - ID: %d, Position: (%d, %d)", id, x, y)
		}

		// Log just released touches
		releasedTouchIDs := inpututil.AppendJustReleasedTouchIDs(nil)
		for _, id := range releasedTouchIDs {
			x, y := ebiten.TouchPosition(id)
			log.Printf("Released Touch - ID: %d, Position: (%d, %d)", id, x, y)
		}

		// Log just pressed touches
		pressedTouchIDs := inpututil.AppendJustPressedTouchIDs(nil)
		for _, id := range pressedTouchIDs {
			x, y := ebiten.TouchPosition(id)
			log.Printf("Pressed Touch - ID: %d, Position: (%d, %d)", id, x, y)
		}
	}
}
