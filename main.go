package main

import (
	"C"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Wieku/gosu-pp/beatmap"
	"github.com/Wieku/gosu-pp/performance/osu"
)

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }

import "github.com/Wieku/gosu-pp/beatmap/difficulty"

type beatmapInfo struct {
	osrPath string
	modsString string

}

func stringToInt(stringInt string) (int, error) {
	convertedStringValue := 0

	// Convert value to int and store it at convertedStringValue
	_, err := fmt.Sscan(stringInt, &convertedStringValue)

	// Return
	return convertedStringValue, err
}

func stringToMods(modsString string) (difficulty.Modifier, error){
	if modsString == "" {
		modsString = "DoubleTime|Hidden"
	}

	// Convert string into int
	allModNumbers := []string{"0", "1", "2", "4", "8", "16", "32", "64", "64", "128", "256", "512", "1024", "2048", "4096", "8192", "16384"}
	mods := []string{"None", "NoMod", "Easy", "TouchDevice", "Hidden", "HardRock", "SuddenDeath", "DoubleTime", "Relax", "HalfTime", "Nightcore", "Flashlight", "AutoPlay", "SpunOut", "Replay2", "Perfect"}

	modsConverted := modsString
	i := 0
	for i = 0 ; i < 16 ; i += 1 {
		modsConverted = strings.Replace(modsConverted, mods[i], allModNumbers[i], 1) 
		
	}
	modsConverted = string(modsConverted)

	// Sum mods 
	modsSplit := strings.Split(modsConverted, "|")
	modsSum := 0
	for _, value := range modsSplit {
		// Create value to store co
		convertedValue, err := stringToInt(value)
		if err != nil {
			log.Fatal(err)
		}
		
		// Sum up modsSum with convertedStringValue
		modsSum += convertedValue
	}
	modsSumDifficulty := difficulty.Modifier(modsSum)
	return modsSumDifficulty, nil
	// SV2 = 536870912
	// PF = 16384 
	// RX2 = 8192 AUTOPILOT
	// SO = 4096
	// AP = 2048
	// FL 1024
	// NC = 512
	// HT = 256
	// RX = 128
	// DT/NC = 64
	// SD = 32
	// HR = 16
	// HD = 8
	// TD = 4
	// EZ = 2
	// NM = 1
	// beatmap.Difficulty.SetMods(difficulty.ScoreV2)
}

func getStars(osuPath string, modsInt difficulty.Modifier) float64{
	// Open osu File
	osuFile, err := os.Open(osuPath) 
	if err != nil {
		log.Fatal(err)
	}
	
	// Parse beatmap
	beatmap, err := beatmap.ParseFromReader(osuFile)
	if err != nil {
		log.Fatal(err)
	}

	// Set Mods
	beatmap.Difficulty.SetMods(modsInt)

	// Calculate stars
	stars := (osu.CalculateSingle(beatmap.HitObjects, beatmap.Difficulty))
	
	// Return star rating
	return stars.Total
}

func getStarsAndPP(osuPath string, modsInt difficulty.Modifier, maxCombo, n300s, n100s, n50s, nmisses int) (float64, float64){

	// Open osu File
	osuFile, err := os.Open(osuPath) 
	if err != nil {
		log.Fatal(err)
	}
	
	// Parse beatmap
	beatmap, err := beatmap.ParseFromReader(osuFile)
	if err != nil {
		log.Fatal(err)
	}

	// Set Mods
	beatmap.Difficulty.SetMods(modsInt)

	// Calculate stars
	stars := (osu.CalculateSingle(beatmap.HitObjects, beatmap.Difficulty))

	if n300s < 0{
		n300s = stars.ObjectCount
	}

	// Get pp
	pp := &osu.PPv2{}
	pp.PPv2x(stars, maxCombo, n300s, n100s, n50s, nmisses, beatmap.Difficulty)

	// Return pp
	return stars.Total, pp.Results.Total
}

//export helloWorld
func helloWorld(namePtr *C.char, ) *C.char{
	name := C.GoString(namePtr)
	hello := ("Hello" + name)
	helloString := C.CString(hello)
	return helloString
}

//export pythonGetStars
func pythonGetStars(pathPtr *C.char, modsPtr *C.char) *C.char{
	pathString := C.GoString(pathPtr)
	modsString := C.GoString(modsPtr)

	if modsString == "" || pathString == "" {
		log.Fatal("String empty")
	}
	modsInt, err := stringToMods(modsString)
	if err != nil{
		log.Fatal(err)
	}

	starRating := getStars(pathString, modsInt)
	starRatingString := fmt.Sprintf("%v", starRating)
	starRatingCString := C.CString(starRatingString)
  
	return starRatingCString
}

//export pythonGetStarsAndPP
func pythonGetStarsAndPP(pathPtr, modsPtr, comboPtr, n300sPtr, n100sPtr, n50sPtr, nmissesPtr *C.char) *C.char{
	pathString := C.GoString(pathPtr)
	modsString := C.GoString(modsPtr)
	comboInteger, _ := stringToInt(C.GoString(comboPtr))
	n300sInteger, _ := stringToInt(C.GoString(n300sPtr))
	n100sInteger, _ := stringToInt(C.GoString(n100sPtr))
	n50sInteger, _ := stringToInt(C.GoString(n50sPtr))
	nmissesInteger, _ := stringToInt(C.GoString(nmissesPtr))

	if modsString == "" || pathString == "" {
		log.Fatal("String empty")
	}

	modsInt, err := stringToMods(modsString)
	if err != nil{
		log.Fatal(err)
	}

	_, ppTotal := getStarsAndPP(pathString, modsInt, comboInteger, n300sInteger, n100sInteger, n50sInteger, nmissesInteger)

	// Convert float to string 
	ppTotalString := fmt.Sprintf("%v", ppTotal)

	// Convert string to CString
	ppTotalCString := C.CString(ppTotalString)
	
	return ppTotalCString
}

func main() {
	mods, _ := stringToMods("DoubleTime|Easy")
	fmt.Println(mods)
	getStars("./gosu-pp/test/Avenged Sevenfold - Save Me (Drummer) [Tragedy].osu", mods)
}