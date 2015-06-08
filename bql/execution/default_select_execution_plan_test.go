package execution

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"pfi/sensorbee/sensorbee/bql/parser"
	"pfi/sensorbee/sensorbee/bql/udf"
	"pfi/sensorbee/sensorbee/core/tuple"
	"testing"
	"time"
)

func getTuples(num int) []*tuple.Tuple {
	tuples := make([]*tuple.Tuple, 0, num)
	for i := 0; i < num; i++ {
		tup := tuple.Tuple{
			Data: tuple.Map{
				"int": tuple.Int(i + 1),
			},
			InputName:     "input",
			Timestamp:     time.Date(2015, time.April, 10, 10, 23, i, 0, time.UTC),
			ProcTimestamp: time.Date(2015, time.April, 10, 10, 24, i, 0, time.UTC),
			BatchID:       7,
		}
		tuples = append(tuples, &tup)
	}
	return tuples
}

func createDefaultSelectPlan(s string, t *testing.T) (ExecutionPlan, error) {
	p := parser.NewBQLParser()
	reg := udf.NewDefaultFunctionRegistry()
	_stmt, err := p.ParseStmt(s)
	So(err, ShouldBeNil)
	So(_stmt, ShouldHaveSameTypeAs, parser.CreateStreamAsSelectStmt{})
	stmt := _stmt.(parser.CreateStreamAsSelectStmt)
	logicalPlan, err := Analyze(&stmt)
	So(err, ShouldBeNil)
	canBuild := CanBuildDefaultSelectExecutionPlan(logicalPlan, reg)
	So(canBuild, ShouldBeTrue)
	return NewDefaultSelectExecutionPlan(logicalPlan, reg)
}

