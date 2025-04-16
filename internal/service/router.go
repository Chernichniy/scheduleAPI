package service

import (
	"context"

	"github.com/Chernichniy/scheduleAPI/internal/service/handlers"
	"github.com/Chernichniy/scheduleAPI/internal/service/handlers/telegram"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Bot's handlers
func (s *service) BotRouterInit() []bot.Option {
	var opts = []bot.Option{
		bot.WithMiddlewares(s.botLogginMiddleware),
		bot.WithErrorsHandler(func(err error) { s.log.Error(err) }),
		bot.WithDebugHandler(func(format string, args ...any) { s.log.Debug(args...) }),
		bot.WithAllowedUpdates(bot.AllowedUpdates{"message"}),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, telegram.BotHelloUser),
		bot.WithMessageTextHandler("/health", bot.MatchTypeExact, telegram.BotHealthMessage),
		bot.WithMessageTextHandler("/new_api_key", bot.MatchTypeExact, telegram.BotCreateAPIKey),
	}
	return opts
}
func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/scheduleAPI", func(r chi.Router) {
		//r.Route("/login", handlers.get_api_key)
		r.Post("/create_schedule", handlers.CreateScheduleRequest)
	})

	return r
}

func (s *service) botLogginMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		s.log.Info("Get message from: ", update.Message.From.Username, " ID: ", update.Message.From.ID, " Message: ", update.Message.Text)
	}
}
