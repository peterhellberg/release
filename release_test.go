package release

import "testing"

var releaseTests = []struct {
	Input      string
	Title      string
	Year       int
	Season     int
	Episode    int
	Type       string
	Category   string
	Resolution string
	Format     string
	Audio      string
	Group      string
}{
	{"Act.of.Valor.2012.720p.BluRay.DTS.x264-HDxT", "Act of Valor", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDxT"},
	{"American.Reunion.UNRATED.720p.BluRay.X264-BLOW", "American Reunion", 0, 0, 0, "UNRATED", "BluRay", "720p", "x264", "", "BLOW"},
	{"Attack.The.Block.2011.720p.BluRay.x264-iNFAMOUS", "Attack The Block", 2011, 0, 0, "", "BluRay", "720p", "x264", "", "iNFAMOUS"},
	{"Avatar.720p.Bluray.x264-CBGB", "Avatar", 0, 0, 0, "", "BluRay", "720p", "x264", "", "CBGB"},
	{"Captain.Phillips.2013.BluRay.720p.DTS.x264-CHD", "Captain Phillips", 2013, 0, 0, "", "BluRay", "720p", "x264", "DTS", "CHD"},
	{"Chronicle.2012.DC.720p.BluRay.x264-REFiNED", "Chronicle", 2012, 0, 0, "DC", "BluRay", "720p", "x264", "", "REFiNED"},
	{"Cleanskin.2012.BluRay.720p.DTS.x264-CHD", "Cleanskin", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "CHD"},
	{"Closed.Circuit.2013.720p.BDRIP.x264.AAC-WiNNy", "Closed Circuit", 2013, 0, 0, "", "BluRay", "720p", "x264", "AAC", "WiNNy"},
	{"Cloud.Atlas.2012.720p.BluRay.x264-BiDA [PublicHD]", "Cloud Atlas", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "BiDA"},
	{"Dead.Man.Down.2013.720p.BluRay.x264-SPARKS", "Dead Man Down", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Detachment.2011.LIMITED.720p.BluRay.x264-TRiPS", "Detachment", 2011, 0, 0, "LIMITED", "BluRay", "720p", "x264", "", "TRiPS"},
	{"Django.Unchained.2012.720p.BluRay.x264-SPARKS", "Django Unchained", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Elysium.2013.720p.WEB-DL.H264-WEBiOS", "Elysium", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "WEBiOS"},
	{"Enders.Game.2014.720p.HDRip.x264-BRENYS", "Enders Game", 2014, 0, 0, "", "HDRip", "720p", "x264", "", "BRENYS"},
	{"Equilibrium.2002.720p.BluRay.DTS.x264-HiDt", "Equilibrium", 2002, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HiDt"},
	{"Fast.and.Furious.6.2013.720p.WEB-DL.H264-HDCLUB", "Fast and Furious 6", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "HDCLUB"},
	{"Flu.2013.720p.BRRip.XviD-TeRRa", "Flu", 2013, 0, 0, "", "BluRay", "720p", "xvid", "", "TeRRa"},
	{"Friends.with.Kids.2011.720p.BluRay.DD5.1.x264-HiDt", "Friends with Kids", 2011, 0, 0, "", "BluRay", "720p", "x264", "DD5.1", "HiDt"},
	{"Get.the.Gringo.2012.BluRay.720p.DTS.x264-CHD", "Get the Gringo", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "CHD"},
	{"Gravity.2013.720p.WEB-DL.H264-PublicHD", "Gravity", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "PublicHD"},
	{"Her.2013.720p.BluRay.x264-WiNNy", "Her", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "WiNNy"},
	{"Inception.720p.BluRay.x264-CROSSBOW", "Inception", 0, 0, 0, "", "BluRay", "720p", "x264", "", "CROSSBOW"},
	{"Intolerable.Cruelty.2003.720p.Bluray.X264-DIMENSION", "Intolerable Cruelty", 2003, 0, 0, "", "BluRay", "720p", "x264", "", "DIMENSION"},
	{"Intouchables.2011.720p.BluRay.x264.DTS-HDChina", "Intouchables", 2011, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDChina"},
	{"Iron.Man.2.720p.BluRay.x264-EXQUiSiTE", "Iron Man 2", 0, 0, 0, "", "BluRay", "720p", "x264", "", "EXQUiSiTE"},
	{"Iron.Man.3.2013.720p.WEB-DL.H264-WEBiOS [PublicHD]", "Iron Man 3", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "WEBiOS"},
	{"Its.All.Gone.Pete.Tong.2004.720p.BluRay.x264-CiNEFiLE", "Its All Gone Pete Tong", 2004, 0, 0, "", "BluRay", "720p", "x264", "", "CiNEFiLE"},
	{"Jack Reacher.2012 BDRip 720p DTS x264-MarGe", "Jack Reacher", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "MarGe"},
	{"Jobs.2013.720p.BluRay.x264-BLOW", "Jobs", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "BLOW"},
	{"Kick-Ass.2.2013.720p.Blu.Ray.x264-VeDeTT", "Kick-Ass 2", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "VeDeTT"},
	{"Looper.2012.720p.BluRay.x264-SPARKS", "Looper", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Man of Tai Chi 2013 720p Webrip x264 AC3-FooKaS (SilverTorrent)", "Man of Tai Chi", 2013, 0, 0, "", "Webrip", "720p", "x264", "AC3", "FooKaS"},
	{"Menace.II.Society.1993.Directors.Cut.720p.BluRay.x264-SiNNERS", "Menace II Society", 1993, 0, 0, "", "BluRay", "720p", "x264", "", "SiNNERS"},
	{"Need.For.Speed.2014.720p.BluRay.x264-SPARKS", "Need For Speed", 2014, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Nymphomaniac.Vol.II.2013.720p.BluRay.X264-AMIABLE", "Nymphomaniac Vol II", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"Oblivion.2013.720p.BluRay.x264-SPARKS", "Oblivion", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Pacific.Rim.2013.720p.WEB-DL.H264-PublicHD", "Pacific Rim", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "PublicHD"},
	{"Pain.and.Gain.2013.720p.BluRay.X264-AMIABLE", "Pain and Gain", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"Pi.1998.LiMiTED.720p.BluRay.x264-SiNNERS", "Pi", 1998, 0, 0, "LIMITED", "BluRay", "720p", "x264", "", "SiNNERS"},
	{"Pornopung.2013.720p.BluRay.x264-SUG", "Pornopung", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "SUG"},
	{"Project.X.2012.Extended.BluRay.720p.DTS.x264-CHD", "Project X", 2012, 0, 0, "EXTENDED", "BluRay", "720p", "x264", "DTS", "CHD"},
	{"Red.2.2013.720p.BDRIP.x264.AAC-WiNNy", "Red 2", 2013, 0, 0, "", "BluRay", "720p", "x264", "AAC", "WiNNy"},
	{"Robocop.1987.4K.Remastered.DC.CEE.720p.BluRay.x264-WiNNy", "Robocop", 1987, 0, 0, "REMASTERED DC", "BluRay", "720p", "x264", "", "WiNNy"},
	{"Robocop.2014.720p.BluRay.x264-WiNNy", "Robocop", 2014, 0, 0, "", "BluRay", "720p", "x264", "", "WiNNy"},
	{"Scott.Pilgrim.vs.The.World.2010.720p.x264-F", "Scott Pilgrim vs The World", 2010, 0, 0, "", "", "720p", "x264", "", "F"},
	{"Secretary.2002.720p.BluRay.x264.DTS-WiKi", "Secretary", 2002, 0, 0, "", "BluRay", "720p", "x264", "DTS", "WiKi"},
	{"Sliding.Doors.1998.720p.BluRay.X264-AMIABLE", "Sliding Doors", 1998, 0, 0, "", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"Snitch.2013.720p.BluRay.DTS.x264-HDWinG", "Snitch", 2013, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDWinG"},
	{"Snowpiercer.2013.720p.BluRay.x264-WiNNy", "Snowpiercer", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "WiNNy"},
	{"Star.Trek.Into.Darkness.2013.720p.WEB-DL.h264-PublicHD", "Star Trek Into Darkness", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "PublicHD"},
	{"Ted.2012.720p.BluRay.x264-DAA", "Ted", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "DAA"},
	{"The Bourne Legacy 2012 BRRip 720p DTS x264 MarGe", "The Bourne Legacy", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "MarGe"},
	{"The Lone Ranger (2013)", "The Lone Ranger", 2013, 0, 0, "", "", "", "", "", ""},
	{"The Way Way Back 2013 BluRay 720p DTS x264 dxva-FraMeSToR", "The Way Way Back", 2013, 0, 0, "", "BluRay", "720p", "x264", "DTS", "FraMeSToR"},
	{"The.Amazing.Spider-Man.2012.720p.BluRay.X264-AMIABLE", "The Amazing Spider-Man", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"The.Best.Exotic.Marigold.Hotel.2011.720p.BluRay.X264-AMIABLE [PublicHD]", "The Best Exotic Marigold Hotel", 2011, 0, 0, "", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"The.Bucket.List.2007.MULTiSubs.720p.BluRay.DTS.x264-BLiNK", "The Bucket List", 2007, 0, 0, "MULTISUBS", "BluRay", "720p", "x264", "DTS", "BLiNK"},
	{"The.Fault.in.Our.Stars.2014.EXTENDED.720p.WEB-DL.H264.AC3-FAS", "The Fault in Our Stars", 2014, 0, 0, "EXTENDED", "WEB-DL", "720p", "h264", "AC3", "FAS"},
	{"The.Five-Year.Engagement.2012.UNRATED.720p.BRRip.XviD.AC3-LEGi0N", "The Five-Year Engagement", 2012, 0, 0, "UNRATED", "BluRay", "720p", "xvid", "AC3", "LEGi0N"},
	{"The.Hand.That.Rocks.The.Cradle.1992.720p.BluRay.x264-HD4U", "The Hand That Rocks The Cradle", 1992, 0, 0, "", "BluRay", "720p", "x264", "", "HD4U"},
	{"The.Hobbit.An.Unexpected.Journey.2012.720p.BluRay.x264-SPARKS", "The Hobbit An Unexpected Journey", 2012, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"The.Hunger.Games.2012.720p.BluRay.x264.DTS-HDChina", "The Hunger Games", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDChina"},
	{"The.Hunger.Games.Catching.Fire.2013.IMAX.EDITION.720p.BluRay.x264-PHD", "The Hunger Games Catching Fire", 2013, 0, 0, "IMAX EDITION", "BluRay", "720p", "x264", "", "PHD"},
	{"The.Lucky.One.2012.720p.Bluray.DTS.x264-HDxT", "The Lucky One", 2012, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDxT"},
	{"The.Machine.2013.720p.BluRay.x264-WiNNy", "The Machine", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "WiNNy"},
	{"The.Magic.Of.Belle.Isle.2012.LiMiTED.720p.BluRay.x264-SPARKS", "The Magic Of Belle Isle", 2012, 0, 0, "LIMITED", "BluRay", "720p", "x264", "", "SPARKS"},
	{"The.Secret.Life.Of.Walter.Mitty.2013.720p.WEB-DL.H264-PublicHD", "The Secret Life Of Walter Mitty", 2013, 0, 0, "", "WEB-DL", "720p", "h264", "", "PublicHD"},
	{"The.Wolf.of.Wall.Street.2013.720p.BluRay.x264-WiNNy", "The Wolf of Wall Street", 2013, 0, 0, "", "BluRay", "720p", "x264", "", "WiNNy"},
	{"The.Wolverine.2013.EXTENDED.720p.BluRay.x264-SPARKS[rarbg]", "The Wolverine", 2013, 0, 0, "EXTENDED", "BluRay", "720p", "x264", "", "SPARKS"},
	{"The.Zero.Theorem.2013.LIMITED.READ.NFO.720p.BluRay.X264-AMIABLE", "The Zero Theorem", 2013, 0, 0, "LIMITED", "BluRay", "720p", "x264", "", "AMIABLE"},
	{"Total.Recall.2012.720p.Extended.BluRay.x264.AC3-HDChina", "Total Recall", 2012, 0, 0, "EXTENDED", "BluRay", "720p", "x264", "AC3", "HDChina"},
	{"Transcendence.2014.720p.BluRay.x264-SPARKS", "Transcendence", 2014, 0, 0, "", "BluRay", "720p", "x264", "", "SPARKS"},
	{"Trolljegeren.2010.NORWEGIAN.PROPER.720p.BluRay.x264-NOHD", "Trolljegeren", 2010, 0, 0, "PROPER", "BluRay", "720p", "x264", "", "NOHD"},
	{"Tron.Legacy.2010.BluRay.720p.DTS.x264-CHD", "Tron Legacy", 2010, 0, 0, "", "BluRay", "720p", "x264", "DTS", "CHD"},
	{"Wanderlust.720p.BluRay.X264-BLOW", "Wanderlust", 0, 0, 0, "", "BluRay", "720p", "x264", "", "BLOW"},
	{"We're the Millers (2013)", "We're the Millers", 2013, 0, 0, "", "", "", "", "", ""},
	{"Weekender.2011.720p.BluRay.X264-7SinS", "Weekender", 2011, 0, 0, "", "BluRay", "720p", "x264", "", "7SinS"},
	{"Welcome to the Punch 2013 720p BluRay DTS x264-TayTO", "Welcome to the Punch", 2013, 0, 0, "", "BluRay", "720p", "x264", "DTS", "TayTO"},
	{"While.You.Were.Sleeping.1995.720p.BluRay.2Audio.DTS.AC3.x264-HDWinG", "While You Were Sleeping", 1995, 0, 0, "", "BluRay", "720p", "x264", "DTS", "HDWinG"},
	{"White.House.Down.2013.720p.BluRay.DTS.x264-PublicHD", "White House Down", 2013, 0, 0, "", "BluRay", "720p", "x264", "DTS", "PublicHD"},
	{"Wild.Bill.2011.720p.BluRay.X264-7SinS", "Wild Bill", 2011, 0, 0, "", "BluRay", "720p", "x264", "", "7SinS"},
	{"World.War.Z.2013.Unrated.Cut.720p.BluRay.x264.DTS-WiKi [PublicHD]", "World War Z", 2013, 0, 0, "UNRATED", "BluRay", "720p", "x264", "DTS", "WiKi"},
	{"The.Handmaids.Tale.S02E01.June.1080p.HULU.WEBRip.AAC2.0.H264-NTb[rarbg]", "The Handmaids Tale S02E01 June", 0, 2, 1, "", "WEBRip", "1080p", "h264", "AAC2.0", "NTb"},
	{"The.Mandalorian.S01E01.INTERNAL.1080p.WEB.H264-NOCROP[rartv]", "The Mandalorian S01E01 INTERNAL", 0, 1, 1, "INTERNAL", "", "1080p", "h264", "", "NOCROP"},
	{"Mr.Robot.S01E07.720p.HDTV.x264-KILLERS[rarbg]", "Mr Robot S01E07", 0, 1, 7, "", "HDTV", "720p", "x264", "", "KILLERS"},
}

func TestParse(t *testing.T) {
	for _, tt := range releaseTests {
		tt := tt

		t.Run(tt.Input, func(t *testing.T) {
			r, err := Parse(tt.Input)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}

			if r.Title != tt.Title {
				t.Errorf("r.Title = %v, want %v", r.Title, tt.Title)
			}

			if r.Year != tt.Year {
				t.Errorf("r.Year = %d, want %d", r.Year, tt.Year)
			}

			if r.Season != tt.Season {
				t.Errorf("r.Season = %d, want %d", r.Season, tt.Season)
			}

			if r.Episode != tt.Episode {
				t.Errorf("r.Episode = %d, want %d", r.Episode, tt.Episode)
			}

			if r.Type != tt.Type {
				t.Errorf("r.Type = %v, want %v", r.Type, tt.Type)
			}

			if r.Category != tt.Category {
				t.Errorf("r.Category = %v, want %v", r.Category, tt.Category)
			}

			if r.Resolution != tt.Resolution {
				t.Errorf("r.Resolution = %v, want %v", r.Resolution, tt.Resolution)
			}

			if r.Format != tt.Format {
				t.Errorf("r.Format = %v, want %v", r.Format, tt.Format)
			}

			if r.Audio != tt.Audio {
				t.Errorf("r.Audio = %v, want %v", r.Audio, tt.Audio)
			}

			if r.Group != tt.Group {
				t.Errorf("r.Group = %v, want %v", r.Group, tt.Group)
			}

			if r.Input != tt.Input {
				t.Errorf("r.Input = %v, want %v", r.Input, tt.Input)
			}
		})
	}
}
