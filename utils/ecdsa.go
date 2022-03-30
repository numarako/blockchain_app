package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	// ECDSAで使用した座標の値等
	R *big.Int
	S *big.Int
}

// Stringメソッドをオーバーライドすることで、Signatureをstringで返した時のフォーマットを定義(構造体の中身を結合して返す)
func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
