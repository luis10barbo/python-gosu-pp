# python-gosu-pp

Little implementation of wieku's gosu-pp on python

With this, you can get updated PP and Star Rating values from maps in python
(btw, i'm not sure if it works without Golang installed)

# Credits

[wieku's gosu-pp](https://github.com/Wieku/gosu-pp) This project is literally a C "bridge" to connect Python and his GO code

# Documentation

def get\*star_rating(osu_path="", mods=""):

```
Get star rating from specific beatmap!
Usage: - osu_path: Type path to map's .osu difficulty file. - mods: Separated by "|", write mods, following this page https://osu.ppy.sh/wiki/en/Client/File_formats/Osr*(file_format) naming convention.
Ex: "Hidden|HardRock", "DoubleTime", "DoubleTime|NightCore", "SuddenDeath|Perfect"
note: NightCore requires Doubletime and Perfect requires SuddenDeath

Returns: rounded by 2 decimals star rating float
```

def get_pp(osu_path="", mods="NoMod", max_combo="-1", n300s="-1", n100s="0", n50s="", nmisses=""):

```
Get PP from specific beatmap!
Usage:
- osu_path: Type path to map's ".osu" difficulty file.
- mods: Separated by "|", write mods, following this page https://osu.ppy.sh/wiki/en/Client/File_formats/Osr_(file_format) naming convention.
Ex: "Hidden|HardRock", "DoubleTime", "DoubleTime|NightCore", "SuddenDeath|Perfect"
note: NightCore requires Doubletime and Perfect requires SuddenDeath
- combo: Type max combo achieved at play as string. Default = -1 , that means, FC.
- n300s: Type number of 300s achieved as string. Default = -1, that means, calculate 300 relative to other hits, starting on SS
- n100s: Type number of 100s achieved as string.
- n50s: Type number of 50s achieved as string.
- nmissess: Type number of misses achieved as string.

Returns: rounded by 2 decimals pp float
```
