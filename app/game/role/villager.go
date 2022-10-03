package role

import (
	"uwwolf/app/enum"
	"uwwolf/app/game/action"
	"uwwolf/app/game/contract"
	"uwwolf/app/types"
)

func newVillager(game contract.Game, playerId types.PlayerId) contract.Role {
	return &role{
		id:           enum.VillagerRoleId,
		factionId:    enum.VillagerFactionId,
		phaseId:      enum.DayPhaseId,
		game:         game,
		player:       game.Player(playerId),
		beginRoundId: enum.FirstRound,
		priority:     1,
		score:        1,
		set:          -1,
		actions: map[uint]contract.Action{
			enum.VoteActionId: action.NewVote(game, &types.VoteActionSetting{
				FactionId: enum.VillagerFactionId,
				PlayerId:  playerId,
				Weight:    1,
			}),
		},
	}
}
