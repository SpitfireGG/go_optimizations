package main

import "fmt"

type SecretHelper interface {
	GenerateRandomSeedNumber(seed int) int
	CheckifUserIsGay(gender string) bool
	CheckIfDogBarks(sound string) string
}

type MockSecretHelper struct{}

func (m *MockSecretHelper) GenerateRandomSeedNumber(seed int) int {
	return (seed * 3) % 100
}

func (m *MockSecretHelper) CheckifUserIsGay(gender string) bool {
	if gender == "male" || gender == "female" {
		return false
	}
	return true
}

func (m *MockSecretHelper) CheckIfDogBarks(sound string) string {
	if sound == "woff" || sound == "woff woff" || sound == "bhau bhau" {
		return sound
	}
	return "sound unknown"
}

func DoSecret(helper SecretHelper) string {
	genRandom := helper.GenerateRandomSeedNumber(11)
	checkGender := helper.CheckifUserIsGay("LGBTQ")
	checkDog := helper.CheckIfDogBarks("woff")
	return fmt.Sprintf("random number: %d, gender is gay: %t, dog says: %s", genRandom, checkGender, checkDog)
}