func TestDefaultSelectExecutionPlan(t *testing.T) {
	// Select constant
	Convey("Given a SELECT clause with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(2) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then that constant should appear in %v", idx), func() {
					if idx == 0 {
						So(len(out), ShouldEqual, 1)
						So(out[0], ShouldResemble,
							tuple.Map{"col_1": tuple.Int(2)})
					} else {
						// nothing should be emitted because no new
						// data appears
						So(len(out), ShouldEqual, 0)
					}
				})
			}

		})
	})

	// Select a column with changing values
	Convey("Given a SELECT clause with only a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(int) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"int": tuple.Int(idx + 1)})
				})
			}

		})
	})

	// Select a non-existing column
	Convey("Given a SELECT clause with a non-existing column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(hoge) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for _, inTup := range tuples {
				_, err := plan.Process(inTup)
				So(err, ShouldNotBeNil) // hoge not found
			}

		})
	})

	Convey("Given a SELECT clause with a non-existing column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(hoge + 1) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for _, inTup := range tuples {
				_, err := plan.Process(inTup)
				So(err, ShouldNotBeNil) // hoge not found
			}

		})
	})

	// Select constant and a column with changing values
	Convey("Given a SELECT clause with a constant and a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(2, int) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"col_1": tuple.Int(2), "int": tuple.Int(idx + 1)})
				})
			}

		})
	})

	// Use alias
	Convey("Given a SELECT clause with a column alias", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(int-1 AS a, int AS b) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"a": tuple.Int(idx), "b": tuple.Int(idx + 1)})
				})
			}

		})
	})

	// Use wildcard
	Convey("Given a SELECT clause with a wildcard", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(*) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"int": tuple.Int(idx + 1)})
				})
			}

		})
	})

	Convey("Given a SELECT clause with a wildcard and an overriding column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(*, (int-1)*2 AS int) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"int": tuple.Int(2 * idx)})
				})
			}

		})
	})

	Convey("Given a SELECT clause with a column and an overriding wildcard", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM((int-1)*2 AS int, *) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"int": tuple.Int(idx + 1)})
				})
			}

		})
	})

	Convey("Given a SELECT clause with an aliased wildcard and an anonymous column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(* AS x, (int-1)*2) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					So(len(out), ShouldEqual, 1)
					So(out[0], ShouldResemble,
						tuple.Map{"col_2": tuple.Int(2 * idx), "x": tuple.Map{"int": tuple.Int(idx + 1)}})
				})
			}

		})
	})

	// Use a filter
	Convey("Given a SELECT clause with a column alias", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(int AS b) FROM src [RANGE 2 SECONDS] 
            WHERE int % 2 = 0`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			for idx, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("Then those values should appear in %v", idx), func() {
					if (idx+1)%2 == 0 {
						So(len(out), ShouldEqual, 1)
						So(out[0], ShouldResemble,
							tuple.Map{"b": tuple.Int(idx + 1)})
					} else {
						So(len(out), ShouldEqual, 0)
					}
				})
			}

		})
	})

	// RSTREAM/2 SECONDS
	Convey("Given an RSTREAM/2 SECONDS statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT RSTREAM(2 AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then the whole state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[1]), ShouldEqual, 2)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[1][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 3)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[2][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[2][2], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[3]), ShouldEqual, 3)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[3][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[3][2], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
			})

		})
	})

	Convey("Given an RSTREAM/2 SECONDS statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT RSTREAM(int AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then the whole state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(len(output[1]), ShouldEqual, 2)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(output[1][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 3)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(output[2][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[2][2], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(len(output[3]), ShouldEqual, 3)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[3][1], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(output[3][2], ShouldResemble, tuple.Map{"a": tuple.Int(4)})
			})

		})
	})

	// RSTREAM/2 TUPLES
	Convey("Given an RSTREAM/2 SECONDS statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT RSTREAM(2 AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then the whole state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[1]), ShouldEqual, 2)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[1][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 2)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[2][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[3]), ShouldEqual, 2)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[3][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
			})

		})
	})

	Convey("Given an RSTREAM/2 SECONDS statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT RSTREAM(int AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then the whole window state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(len(output[1]), ShouldEqual, 2)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(output[1][1], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 2)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(output[2][1], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(len(output[3]), ShouldEqual, 2)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(output[3][1], ShouldResemble, tuple.Map{"a": tuple.Int(4)})
			})

		})
	})

	// ISTREAM/2 SECONDS
	Convey("Given an ISTREAM/2 SECONDS statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(2 AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then new items in state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 0)
				So(len(output[3]), ShouldEqual, 0)
			})

		})
	})

	Convey("Given an ISTREAM/2 SECONDS statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(int AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then new items in state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(len(output[1]), ShouldEqual, 1)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 1)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(len(output[3]), ShouldEqual, 1)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(4)})
			})

		})
	})

	// ISTREAM/2 TUPLES
	Convey("Given an ISTREAM/2 TUPLES statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(2 AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then new items in state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 0)
				So(len(output[3]), ShouldEqual, 0)
			})

		})
	})

	Convey("Given an ISTREAM/2 TUPLES statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT ISTREAM(int AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then new items in state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 1)
				So(output[0][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(len(output[1]), ShouldEqual, 1)
				So(output[1][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
				So(len(output[2]), ShouldEqual, 1)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(3)})
				So(len(output[3]), ShouldEqual, 1)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(4)})
			})

		})
	})

	// DSTREAM/2 SECONDS
	Convey("Given a DSTREAM/2 SECONDS statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT DSTREAM(2 AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then items dropped from state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 0)
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 0)
				So(len(output[3]), ShouldEqual, 0)
			})

		})
	})

	Convey("Given a DSTREAM/2 SECONDS statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT DSTREAM(int AS a) FROM src [RANGE 2 SECONDS]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then items dropped from state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 0)
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 0)
				So(len(output[3]), ShouldEqual, 1)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
			})

		})
	})

	// DSTREAM/2 TUPLES
	Convey("Given a DSTREAM/2 TUPLES statement with a constant", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT DSTREAM(2 AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then items dropped from state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 0)
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 0)
				So(len(output[3]), ShouldEqual, 0)
			})

		})
	})

	Convey("Given a DSTREAM/2 TUPLES statement with a column", t, func() {
		tuples := getTuples(4)
		s := `CREATE STREAM box AS SELECT DSTREAM(int AS a) FROM src [RANGE 2 TUPLES]`
		plan, err := createDefaultSelectPlan(s, t)
		So(err, ShouldBeNil)

		Convey("When feeding it with tuples", func() {
			output := [][]tuple.Map{}
			for _, inTup := range tuples {
				out, err := plan.Process(inTup)
				So(err, ShouldBeNil)
				output = append(output, out)
			}

			Convey("Then items dropped from state should be emitted", func() {
				So(len(output), ShouldEqual, 4)
				So(len(output[0]), ShouldEqual, 0)
				So(len(output[1]), ShouldEqual, 0)
				So(len(output[2]), ShouldEqual, 1)
				So(output[2][0], ShouldResemble, tuple.Map{"a": tuple.Int(1)})
				So(len(output[3]), ShouldEqual, 1)
				So(output[3][0], ShouldResemble, tuple.Map{"a": tuple.Int(2)})
			})

		})
	})
}