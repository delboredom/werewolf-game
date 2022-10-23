package action

import (
	"uwwolf/app/enum"
	"uwwolf/app/game/contract"
	"uwwolf/app/types"
	"uwwolf/app/validator"
)

type action[S any] struct {
	id         types.ActionId
	state      S
	game       contract.Game
	expiration types.Expiration
}

type validateFnc = func(req *types.ActionRequest) string

type executeFnc = func(req *types.ActionRequest) *types.ActionResponse

func (a *action[S]) Id() types.ActionId {
	return a.id
}

func (a *action[S]) Expiration() types.Expiration {
	return a.expiration
}

func (a *action[S]) State() any {
	return a.state
}

func (a *action[S]) Perform(req *types.ActionRequest) *types.ActionResponse {
	return a.perform(a.validate, a.execute, req)
}

// Execute the action request after it passes the validation.
func (a *action[S]) perform(validate validateFnc, execute executeFnc, req *types.ActionRequest) *types.ActionResponse {
	if a.expiration == enum.OutOfTimes {
		return &types.ActionResponse{
			Ok:           false,
			PerformError: "This action exceeds the max number of uses!",
		}
	}

	if err := validator.ValidateStruct(req); err != nil {
		return &types.ActionResponse{
			Ok:              false,
			ValidationError: err,
		}
	}

	// Apply specific validate if general validation is passed
	if alert := validate(req); alert != "" {
		return &types.ActionResponse{
			Ok:           false,
			PerformError: alert,
		}
	}

	if a.expiration != enum.UnlimitedTimes && !req.IsSkipped {
		a.expiration--
	}

	return execute(req)
}

// Validate the action request. Each action has different rules
// for validation. Return empty string if everything is ok.
func (a *action[S]) validate(req *types.ActionRequest) string {
	return ""
}

// Execute action request with receied data. Returning struct with Ok
// field is false means the request could not be fulfilled, otherwise execution
// is successful.
func (a *action[S]) execute(req *types.ActionRequest) *types.ActionResponse {
	return nil
}
