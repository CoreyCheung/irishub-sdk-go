package modules

import (
	"context"
	"fmt"
	"github.com/irisnet/core-sdk-go/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
	"github.com/irisnet/irishub-sdk-go/modules/token"
	"github.com/irisnet/irishub-sdk-go/utils/cache"
	"github.com/tendermint/tendermint/libs/log"
	"strings"
)


type tokenQuery struct {
	q sdk.Queries
	sdk.GRPCClient
	cdc codec.Codec
	log.Logger
	cache.Cache
}

func (l tokenQuery) QueryToken(denom string) (sdk.Token, error) {
	denom = strings.ToLower(denom)
	if t, err := l.Get(l.prefixKey(denom)); err == nil {
		return t.(sdk.Token), nil
	}

	conn, err := l.GenConn()
	defer func() { _ = conn.Close() }()
	if err != nil {
		return sdk.Token{}, sdkerrors.Wrapf(ErrGenConn,err.Error())
	}

	response, err := token.NewQueryClient(conn).Token(
		context.Background(),
		&token.QueryTokenRequest{Denom: denom},
	)
	if err != nil {
		l.Debug("client query token failed",
			" denom ", denom,
			" err ", err.Error())
		return sdk.Token{}, nil
	}

	var srcToken token.TokenInterface
	if err = l.cdc.UnpackAny(response.Token, &srcToken); err != nil {
		return sdk.Token{}, sdkerrors.Wrapf(ErrUnpackAny,err.Error())
	}
	token := srcToken.(*token.Token).Convert().(sdk.Token)
	l.SaveTokens(token)
	return token, nil
}

func (l tokenQuery) SaveTokens(tokens ...sdk.Token) {
	for _, t := range tokens {
		err1 := l.Set(l.prefixKey(t.Symbol), t)
		err2 := l.Set(l.prefixKey(t.MinUnit), t)
		if err1 != nil || err2 != nil {
			l.Debug("cache token failed", "symbol", t.Symbol)
		}
	}
}

func (l tokenQuery) ToMinCoin(coins ...sdk.DecCoin) (dstCoins sdk.Coins, err error) {
	for _, coin := range coins {
		token, err := l.QueryToken(coin.Denom)
		if err != nil {
			return nil, sdkerrors.Wrapf(ErrQueryToken,err.Error())
		}

		minCoin, err := token.GetCoinType().ConvertToMinCoin(coin)
		if err != nil {
			return nil, sdkerrors.Wrapf(ErrConvertToMinCoin,err.Error())
		}
		dstCoins = append(dstCoins, minCoin)
	}
	return dstCoins.Sort(), nil
}

func (l tokenQuery) ToMainCoin(coins ...sdk.Coin) (dstCoins sdk.DecCoins, err error) {
	for _, coin := range coins {
		token, err := l.QueryToken(coin.Denom)
		if err != nil {
			return dstCoins, sdkerrors.Wrapf(ErrToMintCoin,err.Error())
		}

		mainCoin, err := token.GetCoinType().ConvertToMainCoin(coin)
		if err != nil {
			return dstCoins, sdkerrors.Wrapf(ErrToMintCoin,err.Error())
		}
		dstCoins = append(dstCoins, mainCoin)
	}
	return dstCoins.Sort(), nil
}

func (l tokenQuery) prefixKey(symbol string) string {
	return fmt.Sprintf("token:%s", symbol)
}
