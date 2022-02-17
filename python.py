import ctypes

library = ctypes.cdll.LoadLibrary('./main.so')

def get_star_rating(osu_path, mods):
    """
    Get star rating from specific beatmap!
    Usage:
    - osu_path: Type path to map's .osu difficulty file.
    - mods: Separated by "|", write mods, following this page https://osu.ppy.sh/wiki/en/Client/File_formats/Osr_(file_format) naming convention.
    Ex: "Hidden|HardRock", "DoubleTime", "DoubleTime|NightCore", "SuddenDeath|Perfect"
    note: NightCore requires Doubletime and Perfect requires SuddenDeath 
    
    Returns: rounded by 2 decimals star rating float
    """
    
    # Setup bridge
    go_get_star = library.pythonGetStars
    go_get_star.restype = ctypes.c_void_p
    
    # Run commands
    result = go_get_star(osu_path.encode("utf-8"), mods.encode("utf-8"))

    # Transform result
    result_bytes = ctypes.string_at(result)
    result_string = result_bytes.decode("utf-8")
    result_rounded = float("{0:.2f}".format(float(result_string)))
    
    return result_rounded

def get_pp(osu_path, mods, max_combo, n300s, n100s, n50s):
    """
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
    
    Returns: rounded by 2 decimals pp float
    """
    
    # Setup bridge
    go_get_pp = library.pythonGetStarsAndPP
    go_get_pp.restype = ctypes.c_void_p
    
    # Run commands
    result = go_get_pp(osu_path.encode("utf-8"), mods.encode("utf-8"), max_combo.encode("utf-8"), n300s.encode("utf-8"), n100s.encode("utf-8"), n50s.encode("utf-8"))

    # Transform result
    result_bytes = ctypes.string_at(result)
    result_string = result_bytes.decode("utf-8")
    result_rounded = float("{0:.2f}".format(float(result_string)))
    
    return result_rounded

def run():
    b = get_star_rating(osu_path="./gosu-pp/test/Avenged Sevenfold - Save Me (Drummer) [Tragedy].osu", mods="DoubleTime|HardRock")
    print(b, type(b))

    pp = get_pp("./gosu-pp/test/Avenged Sevenfold - Save Me (Drummer) [Tragedy].osu", "DoubleTime|HardRock", "-1", "-1", "0", "0")

    print(pp, type(pp))

if __name__ == "__main__":
    run()