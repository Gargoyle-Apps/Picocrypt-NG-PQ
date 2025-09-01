//go:build pq

package main

import (
    "github.com/cloudflare/circl/kem"
    "github.com/cloudflare/circl/kem/mlkem"
)

func pqAvailable() bool { return true }

func validatePQPublicKey(data []byte) error {
    _, err := mlkem.Scheme768.UnmarshalBinaryPublicKey(data)
    return err
}

func validatePQPrivateKey(data []byte) error {
    _, err := mlkem.Scheme768.UnmarshalBinaryPrivateKey(data)
    return err
}

func encapsulatePQ(pk []byte) ([]byte, []byte, error) {
    pub, err := mlkem.Scheme768.UnmarshalBinaryPublicKey(pk)
    if err != nil { return nil, nil, err }
    scheme := kem.ByName(mlkem.SchemeNameMLKEM768)
    ct, ss, err := scheme.Encapsulate(pub)
    return ct, ss, err
}

func decapsulatePQ(sk []byte, ct []byte) ([]byte, error) {
    priv, err := mlkem.Scheme768.UnmarshalBinaryPrivateKey(sk)
    if err != nil { return nil, err }
    scheme := kem.ByName(mlkem.SchemeNameMLKEM768)
    ss, err := scheme.Decapsulate(priv, ct)
    return ss, err
}


