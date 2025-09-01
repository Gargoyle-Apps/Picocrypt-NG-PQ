//go:build !pq

package main

import "errors"

func pqAvailable() bool { return false }

func validatePQPublicKey(data []byte) error { return errors.New("PQ not available in this build") }

func validatePQPrivateKey(data []byte) error { return errors.New("PQ not available in this build") }

func encapsulatePQ(pk []byte) ([]byte, []byte, error) { return nil, nil, errors.New("PQ not available in this build") }

func decapsulatePQ(sk []byte, ct []byte) ([]byte, error) { return nil, errors.New("PQ not available in this build") }


