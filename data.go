package rando

var (
	// negative-sounding adverbs beget intensity
	adv = []string{
		"abnormally",
		"absurdly",
		"awkwardly",
		"awfully",
		"consitently",
		"deeply",
		"dreadfully",
		"enourmously",
		"extremely",
		"filthy",
		"hopelessly",
		"horribly",
		"incredibly",
		"indescribably",
		"inexcusably",
		"insanely",
		"irredeemably",
		"offensively",
		"outrageously",
		"painfully",
		"preposterously",
		"ridiculously",
		"seriously",
		"really",
		"righteously",
		"terribly",
		"tremendously",
		"truly",
		"unholy",
		"unfortunately",
		"utterly",
		"wildly",
		"woefully",
	}

	// incomplete list of qualities you embody
	adj = []string{
		"abysmal",
		"annoying",
		"appalling",
		"asinine",
		"arrogant",
		"atrocious",
		"belligerent",
		"bitchass",
		"bitter",
		"cantankerous",
		"cretinous",
		"cynical",
		"decrepit",
		"dense",
		"depraved",
		"deplorable",
		"despicable",
		"disgusting",
		"dysfunctional",
		"feckless",
		"gross",
		"grotesque",
		"hermaphroditic",
		"hideous",
		"hopeless",
		"horrendous",
		"ignorant",
		"illiterate",
		"inbred",
		"incompetent",
		"inept",
		"insidious",
		"insignificant",
		"insufferable",
		"intolerable",
		"intolerant",
		"irresponsible",
		"miserable",
		"monsterous",
		"mundane",
		"naive",
		"narcissistic",
		"nasty",
		"nauseating",
		"obnoxious",
		"parasitic",
		"pathetic",
		"patronising",
		"pedestrian",
		"perverted",
		"petty",
		"plain",
		"pompous",
		"raging",
		"reprehensible",
		"repulsive",
		"revolting",
		"sinister",
		"skanky",
		"slutty",
		"smug",
		"spastic",
		"superficial",
		"syphilitic",
		"toxic",
		"ugly",
		"uncivilized",
		"unkempt",
		"unwanted",
		"vile",
		"vindictive",
		"worthless",
	}

	// incomplete list of things you are
	noun = []string{
		"ballsack",
		"bitch",
		"chode",
		"chodemuffin",
		"cockholster",
		"cockjockey",
		"cocksucker",
		"cuck",
		"cumbubble",
		"cumdumpster",
		"cumjunkie",
		"cunt",
		"cuntpuddle",
		"dickhead",
		"dickhole",
		"dickwad",
		"dipshit",
		"douchebag",
		"dumbass",
		"dumpsterslut",
		"faggot",
		"fuckface",
		"fucknugget",
		"fuckstick",
		"fucktard",
		"fuckwad",
		"fuckwit",
		"gutterslut",
		"jackass",
		"jizzmonger",
		"jizzmuffin",
		"lardass",
		"meatflap",
		"parasite",
		"pleb",
		"prick",
		"pussy",
		"queef",
		"retard",
		"schlong",
		"scrotum",
		"shitbird",
		"shithead",
		"shitpouch",
		"shitstain",
		"shitter",
		"skank",
		"slut",
		"spermatorium",
		"sphincter",
		"splooge-vacuum",
		"swine",
		"taint",
		"thundercunt",
		"turdburglar",
		"twat",
		"whore",
	}

	// compound adjectives that come to mind when thinking of you
	compoundAdj = []string{
		"alabama-hot-pocket-loving",
		"bed-wetting",
		"blue-waffle-loving",
		"bug-eyed",
		"chicken-diddling",
		"chromosome-deficient",
		"cock-eyed",
		"cock-juggling",
		"cock-sucking",
		"corpse-fucking",
		"cum-guzzling",
		"dick-infested",
		"dildo-munching",
		"dirty-sanchez-wearing",
		"disease-ridden",
		"drug-addled",
		"enema-addicted",
		"feeble-minded",
		"finger-banging",
		"fudge-packing",
		"glass-bottom-boat-loving",
		"horse-fucking",
		"irredeemably-boring",
		"kid-diddling",
		"narrow-minded",
		"orgasm-faking",
		"pig-fucking",
		"piss-drinking",
		"poon-guzzling",
		"sheep-fucking",
		"shit-eating",
		"simple-minded",
		"small-minded",
		"soul-destroying",
		"weak-willed",
	}

	// the story of your life
	finisher = []string{
		"cesspool of sub-human filth",
		"evangelical crusader of sub-mediocrity",
		"failure whose existence is meaningless and will be utterly forgotten",
		"molester of small furry animals",
		"plotless melodrama of an uneventful life",
		"poster-child for birth control",
		"practitioner of nonconsentual bestiality",
		"<noun> with delusions of adequacy",
		"<noun> with the natural grace of a startled elephant",
		"<noun> who has never had a conversation that didn't result in the other person just shaking their head in disappointment",
		"<noun> who hasn't yet realized that people just tolerate you",
		"<noun> who needs to accept that anyone who ever loved you was wrong",
		"<noun> who reminds everyone a lot like Rapunzel except instead of letting down your hair you just let down everyone you ever came into contact with",
		"<noun> who should go play in traffic or jump off a building",
		"<noun> whose a bigger disappointment than fucking Pluto turned out to be",
		"<noun> whose birth certificate is an apology from the condom factory",
		"<noun> whose life will never amount to anything",
		"<noun> whose like the bottom piece of bread — nobody touches you, and nobody wants you",
		"<noun> whose mother should have swallowed",
	}

	youNoFuckMe = `What the fuck did you just fucking type about my homie, you little bitch?

I’ll have you know I graduated top of my class at MIT, and I’ve been involved
in numerous secret raids with Anonymous, and I have over 300 confirmed DDoSes.
I am trained in the art of online trolling and I’m the top hacker in the
entire world.

You are nothing to me but just another virus host.

I will wipe you the fuck out, with precision the likes of which has never been
seen before on the Internet. Mark my fucking words.

You think you can get away with typing that shit about my friend over the
Internet? Think again, fucker.

As we chat over IRC I'm tracing your IP with my bare hands.
So you better prepare for the storm, maggot. The storm that wipes out the
pathetic little thing you call your computer.

You’re fucking dead, kid.

I can be anywhere, anytime, and hack into your files in over seven hundred ways,
and that’s just with my bare hands.

Not only am I extensively trained in hacking but I have access to the entire
arsenal of every piece of malware ever created, and I will use it to its
fullest extent to wipe your miserable ass off the face of the world wide web,
you little shit.

If only you could have known what unholy retribution your "clever" little
comment was about to bring down upon you, maybe you would have held your
fucking fingers. But you couldn’t, you didn’t, and now you’re paying the price,
you goddamn idiot.

I will shit code all over you and you will drown in it.

You’re fucking dead, kiddo.`

	cleanAdj = []string{
		"angry",
		"banana",
		"beech",
		"beef",
		"bizarre",
		"blue",
		"bookish",
		"chilly",
		"durable",
		"egregious",
		"extra",
		"green",
		"happy",
		"hopping",
		"irate",
		"junior",
		"latent",
		"loud",
		"odd",
		"orange",
		"peeved",
		"red",
		"silent",
		"sleepy",
		"slimy",
		"somber",
		"sticky",
		"strange",
		"united",
		"violent",
		"violet",
		"weird",
		"wobbly",
		"wrong",
		"yellow",
	}

	cleanNoun = []string{
		"ark",
		"artist",
		"auto",
		"bagel",
		"beam",
		"blunder",
		"bounce",
		"calendar",
		"chaser",
		"chef",
		"chipmunk",
		"curtain",
		"deity",
		"entourage",
		"farm",
		"feed",
		"felony",
		"fire",
		"genesis",
		"glee",
		"gopher",
		"iron",
		"llama",
		"master",
		"monkey",
		"monkey",
		"montana",
		"moon",
		"napkin",
		"net",
		"night",
		"panda",
		"penguin",
		"photo",
		"plow",
		"rage",
		"route",
		"scan",
		"seagull",
		"set",
		"source",
		"spatula",
		"spawn",
		"spork",
		"squirrel",
		"sucker",
		"toll",
		"tote",
		"trawl",
		"trinity",
		"trough",
		"waffle",
		"walk",
		"watch",
		"water",
		"whisper",
		"yard",
		"zipper",
	}
)