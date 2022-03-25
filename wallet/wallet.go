package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// walletは鍵を持つ
type Wallet struct {
	// それぞれは構造体(*)
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	w := new(Wallet)
	// 第一引数に指定のアルゴリズム、第二引数に指定のランダム関数を使用
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	// privateKeyのstructにはpublickKeyとD(プライベートキー)いう要素が存在
	w.publicKey = &w.privateKey.PublicKey
	return w
}

// signatureはtransactionをPrivateKeyで暗号化することで取得

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	// 構造体全てを返す
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	// 人間が読みやすい文字列(16進数)に変換
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	// 構造体全てを返す
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	// publicKeyのstructにはx,y(パブリックキー)という要素が存在
	// 人間が読みやすい文字列(16進数)に変換
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
