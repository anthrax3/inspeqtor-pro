package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(3), /* check */
			nil,      /* jobs */
			nil,      /* with */
			nil,      /* name */
			nil,      /* , */
			nil,      /* hour */
			nil,      /* hours */
			nil,      /* minute */
			nil,      /* minutes */
			nil,      /* day */
			nil,      /* days */
			nil,      /* happens */
			nil,      /* every */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* check */
			nil,          /* jobs */
			nil,          /* with */
			nil,          /* name */
			nil,          /* , */
			nil,          /* hour */
			nil,          /* hours */
			nil,          /* minute */
			nil,          /* minutes */
			nil,          /* day */
			nil,          /* days */
			nil,          /* happens */
			nil,          /* every */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* check */
			shift(5), /* jobs */
			nil,      /* with */
			nil,      /* name */
			nil,      /* , */
			nil,      /* hour */
			nil,      /* hours */
			nil,      /* minute */
			nil,      /* minutes */
			nil,      /* day */
			nil,      /* days */
			nil,      /* happens */
			nil,      /* every */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			reduce(3), /* jobs, reduce: Preamble */
			nil,       /* with */
			nil,       /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* check */
			nil,      /* jobs */
			shift(8), /* with */
			shift(9), /* name */
			nil,      /* , */
			nil,      /* hour */
			nil,      /* hours */
			nil,      /* minute */
			nil,      /* minutes */
			nil,      /* day */
			nil,      /* days */
			nil,      /* happens */
			nil,      /* every */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			reduce(4), /* with, reduce: Checktype */
			reduce(4), /* name, reduce: Checktype */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* check */
			nil,      /* jobs */
			nil,      /* with */
			shift(9), /* name */
			nil,      /* , */
			nil,      /* hour */
			nil,      /* hours */
			nil,      /* minute */
			nil,      /* minutes */
			nil,      /* day */
			nil,      /* days */
			nil,      /* happens */
			nil,      /* every */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Check */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(9),  /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(14), /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			nil,       /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			shift(15), /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(17), /* $, reduce: JobList */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(17), /* name, reduce: JobList */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Check */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(9),  /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(18), /* $, reduce: JobList */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(18), /* name, reduce: JobList */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			reduce(5), /* name, reduce: ParameterList */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(16), /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			nil,       /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			shift(17), /* every */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			reduce(7), /* name, reduce: Parameters */
			shift(18), /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(19), /* name */
			nil,       /* , */
			shift(22), /* hour */
			shift(23), /* hours */
			shift(24), /* minute */
			shift(25), /* minutes */
			shift(26), /* day */
			shift(27), /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			shift(14), /* name */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			nil,       /* name */
			nil,       /* , */
			reduce(8), /* hour, reduce: IntAmount */
			reduce(8), /* hours, reduce: IntAmount */
			reduce(8), /* minute, reduce: IntAmount */
			reduce(8), /* minutes, reduce: IntAmount */
			reduce(8), /* day, reduce: IntAmount */
			reduce(8), /* days, reduce: IntAmount */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			nil,       /* name */
			nil,       /* , */
			shift(22), /* hour */
			shift(23), /* hours */
			shift(24), /* minute */
			shift(25), /* minutes */
			shift(26), /* day */
			shift(27), /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(16), /* $, reduce: Job */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(16), /* name, reduce: Job */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: Period */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			reduce(9), /* name, reduce: Period */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: Period */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(10), /* name, reduce: Period */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: Period */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(11), /* name, reduce: Period */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: Period */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(12), /* name, reduce: Period */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: Period */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(13), /* name, reduce: Period */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(14), /* $, reduce: Period */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(14), /* name, reduce: Period */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* check */
			nil,       /* jobs */
			nil,       /* with */
			reduce(6), /* name, reduce: Parameters */
			nil,       /* , */
			nil,       /* hour */
			nil,       /* hours */
			nil,       /* minute */
			nil,       /* minutes */
			nil,       /* day */
			nil,       /* days */
			nil,       /* happens */
			nil,       /* every */

		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(15), /* $, reduce: Job */
			nil,        /* check */
			nil,        /* jobs */
			nil,        /* with */
			reduce(15), /* name, reduce: Job */
			nil,        /* , */
			nil,        /* hour */
			nil,        /* hours */
			nil,        /* minute */
			nil,        /* minutes */
			nil,        /* day */
			nil,        /* days */
			nil,        /* happens */
			nil,        /* every */

		},
	},
}
