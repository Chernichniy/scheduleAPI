package service

import (
	"net"
	"net/http"

	"github.com/Chernichniy/scheduleAPI/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run() error {

	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}
	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}

}

func BotRun(cfg config.Config) {
	if err := newService(cfg).botRun(); err != nil {
		panic(err)
	}
}

func (s *service) botRun() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	//Take options for bot (Errors handler, router, middleware(logs))
	opts := s.BotRouterInit()
	b, err := bot.New( /*os.Getenv(*/ , opts... /*)*/)

	if err != nil {
		panic(err)
	}

	s.log.Info("Bot started\n", b)
	b.Start(ctx)

	return bot.ErrorBadRequest
}
