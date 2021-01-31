package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type Game struct {
	ID uint `gorm:"primaryKey"`
	Publisher int // foreign key to Publisher
	Developer int // foreign key to Developer
	Title string
	Description string
	Price int
	ReleaseDate string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Game{})
	db.AutoMigrate(&Game{})

	p := &Game{}
	p.seed(db)
}

func (p *Game) seed(db *gorm.DB) {
	db.Create(&Game{
		Publisher:   3,
		Developer:   2,
		Title:       "Red Dead Redemption 2",
		Description: "Winner of over 175 Game of the Year Awards and recipient of over 250 perfect scores, RDR2 is the epic tale of outlaw Arthur Morgan and the infamous Van der Linde gang, on the run across America at the dawn of the modern age. Also includes access to the shared living world of Red Dead Online.",
		Price:       690000,
		ReleaseDate: "2019-12-06",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   4,
		Developer:   4,
		Title:       "Cyberpunk 2077",
		Description: "Cyberpunk 2077 is an open-world, action-adventure story set in Night City, a megalopolis obsessed with power, glamour and body modification. You play as V, a mercenary outlaw going after a one-of-a-kind implant that is the key to immortality.",
		Price:       699999,
		ReleaseDate: "2020-12-10",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   2,
		Developer:   4,
		Title:       "Watch Dogs 2",
		Description: "Welcome to San Francisco. Play as Marcus, a brilliant young hacker, and join the most notorious hacker group, DedSec. Your objective: execute the biggest hack of history.",
		Price:       619000,
		ReleaseDate: "2016-11-28",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   3,
		Developer:   4,
		Title:       "Overcooked 2",
		Description: "Overcooked returns with a brand-new helping of chaotic cooking action! Journey back to the Onion Kingdom and assemble your team of chefs in classic couch co-op or online play for up to four players. Hold onto your aprons… it’s time to save the world again!",
		Price:       199000,
		ReleaseDate: "2018-08-07",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   7,
		Developer:   6,
		Title:       "Mafia III: Definitive Edition",
		Description: "After Lincoln Clay's surrogate family, the black mob, is betrayed and killed by the Italian Mafia, Lincoln builds a new family and blazes a path of revenge through the Mafioso responsible.",
		Price:       345000,
		ReleaseDate: "2020-05-19",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   1,
		Developer:   9,
		Title:       "Terraria",
		Description: "Dig, fight, explore, build! Nothing is impossible in this action-packed adventure game. Four Pack also available!",
		Price:       89999,
		ReleaseDate: "2011-05-11",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   4,
		Developer:   8,
		Title:       "Kerbal Space Program",
		Description: "In Kerbal Space Program, take charge of the space program for the alien race known as the Kerbals. You have access to an array of parts to assemble fully-functional spacecraft that flies (or doesn’t) based on realistic aerodynamic and orbital physics.",
		Price:       269999,
		ReleaseDate: "2015-04-15",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   3,
		Developer:   4,
		Title:       "Far Cry® 5",
		Description: "Welcome to Hope County, Montana, home to a fanatical doomsday cult known as Eden’s Gate. Stand up to cult leader Joseph Seed & his siblings, the Heralds, to spark the fires of resistance & liberate the besieged community.",
		Price:       619000,
		ReleaseDate: "2018-03-27",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   10,
		Developer:   7,
		Title:       "Stormworks: Build and Rescue",
		Description: "Join a world where you design, create and pilot your own air sea rescue service. Release your inner hero as you battle fierce storms out at sea to rescue those in need.",
		Price:       119999,
		ReleaseDate: "2020-09-27",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   5,
		Developer:   8,
		Title:       "Need for Speed™ Heat",
		Description: "Hustle by day and risk it all at night in Need for Speed™ Heat Deluxe Edition, a white-knuckle street racer, where the lines of the law fade as the sun starts to set.",
		Price:       599999,
		ReleaseDate: "2019-11-08",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   1,
		Developer:   2,
		Title:       "Phasmophobia",
		Description: "Phasmophobia is a 4 player online co-op psychological horror. Paranormal activity is on the rise and it’s up to you and your team to use all the ghost hunting equipment at your disposal in order to gather as much evidence as you can.",
		Price:       89999,
		ReleaseDate: "2020-09-19",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   5,
		Developer:   10,
		Title:       "World Rally Championship",
		Description: "The most complete and authentic official WRC simulation yet. New physics for all surfaces, a completely redesigned career mode, dynamic weather conditions, 50 teams, 14 countries, 100 tracks, weekly challenges and an eSports...",
		Price:       204000,
		ReleaseDate: "2020-09-08",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   3,
		Developer:   4,
		Title:       "Resident Evil 7 Biohazard",
		Description: "Fear and isolation seep through the walls of an abandoned southern farmhouse. \"7\" marks a new beginning for survival horror with the “Isolated View” of the visceral new first-person perspective.",
		Price:       239999,
		ReleaseDate: "2017-01-14",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   5,
		Developer:   8,
		Title:       "Tom Clancy's Ghost Recon",
		Description: "Create a team with up to 3 friends in Tom Clancy’s Ghost Recon® Wildlands and enjoy the ultimate military shooter experience set in a massive, dangerous, and responsive open world.",
		Price:       515000,
		ReleaseDate: "2017-03-07",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   6,
		Developer:   6,
		Title:       "Unrailed!",
		Description: "Unrailed! is a co-op multiplayer game where you have to work together with your friends to build a train track across endless procedurally generated worlds. Master random encounters with its inhabitants, upgrade your train and keep it from derailing!",
		Price:       599999,
		ReleaseDate: "2020-09-24",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	//
	db.Create(&Game{
		Publisher:   2,
		Developer:   1,
		Title:       "Agent Murphy",
		Description: "Agent Murphy is a fast-paced, hardcore platformer, side-scrolling action shooter.",
		Price:       17499,
		ReleaseDate: "2021-01-20",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   2,
		Developer:   4,
		Title:       "Deep Despair",
		Description: "Deep Despair is a Minecraft-style survival sandbox. Travel, collect resources, craft items, dig, build, grow crops and survive. Freedom of action is limited only by your imagination!",
		Price:       11999,
		ReleaseDate: "2019-09-24",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   5,
		Developer:   2,
		Title:       "Ostrich Runner",
		Description: "Ostrich Runner – is the fascinating animated arcade racing where the main heroes are funny and a little bit crazy ostriches.",
		Price:       52999,
		ReleaseDate: "2005-02-30",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   6,
		Developer:   3,
		Title:       "Dead Cells",
		Description: "Dead Cells is a rogue-lite, metroidvania inspired, action-platformer. You'll explore a sprawling, ever-changing castle... assuming you’re able to fight your way past its keepers in 2D souls-lite combat. No checkpoints. Kill, die, learn, repeat.",
		Price:       149999,
		ReleaseDate: "2108-08-19",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   2,
		Developer:   5,
		Title:       "Night of the Dead",
		Description: "Build defenses against nightly waves of zombies. Survive and escape the island.",
		Price:       119999,
		ReleaseDate: "2020-08-28",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   7,
		Developer:   7,
		Title:       "Assassin's Creed® Unity",
		Description: "Assassin’s Creed® Unity tells the story of Arno, a young man who embarks upon an extraordinary journey to expose the true powers behind the French Revolution. In the brand new co-op mode, you and your friends will also be thrown in the middle of a ruthless struggle for the fate of a nation.",
		Price:       399000,
		ReleaseDate: "2014-11-23",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   3,
		Developer:   4,
		Title:       "Green Hell",
		Description: "Green Hell is a sweltering struggle for survival in the Amazonian rainforest. Clinging to life, the player is set on a journey of durability as the effects of solitude wear heavy not only on the body but also the mind. How long can you survive against the dangers of the unknown?",
		Price:       119999,
		ReleaseDate: "2019-09-05",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   2,
		Developer:   5,
		Title:       "Ori and the Will of the Wisps",
		Description: "Play the critically acclaimed masterpiece. Embark on a new journey in a vast, exotic world where you’ll encounter towering enemies and challenging puzzles on your quest to unravel Ori’s destiny.",
		Price:       149999,
		ReleaseDate: "2020-03-11",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   1,
		Developer:   2,
		Title:       "Pumpkin Jack",
		Description: "Pumpkin Jack is a Spooky Scary 3D platformer in which you embody Jack, the Mythical Pumpkin Lord! Dive into an Epic Adventure through otherworldly landscapes and help the Evil annihilate the Good!",
		Price:       139999,
		ReleaseDate: "2020-10-23",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   4,
		Developer:   6,
		Title:       "RedEx",
		Description: "RedEx is a 2D shooter with Tower Defence mechanics, with a minimalist pixel art stile. You are to fight against mercenaries of underground smuggling organization which pretends to be a logistic corporation called RedExpress.",
		Price:       39999,
		ReleaseDate: "2020-11-26",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
}


