package parser

import "github.com/mperham/inspeqtor-pro/jobs/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Check	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble Checktype ParameterList JobList	<< ast.NewJobCheck(X[3], X[2]), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      1,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewJobCheck(X[3], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble Checktype JobList	<< ast.NewJobCheck(X[2], map[string]string{}), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewJobCheck(X[2], map[string]string{}), nil
		},
	},
	ProdTabEntry{
		String: `Preamble : "check"	<< X[0], nil >>`,
		Id:         "Preamble",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Checktype : "jobs"	<< X[0], nil >>`,
		Id:         "Checktype",
		NTType:     3,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ParameterList : "with" Parameters	<< X[1], nil >>`,
		Id:         "ParameterList",
		NTType:     4,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Parameters : name name "," Parameters	<< ast.AddParam(X[0], X[1], X[3]) >>`,
		Id:         "Parameters",
		NTType:     5,
		Index:      6,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddParam(X[0], X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Parameters : name name	<< ast.AddParam(X[0], X[1], nil) >>`,
		Id:         "Parameters",
		NTType:     5,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddParam(X[0], X[1], nil)
		},
	},
	ProdTabEntry{
		String: `IntAmount : name	<< ast.ToInt64(X[0]) >>`,
		Id:         "IntAmount",
		NTType:     6,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ToInt64(X[0])
		},
	},
	ProdTabEntry{
		String: `Period : "hour"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Period : "hours"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Period : "minute"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Period : "minutes"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Period : "day"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Period : "days"	<< X[0], nil >>`,
		Id:         "Period",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Job : name "happens" "every" IntAmount Period	<< ast.NewJob(X[0], X[3], X[4]), nil >>`,
		Id:         "Job",
		NTType:     8,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewJob(X[0], X[3], X[4]), nil
		},
	},
	ProdTabEntry{
		String: `Job : name "happens" "every" Period	<< ast.NewJob(X[0], int64(1), X[3]), nil >>`,
		Id:         "Job",
		NTType:     8,
		Index:      16,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewJob(X[0], int64(1), X[3]), nil
		},
	},
	ProdTabEntry{
		String: `JobList : Job	<< ast.NewJobList(X[0]), nil >>`,
		Id:         "JobList",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewJobList(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `JobList : JobList Job	<< ast.AppendJob(X[0], X[1]), nil >>`,
		Id:         "JobList",
		NTType:     9,
		Index:      18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendJob(X[0], X[1]), nil
		},
	},
}
