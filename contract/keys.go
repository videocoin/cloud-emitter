package contract

import (
	"math/rand"
	"time"
)

var managerKSPool, validatorKSPool RoundRobin

type KSItem struct {
	Key    string
	Secret string
}

var managerKS = []*KSItem{
	{
		Key:    "{\"address\": \"b221a44420c2b286f0deb37119b19af6d8104ca4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"66faacc55803f4c8f6fe09fb1841cce0\"}, \"ciphertext\": \"b31e39140abae2bc9ba0ea88dc81b042d92e9b183a63ded90edbd6fb2d3ef4f3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5d279fd8967a9466d32363bfec6471d9\"}, \"mac\": \"2b9d5ac49340f7c081bf5bddeb6f0310ff6b132a8119071336da47eec0b969bd\"}, \"id\": \"6c33bfeb-f598-4e02-a52c-5eaafb315bf1\", \"version\": 3}",
		Secret: "p9zy5raygm6tahstl8",
	},
	{
		Key:    "{\"address\": \"5c940245fb5565eee4f9d6ac85ec8072bd08aa8b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"962be18f0f771de313cb1421b2b770bc\"}, \"ciphertext\": \"61382e8d1d3b4196dc9246546e6fdf5f0207a054ded931b38e08611eff8c874e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"79f649fd8a963040e829339a9f20573f\"}, \"mac\": \"a7bafd22e0c2caacdebd022dd751a6ddcba324e9abb557e93571344be5d57825\"}, \"id\": \"36299d5c-a282-4e77-9df0-ab4560ed585b\", \"version\": 3}",
		Secret: "eg1wgg3l9jmojtha38",
	},
	{
		Key:    "{\"address\": \"89309476f86ccfbc5c3b163d6fb672295392c518\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"56a1b5821eef711805ecf9137e5e6bcb\"}, \"ciphertext\": \"64e879b8d5e4276e01da1c8fcb894a08c00b8cb3cef1f0440074b9f4133cbc67\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"dd992614325fe1ead0bb73a27bc7e83a\"}, \"mac\": \"39d3c97cdbe04b442e43b75d39a847d85f43d8b8913b39db332ab6f2d814a074\"}, \"id\": \"e018a93d-7c86-4db5-9b12-b4a25bda0cfa\", \"version\": 3}",
		Secret: "v1l8rs72gbykr72xxo",
	},
	{
		Key:    "{\"address\": \"8c635d8bf0be948de71c02f14f2c26170d50327a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8db6a302bd0f9b2114f33434ef837d3a\"}, \"ciphertext\": \"6c3f18bff8f588db018a911ffb8b05fab4351fd85b90609f3d72fb0f807923ed\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"da5cbdef847ec26f96eeafa2c651e2b2\"}, \"mac\": \"47d2a35ba45bef75e6714924b5a412e4882439c7f7cfaaf55457631fce27234a\"}, \"id\": \"a1c22561-3395-477b-853b-aa974df84975\", \"version\": 3}",
		Secret: "cfd8h8suu9tu1t4pdn",
	},
	{
		Key:    "{\"address\": \"f4f32afb40ff1ad7e317e8a635f204a4dfc4a4c5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"110288792b6a16383972be7d7798573f\"}, \"ciphertext\": \"91b7bf63e1dd15da8c8b2d05f6a2f8cd2062b382eea804ed20b9ec7ecccb3e92\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"73765ef0ff4986f6e4b4279ae12d5e39\"}, \"mac\": \"79d0da26b2148f567f062053677f99b670a7c5235ba6629d595069092e1e1691\"}, \"id\": \"4bfc0eee-c87d-4c93-996a-00d50557b993\", \"version\": 3}",
		Secret: "curacbmxvo18caxhu9",
	},
	{
		Key:    "{\"address\": \"963b0d6d493770c42d03c31825468f092d99f389\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"59b1ed195b752b7f14d6a99f9834e156\"}, \"ciphertext\": \"91ac6d5eb5721d46dbd14fd5f63c0488fbbd9bb0f8afa47f6827a5c69ab37599\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"42c2284833df95fbc5d7e4ca6aef7417\"}, \"mac\": \"8d0120d4a6b19b2625e26a88bde5bf353f9c86096af215f4baed1d5a7df037f0\"}, \"id\": \"e6eeda38-5939-446e-8783-cd2b3b2bd1fe\", \"version\": 3}",
		Secret: "yar6i9nrwevowqxgkt",
	},
	{
		Key:    "{\"address\": \"60317a60681098471cb699a7aeacc5f73790aa3c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dcdd4dc03c64fcc6a8e113d518e57c98\"}, \"ciphertext\": \"f5d92825e6cb35cc2eb2f6a206f96d32dc08634b98dcd53e8a3b4ab82e03cd30\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e218d7992c15e52281817ee45a38238f\"}, \"mac\": \"a89cf2c9f8d64bbd765bf0a808f5c190cf7886b21d448cd9c64d33c7801f9e72\"}, \"id\": \"ab443ec4-4030-4e49-8324-02bf9c2c83c5\", \"version\": 3}",
		Secret: "t5r112k22vurnprcs0",
	},
	{
		Key:    "{\"address\": \"5b8f36119f2a607e90074bc0e7d10fced12d5e7b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"acf9558daf4de49b6b97e7a32ae4f89a\"}, \"ciphertext\": \"cf99bc393a3024e0b8015e86a98c019d15b99ec876000962db25eedbd66626cd\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a3e2d7cc52b6dcdc12a8676fa96dc044\"}, \"mac\": \"ebc541915f6e55732992ed33d4c3679be44b0e89f47c6b695f0bccd0c8cb165f\"}, \"id\": \"0792cc28-95f0-408c-8eec-d9f23867b737\", \"version\": 3}",
		Secret: "ciibohxanjq40xtuak",
	},
	{
		Key:    "{\"address\": \"165cbcbcdd47bb05c521d989a93c5156edc73e4c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2af20cffae6a38a47ed6779d75f9609c\"}, \"ciphertext\": \"eb1f20f3714cdbb3088b30b16048028fba96afba34a5e8b9130c834617b21add\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4796dd37bde8cac52a7a44393e4eb2c1\"}, \"mac\": \"d1eccbb20cfe46db402a55e5af70931ffe1c6ec207421da3b6348cc267a1d9bf\"}, \"id\": \"afc970d1-8d72-497a-a9ff-e96626095416\", \"version\": 3}",
		Secret: "hh1y7n4j34xtd94vp9",
	},
	{
		Key:    "{\"address\": \"37ab9237bdfef1bae2c7325864f9ff947d36b2e6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4aab5e74ba92cd79c3307a926cbffbb2\"}, \"ciphertext\": \"d657090e07585be7fc83b754bb994117cf67bf28ad896a367cd9a4ea74edc3fc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"24662d5aa601d52e64a0d361b9afa1e6\"}, \"mac\": \"dae009c396989edcfb53d44c34d02f8c04a7f6e02381b426948a06759200e8c7\"}, \"id\": \"73564d5e-ed3e-4f58-958f-2cd73870f51d\", \"version\": 3}",
		Secret: "epay3arbyp25gqbwjz",
	},
	{
		Key:    "{\"address\": \"861eb30af4922c7397c88a3f47d759fba9db738e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dc42c0f8e5bc672dd9e6b3832b66f7a4\"}, \"ciphertext\": \"a94c8ecee23ef382f89717a623e5fa2d5651a81dd9ba5c49da381d77e5fcd8b5\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4bb6db9e8deb678f17565945f8ea6c62\"}, \"mac\": \"1ec738932818962f8926167acdb93a3be2be92e469e518e5904c40f9c2fb0c0c\"}, \"id\": \"9982f347-2f87-4c6c-90ad-79bc212ed308\", \"version\": 3}",
		Secret: "1wdqga5ledv5dva8vv",
	},
	{
		Key:    "{\"address\": \"597598a2051f4353c4ffb2d500e06b50fdf3a828\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4319809b13298659994fc76974f6711d\"}, \"ciphertext\": \"8c7acb52d9ca95ecf3bbef70b0e428ff767f89da4936cee86e26301d52c88997\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bf6ecbdec83ca984f8026b87b65cbfbf\"}, \"mac\": \"cb58b156e57e9d9648800838349ff09410728da8872186697b78c8ed47ededee\"}, \"id\": \"aeb55002-91b4-4550-99ff-837174492bdc\", \"version\": 3}",
		Secret: "7onipx0imljfvf8cr3",
	},
	{
		Key:    "{\"address\": \"8c6da48e78f33bf1cc80ce26a42dae262d63c91c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3a42d69b400c6a196a9ee1fb4363c6ac\"}, \"ciphertext\": \"d286a98702503c17ae8cc99a79cfe5e8b66367f134c2c5bfacf73e3cd35ae047\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e4ec8ccbf8044d67f12c71f96ea9a0a3\"}, \"mac\": \"895a2655b78d20c4e14fb6a38b12762c706df0311a70e1ee7f4ff9deb01539c3\"}, \"id\": \"68c78af2-0aa2-480d-bd85-46b9ca59187f\", \"version\": 3}",
		Secret: "o2zqkvbcmzldpxfxna",
	},
	{
		Key:    "{\"address\": \"3f5b2e69d573ff21190d649b2aa84f6ef608f98d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"283621436893ff42ba6be6f1cdc87d92\"}, \"ciphertext\": \"706a71ec16e79c19e3d0182302aa02e6bb2c9829434fb84e7397342a5d1af8ae\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"81b61999d2bbae08839a01b4118487f2\"}, \"mac\": \"c61896cb9e83136eabf90b82ba3cc75dc2e36dd8cef0cac089a853fdc62cd386\"}, \"id\": \"fd2126f7-a1fa-4041-9110-3bc3c511b736\", \"version\": 3}",
		Secret: "1jp8iux3xm5p53kuif",
	},
	{
		Key:    "{\"address\": \"8d9f120704682f7c962bea9e557118934964c16d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5a57e0179c3c21a2a9528833f862be31\"}, \"ciphertext\": \"3ac20d6f213c7727b5072d0d055c1eb106792486fd9bd938eea99ed19205d89f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c416761c3052e3e0f0704d3007d7b3c9\"}, \"mac\": \"caf1d5a9810ad0b10da5ed7d209429fd6f2a04473f404a0d8a9c9ce8c1006c8c\"}, \"id\": \"68170ef0-178d-4018-b6ee-7eeb75a6de4e\", \"version\": 3}",
		Secret: "bjsoma3mg4utypss0y",
	},
	{
		Key:    "{\"address\": \"e3687eadbc6071e1f9f22785f7f949b060df6e89\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"17888296afab6603ee78c2ecec0ec182\"}, \"ciphertext\": \"0cf0d1db3207deec65be1ddccca5a2d9292a6d751466de7ace8860b586ee8bcc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"80d6aa2dca6548ad4d9394fc5893f5ec\"}, \"mac\": \"656313ba3bad6c92f52e414bab3e1b3a8bcb7a1cd461f00545ce8f36aaa4870f\"}, \"id\": \"9c19bf66-bb16-4c4f-a48b-eab48c325eb0\", \"version\": 3}",
		Secret: "he2brahdst6bzbmhzs",
	},
	{
		Key:    "{\"address\": \"71bb6dde13a45cf54b59114018cd80da95fc6f67\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ebba8017c7663060d5aeba9a7929ff81\"}, \"ciphertext\": \"8ab5dbedc7d6ebe9e85f37116e6d090e081d280d615acd762b25d8e50f94a562\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"fe3712bbe99aa28604435ef0e41869bf\"}, \"mac\": \"3b1ae89b4668f1644d87747af8add8ded47305095c83821a859f9fc2cfc7cd62\"}, \"id\": \"8dced6cd-4303-4f9d-841b-a93293af577e\", \"version\": 3}",
		Secret: "alppsdy3dhvoggf316",
	},
	{
		Key:    "{\"address\": \"03da6764455fc2fa0cc0b84cb847e417b60f84e6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f9d26b3d2f1c80e1cd6d692c90589e09\"}, \"ciphertext\": \"04cf075ded32d9af17bb0f0f57deeb1053b6b7439a69c45b80d402290d75c1b1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a79bc417e483dbf6dda7b5b86bef6e1e\"}, \"mac\": \"99c8e212efce2c1978bf1fac44868ad72526133d97c38d9f53aa17037d464774\"}, \"id\": \"32072559-27b6-4004-b5bb-da72b5453367\", \"version\": 3}",
		Secret: "7g2bbq16rzc9t0d9t7",
	},
	{
		Key:    "{\"address\": \"35dfba840ea636163c14efec28ad114454de4ab0\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ebe8278de1d725fb45db559d7fb03e05\"}, \"ciphertext\": \"d0455ba4f1b81f570da1e484b182a3b98181f28b752a5217c374b882434ed7c5\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"32b1498dc6f77c8bf3fca4629d79d863\"}, \"mac\": \"86700e52d9df65ae5efbcc5e927e01b635ca1a87ba567f5dde4fe590d85fb330\"}, \"id\": \"ed44ea9c-8217-4167-a717-bd363b7d7dc2\", \"version\": 3}",
		Secret: "nzz1rpxm0usy02ptev",
	},
	{
		Key:    "{\"address\": \"21f1d58fe768419177b8c38d89a45ddfaa6ac9b0\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"78ea7e9bc595d8bb2e3cb48b4a1fbcb9\"}, \"ciphertext\": \"b9b20188ef243d6ece712b8f067a64498fc1c13936d3b6a1820d41589f4bf4b0\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"06c1ca0d98e6bb477ec7eadaa66b6de1\"}, \"mac\": \"d40ff0f537628d0efbd5e90496ff2451886156a0b9735dcebb1b5a6294dc4c9e\"}, \"id\": \"19928d22-cb4c-41d2-8e59-faba0fb5a499\", \"version\": 3}",
		Secret: "5bzzrn6fd830jwaxjc",
	},
	{
		Key:    "{\"address\": \"34d66db4ff869d1df08575e33c79bd0ec6eea8bc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f372ef1e7d426b31cb52198c989e8cde\"}, \"ciphertext\": \"4c5376716649071f7fafaac9f1b3a72311c3ac28aace1ef084fa1a782b304beb\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b0282fa212a74fc85ee9916491498087\"}, \"mac\": \"4d1de14958c7f0932eddb6119e64edcccc9190f3aaa9ca4af370594b0860a2ab\"}, \"id\": \"cfa04613-91ca-464e-ba8b-0eacf6b9768f\", \"version\": 3}",
		Secret: "kqd13jgli2o9um5kif",
	},
	{
		Key:    "{\"address\": \"26549f684739c53a5820af761d933a32ac747d00\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4db0653b330f3ed25526ec855b315886\"}, \"ciphertext\": \"93f2c3d015b6b910b36f6e9a85d55a2e8ab6e4cb262844241dc75ae71b12847e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ae1bbff9831dbde29fc60ad3b8fc6267\"}, \"mac\": \"39e8e149cb38d5e69ff2883983af70899e96192542d0750262e57c584a3944b4\"}, \"id\": \"08f1c59d-0076-449c-be90-4b3d7e6b70ea\", \"version\": 3}",
		Secret: "57kal25bwcxcyjgn7v",
	},
	{
		Key:    "{\"address\": \"c54d0573161b4b97537e7ae8d8490cef9915e448\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ab388edbe25bad4f10d0534c92270879\"}, \"ciphertext\": \"057e947b3f2a3699933e28cb3e12e59d9700726297e214d97d4ea46bd1717c39\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f098ba06dcd77d64bc8e514ec0ef346a\"}, \"mac\": \"632e52b7451f494544f37b6dcc44fc8ab9632ea802c5b430847cb927ec32289d\"}, \"id\": \"6176bce8-d3b2-42ae-a1c9-99cc1829e933\", \"version\": 3}",
		Secret: "zztq09662eo0wqi2yd",
	},
	{
		Key:    "{\"address\": \"0f4970235618016c2b4b0b95c9527b31eb19b4b8\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e52b518cdf35de24cffcb02252688c05\"}, \"ciphertext\": \"701b217e4c2bb04c7e17664756f6107b3e4f465a0ad9eba4345a508258f77649\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"801aeb9cbb3ecc09dda9507688ab2424\"}, \"mac\": \"836bd70ddad9a7ce23c1a674897c143ad5ab480b626f0d8b073ad0b1bbdd7596\"}, \"id\": \"3d1f7d82-fe4b-4165-8087-3e0295fa733f\", \"version\": 3}",
		Secret: "xdygnye4nhilduf8ac",
	},
	{
		Key:    "{\"address\": \"d20effe52d9e3332f53486baa3a1b365423c6fff\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"62a7ff9fde5a05600a7ae44f491e5799\"}, \"ciphertext\": \"ae469cb45d210ce1b9d03ab25362d5ccca280a7989ca41f2e8feb381056ed150\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"996a5525c8b3d51b536aa3c189cede80\"}, \"mac\": \"b6a61a321f3159d9eac181f7652eaa92aea2d2c0e5dfd540b242cb17ccc686d9\"}, \"id\": \"127e2331-da69-43f3-aeb3-9dcabefc11f4\", \"version\": 3}",
		Secret: "0dv6lm1uo5xxycw5a8",
	},
	{
		Key:    "{\"address\": \"96a4f0f3ee4fef4d09071c451d43e9913f67093a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8463118e09c9af48c32829c22160da68\"}, \"ciphertext\": \"3d1a321fd47f54ac94dc029e28165b3dad79fce15ecfa30c228167313a762322\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"db18bd82cb139301abf5541b8630c667\"}, \"mac\": \"d69a6fa170027df0119cafe6e29a27a546227ffe529ec10ad459af7c43d49444\"}, \"id\": \"f3028e16-c483-4b7a-9086-6630970dd6fa\", \"version\": 3}",
		Secret: "xe7vvgaheu6ojpgf96",
	},
	{
		Key:    "{\"address\": \"03777de141eb5c20b937bab764e1bedde7be3407\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"39518d6e97fa55b8313c28546bb5f29e\"}, \"ciphertext\": \"1d0c1e207c1a3e602f4fecec6ecfbcd6e28e2da330e8883d93926944bf484189\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5d1c17b496c420b9504c2fd26630217e\"}, \"mac\": \"229d571ab9549cbcdba95eb306bc878ee5eda8ff589dd3d6bb157812955f400b\"}, \"id\": \"9438de91-4e68-4da4-9e8a-f85802627907\", \"version\": 3}",
		Secret: "btc7akmvofs0q9u46z",
	},
	{
		Key:    "{\"address\": \"f127067cadc2c41e38faba761227bd08af0215af\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"45d0708e69c4327a38e99a902481af7a\"}, \"ciphertext\": \"8ece28a86e4f167471898b8c006b644c73614f1dac2f9ea6194785bc9ca4e797\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1728dc740ef3407c0f62594976da892f\"}, \"mac\": \"25212c312c7796cbab3a6f30475e3213269c101350f1afd7a7f7dd5320e75a71\"}, \"id\": \"fbf0b2d1-b9d9-4350-a2b2-31e52bea4efd\", \"version\": 3}",
		Secret: "yrqlxytnnnljcwlkbi",
	},
	{
		Key:    "{\"address\": \"6cec1150f26fd484c305791411e32282cc8aa26b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"63e34e8a60548fc93df50e5733f349a1\"}, \"ciphertext\": \"a613d9ba28259e998867f77e553e45f611b0b4d4e4c5ce764968be8e72cd4541\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f78e375e9146b2ac7e43dcf04035de58\"}, \"mac\": \"98c658c8ef8a982acaf47bf90415c91e75528f8b845eeff40faef32ac8f9e317\"}, \"id\": \"6e59cfb3-ad55-4c29-b12d-3d2bd3cd43fb\", \"version\": 3}",
		Secret: "hbjd1vadlpgb8jpc91",
	},
	{
		Key:    "{\"address\": \"8b81cfcd4505a10dc0d423eabede18708ec019e0\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"acbc6be1c8ab5e91b77d7cd0b97ea7f7\"}, \"ciphertext\": \"e62d507ed2909c5760e6b7bd9346890d9039f0231002ca7c4479651e1362f435\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6fc1df391f49166f96105ed329a1d28b\"}, \"mac\": \"97bc19e88b16f73fdcc520a3ddf2d3b8b3c542cc793d803057453b4ae75df689\"}, \"id\": \"75e0b825-9fdc-474f-a504-ceb788882f6c\", \"version\": 3}",
		Secret: "o9xjo45xj4ex797xh2",
	},
	{
		Key:    "{\"address\": \"cd2b8f3172065d4f2d63217cd1b8be16049328b2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5fababbe69677e6cb339b693bc2e1370\"}, \"ciphertext\": \"bc89d4797fdb512238b70c17982c17b495d3cf2d2429fbd1752124a86ca876b0\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e21da4df631c62832f2fe44c9ee0d7da\"}, \"mac\": \"3fbdad714a27988645508db0a9dfb1ec379b29773a85e0810f548c0d839141d0\"}, \"id\": \"1f09a7d4-bb1d-4cc7-95b6-40fb5b147acc\", \"version\": 3}",
		Secret: "3ghrmrkcij2r2pvkfc",
	},
	{
		Key:    "{\"address\": \"6925a872bd44d922761016d0e492fe90f2fbf3db\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ddbf6a7cb48dda626e12dc9f5ce442b2\"}, \"ciphertext\": \"e725c31099842589195f495a051800b6738d16f649b9a64ef9ff73fd36b4ac88\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b11348ca037f43169158507f398a6dfe\"}, \"mac\": \"f295a81c41ef7dfaf91daf37ae3126635c841e8a751444d1b671b107b372064e\"}, \"id\": \"04000712-5f91-4a08-8ab7-5006e78e5205\", \"version\": 3}",
		Secret: "femlawwp2qtitb20xy",
	},
	{
		Key:    "{\"address\": \"23507af101b5c917d424108188c35eeef68b0e05\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"7590421945d08b2dc6b94689506541c4\"}, \"ciphertext\": \"ba955e9eca85a2c38225c7e53b7208fe21fe4a0e49b6735882e90ff4c9acd749\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ae797ec7cc3a0c26d463d098aa2b661c\"}, \"mac\": \"e9904333b2a92f9633572482012a76feff2e4b25213ac45f24081374a26ff9bf\"}, \"id\": \"de3d37c5-1a28-414b-8b18-ce3f41e8666b\", \"version\": 3}",
		Secret: "mk7q2p7pdxzaw8wb0h",
	},
	{
		Key:    "{\"address\": \"8c43641f934cd9d4e7f5357e43e7da617c6bb66e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"462ac0d68bb227dad848b2dd93428f55\"}, \"ciphertext\": \"33c864f6f3cbcf1d78f2957989d0abc408c7a815e7435ecac05ea32db3ebecf0\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2c1a07ebc97bdfa9fe94bf545b1638da\"}, \"mac\": \"7833c3c9bfeacf8fcaf9d7ae5e9a5a33e43a21ac3d3f3f84ef9488046ecc0693\"}, \"id\": \"f810627e-c782-4910-8670-c50106561875\", \"version\": 3}",
		Secret: "xez96nf3qdv4ckqswi",
	},
	{
		Key:    "{\"address\": \"1dff13866f94579b6813c4dfe538ba6d946e5103\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f78cab5e6faa8d25ed096bb4e74e4989\"}, \"ciphertext\": \"c3fff3bc7d52e0aafe24c71dd38d6c89eee46c0a38e2c48d9238e2abb2df2263\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"10940b56d4ae00a0930a74c0fc5a1f23\"}, \"mac\": \"9a848ab381e76e4dbd29e62d2e5a7b42138c98f184e128838a100aca3925163b\"}, \"id\": \"5875bbba-de5e-42c1-9cdf-e26642e399c0\", \"version\": 3}",
		Secret: "ngafpbg36sklzd5z20",
	},
	{
		Key:    "{\"address\": \"de64f48677d16dbdb2e45083626350b70641cfb3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f6733f6b7143df5e2d91418e943523b8\"}, \"ciphertext\": \"7c5a1c00263236ebd2590c0102330e3a64c608993b18739508ac2094c10eb8d1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2fc07a6f6cc4c9342e110688b7615944\"}, \"mac\": \"141b7974147f47579d090177fd35db3e15103c563d9ff8dedc31742aec35f055\"}, \"id\": \"089b8d5a-7484-4112-badd-abebf172caba\", \"version\": 3}",
		Secret: "idcicoi4nnluf7e1yu",
	},
	{
		Key:    "{\"address\": \"fc9f767b87b160cc99b2187faac6da9232600be5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f4f2574f52842cf0d1e65080dc98b7cf\"}, \"ciphertext\": \"bf0a516eb3c9cb5303716234ef4bdfef8cb853a75fd634792d84f34492f12e84\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f37419d3af0259bab39879a61c3c9085\"}, \"mac\": \"11e60f4e1a5430e24c2542df812cf402fb912808575b02e82d638b24bf9d7419\"}, \"id\": \"1f094558-891d-4d16-9d8a-6c938e71c382\", \"version\": 3}",
		Secret: "w17w3xge9wk9u4ijsc",
	},
	{
		Key:    "{\"address\": \"af43144e9c147973f0e9f06677cc5a13cefb2091\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5f154c87d5baf06afb4c57e05c83f548\"}, \"ciphertext\": \"d2c8a02eb4c803b343ed74367c0d27b63177c2e76193f25c7deaedae1683cf9d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f9b33710bf4d34d419d870724498529d\"}, \"mac\": \"ad9bb445d81669db566f8812bb8332a49224da5ede9acbca4e39ac66205b9b77\"}, \"id\": \"db210ecf-eee2-4db4-acbb-176bf39d416d\", \"version\": 3}",
		Secret: "zduzcls1z6zhw5kfyx",
	},
	{
		Key:    "{\"address\": \"e102abadd15957de92e69034fd712d9ae04f92ce\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"79e6807347b64676409eb501afb76025\"}, \"ciphertext\": \"b35e007ffa2afc4d33fded23948ca36723835765441b08b9272d7c24ba2c5a81\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4aad157f828906c881614fad96337edd\"}, \"mac\": \"7229839215328df7529576f6a5c4eaeacfe94158740630d373bf3808b411785b\"}, \"id\": \"1a04e664-3367-4342-a99d-a84c2ce423b3\", \"version\": 3}",
		Secret: "jhj13j9tuk5nhsfxjl",
	},
	{
		Key:    "{\"address\": \"fd38658fa2ed4a2560305078e5b22a9110be5151\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ccabf93ed01f73185597067c4d1a161b\"}, \"ciphertext\": \"4fd9e49f91fe0bb3d36914ee9fb1188713e34f3f74bcc427ae5ff2dc7e6f6953\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a121bb8b572ea11b2a8f1340d7b9f631\"}, \"mac\": \"861aa98499ee347e87d48d4fe0ab62f17598fe8515ccd6da94190bf111c9894f\"}, \"id\": \"38730137-76d7-44d0-949f-2d96610e04a6\", \"version\": 3}",
		Secret: "qwk9kqtqez3qm1tbqi",
	},
	{
		Key:    "{\"address\": \"043cc0cf1fda40d6307cebf3bdf52ea977dfa622\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"20d04ac23c808a677d6e36a1cfc5f79d\"}, \"ciphertext\": \"cb472b5e0bdc19de0b1b5b637cf3aee0d0da96747fcffb78bc65933fca0fb629\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"531187d9029e9681aba1a5d8f0fceea0\"}, \"mac\": \"82cfbc6c71070a8746a54a8e0827426da768af0a969a0e31d176a4b200d00076\"}, \"id\": \"24629e3e-2da4-4aa0-9ad2-b7dc5d6eb9c0\", \"version\": 3}",
		Secret: "sglafegu2iozqmd4e0",
	},
	{
		Key:    "{\"address\": \"3f03fadaecdf51a12eaf67cc65d7fae07cc85c17\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"04fef262a197608ec3aa8799fb4cdcf5\"}, \"ciphertext\": \"ab82c7eb819ce83042d9fd367bf197f19b2339d823b4d2d527020ad8679e257b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bea1dae3072706d57fdd4f21c71c64cd\"}, \"mac\": \"7c7f605278edd0c7757ca9368f1164549a894158129241c057d9814fddac3d29\"}, \"id\": \"ee375583-c2be-44ef-be94-d91056b023d5\", \"version\": 3}",
		Secret: "c04m62nclbxtmo9crd",
	},
	{
		Key:    "{\"address\": \"180b46749d31c306a6d760e4a933b2bce1ebd9f2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9f79e66c466e52808457778a8514ee72\"}, \"ciphertext\": \"ca77cac60655f3d21b31fb3b98688d76fa44225c0b6ce3f7989f0b76ede9f80b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9b699d3a7d538c0f6506a524c879dc19\"}, \"mac\": \"f74ab817688a9a28bef4456bb3b6509096b279747e844228cb8157ffbb59adbf\"}, \"id\": \"498f0779-2d36-4f30-b778-74cbac5f07d9\", \"version\": 3}",
		Secret: "x7wnyj5psw7g989xbf",
	},
	{
		Key:    "{\"address\": \"aa4de58e58643a50f0be5bd36372c594e190c9f9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a03ea497b1d4f9ba06ea094d27a7aeee\"}, \"ciphertext\": \"9eecc1fe0677b4153977a2af020e066b5f5c2a38b3cd00d6ed19e56e711b1498\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1710e10ad9133bd64e1509d274e6a24a\"}, \"mac\": \"250ce983c220e5e4900f317fbf55440b002d80d9d7010c273254e830eff73640\"}, \"id\": \"0653bb65-826a-41b8-8bf9-a56e695ed7a4\", \"version\": 3}",
		Secret: "urlk8mgwvae9uax0td",
	},
	{
		Key:    "{\"address\": \"572b07e8a58808a596f19b5bd33dfb27835a6448\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8674303ee04c3437de538dcff19687f1\"}, \"ciphertext\": \"9db7daad5f1f84f726d40b8a884ef011d2bfabf76366b4d5ea5f8ca0dc7ff641\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5c0ec68354db94a3eda020a3d5e456c9\"}, \"mac\": \"6ea736dfe31fe5e7436f5810c8524adee5c2a9292a190b6f74f6b2bc2d22052f\"}, \"id\": \"a0437760-2123-4b62-816a-5fc6a582ce6e\", \"version\": 3}",
		Secret: "51u0ajq6fclkiev2dk",
	},
	{
		Key:    "{\"address\": \"b4214437903274ffc7eb843879b465716c5f3a72\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f9f670b49aefd70f1fd89e2ce438b454\"}, \"ciphertext\": \"2eb4759f9ed26f9665ad405df1bd37e9a82005af244d40c38d4bbe7bb92705b1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0cc25f00b99c39334560cb77f780e045\"}, \"mac\": \"2403e9ffeae90e197e0bfa12fe78b761fdf1d81e55fe3f212ca458bf558f7dee\"}, \"id\": \"dbe21e27-ad48-4ab9-9f24-22cdaf4fdcff\", \"version\": 3}",
		Secret: "s7bo4fs9q9aa1odjpb",
	},
	{
		Key:    "{\"address\": \"8f24f2900834d765fefe1dbc9c898b0d070aadb2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"02cb0c1628ae778342a39d51b7571810\"}, \"ciphertext\": \"bb5e3d04bd526bb9b20942e11783ab4a3b8691af017e50a310fcd50fc6672e35\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bb1370c2a09bdfbfda4962d09780fc3e\"}, \"mac\": \"7abd5c4c9c4b1a7757356e4f41af7834cb123069eae8dae580f148e65b912005\"}, \"id\": \"e6988ff1-77d1-4981-8509-83b9df65f1d2\", \"version\": 3}",
		Secret: "1bek8ldad29ari8f6n",
	},
	{
		Key:    "{\"address\": \"d22e8951ad4566d72a4c487fd09316ca0c5397cd\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"bceccca7e334fa2ce42bfa20ab3e6505\"}, \"ciphertext\": \"69186f000debf9befa667b5b738a954b95672a1e24aedf6bf96cce08e808d008\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c01d0dc9ca49f62376ba5187a02260db\"}, \"mac\": \"ebe97112c8598f06823189b435fcacb53f6ea9462427ddb481ded6858217b7fb\"}, \"id\": \"76ca892b-8bed-4532-81bc-108442c075ab\", \"version\": 3}",
		Secret: "jmtgcfq2p6j3h9vf4g",
	},
	{
		Key:    "{\"address\": \"d6304ea8f0f7189209e35225ddd55581f8f1b9de\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6ebe908f1bf1a4fe76a5e2beebf9ea35\"}, \"ciphertext\": \"dc71174309cef4b910904035b5a4d20199ea8af09ee701f93323dab6dee347c2\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7265c8d2828163efa615376ea4fa2296\"}, \"mac\": \"2f10f45a17a556c75ddf56f971e9921fef4bfa3464767202104f5eb5af15a4f5\"}, \"id\": \"9975a81f-9c6e-4e7c-887f-28c2ca6641a5\", \"version\": 3}",
		Secret: "3ve0vefvu4zb9ibsra",
	},
	{
		Key:    "{\"address\": \"fa65504f92bb0b2e87f7905c504a08a1b5502e44\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"38004ae56b47b3236be5f81216285736\"}, \"ciphertext\": \"d36fe18b8daf22add463687a9dd3283027792656347192dd1d3cf14e2d948b81\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"34d99df6bb1c2a9c207c046c46e6ae60\"}, \"mac\": \"2a7fb1258e99924c6fb6e8b4c39ae4da75a6c11414dd30ad4cbf732a7cc4ba57\"}, \"id\": \"08689a73-2b39-4f05-b841-61e17ca1df0f\", \"version\": 3}",
		Secret: "pqijpozz6kxusuo6aj",
	},
}

