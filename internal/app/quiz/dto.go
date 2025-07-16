package quiz

type SubmitRequest struct {
	Answers []int `json:"answers" binding:"required"`
}

type QuizResponse struct {
	Success bool     `json:"success"`
	Data    QuizData `json:"data"`
}

type QuizData struct {
	Questions []QuestionData `json:"questions"`
}

type QuestionData struct {
	ID            int      `json:"id"`
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectAnswer int      `json:"correctAnswer"`
}

type SubmitResponse struct {
	Success bool       `json:"success"`
	Data    SubmitData `json:"data"`
}

type SubmitData struct {
	Score           int              `json:"score"`
	CorrectCount    int              `json:"correctCount"`
	TotalQuestions  int              `json:"totalQuestions"`
	AllCorrect      bool             `json:"allCorrect"`
	CanRetry        bool             `json:"canRetry"`
	IncorrectAnswers []IncorrectAnswer `json:"incorrectAnswers,omitempty"`
	NewBadge        *BadgeData       `json:"newBadge,omitempty"`
}

type BadgeData struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type IncorrectAnswer struct {
	QuestionID     int    `json:"questionId"`
	UserAnswer     int    `json:"userAnswer"`
	CorrectAnswer  int    `json:"correctAnswer"`
	QuestionText   string `json:"questionText"`
}