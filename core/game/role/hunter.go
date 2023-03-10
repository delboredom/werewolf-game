package role

import (
	"uwwolf/game/action"
	"uwwolf/game/contract"
	"uwwolf/game/types"
	"uwwolf/game/vars"
)

type hunter struct {
	*role
}

func NewHunter(game contract.Game, playerID types.PlayerID) (contract.Role, error) {
	return &hunter{
			role: &role{
				id:           vars.HunterRoleID,
				phaseID:      vars.DayPhaseID,
				factionID:    vars.VillagerFactionID,
				beginRoundID: vars.FirstRound,
				turnID:       vars.HunterTurnID,
				game:         game,
				player:       game.Player(playerID),
				abilities: []*ability{
					{
						action:      action.NewKill(game),
						activeLimit: vars.OutOfTimes,
					},
				},
			},
		},
		nil
}

// OnAssign is triggered when the role is assigned to a player.
func (h *hunter) OnAssign() {
	//
}

// OnAfterDeath is triggered after killing this role.
func (h *hunter) OnAfterDeath() {
	diedAtPhaseID := h.game.Scheduler().PhaseID()

	// Ability is disabled if current round is too early
	if h.game.Scheduler().RoundID() < h.beginRoundID {
		return
	}

	// This turn can be only played in the current round
	slot := &types.NewTurnSlot{
		PhaseID:       h.phaseID,
		PlayerID:      h.player.ID(),
		RoleID:        h.id,
		PlayedRoundID: h.game.Scheduler().RoundID(),
	}

	if diedAtPhaseID == h.phaseID {
		// Play in next turn if he dies at his phase
		slot.TurnID = h.game.Scheduler().TurnID() + 1
	} else {
		// Play in his turn of the next day if he dies at
		// a time is not his phase
		slot.TurnID = h.turnID
	}

	h.abilities[0].activeLimit = vars.Once
	h.game.Scheduler().AddSlot(slot)
}
