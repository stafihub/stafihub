package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// F = fis Balance (before)
// R = rToken Balance (before)
// f = fis added;
// r = rToken added
// P = existing Pool Units
// slipAdjustment = (1 - ABS((F r - f R)/((f + F) (r + R))))
// units = ((P (r F + R f))/(2 R F))*slipAdjustment
func calPoolUnit(oldPoolUnit, fisBalance, rTokenBalance, fisAmount, rTokenAmount sdk.Int) (totalUnit, addUnit sdk.Int) {
	if fisAmount.Equal(sdk.ZeroInt()) || rTokenAmount.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}
	if fisBalance.Add(fisAmount).Equal(sdk.ZeroInt()) || rTokenBalance.Add(rTokenAmount).Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	if fisBalance.Equal(sdk.ZeroInt()) || rTokenBalance.Equal(sdk.ZeroInt()) {
		return fisAmount, fisAmount
	}

	P := oldPoolUnit
	F := fisBalance
	R := rTokenBalance
	f := fisAmount
	r := rTokenAmount

	numerator := F.Mul(r).Add(f.Mul(R))
	rawUnit := numerator.Mul(P).Quo(R.Mul(F).Mul(sdk.NewInt(2)))
	slipAdjDenominator := F.Add(f).Mul(R.Add(r))
	adjUnit := F.Mul(r).Sub(f.Mul(R)).Abs().Quo(slipAdjDenominator)

	addUnit = rawUnit.Sub(adjUnit)
	totalUnit = P.Add(addUnit)
	return
}

// y = (x * X * Y) / (x + X)^2
// fee = (x^2 * Y)/(x + X)^2
func calSwapResult(fisBalance, rTokenBalance, inputAmount sdk.Int, inputIsFis bool) (y, fee sdk.Int) {
	if fisBalance.Equal(sdk.ZeroInt()) || rTokenBalance.Equal(sdk.ZeroInt()) || inputAmount.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	x := inputAmount
	X := rTokenBalance
	Y := fisBalance
	if inputIsFis {
		X = fisBalance
		Y = rTokenBalance
	}

	t := x.Add(X)
	denominator := t.Mul(t)
	y = x.Mul(X).Mul(Y).Quo(denominator)
	fee = x.Mul(x).Mul(Y).Quo(denominator)

	return
}

func calRemoveAmount(poolUnit, rmUnit, swapUnit, fisBalance, rtokenBalance sdk.Int, inputIsFis bool) (fisAmount, rtokenAmount, swapAmount sdk.Int) {
	if poolUnit.IsZero() || rmUnit.IsZero() {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	}
	if rmUnit.GT(poolUnit) {
		rmUnit = poolUnit
	}
	fisAmount = fisBalance.Mul(rmUnit).Quo(poolUnit)
	rtokenAmount = rtokenBalance.Mul(rmUnit).Quo(poolUnit)

	if inputIsFis {
		swapAmount = fisBalance.Mul(swapUnit).Quo(poolUnit)
	} else {
		swapAmount = rtokenBalance.Mul(swapUnit).Quo(poolUnit)
	}
	return
}
