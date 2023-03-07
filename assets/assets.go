// Package assets contains font, image & sound resources needed by the game
package assets

import (
	"bytes"
	"fish/assets/fonts"
	"fish/assets/images"
	"fish/assets/sounds"
	"fish/tool"
	"image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

type Characters struct {
	Pacman *ebiten.Image
	Ghost1 *ebiten.Image
	Ghost2 *ebiten.Image
	Ghost3 *ebiten.Image
	Ghost4 *ebiten.Image
}

type Powers struct {
	Life          *ebiten.Image
	Invincibility *ebiten.Image
}

type Walls struct {
	ActiveCorner   *ebiten.Image
	ActiveSide     *ebiten.Image
	InActiveCorner *ebiten.Image
	InActiveSide   *ebiten.Image
}

type Assets struct {
	ArcadeFont *truetype.Font
	Skin       *ebiten.Image
	Characters *Characters
	Powers     *Powers
	Walls      *Walls
}

// LoadAssets converts the character images(png, jpg, ...) to
// ebiten image format and loads fonts.
func LoadAssets() (*Assets, error) {
	skin, skinErr := loadSkin()
	if skinErr != nil {
		return nil, skinErr
	}

	font, fontErr := loadArcadeFont()
	if fontErr != nil {
		return nil, fontErr
	}

	characters, charactersErr := loadCharacters()
	if charactersErr != nil {
		return nil, charactersErr
	}

	powers, powersErr := loadPowers()
	if powersErr != nil {
		return nil, powersErr
	}

	walls, wallsErr := loadWalls()
	if wallsErr != nil {
		return nil, wallsErr
	}

	return &Assets{
		ArcadeFont: font,
		Skin:       skin,
		Characters: characters,
		Powers:     powers,
		Walls:      walls,
	}, nil
}

func loadSkin() (*ebiten.Image, error) {
	sImage, sImageErr := png.Decode(bytes.NewReader(images.SkinPng))
	if sImageErr != nil {
		return nil, sImageErr
	}

	skin, skinErr := ebiten.NewImageFromImage(sImage, ebiten.FilterDefault)
	if skinErr != nil {
		return nil, skinErr
	}

	return skin, nil
}

func loadArcadeFont() (*truetype.Font, error) {
	return truetype.Parse(fonts.ArcadeTTF)
}

func loadCharacters() (*Characters, error) {
	cImage, cImageErr := png.Decode(bytes.NewReader(images.CharactersPng))
	if cImageErr != nil {
		return nil, cImageErr
	}

	characters, charactersErr := ebiten.NewImageFromImage(cImage, ebiten.FilterDefault)
	if charactersErr != nil {
		return nil, charactersErr
	}

	pacman, pacmanErr := tool.GetSprite(61, 64, 0, 0, characters)
	if pacmanErr != nil {
		return nil, pacmanErr
	}

	ghost1, ghost1Err := tool.GetSprite(56, 64, 66, 0, characters)
	if ghost1Err != nil {
		return nil, ghost1Err
	}

	ghost2, ghost2Err := tool.GetSprite(56, 64, 125, 0, characters)
	if ghost2Err != nil {
		return nil, ghost2Err
	}

	ghost3, ghost3Err := tool.GetSprite(56, 64, 185, 0, characters)
	if ghost3Err != nil {
		return nil, ghost3Err
	}

	ghost4, ghost4Err := tool.GetSprite(56, 64, 244, 0, characters)
	if ghost4Err != nil {
		return nil, ghost4Err
	}

	return &Characters{
		Pacman: pacman,
		Ghost1: ghost1,
		Ghost2: ghost2,
		Ghost3: ghost3,
		Ghost4: ghost4,
	}, nil
}

func loadPowers() (*Powers, error) {
	pImage, pImageErr := png.Decode(bytes.NewReader(images.PowersPng))
	if pImageErr != nil {
		return nil, pImageErr
	}

	powers, powersErr := ebiten.NewImageFromImage(pImage, ebiten.FilterDefault)
	if powersErr != nil {
		return nil, powersErr
	}

	life, lifeErr := tool.GetSprite(64, 64, 0, 0, powers)
	if lifeErr != nil {
		return nil, lifeErr
	}

	invinc, invincErr := tool.GetSprite(64, 64, 67, 0, powers)
	if invincErr != nil {
		return nil, invincErr
	}

	return &Powers{
		Life:          life,
		Invincibility: invinc,
	}, nil
}

func loadWalls() (*Walls, error) {
	wImage, wImageErr := png.Decode(bytes.NewReader(images.WallsPng))
	if wImageErr != nil {
		return nil, wImageErr
	}

	walls, wallsErr := ebiten.NewImageFromImage(wImage, ebiten.FilterDefault)
	if wallsErr != nil {
		return nil, wallsErr
	}

	inactiveCorner, inactiveCornerErr := tool.GetSprite(12, 12, 0, 0, walls)
	if inactiveCornerErr != nil {
		return nil, inactiveCornerErr
	}

	inactiveSide, inactiveSideErr := tool.GetSprite(40, 12, 12, 0, walls)
	if inactiveSideErr != nil {
		return nil, inactiveSideErr
	}

	activeCorner, activeCornerErr := tool.GetSprite(12, 12, 52, 0, walls)
	if activeCornerErr != nil {
		return nil, activeCornerErr
	}

	activeSide, activeSideErr := tool.GetSprite(40, 12, 64, 0, walls)
	if activeSideErr != nil {
		return nil, activeSideErr
	}

	return &Walls{
		ActiveCorner:   activeCorner,
		ActiveSide:     activeSide,
		InActiveCorner: inactiveCorner,
		InActiveSide:   inactiveSide,
	}, nil
}

type Sounds struct {
	Beginning *mp3.Stream
	Chomp     *mp3.Stream
	Death     *mp3.Stream
	EatFlask  *mp3.Stream
	EatGhost  *mp3.Stream
	ExtraPac  *mp3.Stream
}

// LoadSounds returns a struct with wav files decoded
// for the provided audio context.
func LoadSounds(ctx *audio.Context) (*Sounds, error) {
	beginning, beginningErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.BeginningMp3))
	if beginningErr != nil {
		return nil, beginningErr
	}

	chomp, chompErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.ChompMp3))
	if chompErr != nil {
		return nil, chompErr
	}

	death, deathErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.DeathMp3))
	if deathErr != nil {
		return nil, deathErr
	}

	eatFlask, eatFlaskErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.EatFlaskMp3))
	if eatFlaskErr != nil {
		return nil, eatFlaskErr
	}

	eatGhost, eatGhostErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.EatGhostMp3))
	if eatGhostErr != nil {
		return nil, eatGhostErr
	}

	extraPac, extraPacErr := mp3.Decode(ctx, audio.BytesReadSeekCloser(sounds.ExtraPacMp3))
	if extraPacErr != nil {
		return nil, extraPacErr
	}

	return &Sounds{
		Beginning: beginning,
		Chomp:     chomp,
		Death:     death,
		EatFlask:  eatFlask,
		EatGhost:  eatGhost,
		ExtraPac:  extraPac,
	}, nil
}