var validatorKS = []*KSItem{
	{
		Key:    "{\"address\": \"4f8b98d55defca521a9ea3b6aac9aad667e4fca9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8aceff858b3649ca9e75661dcc68cadb\"}, \"ciphertext\": \"4f12882e2300d46061dbc4f81accf6b05ad05a6863fdf0f7e0452c186878f3b6\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"da6b11fee1cac882642547dd445b03b1\"}, \"mac\": \"7f336228d1983ef21356f0c7caf6e0851d920a4fad30d3944918ed0db6d3f253\"}, \"id\": \"351944f2-0933-4dfa-93db-03d1e8f2d388\", \"version\": 3}",
		Secret: "2ursnwjrd6mfkvbcxx",
	},
	{
		Key:    "{\"address\": \"2138b6bef719949fd6d051a96f353eca1d41a119\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"562c8cdc43dae67672bece0a638f5b65\"}, \"ciphertext\": \"9aad4b0185416c3d5011b27ca3c8e3eca0fad24071378bf96a61d5a7d94c04b4\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"34a858964b4b97be1315365a0c4c5d11\"}, \"mac\": \"84b6a5a6e584780eaef70352758c039083d7fa2b1c9a9066c2b07f0584b8bf9f\"}, \"id\": \"058b61c9-7187-4bda-89eb-729c231f4b70\", \"version\": 3}",
		Secret: "qc6hv7xthw82n4a0wc",
	},
	{
		Key:    "{\"address\": \"d18edaf6f00f2ed69890b592d626981b825818fb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"75e62b888fa7440ed54e4916649d930f\"}, \"ciphertext\": \"1d17f8450af924d4cf486f37c07e1ca6e3a7f57b45b524b781e7ad2e7b220d16\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"76a03ef3f91823114714dc56d0cea00d\"}, \"mac\": \"dff259b442b1e70e11b9b53f86fdc711a8100430e411b472a89b2ac531e32fe2\"}, \"id\": \"2e5818dd-b0ea-485b-b7b3-8ab8bc43d6d4\", \"version\": 3}",
		Secret: "wu2hhxxoog9022uc4d",
	},
	{
		Key:    "{\"address\": \"048e865b699f9ac481a792e8326a7dcb0141222c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"647bd59fa5bc1ade77aa175cd897aafa\"}, \"ciphertext\": \"42a28dede5036dc4a3aa6d5845853b57c6e27165ded9b767d98a4036c4662a23\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ef4c4e9078a0c03249e3986993d6715f\"}, \"mac\": \"d36efcfb30cadd63e634ae11454201394b742747bb0af5103454b1afd661c257\"}, \"id\": \"5c761ea5-872a-4afa-9525-2b86a8d9db11\", \"version\": 3}",
		Secret: "q0c2yhva6lsaieb8pu",
	},
	{
		Key:    "{\"address\": \"354ceec3d0dddc680bd63ce18ddf44dd1289c812\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"cf6f9e932bf6b2906e94dcd51d879875\"}, \"ciphertext\": \"7695f1ba4a780647fc84bfd45f63684e761c971ed434ddaff2a68d75cc6def2e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4725bd576eb6cdc7d5cc5d195035555a\"}, \"mac\": \"15fa93555e1ae54fd4a7edd7b664ab15ff8fc1a1224a241a7ebe58438a99038a\"}, \"id\": \"a37b05b4-1e9f-45b4-bba5-cff66890c55e\", \"version\": 3}",
		Secret: "u0lcub4lkpnl8qx45j",
	},
	{
		Key:    "{\"address\": \"9c34c22e29bde211037a46513f3c6367cb2f9293\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c3fd6169fd7316a1ac39af33aeba7b79\"}, \"ciphertext\": \"041491f92b7f5540f6c044df56f80d526132357e99241a4b78e2e5a056367dfe\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"cadee3e70371eceed99e7dc35feeb66a\"}, \"mac\": \"567923ef5d32435b09cfc6419085df7fd9648a3884e37f4116c7772836ed4254\"}, \"id\": \"05ef1b42-4916-4f69-b21b-bf1725b0741f\", \"version\": 3}",
		Secret: "e8gvxqb2z60f8usa40",
	},
	{
		Key:    "{\"address\": \"f01f16fb122289c7dbaa35f1860754c01e8cf989\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f59c48fb6be07499ba7f05e8ff66f62b\"}, \"ciphertext\": \"13d920c3b64d4f46fd5c05160b19f8271cf0891dc6ead2937f2f7483bb59e995\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f04c065c6de655035df1ce1468a44262\"}, \"mac\": \"56e9930dbcb7c1d19b6ca87004c998582127904cc0750bd865ea7f7fc59c148f\"}, \"id\": \"38a6e60a-d1ef-4c06-9ae2-c5249459e8c4\", \"version\": 3}",
		Secret: "zxw4z4dvh87pr9l1vu",
	},
	{
		Key:    "{\"address\": \"86d16dff9a70b21f824667de6cf56c017a15428d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"90e4f7b1898836380a20ea89f904330c\"}, \"ciphertext\": \"8789fc98750c9b09d33123bd5c4d9ba615d908a04f3e6e1da37bda82a167768c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f431b6fb44e161bfeb4e9f0fe1c7f19e\"}, \"mac\": \"80f69105a9f5288b00010a5425d3bbaa4c8b80b06c816af33f1956955612816b\"}, \"id\": \"daa291ac-e0fd-419a-9f87-60b59d04e6e8\", \"version\": 3}",
		Secret: "2w55klohris3gi0tkm",
	},
	{
		Key:    "{\"address\": \"1298efcc15bb7088066c429232c7f321bfbffa4f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4ffa3028d6ef0b69eb31159f0c090994\"}, \"ciphertext\": \"1c9a247ea0ccb47553c121cc6a5fdcc98ee3962248400e7d30ad0452b6cec21b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bc4d2c03bf6dce338d0477f268c49cf7\"}, \"mac\": \"cad97c309ea87f3ebc18e76829a5f82fbada1cc863e093d8bc2f90b9536ef930\"}, \"id\": \"398ad8dc-799d-4b13-aede-1f77b9269c4d\", \"version\": 3}",
		Secret: "k7smkn57aeniui3be7",
	},
	{
		Key:    "{\"address\": \"10d45d4014343062ba3504ed497f03b5490c4941\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"06de0a75eba73c87787a56d69994e759\"}, \"ciphertext\": \"b81d6c9f1ccff9e125fe34110e8cea52608bd934163dbaca7aec713531aff35b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1f1e91dbeebd711500ba69ef3a256fc3\"}, \"mac\": \"005541a14a13b770a771e52a9e1cdb63de0b6f608a65adcb9a78e49bfbbedee2\"}, \"id\": \"05dbdc51-03e3-4a4a-b5d9-4496b07a4861\", \"version\": 3}",
		Secret: "0fhjxtxs35jwbyar17",
	},
	{
		Key:    "{\"address\": \"1d0f6290c591a5b2328fc53bd1814adfa81371a3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8300499baec98cadf16fff1c6f38c5ac\"}, \"ciphertext\": \"dbb72bca854572a641c29a6470f43b655df1d745800aafc9b945ff3ad743dd56\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d2633edc947b272308400c96714c1256\"}, \"mac\": \"e1bbfdead64c21acd44ed2058d2cc496285c0e1848098bf70a201e3ca75b47d1\"}, \"id\": \"9c7b2396-d8d8-4bac-afd9-306ba6975b16\", \"version\": 3}",
		Secret: "o81xi8i9ozweo8haqi",
	},
	{
		Key:    "{\"address\": \"2db4581fb9a8abb72d63f9609e642ff116db0e93\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b144d8d823c45bf40f588068b1cfe82b\"}, \"ciphertext\": \"37a9c3d99d85f5e0ecb549020bfd07a68c6086936d68abcd776a64a41e2c318a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d469ebbb1e36ad51bfcf3b6e86c018b7\"}, \"mac\": \"c98920056874633b095ee75d4e5a0d555a079c5cf32bb37de95884f596796572\"}, \"id\": \"7f81f1d6-432d-44f9-bc09-d58353c013da\", \"version\": 3}",
		Secret: "r594dp5wfx63ibujoo",
	},
	{
		Key:    "{\"address\": \"0bc40fcb4d5edd127ba4efa9cf6919e3fea4dcf9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3f47743071056f28d915916c2de7a3\"}, \"ciphertext\": \"43003e08993ec72bfc3e6f4ea1fb213daf7887e00bc519c8fbd80b5020b405f8\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a3c52d62eadfa679fb8d275c9e48f75d\"}, \"mac\": \"789591281a764dee0e0e37d28fd90cd6afb480052313248389b19d23808e37fb\"}, \"id\": \"cc2db42f-0277-4338-8e4a-b3f81232e560\", \"version\": 3}",
		Secret: "7nsi0azl5488af7usy",
	},
	{
		Key:    "{\"address\": \"df410306600c9219eddc13d9258002b820ec133a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f80fac115528237312928471378c8e2b\"}, \"ciphertext\": \"a3ed14222649fbac6085cda17831061038001839e9fc93e9188fb888c312c873\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3a82e2c0a8494a3629f5b008c048f17b\"}, \"mac\": \"5a3a1195ccca6133ef93cc22579791b072095cd4a86b81a49774b0eb80b6127e\"}, \"id\": \"ee75f6ae-4942-4288-a099-a3c619918e71\", \"version\": 3}",
		Secret: "b9mb8nc776swztge02",
	},
	{
		Key:    "{\"address\": \"c9eaadad232326edc3be13a57c406aa18ad66654\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1cbbdc3c106d3ed99bb24bc5d5e7d7e4\"}, \"ciphertext\": \"aa0afb7fea43432ef11d0ae2956295d42465a9ee19901205ede32f7e7ae19849\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"931ef892c4cb49aa5e785d22c13c234d\"}, \"mac\": \"15df55700ba66b714b03aba82e67fb8271e14fbd8d3d0a540c30f4ce29e1da01\"}, \"id\": \"3c440baf-99bc-4893-96d2-f107d49a55da\", \"version\": 3}",
		Secret: "oou6no6qsou71q1iol",
	},
	{
		Key:    "{\"address\": \"15219aab059a4988912cc091bb818f9fe2c83c35\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8e22497ea4ff9e50c871f6a312d5f33a\"}, \"ciphertext\": \"ae4342beba53ee206789935cbedc1b7eac42607710cbf5d3d664742145dff285\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"dbe10d1d5a007c8f3cf84b8df6fa5617\"}, \"mac\": \"3415a435ccef80aa5f43b2d0594580a90cc6275bbde2e5bcd3058656b2778380\"}, \"id\": \"0c371823-2508-468e-a3a8-0db8b7414c28\", \"version\": 3}",
		Secret: "lzjitp91xk0j4rdmtl",
	},
	{
		Key:    "{\"address\": \"f67283803f9e7185bc43588a75ec71d493d5b996\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"202cbda722dd53fe82f6550b8760eb31\"}, \"ciphertext\": \"a1cdf439edc1ff4115bdead734de57a720deb8c751ea8ea52f84b7fec966eb74\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a8443b44427703bf1f8f62aac4753310\"}, \"mac\": \"c6bd185afa315c72e4dea75498d9e3da482cd3ed216a958bf1107ed84df412d8\"}, \"id\": \"0771b812-6cc2-404e-a4c7-e9a618bde4b4\", \"version\": 3}",
		Secret: "bqrvopalq8mm130u8i",
	},
	{
		Key:    "{\"address\": \"69802148b1f74b5db761f615eaa0cbfb7c2b4f3f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4ef82268be6cae3f0ec52d22bde12400\"}, \"ciphertext\": \"960923bb5949d46a51c82a27efb26bcab56f0aeba931ca029aa6e01f80e68154\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ca2f04ca2e23ef3ac6bdc801f72a7fde\"}, \"mac\": \"236c6cd2da2ed32759c3b49b4ae97131858cf94d208453886893a16bfd5b42e5\"}, \"id\": \"9210f436-8d13-4908-b602-32e19c676854\", \"version\": 3}",
		Secret: "rt178gydgtg58wg23t",
	},
	{
		Key:    "{\"address\": \"3cb1995e28fbae3308fac33ee237ee7c3867dfb5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ae7129c043ca07aa044614e9257720fd\"}, \"ciphertext\": \"017c6b39a7a7c7ddca0ac89885060b07123da1cd6c48d21f9881f135697472af\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"899a323b3da2e1f5c27d9190fdebba8f\"}, \"mac\": \"365e6d22d3a0f04445008ed86e2430a0165caadfbb54ecf802f86653b1b30669\"}, \"id\": \"e3cbb037-e97a-4cce-b2df-a5fa0da52c36\", \"version\": 3}",
		Secret: "htz714fu3yde2t7mbj",
	},
	{
		Key:    "{\"address\": \"c3b7441d3342c06fe7f2899908ce5742a250ca11\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"70cccd9fc7f0bc8dfd0f0333479b59a8\"}, \"ciphertext\": \"2833a92e87d8ce7d9fdacaca441a262c214261bb475e177ac310d82c22c6fbbc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a590ed544e618b0a7534c93887c11660\"}, \"mac\": \"ba63d753d4139753653f64c11d69b940e65fd2f4a25108887430bd15ce131e74\"}, \"id\": \"6f3e1d03-4231-4f62-ae4c-887ba9ac2a9c\", \"version\": 3}",
		Secret: "pxcgi1o0ykmxadieku",
	},
	{
		Key:    "{\"address\": \"13aff5627d79c61f47e4052bb8aadea24579ef24\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a8da0d4da3bdbe0ec6a6c18a01ea68c1\"}, \"ciphertext\": \"d2f64e3e2823651e8cde2a18648dcf4369e7edcf8820c8e7f8d47e3440238f21\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"82bd07b3219c06a5d835aee698c0af4b\"}, \"mac\": \"70009fc6b3bb5ca3eb98af0310582de52060829a1244eaf3ac1293b7249973ab\"}, \"id\": \"32e9df88-0f0e-4df8-9742-b316d5af8fc1\", \"version\": 3}",
		Secret: "jk4veqzd3f8vp4xcl6",
	},
	{
		Key:    "{\"address\": \"2dd91affc87711bb51f942b127921c4841e77ae9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a1e2a0a4a7c054eb549a837e104421e5\"}, \"ciphertext\": \"ee8cccc8b61b0c7029e223237f439f138d7644e4bee50da136d75e586315956f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ef67cc2c953f4c48bbf4f29d2deb7dd3\"}, \"mac\": \"893c552e297c66bf2afb0d19ac08ab58a7b00bb65d9e6375754b8b10ddbe1875\"}, \"id\": \"63eaacb5-9900-47c3-bfa4-9142685a6d30\", \"version\": 3}",
		Secret: "5tr2h34idr4jrfxgvc",
	},
	{
		Key:    "{\"address\": \"eee6c52db754a632c4057b3e58ee675cab7d7706\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d5825dad265249b9fabe19e7cb3489db\"}, \"ciphertext\": \"a9e4b36758d701d57f61b583ed6200e979c459fd7b14980d0f8d25a1c6ccaf25\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0af7c8c6d7066fbaf833fdf45ec85794\"}, \"mac\": \"21381ea2ce93d40242e9d038a2a0099b49d9f43ca496083608abe7b77cfe3467\"}, \"id\": \"fce75670-8bd6-43bc-97c3-d2ed4d1db9f2\", \"version\": 3}",
		Secret: "ypkkr693hn2wkp738v",
	},
	{
		Key:    "{\"address\": \"35e470d5460d2397359288d2d8f1ec4a5fe2dbda\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"be342db16b4e06907a1ebb2ac44612e4\"}, \"ciphertext\": \"c5de210568b980684b4b3d3175978abc75abdeb98c15e20347a5c723a572bc8a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4fe964e4ae4a50e2b28707d091fbbdf2\"}, \"mac\": \"7202e394b7bf9829e76506f738b04cef9c44edd5ce9c0c73122d23498730427e\"}, \"id\": \"7de1d9cd-9edd-4097-ab31-8605a4ac2e97\", \"version\": 3}",
		Secret: "p24nc2hwxvr0qwyw3e",
	},
	{
		Key:    "{\"address\": \"ace172bfd3375a887d0b823fef7d078d3b6d1691\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b5a0b60a74dfa88c57e70593478a05c0\"}, \"ciphertext\": \"600a65a0958cf8a46344056b16302c0b6315934ffc95f7191ec082e68e041bff\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b2a4a9a13d9f40c9ebc69c53a2ef2e96\"}, \"mac\": \"9cb6b69cb1b270eafc2d52a467454fa2185c748911c0cd94034aa9cfc6fca615\"}, \"id\": \"75f21d7b-553d-45d8-be53-d9703fe3953c\", \"version\": 3}",
		Secret: "sljo31l8o7hinbuwwt",
	},
	{
		Key:    "{\"address\": \"3c29d4867b61c4590aed35e9becc800b8d58f21f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0f8d06027ac0472196fdc7cc81416988\"}, \"ciphertext\": \"f5c01db04bc0d69f585af0abbf44d9a31cabeee9c15ccb1e5cc0be67990a7498\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f2ee18cc0d784b563b02bb58d33fbff7\"}, \"mac\": \"7a684ba0e376069f3a3369c681a2a867a66c5d878edcca63be3ce79cf1b71a22\"}, \"id\": \"4f5e83e3-fc1c-4a7f-817b-28cdc64c3e99\", \"version\": 3}",
		Secret: "m9h5k0whe7f8jdjxy8",
	},
	{
		Key:    "{\"address\": \"959d9c141f21a1dce57881de750b6d93e300b5ba\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3eb0e9fe299b94512b5c3fc9bd1f454c\"}, \"ciphertext\": \"455e687b9897f69894f27376faa6cb6d0b8cf617c1347556446b80d4abfa142d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"093e2f01d134fcb95aa5dfdc7cb5d41c\"}, \"mac\": \"6cee978aa6e2a144126220515acf757c5c4508766953dde537e80bf09f417923\"}, \"id\": \"0ab1ee84-1236-4d4e-9ffd-1c47f85d9fa5\", \"version\": 3}",
		Secret: "dq6npqopci89mx9j0r",
	},
	{
		Key:    "{\"address\": \"a97900cc7ddaab7be7c53da48e0a1ff0cfbe541d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f05b9eb09ca9c0cc6f64dcaadcc1c67c\"}, \"ciphertext\": \"ccce6198480fe2930c033e1870cd9fb3256f017a6e396278168d507b40a9c099\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f66df3b8d16d96e7d3f4385b15175669\"}, \"mac\": \"b26d6289e485fc396773f32fa2bd5424bd6638f60fd76659100f50ff63aba7ff\"}, \"id\": \"42dac564-284b-4081-9f2a-ce624d737309\", \"version\": 3}",
		Secret: "3fsrtfjkbhxly6cdv6",
	},
	{
		Key:    "{\"address\": \"6f545d4aa43e36883acb2d068ec73f544499071b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e09a05eec4451b47374c373b9a376105\"}, \"ciphertext\": \"6871ce33cff5752cd83e31921e36a868064bee46f6081c866c1debd702fe8b0d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"010dbe1e4e632295819adcdef4894852\"}, \"mac\": \"43796253eab377296be347f11f92b4506a452c22722808078d0000a8001e29c9\"}, \"id\": \"79a2a47d-8f32-4a58-b59f-a2f760e6f121\", \"version\": 3}",
		Secret: "2e1k06d2qi609wc0z8",
	},
	{
		Key:    "{\"address\": \"a40cfcdd925386e4e9e692879ee7c07cb208638a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a48f9b71e8c73f7409406b05fc430197\"}, \"ciphertext\": \"c7b8972520edbd312a6b0fff5d0e179acd1a6fb76ff93a669f6b804ab94a0a2a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"46f347bae0f6ec277e9afde28e63ac2a\"}, \"mac\": \"5be99579d81f14d13294804f72bc670a1f5644aceaf6f1649adeda04d6445c13\"}, \"id\": \"46267d1f-6433-403f-a5af-fff5b2c17add\", \"version\": 3}",
		Secret: "ce1qdmhv33tcr2ztko",
	},
	{
		Key:    "{\"address\": \"15ebfd33d05b66d2c793a253fc328732533acd31\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"bb24e9ec3ff3468f9f003c39f276fc8f\"}, \"ciphertext\": \"722f0b1e3f72355c9a4b554ff3a2bab19a55fba8b3d32e353ab1651c2a23a971\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"66749c6a2f923ed4863a29b368448086\"}, \"mac\": \"6c5869a122c3d176b2eb5f9a2371dc0f140d284217d67a218ed950b6a7707e62\"}, \"id\": \"076931d6-88f1-4b43-8b59-1b8511b04466\", \"version\": 3}",
		Secret: "b8i3900g75w46i0mhz",
	},
	{
		Key:    "{\"address\": \"3591374ec6488a34e225e155e239dc9e05ee773e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1636ffd6502e41137fd68057c41e6aef\"}, \"ciphertext\": \"f448459bbf30223da025b8fcc3ed98698f5491f9f05e26bab59c35597474ac79\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"67df08401d4344ab8818960a03a75334\"}, \"mac\": \"a09de0f2e059ee09ee6d9bea05fc69e9aa376643b1b428f66d06250bc8492eb9\"}, \"id\": \"2e472e7d-8fab-4876-8b27-92d1953300a4\", \"version\": 3}",
		Secret: "zvil1g20oi1uy9f98v",
	},
	{
		Key:    "{\"address\": \"fc2f3d8294540611153445e0344d3d43e02b7154\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e0ae0aa89222bde6abe706a7dc13cd69\"}, \"ciphertext\": \"26d6d6a5441f31830de19b8db954a35d26a8ca6abe8b107ae6a012bdcfb2dea4\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e57098fd3ed631040d117bdbdb75a33b\"}, \"mac\": \"df64a8ce21b26caf00d66881c36397ef64068a07f86c73cc8e1d9be4c4e67bcc\"}, \"id\": \"ce938feb-048d-4de6-80bf-5616a5be0a23\", \"version\": 3}",
		Secret: "3syo7s8957pbaedzls",
	},
	{
		Key:    "{\"address\": \"c2c18b02df3ef32b03e91daf4a8035378b80e625\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"08e501a5b2d3a117f4a4e086d141097b\"}, \"ciphertext\": \"8d77abfad44f51fe4046cee8f5764c710636b90db380387edaf34ac69d84dda6\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5b86f5c706dc41d3a753184726846500\"}, \"mac\": \"f026a1fb5c401a21b4edbe120a302bdaaaf06ac00dd15590ff4eac678ff1efa6\"}, \"id\": \"2d2e3d34-9d96-4092-8945-9d7851210996\", \"version\": 3}",
		Secret: "4vsyxu4wiieofrpbpe",
	},
	{
		Key:    "{\"address\": \"9f7440d28028bab9e2a909b37e9def6848c8aa8d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ce5965bad26269794bbcbd52aaac2142\"}, \"ciphertext\": \"02919f55cc4ed6fbe66b30e6d068db0f07ded953454d8b5a4ba66fb3431a0896\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"dca9f1325bfff5dbb02a544555f16b35\"}, \"mac\": \"b7853684917e1006cc3cec5e92823fe13bf8c73e2e6653302b4fa51892739d22\"}, \"id\": \"f1dd4296-9590-4a93-b530-d0b499b7ffe0\", \"version\": 3}",
		Secret: "xenspfo1z6uwjuj73t",
	},
	{
		Key:    "{\"address\": \"41266d3a12de89a0461b0ab2fac9d4262990202a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5a319eac728bb181829609b34ff89894\"}, \"ciphertext\": \"e15026f9ffbfa76353a8e1baae9047455d743ecd2d7c81287559ab9e38c4db44\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d8af0fcdcd692d4980c0ca6c6e5b4224\"}, \"mac\": \"e17d5013611d519591fd59aac7defcf997248a8652772083f5ef63b2b8968e34\"}, \"id\": \"3cde7e6f-100e-4c79-bb96-534d4877ede8\", \"version\": 3}",
		Secret: "rxqyth0ru878l6kb1j",
	},
	{
		Key:    "{\"address\": \"52bc7ddaca4e617d1e102d8268e723575dac0ce6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e4c099b8ca2d264fbb3e135c28c53053\"}, \"ciphertext\": \"00d5ea298b4c05085ed3859b14129ea8d893a0c659984edc2f7ed0847bbd300c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"db9629dd681658583fe2c248ff256230\"}, \"mac\": \"0398681ef67aef0cc95d4d6c02f44856f93e0f21edf2aa2b2f16a8d0aa23b183\"}, \"id\": \"2d01b02b-c203-41d8-829d-eeaa63614749\", \"version\": 3}",
		Secret: "9wqabaqoeoo3049psa",
	},
	{
		Key:    "{\"address\": \"577e11011801ca9f0872e49c1997d0f1fe84f694\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ba09bdc897c5b7c4aa11d83643210401\"}, \"ciphertext\": \"b0b04b5beeb6dcc0ef4cf5da93067071d0b4a3e3d44d497eca015c1ec45bbd77\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"07258350ec0a76fa156506a25556dafc\"}, \"mac\": \"e89ac530a6cffc491c99c071158f4e796359c5dc44f400ff9ea90137f2f45d8a\"}, \"id\": \"7b190754-fbbf-4c7f-b0e6-13f52ce36e57\", \"version\": 3}",
		Secret: "klwgj4l42hwbstu456",
	},
	{
		Key:    "{\"address\": \"a4b0840349e2dcea854b824a3f2dd5f34a06b8ba\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8685dd131d67d3d009c7c5f6cbd14a2b\"}, \"ciphertext\": \"66e24edab14866717f563ad9c8146f14d6dfc2ddc65355292585717ae3a57433\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"99817ada74a78c5bf5e25270ca5fa6e9\"}, \"mac\": \"c72c69d5ed1e105e97861128c27fd2acb9b077b1e6233f0cc62340512e39a22f\"}, \"id\": \"bcdc846c-6b7c-4657-833f-c0ba6c396fd3\", \"version\": 3}",
		Secret: "yrny0ce98cbohm844k",
	},
	{
		Key:    "{\"address\": \"420786de401ef0e3ca7a3a72e9cb5db6b60ae668\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9e98765758b43570ebef6ac9579775b1\"}, \"ciphertext\": \"d7bef446f0a2a83b9d62177be7a2a9762ab5e4f53115e1f4d30314371e3de12a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9a37b3d362562e517203cb2e927db755\"}, \"mac\": \"68d6ffffa35aa40522aae0f74582c273d4c73771fd87d4d31be50d3747e85668\"}, \"id\": \"4f4bb3d4-71c9-435c-8e1e-72012c095f3b\", \"version\": 3}",
		Secret: "nyx2rm39py7f5th8hc",
	},
	{
		Key:    "{\"address\": \"a9673935f0e318d1c393f717e0b71f82707ac5dd\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0d374e716252acf47e01419a557d1327\"}, \"ciphertext\": \"8e235f8c00b282dac8f8167b06fff91f38174e1754b8dd562e3abf8d0547326c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"764502ee1921378f25b53f045aacbf7b\"}, \"mac\": \"4438d00ecdca6fc8ba306fc9adf1ea3abf9c039b2ed1ca3d9a42e2291fe0f286\"}, \"id\": \"0fe59b81-1656-4046-b644-634e42fa6cd5\", \"version\": 3}",
		Secret: "tw8g3xu595o1btvff2",
	},
	{
		Key:    "{\"address\": \"2181cc0e0b6474e9efc8e04fca2ae7a24e297323\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5231c4ecb4b45bf9c7ac6e6f091a29a7\"}, \"ciphertext\": \"9b8dfe2fa6ba47931bb8432968f78ccb41e35ae3d224826fc944dfe30c5ad356\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e1793f2e5dee8fb64f5a97b5a69e263b\"}, \"mac\": \"abcb0d710c4bf545ff2ce43d14c5a1e3759c3c2e5907079edfe449c122f02e7d\"}, \"id\": \"dc0e9ab9-6802-4dd3-a6e3-d0ea715e1e26\", \"version\": 3}",
		Secret: "17ri1tnsvtazvp3wrq",
	},
	{
		Key:    "{\"address\": \"8756bb86e541dc6605f57f877f5505aae93cd455\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"df28e8541fb97d5ce749855c32e569e0\"}, \"ciphertext\": \"ff9321e44d3b11eb03304217af757e214384315de053e53f6f1632eea557bf20\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e45cb1d15fdc5f46e6e4bb949f996afa\"}, \"mac\": \"ca2d4df9ffb36bb01ae0dcb86c97420df22e4f12127ee5b807aee053305ee40b\"}, \"id\": \"c45524e5-e295-4b77-8d62-c87b9b9c7501\", \"version\": 3}",
		Secret: "tmgrye043kgqz470fx",
	},
	{
		Key:    "{\"address\": \"f97f2c844a9005f30935e7d924b1807acc33425d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"7a5bdcbec195bcbe2d27db3bf63cfbcb\"}, \"ciphertext\": \"a258dd53b188289bcf3661b3077ee004a8d52ca5726ddb9bc6b4f109cca77900\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1697453d015a8b1a63abd7e071e53c3e\"}, \"mac\": \"1b6553843c1bb31e7e7f35acb31c256d20f4beb539cf187bcc474f5907ea288d\"}, \"id\": \"1187b3d8-eb79-4b98-ad31-63e174c77ad9\", \"version\": 3}",
		Secret: "x0bfqdubsm6brdgwyr",
	},
	{
		Key:    "{\"address\": \"45b88950ccd93f43c14161a5e29b56ebcdfc001a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e2cdf5139007b477ee6a218bddf22a52\"}, \"ciphertext\": \"dcd78a2819e84ea31670ec421feec34521cd01f00e26b104d0728b080fb5c61b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7f3c351738dee3a04c9a23bc6e57afb5\"}, \"mac\": \"f0cf99f1431f03a14d2eff73407f94806d85035bcbff9579961fd92c23b63860\"}, \"id\": \"f1362b6b-02ec-4c28-88e4-27633e2fb199\", \"version\": 3}",
		Secret: "dhl7kqycmshpkf1wdb",
	},
	{
		Key:    "{\"address\": \"8540292f79e943ccca91fbf7dc219ff48c73d99e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9b0c796751aa82273ae4d1ed4e7da3cd\"}, \"ciphertext\": \"2fc45800a168f172804e291c0dc0516e86b7f1ffa4c33dc6e37204e86dc30ee7\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c965ec0075b4da607106c68b81f1c3df\"}, \"mac\": \"9fd6fc97186089745a1492b519cc2afb5e523ad823b5bb8d3cc6363be08c7eac\"}, \"id\": \"70cbcbe6-8581-4226-a12b-51e1323db9db\", \"version\": 3}",
		Secret: "w4du8iig8odaqwdlpy",
	},
	{
		Key:    "{\"address\": \"4b153afbf276d45ab0ccdb586da9849a5d7e1be8\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c39d217f16fb6f797ab421f05ff9af2e\"}, \"ciphertext\": \"d12bc81386a0c3c385241fcf626ca9091d72971c8caf73d945318ce15e8bb3ef\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b3b8ad5b8e652050d1ef35ce00c8831d\"}, \"mac\": \"ba9831b43109de12a6ae09fee7bcc8994f90b9e70aff9514f4d1c03c5041a543\"}, \"id\": \"b994db81-9a16-4ba5-a2fe-21f5660405bb\", \"version\": 3}",
		Secret: "4llp8yjyg9wmd3k17j",
	},
	{
		Key:    "{\"address\": \"79e195ef6e4abe76b637394794b9286778b8b441\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b6b8265ce82e7ded9eb21193934f7f81\"}, \"ciphertext\": \"5eb7287072116a351187cea402d06214bcf0ab8ab770fa4d4b2a555450cbdfdb\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bd20434d06b47f49d3173df6fc6d0b9f\"}, \"mac\": \"01ad8945381aa5a5f5cc0e21622a9d81756ee4d382c93d3cfbecab938ad4857b\"}, \"id\": \"8f19c540-a97c-47f5-b607-bb680b1f52df\", \"version\": 3}",
		Secret: "k612mze7nqozhkftzu",
	},
	{
		Key:    "{\"address\": \"2136492bd0a805c1b952aa379cd3f7cf22edc862\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"659496e90849cfb35acea740529dca5f\"}, \"ciphertext\": \"02e308f8bb27ec541360c5998c6d475507e72c3ab0a302fb500ded3a290ad98f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7dc6381fc87d12cfbc4c221351d3b1bf\"}, \"mac\": \"45d1091dff610c134c69cc003b4805ed7f26bac768da13bc02dd15deaac2f53a\"}, \"id\": \"72b77094-15f5-44ec-ab69-3d5fb9b19039\", \"version\": 3}",
		Secret: "cbcq6scxcoi31e2eup",
	},
	{
		Key:    "{\"address\": \"38cb7cc0d377d58e983a53b46d1bc0c59e08a5d9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f9a95626a67ba7ab20a0699485a48a8d\"}, \"ciphertext\": \"b9b8d8a39ddcfa29b09cad53de6c33fabed5c177ff51a50e1783d09326def897\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d02baacc34763f77dd32d4d442a1c4cd\"}, \"mac\": \"bd21011b3fb475624a2470e3eccef510d0d4d2f6925ee73d582708da1f1ddff9\"}, \"id\": \"b9d7122f-1d76-4d17-8e0c-c87aa7a3d6eb\", \"version\": 3}",
		Secret: "gpuohtxlqs5w31z3c3",
	},
}

func init() {
	managerKS = shuffle(managerKS)
	validatorKS = shuffle(validatorKS)
	managerKSPool, _ = NewRoundRobin(managerKS)
	validatorKSPool, _ = NewRoundRobin(validatorKS)
}

func shuffle(vals []*KSItem) []*KSItem {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]*KSItem, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func GetManagerKS() ([]byte, string) {
	item := managerKSPool.Next()
	return []byte(item.Key), item.Secret
}

func GetValidatorKS() ([]byte, string) {
	item := validatorKSPool.Next()
	return []byte(item.Key), item.Secret
}
