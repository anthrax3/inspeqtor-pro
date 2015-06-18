/*
 */
package parser

const numNTSymbols = 10

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // Check
		2,  // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Check
		-1, // Preamble
		4,  // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		6,  // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		10, // Job
		7,  // JobList

	},
	gotoRow{ // S5

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S6

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		10, // Job
		11, // JobList

	},
	gotoRow{ // S7

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		12, // Job
		-1, // JobList

	},
	gotoRow{ // S8

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		13, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S9

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S10

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S11

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		12, // Job
		-1, // JobList

	},
	gotoRow{ // S12

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S13

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S14

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S15

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S16

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S17

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		20, // IntAmount
		21, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S18

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		28, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S19

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S20

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		29, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S21

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S22

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S23

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S24

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S25

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S26

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S27

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S28

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
	gotoRow{ // S29

		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // Period
		-1, // Job
		-1, // JobList

	},
}
