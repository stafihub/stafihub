package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// F = pool baseToken Balance (before)
// R = poo token Balance (before)
// f = baseToken added;
// r = token added
// P = existing Pool Units
// slipAdjustment = (1 - ABS((F r - f R)/((f + F) (r + R))))
// units = ((P (r F + R f))/(2 R F))*slipAdjustment
func CalPoolUnit(oldPoolUnit, baseTokenBalance, tokenBalance, baseTokenAmount, tokenAmount sdk.Int) (totalUnit, addUnit sdk.Int) {
	if oldPoolUnit.IsNegative() || baseTokenBalance.IsNegative() || tokenBalance.IsNegative() ||
		baseTokenAmount.IsNegative() || tokenAmount.IsNegative() {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	if baseTokenAmount.IsZero() && tokenAmount.IsZero() {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}
	if baseTokenBalance.Add(baseTokenAmount).IsZero() || tokenBalance.Add(tokenAmount).IsZero() {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	if baseTokenBalance.IsZero() || tokenBalance.IsZero() {
		return baseTokenAmount, baseTokenAmount
	}

	P := sdk.NewDecFromInt(oldPoolUnit)
	F := sdk.NewDecFromInt(baseTokenBalance)
	R := sdk.NewDecFromInt(tokenBalance)
	f := sdk.NewDecFromInt(baseTokenAmount)
	r := sdk.NewDecFromInt(tokenAmount)

	// P(r F + R f)
	numerator := P.Mul(F.Mul(r).Add(f.Mul(R)))
	// (P (r F + R f))/(2 R F)
	rawUnit := numerator.Quo(R.Mul(F).Mul(sdk.NewDec(2)))

	// (f + F) (r + R))
	slipAdjDenominator := F.Add(f).Mul(R.Add(r))
	// ABS(F r - f R)
	slipAdjNumerator := F.Mul(r).Sub(f.Mul(R)).Abs()
	slipAdjustment := sdk.OneDec().Sub(slipAdjNumerator.Quo(slipAdjDenominator))

	addUnit = rawUnit.Mul(slipAdjustment).TruncateInt()
	totalUnit = oldPoolUnit.Add(addUnit)
	return
}

// y = (x * X * Y) / (x + X)^2
// fee = (x^2 * Y)/(x + X)^2
func CalSwapResult(baseTokenBalance, tokenBalance, inputAmount sdk.Int, inputIsBase bool) (y, fee sdk.Int) {
	if !baseTokenBalance.IsPositive() || !tokenBalance.IsPositive() || !inputAmount.IsPositive() {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	x := inputAmount
	X := tokenBalance
	Y := baseTokenBalance
	if inputIsBase {
		X = baseTokenBalance
		Y = tokenBalance
	}

	t := x.Add(X)
	denominator := t.Mul(t)
	y = x.Mul(X).Mul(Y).Quo(denominator)
	fee = x.Mul(x).Mul(Y).Quo(denominator)

	return
}

func CalRemoveAmount(poolUnit, rmUnit, swapUnit, baseTokenBalance, tokenBalance sdk.Int, inputIsBase bool) (baseTokenAmount, tokenAmount, swapAmount sdk.Int) {
	if swapUnit.IsNegative() || baseTokenBalance.IsNegative() || tokenBalance.IsNegative() {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	}

	if !poolUnit.IsPositive() || !rmUnit.IsPositive() {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	}
	if rmUnit.GT(poolUnit) {
		rmUnit = poolUnit
	}
	baseTokenAmount = baseTokenBalance.Mul(rmUnit).Quo(poolUnit)
	tokenAmount = tokenBalance.Mul(rmUnit).Quo(poolUnit)

	if inputIsBase {
		swapAmount = baseTokenBalance.Mul(swapUnit).Quo(poolUnit)
	} else {
		swapAmount = tokenBalance.Mul(swapUnit).Quo(poolUnit)
	}
	return
}
