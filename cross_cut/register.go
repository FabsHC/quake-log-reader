package cross_cut

import "quake-log-reader/internal/application/usecase"

type Register struct {
	ProcessEventUseCase *usecase.ProcessEventUseCase
}

func NewRegister() *Register {
	return &Register{
		ProcessEventUseCase: usecase.NewProcessEventUseCase(),
	}
}
