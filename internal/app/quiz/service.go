package quiz

import (
	"context"

	"hackathon-dsm-server/internal/pkg/utils/errors"
)

type Service interface {
	GetQuiz(ctx context.Context, heritageID int) (*QuizData, error)
	SubmitQuiz(ctx context.Context, userID, heritageID int, answers []int) (*SubmitData, error)
}

type QuizService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &QuizService{repo: repo}
}

func (s *QuizService) GetQuiz(ctx context.Context, heritageID int) (*QuizData, error) {
	quiz, err := s.repo.GetQuizByHeritageID(ctx, heritageID)
	if err != nil {
		return nil, err
	}

	questions, err := s.repo.GetQuestionsByQuizID(ctx, quiz.QuizID)
	if err != nil {
		return nil, err
	}

	var questionData []QuestionData
	for _, q := range questions {
		questionData = append(questionData, QuestionData{
			ID:            q.QuestionID,
			Question:      q.QuestionText,
			Options:       []string{q.Option1, q.Option2, q.Option3, q.Option4},
			CorrectAnswer: q.CorrectAnswer - 1, // Convert to 0-based index
		})
	}

	return &QuizData{
		Questions: questionData,
	}, nil
}

func (s *QuizService) SubmitQuiz(ctx context.Context, userID, heritageID int, answers []int) (*SubmitData, error) {
	// Check user's visit status first
	visitStatus, err := s.repo.GetUserVisitStatus(ctx, userID, heritageID)
	if err != nil {
		return nil, err
	}

	// Allow quiz retries if not completed with perfect score
	if visitStatus.QuizCompleted && visitStatus.QuizScore == 100 {
		return nil, errors.BadRequest("Quiz already completed with perfect score")
	}

	quiz, err := s.repo.GetQuizByHeritageID(ctx, heritageID)
	if err != nil {
		return nil, err
	}

	questions, err := s.repo.GetQuestionsByQuizID(ctx, quiz.QuizID)
	if err != nil {
		return nil, err
	}

	if len(answers) != len(questions) {
		return nil, errors.BadRequest("Number of answers must match number of questions")
	}

	// Validate answer indices
	for _, answer := range answers {
		if answer < 0 || answer > 3 {
			return nil, errors.BadRequest("Answer indices must be between 0 and 3")
		}
	}

	correctCount := 0
	var incorrectAnswers []IncorrectAnswer
	
	for i, question := range questions {
		if i < len(answers) {
			if answers[i] == question.CorrectAnswer-1 { // Convert to 0-based
				correctCount++
			} else {
				// Record incorrect answer for feedback
				incorrectAnswers = append(incorrectAnswers, IncorrectAnswer{
					QuestionID:    question.QuestionID,
					UserAnswer:    answers[i],
					CorrectAnswer: question.CorrectAnswer - 1, // Convert to 0-based
					QuestionText:  question.QuestionText,
				})
			}
		}
	}

	score := (correctCount * 100) / len(questions)
	allCorrect := correctCount == len(questions)
	canRetry := !allCorrect

	// Always update quiz score (allows retries)
	if err := s.repo.UpdateUserVisitScore(ctx, userID, heritageID, score); err != nil {
		return nil, err
	}

	var newBadge *BadgeData
	if allCorrect {
		badge, err := s.repo.GetBadgeByHeritageID(ctx, heritageID)
		if err != nil {
			return nil, err
		}

		if badge != nil {
			exists, err := s.repo.CheckUserBadgeExists(ctx, userID, badge.BadgeID)
			if err != nil {
				return nil, err
			}

			if !exists {
				if err := s.repo.CreateUserBadge(ctx, userID, badge.BadgeID); err != nil {
					return nil, err
				}

				newBadge = &BadgeData{
					Name:     badge.Name,
					ImageURL: badge.ImageURL,
				}
			}
		}
	}

	return &SubmitData{
		Score:            score,
		CorrectCount:     correctCount,
		TotalQuestions:   len(questions),
		AllCorrect:       allCorrect,
		CanRetry:         canRetry,
		IncorrectAnswers: incorrectAnswers,
		NewBadge:         newBadge,
	}, nil
}