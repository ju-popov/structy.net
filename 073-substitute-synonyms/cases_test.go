package substitutesynonyms_test

type testCase struct {
	name     string
	sentance string
	synonyms map[string][]string
	expected []string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		sentance: "follow the yellow brick road",
		synonyms: map[string][]string{
			"follow": {"chase", "pursue"},
			"yellow": {"gold", "amber", "lemon"},
		},
		expected: []string{
			"chase the gold brick road",
			"chase the amber brick road",
			"chase the lemon brick road",
			"pursue the gold brick road",
			"pursue the amber brick road",
			"pursue the lemon brick road",
		},
	},
	{
		name:     "test_01",
		sentance: "I think it\"s gonna be a long long time",
		synonyms: map[string][]string{
			"think": {"believe", "reckon"},
			"long":  {"lengthy", "prolonged"},
		},
		expected: []string{
			"I believe it\"s gonna be a lengthy lengthy time",
			"I believe it\"s gonna be a lengthy prolonged time",
			"I believe it\"s gonna be a prolonged lengthy time",
			"I believe it\"s gonna be a prolonged prolonged time",
			"I reckon it\"s gonna be a lengthy lengthy time",
			"I reckon it\"s gonna be a lengthy prolonged time",
			"I reckon it\"s gonna be a prolonged lengthy time",
			"I reckon it\"s gonna be a prolonged prolonged time",
		},
	},
	{
		name:     "test_02",
		sentance: "palms sweaty knees weak arms heavy",
		synonyms: map[string][]string{
			"palms": {"hands", "fists"},
			"heavy": {"weighty", "hefty", "burdensome"},
			"weak":  {"fragile", "feeble", "frail", "sickly"},
		},
		expected: []string{
			"hands sweaty knees fragile arms weighty",
			"hands sweaty knees fragile arms hefty",
			"hands sweaty knees fragile arms burdensome",
			"hands sweaty knees feeble arms weighty",
			"hands sweaty knees feeble arms hefty",
			"hands sweaty knees feeble arms burdensome",
			"hands sweaty knees frail arms weighty",
			"hands sweaty knees frail arms hefty",
			"hands sweaty knees frail arms burdensome",
			"hands sweaty knees sickly arms weighty",
			"hands sweaty knees sickly arms hefty",
			"hands sweaty knees sickly arms burdensome",
			"fists sweaty knees fragile arms weighty",
			"fists sweaty knees fragile arms hefty",
			"fists sweaty knees fragile arms burdensome",
			"fists sweaty knees feeble arms weighty",
			"fists sweaty knees feeble arms hefty",
			"fists sweaty knees feeble arms burdensome",
			"fists sweaty knees frail arms weighty",
			"fists sweaty knees frail arms hefty",
			"fists sweaty knees frail arms burdensome",
			"fists sweaty knees sickly arms weighty",
			"fists sweaty knees sickly arms hefty",
			"fists sweaty knees sickly arms burdensome",
		},
	},
}
