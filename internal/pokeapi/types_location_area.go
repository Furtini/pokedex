package pokeapi

type LocationAreaResult struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            string `json:"name_index"`
	EncounterMethodRates []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEntountes []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int `json:"min_level"`
				MaxLevel        int `json:"max_level"`
				ConditionValues []struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		}
	} `json:"pokemon_encounters"`
}
