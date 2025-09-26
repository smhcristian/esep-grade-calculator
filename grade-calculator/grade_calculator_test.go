package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//test if the final grade is C when the averages of assignments, exams and essays are between 70 and 79
func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 75, Exam)
	gradeCalculator.AddGrade("essay", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//test if multiple assignments are added, the average should be calculated correctly
func TestMultipleAssignments(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 100, Assignment)
	gradeCalculator.AddGrade("assignment 2", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 85, Exam)
	gradeCalculator.AddGrade("essay", 88, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//test if no grades have been added, the final grade should be "F"
func TestNoGrades(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()
	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

//test the stringer method of the GradeType enum
func TestGradeTypeStringer(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment.String() to return 'assignment'; got '%s' instead", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected Exam.String() to return 'exam'; got '%s' instead", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected Essay.String() to return 'essay'; got '%s' instead", Essay.String())
	}
}
func TestPassFail_Pass(t *testing.T) {
	expected := "Pass"
	// Passing is defined as a C (70) or greater.
	gradeCalculator := NewPassFailGradeCalculator()

	gradeCalculator.AddGrade("assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam", 75, Exam)
	gradeCalculator.AddGrade("essay", 70, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected grade to be '%s' but got '%s'", expected, actual)
	}
}

// Test a clear failing grade in Pass/Fail mode.
func TestPassFail_Fail(t *testing.T) {
	expected := "Fail"
	gradeCalculator := NewPassFailGradeCalculator()

	// This combination results in a weighted average of 66.
	gradeCalculator.AddGrade("assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam", 75, Exam)
	gradeCalculator.AddGrade("essay", 68, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected grade to be '%s' but got '%s'", expected, actual)
	}
}

// Test the boundary condition for passing (exactly 70).
func TestPassFail_BoundaryPass(t *testing.T) {
	expected := "Pass"
	gradeCalculator := NewPassFailGradeCalculator()

	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 70, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected grade to be '%s' but got '%s'", expected, actual)
	}
}
