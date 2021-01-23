// Copyright 2020 FOSS GmbH. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package srt

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestUnwrap(t *testing.T) {
	c, err := newCrypto(16)
	if err != nil {
		t.Fatal("failed to create crypto context: ", err)
	}

	km := &cifKM{}

	salt, _ := hex.DecodeString("fe07286ed3b335728453a063917d5efe")
	wrap, _ := hex.DecodeString("e613ff213db50c951e5b07558bf367245757f554046b5c33")

	km.keyBasedEncryption = evenKeyEncrypted
	km.salt = salt
	km.wrap = wrap

	if err := c.UnmarshalKM(km, "foobarfoobar"); err != nil {
		t.Fatal("failed to unwrap: ", err)
	}
}

func TestDecode(t *testing.T) {
	c, err := newCrypto(16)
	if err != nil {
		t.Fatal("failed to create crypto context: ", err)
	}

	km := &cifKM{}

	salt, _ := hex.DecodeString("b56025b946130d018581bef8fe3d6dc7")
	wrap, _ := hex.DecodeString("1cf74f15f9d4318377f2ad6e6bf44f3dfa80e5527fa8319b")

	km.keyBasedEncryption = evenKeyEncrypted
	km.salt = salt
	km.wrap = wrap

	if err := c.UnmarshalKM(km, "foobarfoobar"); err != nil {
		t.Fatal("failed to unwrap: ", err)
	}

	packetSequenceNumber := uint32(0x79ee189e)
	data, _ := hex.DecodeString("ecc330219818e4b626aaaeb4e61c6e27e0cef3ceab08505012b5d31ff414bed4a1b646e53da0a68b0cfb11d2ae06576d4b600b31c1257b5a1787757caf25f7e98c459e70ac55a21d8b5d9b5a9d7230ba3a272817e017acd5577808828990728ae6bb7b694a3d154ac5a37ffa430af02fef0f4706c728fa36e8593cfb3ec85dba8af01e23b3dab8ad8e949d63de8a8851054eccf42f30990eed9f1c988f66c9e2bfab1deae231a21ed068664dfa31507e3abfc376da003503613835e8ec534c1eac24d01f6f3168880732554ae2f5daaa9bdb2aee49e6ac07744b0e12a435e08e00c94cec3c968141a18190e882dee01d043c434eef721f20fbd5b13363515a829fada86a49356c191a66cdad6b5b2211eb90df8d6cf57bcdf6e8c9290e6b8dc05153bc6bfe06715828996846d4ba4397ef4faf39bf17f757f1387dc3736b1bffe69e6e534f0bdf7537e347c42aaf8f295de6d1df531bf4ad15f9466af8f8aec94bba41e8409506777cbbd13652193769bb21ba6d7abc96dbe66b087f1fc734c9be6cee4689e567826bf609f59826acd3d677616ef6c608bc6e8382b3667b800a00adeb5c0c908ff8867f6de7b0ba45b31ade65da6253a188e7256cc1e3ef469f430b14380621e82487e8c3e01d4c373a742ca88abbfedb933dea30dc895823be4aca758ded22ecbbda9e24563e75e0afc99d8a9b4809e267c2b71e7b840c6ee4c422218012251e27f662c194df842b6e72276f2c3a14f2a07b852e1c6f833b767ded589e50ae1717f28dec41e475bbf717c7399571a3df6d6cb5836ca2cd3a8f624a1be0d226535a4c7b16dec0c0888d4d2ef15a6820dced72dbcb5319b3d8a80c7bd31afec93819adab33fd4d11bbdf9a8009aad39a6c125102689c80334c8ff215a41de69ef77a43fd9854600f237988c3fb69ab1233c84afa6417fe7363b9024b69ce49a618209dfa4d19aa8e8cfbebc65e78c71c888013d04defc9af0a44eeb0901bbc2ab756a8448daf5b7fb3d4d868c87e359ce24a14545a7933abe85d5acc11f643948ef581565796ce020f94492b4b618c7e14cfb6e9dff6e17cb8f1d74cc213bc03508a33fd91cf13afba37847ef5c9a17145960557d07b7ba3ba1f10abca3f9b3a7bbc9e710722f1dcc87658507d204146a55ed422d205d9486123d0ad5ed0bc199764c9c3207a33f43e47e37f980c01341b38b7d6a95ee2fa2439d0d6553e189f298cc2c63748a2e94fa853de77b41ee62af265f1a69c1a4332bef252c4377baf52e4528501af9a273639dc41b8c8049df5272c3461a4410ded64b287822e72b37c98f180e8fe24e12400ce332947a6bfeec178507e89a891b220a723a79e43c763af785ae7741a042efdf86d22dc57f3f4b2fe7e992c44aaf36afd57cc9751b7108ae6855e63e92e19db2fad00948ffe52463256710dfb69d8883326a498d0cefa3af7d81596885919e675fe30ee4883d8b2d80323d75c28fe19c7b7682e6a93202ae83eec75c67cde71f899bd69d90155cc67242897b7b3f09bd3f83a773c435bc2656b5e959137565188261a1905c04ca1cbd31d47a1d3dab038de29b9aa15afdfa8f155952ec56495487ab7db81c561418add328fe2b457b5a66d19d2c510caab1e3dda8a37a8ac7066637560373997350a64de0a20805432af6aa81daa4603c0141074f4b75e8a4091023aacffe20dbba66587869af7e8cc3e6649ff031e64e5b2943446d554d0084ba26709af9b9bb8af81cf00c1d5df46cebae5e8c17a90a8606500b828e79201ef05cb58aeb94592bb00b875528a82abb0497a7b1c0afc77cf8de141cc3512e822c002658cbb170e65f54e8b6bada79286d64a34")
	decrypted, _ := hex.DecodeString("47410011000001e00000808005211651aa410000000109f000000001419a2088dc6f8003fd0320993a9d4f33dd3a34628addfddfeffd03bf8f8677d149a74e8b1e56fd8208df79622c98c2edd137b6ff2331378c2c1299a2643d4048dd5e7062159a71fceacc10e61138f1d4e8788051c2f0f000fd1213d43ad36bd05682297ec6071bbfa83864a3436e813118a71e44797f3e12c22fc9f77e03455d5bc1405a001cae497adbff3f36004d122e5b77bfba7d391a6a830ef2235cf40747010012fe062c21e108d840187080300f6197be0c04ecb77bc455f686e23867fc2f802fe36677abf008cc43760db57d600afafd3fdad2f523ce844cf7ba3668f3e2ffff085b9c20baf084b810031e0400e7a094d9f114806e003ef32a9ef7684ddbef9b0007277e9524458b687c0801ae379e360c06f9c73130e90da9fff0201e3e81f7f5f0f35ff81016b8435c12e11d52f8342465d3a0de082f87381e0ccc83c1999040918a100908fc0e0428897c0e010e225e18060083010cc847010013cab57e091b4ff7f20697809cc00441b4ec99a00906ee73f795577dd2fef700111bbb5b8002f73beeb9d8a5debd5dd821020066420423209d2222ad5f8499e05ac0a6736b460c5d3c38312fb98cbcbc3e08023f0ffeabc09c285104a5cf89d002fe006c4dfa933a12aeabe60e0016689962dabb48ffefc3011e19c3f8200c7003b226030ae2036d41f4d00c23c8ad501344ee4ad41c4cd808210406da830301a7011c007977125ea5e4b3ef78783881e52841bfaf80065a6b47010014e72d60801017eb5c10020573e36020a4947bfd7c5d66fa85882e06ac5602bb987c92f0446c1c21883013b98ea1db881fa00bd7a5326b80b68c14c43856cafff000cebf7bc7ee3f000f38a4ccbb0ab620f78630085f575d4c85907f20276aa3df30835a0fc1673d74bba557609fabafdefe1f82dac64556e4155409b77275bfbff072105b9c100478679de1c82102070400c0bddf0400c382085eb5ae0084c88020e4915a80f4cc1c0026999a12652ab37e11863f00800e264701001546021dc486dab21021d062c81e8fff8c9461b27a997aa6d9a88187003e3aacd93a67bdeb5f249fc38a002d3607c66dd9b7d7c200132200133c811f612017e2004ce2004cf030001898220002010870e0104112c300192418a81f80885b482ed470a4907da0380108225f60a44452123200491a1b9486026ed55e484c42180bbb5e1983d25dfe03b5f7ffffff6830fbf17b3d9b33030f3f9fcef0cc200c3821020028e6c59b05e62e080180b91517eef77c210af0038ed823470100161881b6af0045260ce232b57f7104d0ca0007cd483569fe4cbc44b20992fff750c300186bebfb0677fb4fdfabffc37e0074a6606339e1b6a0972434060014f49331dcc8ee301df808bfd3e77fbbdeb9eb0f0c347f043cefe0c782085604e5a6825cfc0554d4cbe411286bb3f2cc06ab7fdfbf813ec7e45fe648fb8a216b6a26bfdf0e7f3f9de17820061c10030e1085f0042d20215acad5eef733834006419dd5def7e0032daeeb5deff867e0085a6085636b57e6418fdd354701001767fbafd3063f3a74cff7fb2c188406c5d5555545d45c998b8e32e3b1de39355fe10866eef5f849efbd00fc210d600a740661595abe01244c1800622d32398baa43c11fe769cf0843540c075b7dd0170089bebe7ff80073446b69b22df7f08433a0c07a4bba03ad7ef5f88ffe05d9a71348ce7dff9d73ac290400c3820060b5e0c017f03820416cb06009597c3189040f89003613152b3b7da0033850ab2322fb41c02f7353c002a9089cc4111f683c210ce008a6c188c5d6")

	if err = c.EncryptOrDecryptPayload(data, evenKeyEncrypted, packetSequenceNumber); err != nil {
		t.Fatal("failed to decode: ", err)
	}

	if x := bytes.Compare(data, decrypted); x != 0 {
		t.Fatal("unexpected decrypted data")
	}
}
