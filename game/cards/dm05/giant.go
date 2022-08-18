package dm05

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)

// AvalancheGiant ...
func AvalancheGiant(c *match.Card) {

	c.Name = "Avalanche Giant"
	c.Power = 8000
	c.Civ = civ.Nature
	c.Family = family.Giant
	c.ManaCost = 6
	c.ManaRequirement = []string{civ.Nature}

	c.Use(fx.Creature, fx.Doublebreaker, fx.CantAttackCreatures, func(card *match.Card, ctx *match.Context) {

		if event, ok := ctx.Event.(*match.Battle); ok {
			if !event.Blocked || event.Attacker != card {
				return
			}

			opponent := ctx.Match.Opponent(card.Player)

			ctx.Match.BreakShields(fx.Select(
				card.Player,
				ctx.Match,
				opponent,
				match.SHIELDZONE,
				"Avalanche Giant: select shield to break",
				1,
				1,
				false,
			))

			ctx.Match.Chat("Server", fmt.Sprintf("Avalanche Giant broke one of %s's shield", opponent.Username()))

		}
	})

}
