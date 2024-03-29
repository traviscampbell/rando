package rando

var (
	// negative-sounding adverbs beget intensity
	adv = []string{
		"absurdly",
		"annoyingly",
		"astonishingly",
		"brutally",
		"deliberately",
		"dreadfully",
		"exceptionally",
		"extraordinarily",
		"extremely",
		"filthy",
		"horribly",
		"incredibly",
		"indescribably",
		"inexcusably",
		"inexplicably",
		"insanely",
		"irredeemably",
		"offensively",
		"outrageously",
		"painfully",
		"preposterously",
		"remarkably",
		"ridiculously",
		"righteously",
		"seriously",
		"suprisingly",
		"terribly",
		"tremendously",
		"truly",
		"unexpectedly",
		"unfortunately",
		"unholy",
		"unsettlingly",
		"utterly",
		"wildly",
		"woefully",
	}

	// incomplete list of qualities you embody
	adj = []string{
		"abhorrent",
		"absurd",
		"abysmal",
		"annoying",
		"appalling",
		"asinine",
		"arrogant",
		"atrocious",
		"belligerent",
		"bitchy",
		"bitter",
		"blundering",
		"callous",
		"cantankerous",
		"careless",
		"clumsy",
		"cockamamie",
		"conceited",
		"confused",
		"cowardly",
		"crazy",
		"creepy",
		"cretinous",
		"cruel",
		"cynical",
		"decrepit",
		"delirious",
		"dense",
		"depraved",
		"deplorable",
		"depressed",
		"demented",
		"deranged",
		"despicable",
		"dirty",
		"disgusting",
		"disillusioned",
		"disturbed",
		"dysfunctional",
		"envious",
		"evil",
		"filthy",
		"flabby",
		"flacid",
		"floppy",
		"foolish",
		"feckless",
		"fragile",
		"frail",
		"frightened",
		"futile",
		"gross",
		"grotesque",
		"helpless",
		"hermaphroditic",
		"hideous",
		"hopeless",
		"horrendous",
		"horrible",
		"hostile",
		"idiotic",
		"ignorant",
		"illiterate",
		"impatient",
		"impulsive",
		"inbred",
		"incompetent",
		"inconsequential",
		"inept",
		"insidious",
		"insignificant",
		"insufferable",
		"intolerable",
		"intolerant",
		"irrational",
		"irresponsible",
		"irritating",
		"lonely",
		"ludicrous",
		"malicious",
		"manipulative",
		"miserable",
		"misinformed",
		"monsterous",
		"mundane",
		"naive",
		"narcissistic",
		"nasty",
		"nauseating",
		"neurotic",
		"nonsensical",
		"oblivious",
		"obnoxious",
		"oppressive",
		"overbearing",
		"pathetic",
		"patronising",
		"pernicious",
		"perverted",
		"pessimistic",
		"pestiferous",
		"petty",
		"pompous",
		"preposterous",
		"presumptuous",
		"prickly",
		"raging",
		"reckless",
		"repulsive",
		"resentful",
		"revolting",
		"sadistic",
		"selfish",
		"shallow",
		"shiftless",
		"shoddy",
		"sinister",
		"sloppy",
		"slutty",
		"smug",
		"sour",
		"spastic",
		"stingy",
		"stupid",
		"superficial",
		"syphilitic",
		"terrible",
		"thoughtless",
		"thunderous",
		"timid",
		"toxic",
		"trifling",
		"trivial",
		"troubled",
		"twisted",
		"ugly",
		"uncivilized",
		"uncultured",
		"uneducated",
		"uninformed",
		"uninspired",
		"unkempt",
		"unreasonable",
		"unsightly",
		"unwanted",
		"unwise",
		"unworthy",
		"uptight",
		"vile",
		"vindictive",
		"worthless",
		"yellowbellied",
	}

	// incomplete list of things you are
	noun = []string{
		"assclown",
		"asshat",
		"asshole",
		"ass-Kisser",
		"asswipe",
		"abomination",
		"ballsack",
		"bastard",
		"bitch",
		"bumfuck",
		"butt-Muncher",
		"carpet-Muncher",
		"charlatan",
		"chode",
		"chode-Muffin",
		"chump",
		"cockholster",
		"cock-Jockey",
		"cocksucker",
		"coward",
		"crackhead",
		"crackwhore",
		"cuck",
		"cum-Bubble",
		"cum-Dumpster",
		"cum-Junkie",
		"cumrag",
		"cunt",
		"cunt-Puddle",
		"dickhead",
		"dickhole",
		"dick-For-Brains",
		"dickwad",
		"ding-A-Ling",
		"dipshit",
		"douchebag",
		"dumbass",
		"dumbbitch",
		"dumbfuck",
		"dumpsterslut",
		"failure",
		"faggot",
		"feltcher",
		"flake",
		"fuckface",
		"fucknugget",
		"fuckpig",
		"fuckstick",
		"fucktard",
		"fuckwad",
		"fuckwit",
		"gutterslut",
		"hillbilly",
		"ignoramous",
		"imbicile",
		"incel",
		"jackass",
		"jizz-Monger",
		"jizz-Muffin",
		"lowlife",
		"meatflap",
		"moron",
		"parasite",
		"pervert",
		"prick",
		"pussy",
		"queef",
		"retard",
		"schlong",
		"scrotum",
		"scumbag",
		"shitpouch",
		"shitstain",
		"shitter",
		"simp",
		"sissy",
		"skank",
		"skank-Whore",
		"snowflake",
		"spermatorium",
		"sphincter",
		"splooge-Vacuum",
		"taint",
		"thunder-Cunt",
		"tryhard",
		"turd-Burglar",
		"twat",
		"virgin",
		"weirdo",
	}

	// compound adjectives that come to mind when thinking of you
	// suppose these could be made on the fly by using a formula like
	// <noun>-<adj>... idk I'm have to learn way too much engrish for this shit.
	compoundAdj = []string{
		"alabama-hotpocket-loving",
		"ball-chortling",
		"bean-flicking",
		"bed-wetting",
		"bitch-ass",
		"blue-waffle-loving",
		"bow-legged",
		"bug-eyed",
		"chicken-diddling",
		"chromosome-deficient",
		"cock-eyed",
		"cock-infested",
		"cock-juggling",
		"cock-sucking",
		"corpse-fucking",
		"cum-guzzling",
		"dick-infested",
		"chode-infested",
		"dildo-munching",
		"dirty-sanchez-wearing",
		"disease-ridden",
		"drug-addled",
		"duckbutter-loving",
		"enema-addicted",
		"finger-banging",
		"foopa-fondling",
		"fudge-packing",
		"glass-bottom-boat-loving",
		"hippo-humping",
		"horse-fucking",
		"kid-diddling",
		"monkey-spanking",
		"orgasm-faking",
		"pig-fucking",
		"pigeon-toed",
		"piss-drinking",
		"poon-guzzling",
		"salad-tossing",
		"self-centered",
		"short-tempered",
		"sheep-fucking",
		"shit-eating",
		"soul-destroying",
		"stim-fapping",
		"taint-tickling",
		"tampon-sucking",
		"trick-ass",
		"trick-turning",
		"weak-willed",
	}

	// the story of your life
	finisher = []string{
		"cesspool of sub-human filth",
		"evangelical crusader of sub-mediocrity",
		"molester of small furry animals",
		"plotless melodrama of an uneventful life",
		"poster-child for birth control",
		"practitioner of nonconsentual bestiality",
		"<noun> with delusions of adequacy",
		"<noun> who hasn't realized that people just tolerate <you><them>",
		"<noun> who looks like two pounds of shit in a one pound bag",
		"<noun> who needs to accept that anyone who ever loved <you><them> regrets it",
		"<noun> who reminds me of Rapunzel, not because <you><they> let down <your><their> hair but because <you><they> let down everyone in <your><their> life <you><they> come into contact with",
		"<noun> who should go step in front of a train on <your><their> way to jump off the nearest tall building",
		"<noun> who should slip into something more comfortable... like a coma",
		"<noun> who thought HD-DVD was really going to take off",
		"<noun> who's a bigger disappointment than fucking Pluto turned out to be",
		"<noun> who's like the first piece of bread -- everyone touches <you><them> but no one wants <you><them>",
		"<noun> who's like the last piece of bread -- nobody wants <you><them> and the closer <you><they> get the sadder they become",
		"<noun> who's so fake <you><they> make Barbie jealous",
		"<noun> who's so ugly that <you><they> scare the shit out of toilets",
		"<noun> whose birth certificate is an apology from the condom factory",
		"<noun> whose existence is meaningless and only brings pain and suffering to those around <you><them>",
		"<noun> whose life will never amount to anything",
		"<noun> whose mother should have swallowed",
		"<noun> whose never had a conversation with anyone that didn't end with the other person shaking their head in disappointment",
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
So you better prepare for the storm, maggot.
The storm that wipes out the pathetic little thing you call your computer.

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
		"absurd",
		"angry",
		"bizarre",
		"blatant",
		"blue",
		"bookish",
		"broken",
		"calm",
		"chilly",
		"chubby",
		"comical",
		"creepy",
		"dumb",
		"durable",
		"egregious",
		"extra",
		"farcical",
		"flacid",
		"green",
		"happy",
		"hopping",
		"hungry",
		"hurt",
		"ill",
		"inane",
		"inexpensive",
		"irate",
		"jealous",
		"junior",
		"kinky",
		"kooky",
		"lame",
		"latent",
		"laughable",
		"lazy",
		"limp",
		"loud",
		"lumpy",
		"mundane",
		"nosy",
		"obtuse",
		"orange",
		"outrageous",
		"peevish",
		"petite",
		"puny",
		"quirky",
		"rancid",
		"rapid",
		"red",
		"silent",
		"silly",
		"sleepy",
		"slimy",
		"somber",
		"sticky",
		"story",
		"strange",
		"stupid",
		"timid",
		"tough",
		"united",
		"violent",
		"violet",
		"vivid",
		"wimpy",
		"whacky",
		"wobbly",
		"wonky",
		"wrong",
		"yellow",
		"zany",
	}

	cleanNoun = []string{
		"alpaca",
		"airhead",
		"ark",
		"artist",
		"auto",
		"bagel",
		"banana",
		"beach",
		"beef",
		"beam",
		"beer",
		"bird",
		"bladder",
		"blunder",
		"book",
		"bottom",
		"bounce",
		"brat",
		"burrito",
		"butt",
		"calendar",
		"chef",
		"cheese",
		"chipmunk",
		"chump",
		"coin",
		"cornhole",
		"couch",
		"crack",
		"crackpot",
		"curiosity",
		"curtain",
		"deceit",
		"deity",
		"devil",
		"dingbat",
		"dingdong",
		"dingleberry",
		"dipstick",
		"doofus",
		"donkey",
		"dork",
		"dummy",
		"empathy",
		"ending",
		"entourage",
		"eye",
		"farm",
		"fathead",
		"feed",
		"feet",
		"felony",
		"fire",
		"flake",
		"flower",
		"fool",
		"genesis",
		"glee",
		"goof",
		"goon",
		"gopher",
		"hatred",
		"head",
		"hippo",
		"hypotenuse",
		"iron",
		"jerk",
		"jerkface",
		"knucklehead",
		"ladder",
		"llama",
		"master",
		"monkey",
		"moon",
		"napkin",
		"nerd",
		"net",
		"night",
		"nincompoop",
		"nitwit",
		"noodle",
		"nut",
		"oddball",
		"panda",
		"paperclip",
		"pasta",
		"penguin",
		"phony",
		"photo",
		"pillow",
		"pinhead",
		"plow",
		"poot",
		"pouch",
		"quack",
		"rain",
		"rage",
		"route",
		"salt",
		"satisfaction",
		"schlub",
		"schmuck",
		"skunk",
		"slobe",
		"snowflake",
		"spatula",
		"spawn",
		"spork",
		"sqeaker",
		"squirrel",
		"staircase",
		"stinker",
		"steak",
		"sucker",
		"suitcase",
		"sympathy",
		"toad",
		"toot",
		"top",
		"tote",
		"trinity",
		"troll",
		"tryhard",
		"turd",
		"turkey",
		"ulcer",
		"urine",
		"vanity",
		"vapor",
		"waffle",
		"walk",
		"watch",
		"water",
		"weiner",
		"weirdo",
		"whisker",
		"whisper",
		"wimp",
		"wuss",
		"yard",
		"zealot",
		"zipper",
	}

	// accidents happen. i mean just look at you...
	compliment = []string{
		"must be part beaver because, DAAAMMMMNN",
		"is more fun to be around than a bag of blow and a hooker",
		"must have a mirror in <your><their> pocket, because I can see myself in <your><their> pants",
	}
)
