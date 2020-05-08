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
		Key:    "{\"address\": \"c41a4b13615f501870a1292567657dc4712f333e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5fe3e58851bebf1673c1b8bb26d673bf\"}, \"ciphertext\": \"9a11818c562f15416cdc56e6f0dfb316cda90ff83d8d6fa3edb370e99abf7d3d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1ff0f730523f2290749b03de8ad53f27\"}, \"mac\": \"b74f81c20ced379cfa3be5645e1bb1e532fda3a81bc894cdcb613c04cc3a5166\"}, \"id\": \"b49e390b-e55b-409d-96c3-5fb07c75d562\", \"version\": 3}",
		Secret: "q5tgdklgq80dwfhvsq",
	},
	{
		Key:    "{\"address\": \"eb64ef1b83d90816642823bffcc933492ea5ea5a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"cd45ead0d5bc93d245be450893762549\"}, \"ciphertext\": \"2f69c3eeae6dad9c6a00832325d45e47f5ade3fc722ef1ad4ce7fdb30987e279\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f8b5917c481729255fc2d73c852da0c6\"}, \"mac\": \"90682d497a0cb0ac4fb1b5658fa40d83a20a45c65195038a3670337b7aea05f3\"}, \"id\": \"3a855663-e707-4202-931c-4321fdfdc43f\", \"version\": 3}",
		Secret: "97nwxjra8ti9v7c435",
	},
	{
		Key:    "{\"address\": \"06f26668ca2ae4ba5de193774574479ab01b5db1\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"75641b4af762e560eaffd9879b4fff40\"}, \"ciphertext\": \"d40f7c385be1a41e2431e89c685657f4a0f6b3a7fb8fe2d4b3d05319fea9e011\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a0354a2ed68e1871b91f9b73ff45cd2c\"}, \"mac\": \"eb0d0baa15d147f5a6306e4ca38f4f6d6d8295c217ab6692dd95252a6da6b9e0\"}, \"id\": \"ffa744a4-fdd5-4baf-870b-377ca9540b43\", \"version\": 3}",
		Secret: "axcepnw45err14n20m",
	},
	{
		Key:    "{\"address\": \"2ee7540e0d413ef6c37d6bd1c257417a315d2a3f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"880388da8c13ef1778b9e0fb3d84e442\"}, \"ciphertext\": \"f7f49315c421346e2007592e08f9333eacd1d57ce45ffc6f349b280a76d6ab1c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"69c17127ab0e9fd94c76bd13d47e0000\"}, \"mac\": \"6ddb9edb7a64455db56728f80bdb06138dfe470bd4639fd4fb4b55a4a21fcd18\"}, \"id\": \"a31d249e-c917-4cf7-870d-e6cb476bf15f\", \"version\": 3}",
		Secret: "tyggrvwpghs2ro0byt",
	},
	{
		Key:    "{\"address\": \"effcbd6388334cc3c4fb4b71d9cf9b1aef9519a4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"22604e9e725782f82f0f56b8b04fd384\"}, \"ciphertext\": \"8a9bf7396902d50570a3ad7cbcb268b2c7b3d360ffcca26f24ff5a4a76bda91a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"106a41cee5649d5c3155336ed45a19fc\"}, \"mac\": \"a9f5e98978abe605811f1bdd58ac5e45396dc663bf1641591fd000af90fa763b\"}, \"id\": \"9fdcaf42-bf61-4a94-9765-60e772caf19d\", \"version\": 3}",
		Secret: "9ewqssa57djoxvtnpa",
	},
	{
		Key:    "{\"address\": \"135660133278906c4300bcaba1a27c5a48752913\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2eb350008a9ae46fc0566aa834924c52\"}, \"ciphertext\": \"d4639cf5e99fbf01ef0f7bb660cc8f3352c7fb1bb3babc2e2ae419acd610af60\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c9718da7315b7add0707a50a1564d4f9\"}, \"mac\": \"ea83c9d508b7152931ceb32f1b3512df4520f80f5d52f52b18e7e6d0f4d47e54\"}, \"id\": \"6698097b-d85b-4887-85c0-3203913be923\", \"version\": 3}",
		Secret: "uvb3n5sj51l7cyuj9y",
	},
	{
		Key:    "{\"address\": \"1807b5a94253cff20db7eb1c34d012dd61afe117\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c7f8173f7bac1074641dee70320f1e06\"}, \"ciphertext\": \"0fd952e9d8e403aea1d5315bd21c30854bc0d15c096667764f63f2613ad5b5f5\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9c41b8fcc698fdeae89b977ff7b66dee\"}, \"mac\": \"0e4ee2d721c4441c8595937ebd0453dd333b19d02764be1e6c12bedc50f29ada\"}, \"id\": \"a72798dc-f94d-400a-81a9-1c5de5b33f8b\", \"version\": 3}",
		Secret: "t9hwuwh2ajkmp2p1bu",
	},
	{
		Key:    "{\"address\": \"f36a2ab6049a22abb9b3136dce96b51085b5fb3a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5c415b968e6a59e7389d7efd3ba58ed4\"}, \"ciphertext\": \"17451096ffe5718c1f754e56fca74482871e560de4368bd919a1b086de0bf522\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d8b97eaebc07b4266dab62574ec90b5e\"}, \"mac\": \"584f078e41400d139896e656e35362c9d64cad04b0615470dee10bac1b3c9b32\"}, \"id\": \"8040188a-a0f1-4a02-913d-d2cd8c0f48e3\", \"version\": 3}",
		Secret: "3pqe3yijon4d7zp8kx",
	},
	{
		Key:    "{\"address\": \"98c547d7121d692f7ce3fbc94198e60a35e97f9b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dd6cc1078b0959c2edb9abd5b5f9e7cd\"}, \"ciphertext\": \"5212db9453f9ca04144b9f1b7140ae8804e6bf7786a22afe418caafde44fd29e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3c2c1a09e463803d7c8e1b1f1c9266d0\"}, \"mac\": \"c09e379058b2cb8e76cea115e1cd8cba06edbf2baf86e6b7cada15ce6151d3a1\"}, \"id\": \"01cec965-8483-444b-95be-360299e2bd04\", \"version\": 3}",
		Secret: "8o8llfbg35zw73l5g1",
	},
	{
		Key:    "{\"address\": \"5ce38d3ad1d897f1fb5fc7b141963e1a7cc068f5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f5fe11a55a6f737f7f8634831d8b8fe5\"}, \"ciphertext\": \"989a2d47d0c61b1d4e893c841ecee6a0d4c9b18007d5d3f51c03f16c2df179ac\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"291e91fd339ce1047e6ac5a276c62f77\"}, \"mac\": \"656a614f071c5ec86d9d1db660b97efdc1833782c06e38c1389806e70ad4dcca\"}, \"id\": \"b7a95e10-ea4c-4b19-82e9-b4f0d1a1e6d9\", \"version\": 3}",
		Secret: "s5ynctze3hjf53wa96",
	},
	{
		Key:    "{\"address\": \"63436c6536d9f786d9c58856b794189472efa007\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"658259420ab03ea38823c2dda7f037bb\"}, \"ciphertext\": \"208311839f93fa68787129bb46185e52c61b9b7f5a470e35e62865d1adad1b1d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6ff327f62764b487b78ae22c6ef33a7b\"}, \"mac\": \"1176a53a98c5678cd0460440cfc1e528a85e77ac33710b08db935f39926763d8\"}, \"id\": \"8d967e57-b7ff-4508-aa91-2580ba9613f0\", \"version\": 3}",
		Secret: "84xv62tkyyps2iwc7i",
	},
	{
		Key:    "{\"address\": \"40aeca58e637b513f129c0d4bb5956d925736adb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"535d8bc1a98c1656ace0cb3e5e700544\"}, \"ciphertext\": \"350da34f52d2a5825284e1e49dd7ce7349cb0376a06d68a4de7caa98bf6f28b5\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"448686460e8b2ee6119a94d7f1d800df\"}, \"mac\": \"7b469f7bb2c18598f75c67c16250cb85e7d4bf214f55ef1c3fc29c12485a25c2\"}, \"id\": \"a154d780-cd40-4f6a-8c08-4a5e4ca54842\", \"version\": 3}",
		Secret: "kqgscjbxhnnz0qy6ux",
	},
	{
		Key:    "{\"address\": \"ada56b08ded809e9beb19b804d20d443305d5221\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"440e2e61acd60935adc01bfa202161f8\"}, \"ciphertext\": \"c6dedc447a85f67609ac82861f77ba2ee593a156116d0adde85f14769b21fe09\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bfd90ea2338902f333ff8e1e64514403\"}, \"mac\": \"255b07a437ef8fa431b33353cfabb4a4a7919b8807e52ca430c01123d6fc7612\"}, \"id\": \"a1973112-5a5b-447d-9a59-5ab55d58bfe2\", \"version\": 3}",
		Secret: "ox2f4jddqb5wiosbcc",
	},
	{
		Key:    "{\"address\": \"ee47e4110f409678125006920893f9384f512c77\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d2d38b4d32ce0a7d5160b14e5cb9218f\"}, \"ciphertext\": \"9b0038d1d4d688f3fde62942d7a75da03a69c1b19c2813824289c5e12f32b6d4\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"281effa00b5ce07aaa8ea2c458d6c3e4\"}, \"mac\": \"4f5110908822fc233bcd013b209c9dd64c7fa88eaa3a5a285592cd8d141ee1e8\"}, \"id\": \"17296546-9473-4161-84e1-06e197c05a19\", \"version\": 3}",
		Secret: "4xb87s1trsu970omia",
	},
	{
		Key:    "{\"address\": \"2ffeebb01e30597f36622eb63f69ac3bae6f6797\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"01ee91cee5495fc9ea1fe1cea2896850\"}, \"ciphertext\": \"d62166b4e355b76814c5ba17d9eb1c55a3cac039f53931e12209f82f746c1f1f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"61343d75c2f9a5f58a76e5c62c74d502\"}, \"mac\": \"25a46ed89d13022a8be19e0267a52e0ea59ae5f54a33580ffef2efd772d94663\"}, \"id\": \"61d51013-cf63-47e5-99ba-faf2315927b7\", \"version\": 3}",
		Secret: "q1aagxst8phgbyczi5",
	},
	{
		Key:    "{\"address\": \"5fcdc932f428e082040452efb0e677a37dd7292a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"7f3d95bf27567c55255aaddf7a5b1dcc\"}, \"ciphertext\": \"f55a02e6f37323516eb52f55b5ee1c405d25aec520b6a83e3a975c53b9e3aa41\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6cd671f18a917ae79e1c7ca7b4fda86c\"}, \"mac\": \"eaea0c85a745c91a5fab4163e49d5fb852dd26759b527468fa058b17a4d89d5b\"}, \"id\": \"c4ec094c-d8f5-4e90-b8ff-c35bf1eafca7\", \"version\": 3}",
		Secret: "2ytq1llhukao2wmshi",
	},
	{
		Key:    "{\"address\": \"2e5b6bba0c62336139529926e6541fbc4e2b7c58\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"cffb6a22f25385b139887b126a7e4e7d\"}, \"ciphertext\": \"08b1dff14c92cb4a43b051cecb91c132023864407dd3c8adddc3971ffae35e6c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0a4806c09a8f342ef0e2e564c9060373\"}, \"mac\": \"ac44725739045e2f469b0d9b2b9083c685220956a232069f425b0e44d7633441\"}, \"id\": \"40a415a7-43da-42ce-83c3-0279f61fe250\", \"version\": 3}",
		Secret: "0iovgfrvwfcwdvkfhb",
	},
	{
		Key:    "{\"address\": \"1db4fed948c4ddedaadd92da76e2dbbd0692b157\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"10a96ac80d3f4d0786f9dae3d6f77235\"}, \"ciphertext\": \"f6434a0721600e92e0df15ae21053598355f675a53a03dfb4c580afed511ac8d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ffb8fe4754e815e3d13bd0cf05461242\"}, \"mac\": \"dbe12f2148b3b16b7bd8d8a59d11b5394ac53ab07d5ab9b34cb413d3d13e418d\"}, \"id\": \"6fb02795-b872-44a8-9434-907e33bd1831\", \"version\": 3}",
		Secret: "2pnsm5tubhldle09hc",
	},
	{
		Key:    "{\"address\": \"b0bd05d0b9b2b3838222405b429da0643d1f3c04\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d94481d6f41f422a07a5094fd6c3876d\"}, \"ciphertext\": \"6aa44a841d24ddd1b3ab6b1c3492ea85d9ce6b91d7135127c64d4e6d0ae64261\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0619d46c6fee82276bbaf321000c7c02\"}, \"mac\": \"524310b371cb191a2175a41e3922aab1e6c9b6e3f89bd59545c947a4d4d95e58\"}, \"id\": \"5ce3b868-b330-484a-ab1d-7dcc8680175b\", \"version\": 3}",
		Secret: "f9fyc6lwn6pgst99rw",
	},
	{
		Key:    "{\"address\": \"ee0164393b4952c230a719cde83e3742a1f483f8\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3dc3d4044af61a259e62f4750b0a4de1\"}, \"ciphertext\": \"fa1abf759cfb74a058b74b9dc36845761473acf5b46573dcfddbae26d221f2a2\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1e80276e7b6dbe4e2aeb35e5373e052f\"}, \"mac\": \"3c07b0432a0bc73fbc570dd9916e53a72d685e416c88468b6aa5c9edad0b9912\"}, \"id\": \"f4922cb3-aa0d-4001-8031-9f7801e9e9fe\", \"version\": 3}",
		Secret: "yk6v7hnrx6lubjwmcw",
	},
	{
		Key:    "{\"address\": \"c5aff216bc2af571813b4829a2e53e81b9e02e50\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e092fd5c91e5329149bf4e7572442ca2\"}, \"ciphertext\": \"d0265e2d0832b7c6e2b4ddf20bf48f103750d1a1087d47bc0b4c08d57e30722e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c300c53bc76f6f6b8bbdca9191141ca3\"}, \"mac\": \"f7f936e6bc75f45758cd072e6df327df652c75fdf11a632e53b1c09f40d8acb5\"}, \"id\": \"bfa219a4-90fc-47e4-81f4-407e9757e195\", \"version\": 3}",
		Secret: "fsb5xuzekjqqw3x9wr",
	},
	{
		Key:    "{\"address\": \"e33dc17789cce0db95c3e399e322799a30d82134\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"677241dc1effd4b3d83dd50123a6fa39\"}, \"ciphertext\": \"06a774aa2aaab2dc9e36cc5f28125b9af1e16a3de6daaea7eac7494bbf8f6d9c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"78cad94cf006fa33edfe49aa6d3fd331\"}, \"mac\": \"745cb56b53cb695631884212c026fbc3e5615bb0ec37ebaea698735abba9a5c5\"}, \"id\": \"e0030b2b-9147-4c9b-b86c-badd8b67f289\", \"version\": 3}",
		Secret: "hjxb0kmtlkaxkna4fb",
	},
	{
		Key:    "{\"address\": \"276f04c1db35d19a3c3cd6193c95aa70f75d0cef\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"73973e075b61e5050ada54660ae664de\"}, \"ciphertext\": \"575026af59b25b069c168d9e3a42663b58fd09aa5379c00899df650473caa8ae\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e123d495438a55977718deda03b0cff6\"}, \"mac\": \"7f16957f1daa0591c2ce1ed8fa8beea7554b989e47962c1729bb55f861d54e35\"}, \"id\": \"3a397b8a-2785-466c-add4-41c6cab149ca\", \"version\": 3}",
		Secret: "ae6m3g63cgfof2r9lb",
	},
	{
		Key:    "{\"address\": \"1ade46b66dfaab25a444bff623c2fb08494ce3d7\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c447cd7a7767d18e1ed99b5a92e20e40\"}, \"ciphertext\": \"ce584f1da41cb02d3ce0b56e38dc53c3c46eb205b46d61484c1d26939eb2297d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"20fd2d8824d84d9ce0551cfe1001ad1a\"}, \"mac\": \"888037f967a676d00856387f1d74a7040bf86a1652a78cea6b4fd2e860e33137\"}, \"id\": \"e9a65918-61fc-4162-87bf-85b2330165a2\", \"version\": 3}",
		Secret: "sst2jm8n8qwqu8vk50",
	},
	{
		Key:    "{\"address\": \"e74e744e3149b1168321b1b6c0dbbd7e64897707\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"aaa12e810368e065a662c01da49f30b9\"}, \"ciphertext\": \"5ee5513a8b1b788dced05f6efb0ab4e80e665cd207cdf1d86990becdb0ccfd7e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e495ad383c4212d97231b5bce30dc319\"}, \"mac\": \"e6239965c050a5f172c28fe7de80aeed17a12a73949ec4bfda197399c44b6c69\"}, \"id\": \"92b7e5b7-5de6-4566-b83e-1af861d6fb3a\", \"version\": 3}",
		Secret: "wd6krlb4kgmjvnkyxo",
	},
	{
		Key:    "{\"address\": \"3a3699df4829dbbaddf71d7c920f1725f435f886\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"416ce90e5b0d2390371464bfddc42e9c\"}, \"ciphertext\": \"2f9bfb224be1e6f1f49f628998317f3ecda247b5d02e69919bd29000333fae54\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8c4dfbcbb70ca4b91199e1f1ffd21b0e\"}, \"mac\": \"95d3fba11e261c0c13d020e48ed646ec75f0a12c213756e6f806985627a7e1a2\"}, \"id\": \"c4b4a3ad-502d-49b2-86e0-d1e73119dbbc\", \"version\": 3}",
		Secret: "wdncv9mgqtpb2z14sx",
	},
	{
		Key:    "{\"address\": \"a421fd97feb6b4eaad1befa7b450db27dccfbed3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f55a825f6f0d080f916baff9fb04b7cb\"}, \"ciphertext\": \"6fdafe97ea695aa45c5ff13a4ff83fd3863e5717e10a382e72bca1c0b9627ffd\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4bc411d650cbfa4fe10628bf4d08b062\"}, \"mac\": \"9127f7ac99e6c9063bb646897feeed43207fd1ed3df512c404b58cb1c42cc200\"}, \"id\": \"39ccf477-4853-40a9-9e06-4c47bade30e0\", \"version\": 3}",
		Secret: "cbw77kb4a8t6tg8vvq",
	},
	{
		Key:    "{\"address\": \"477ab629bc7ef7ba9c433bb2de225f289ad8b79f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"39c4388b437d0c55c0b91fbaa9b7bb05\"}, \"ciphertext\": \"abf282fe500c948245954a58ce8d09bc50da77e0eba889df0b4869c41395ccd8\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7bb144fd7a2522080dffd319d3f905e6\"}, \"mac\": \"634742c09ece0916b3af059c57cd15fd5f53323f2c3609d35280b36c0a32de06\"}, \"id\": \"f48cd756-e51b-41b9-b97c-568ff72dc8bf\", \"version\": 3}",
		Secret: "07nzz89todk9vk7m9m",
	},
	{
		Key:    "{\"address\": \"15e274b1b5f107a19a9417500335b5e3cc05581a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"afb5c236c8709eb851a924195bd03f90\"}, \"ciphertext\": \"655a927ae06d10427b591c022e98280012c4c3be4bb25b2dbc5714bbe5ff031a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c684b13dce431c469fa399671b91f9a3\"}, \"mac\": \"11d470a4bed709e5969c2f18d40835262d1d3ca919c0d8412601cce13d220d3f\"}, \"id\": \"d7f7d5e3-ef4c-44e6-9d25-87de6feb9f18\", \"version\": 3}",
		Secret: "dctw5t71btzh2lc2zt",
	},
	{
		Key:    "{\"address\": \"46ffa86287b1203d795b1f15d3576435550f99fc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"eaaaa4f91095b16c2cee867937699b6e\"}, \"ciphertext\": \"3a8c4fcc34103ebd0b23e141ef742f1026ea617d9c0d6f5d20684d4b02f21222\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ceeb2b8319068ec7c3ea061222fb9a44\"}, \"mac\": \"81912eef4504b07c29e9774dbee331bc10983dfe9bc7ef16be0a29eb93d7614f\"}, \"id\": \"7d41c16a-e7ea-46df-b2fa-23dc593c460f\", \"version\": 3}",
		Secret: "tujvma9ki598bykymk",
	},
	{
		Key:    "{\"address\": \"faa54c0cc26a26e5a7f38c2e198b7f4b137134b4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f889a1dfd6c619e1f930eab75f9de6cc\"}, \"ciphertext\": \"2c7889d264fd7f50e9ac37e0a9f6fd0f7b52b6eeb8b749c25ce2c7d15c16ab76\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"675710b83c7ffdd98ac10d55647b0bbb\"}, \"mac\": \"bc15888ec1eec1f219badbfc01e15cd1f404ef98cc96da1789292c7f3fd0f7a5\"}, \"id\": \"38965933-8414-4547-b15f-1fd08a6cf22d\", \"version\": 3}",
		Secret: "hg9r723uy08072v2pc",
	},
	{
		Key:    "{\"address\": \"0c7617c66ef1f1b67e5b8adf51bb739514520680\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9953f4fe081222be7ca22614c1dab401\"}, \"ciphertext\": \"7eaa9d71182cad90afdbd8ce61b2a9ccd167b905a604cc51cb2253ef41b9d974\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6ab606e84b22ddc1c2ed5f93cd082805\"}, \"mac\": \"53961004bf605f3fa84b7a81ee1fa4ea08b212488466d9d1485404afbacb292e\"}, \"id\": \"865d1ab9-6c0b-45a9-95fa-8bd029c53fbe\", \"version\": 3}",
		Secret: "mpiw0zax1s1nhwhsi8",
	},
	{
		Key:    "{\"address\": \"08123746f720d76607208c40b1ef0feb92156792\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b634759084d0edb65eee5c68f96ef223\"}, \"ciphertext\": \"b176713c67a530408fa2d5cdb3126a13b01d2aeeafe2c18b549ce3f6d6a5fa21\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"aee3ea51f06107c16032e51ec4a0cf10\"}, \"mac\": \"0685168eb7898aac141eaa906136c0598efa5fa6caabbc5407af3247f8135c22\"}, \"id\": \"c0c86cda-eb5c-446d-a266-b06aeed605a3\", \"version\": 3}",
		Secret: "mkj14ri5ewjicv1ayo",
	},
	{
		Key:    "{\"address\": \"a93da4fbd6c56954223c9595869a55fa44a4af42\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f13e573e544d4a6e58033a9db8536d23\"}, \"ciphertext\": \"49a59e35e21ef86b0c0e7fb8eeda8b5a47bdd40e74695a04d23b4b5bea9a0815\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3e5e015ce338c6befbf068c3b31ce5d3\"}, \"mac\": \"63e377c6531413ecf7345787b3fe1cf96e3f50ddccc749e38b5e20cfc23bbccd\"}, \"id\": \"3a7cf8c4-e231-42af-89b5-2f9f1d3315de\", \"version\": 3}",
		Secret: "8cilpv6efxjjtjgewf",
	},
	{
		Key:    "{\"address\": \"4ed3b3aee91daf881b9ee9b89350d4f025033f38\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1d2619d07dc92bd485e4026f331bb564\"}, \"ciphertext\": \"b9a59b02ddb2590a9973d382c77b070df64b2eecb7e0dd948aa30d8b3541b781\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"212ed3a6524b7d17375e029f4ee7096d\"}, \"mac\": \"f7a3c7ce010f0b53702fcf139fb2b5df722aeb6530e99dc60b6eaa1097f47195\"}, \"id\": \"2a42febd-c8fb-4a8f-9969-38cfc8468226\", \"version\": 3}",
		Secret: "648x2n107pwquwnqoe",
	},
	{
		Key:    "{\"address\": \"79046bfaf9bcc17bb45608e847e4e6e8f6b957cc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5bebf7e8f3de5a782f65a55780044e7c\"}, \"ciphertext\": \"a705e154bfa928d8e8ffa0f44f5d70d3e6e052f0392ebd25a3b45bb8e8d2bcbe\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"66e6edc620fc302e6eb58b3ef623916f\"}, \"mac\": \"169a1dda7833d358b1045bbb4e4fee5fe3c9598c6c71196110f179e27fb10355\"}, \"id\": \"3f9ac1d6-ecbb-4e5f-98c2-08265afd47f1\", \"version\": 3}",
		Secret: "g5xgqotwmv76ophev7",
	},
	{
		Key:    "{\"address\": \"10973106d1916dda937a476f6689c4043413ab34\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1ceb0d050bdf698596210f2894e39d7c\"}, \"ciphertext\": \"e9bba8b4ee1d796f8570bb4c655e04f76adb4bfb5c31ce90f6e2dac16510837c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2954455b1cc4fea8989aac32a52b278f\"}, \"mac\": \"8eab14a63a60ee4bddf116cdd878034fa036d6d681542f0ecfda37984aff1e0e\"}, \"id\": \"b3cdaf75-83ae-4557-a478-dd052983c720\", \"version\": 3}",
		Secret: "42auodghohb8zib2u6",
	},
	{
		Key:    "{\"address\": \"b90ae75462179f09bac3e6fc697e4559f8d81e3f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"117adb13f9c550f0d7b85c3853c63d5e\"}, \"ciphertext\": \"b8d090f127579558ea4a9658b90f1a6bc82a55d7ec4d9e7852684f0d24f19ca0\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4057fda59a3c60e9d830d26c38f1967f\"}, \"mac\": \"2276e5cfe8e344b7cb596aa8e3b49e8a206a89a1cb2346eabbb5740080ec16ed\"}, \"id\": \"91b7c3e6-5ae3-44ef-b851-85c22a7a3e79\", \"version\": 3}",
		Secret: "w46a29y6gt3d790cmo",
	},
	{
		Key:    "{\"address\": \"44f222fd8de19c2872bf251ca5e28c8d8867a9c6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e6b806631bfcd9590634db99cd524687\"}, \"ciphertext\": \"a7be82976ee95997a5af71c94bd76c2d3040704ebfa87c3c7a5fc4f2ec44ad6a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4f1c7bf8b83021795678bdc04ce0a645\"}, \"mac\": \"2a2e537c2b954402dcf3af70c1b64cad2a951b1f5d3f20888471eba6fad0794e\"}, \"id\": \"aa86c513-c779-4d88-bc9d-4a6afa389f57\", \"version\": 3}",
		Secret: "st8ya06ggbzbq3yd8j",
	},
	{
		Key:    "{\"address\": \"2faf8a54924ca4c8fef687b64077706894e14618\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8806e2954e12a1dcc87705248301db76\"}, \"ciphertext\": \"ada0ceeebc5f835f5fe64c0e3a9e19eb4dad1a743a99cac5b46e50d58691fe98\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"455de114a4714c2e138e431e2762ed19\"}, \"mac\": \"d0b54e3f4dd440f0d47ae636431eba7169db87e42418321d127d33e3212949a7\"}, \"id\": \"67355398-6f62-4855-80d3-438df22f1c7e\", \"version\": 3}",
		Secret: "9qd65clo1s0rpse4ox",
	},
	{
		Key:    "{\"address\": \"2f93eb9305ad7188bde5507f6bd0fdd7c2f2da0d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b6e84fbdeb77207f4af96b7839930dc6\"}, \"ciphertext\": \"d6e6348953914956928817aed88711b91b47c8827280fbc87bbee2f843eedbc1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"08df9a1bef3e5a5333d6bb0082de4da8\"}, \"mac\": \"19044e72f650d570a32866a399912123581d51bd27eea8d82695bc80bd02a62f\"}, \"id\": \"ddc02d25-205b-4f48-ae6d-d84c1275b9ef\", \"version\": 3}",
		Secret: "ndkyculp4n1hxbdk79",
	},
	{
		Key:    "{\"address\": \"b2b28b9d1fa16c97923900b81d6205714e3e6afe\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"68ecdb150cd9a47c543bbb2454b69cf0\"}, \"ciphertext\": \"02596951a18f4afda7b1082f4be42f3cc63666246200f1d7fb62549e77176181\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6bc74684ebf9885827f203fafef02003\"}, \"mac\": \"08c693b34b1f22c528b7f8c2d92ab5c0294d8af957875e7f65d6c9ead55d5070\"}, \"id\": \"df6d97a4-7e95-4850-bd28-ed30de4dd15d\", \"version\": 3}",
		Secret: "5sam4l03ljl41gvxuh",
	},
	{
		Key:    "{\"address\": \"f751b926e64f25ac7432679b7690a90b3753b09a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"61e4bdcc72fddb8e565bb6c9bac0d1d1\"}, \"ciphertext\": \"505b5cdc2e573534f7f8be4c7ab36e1c02fb74f79055645abe3fab5b53010722\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f181e52e5dae0561fd366b619ab84233\"}, \"mac\": \"344503dd6b59dbde2acbbbfbdd37422dbfa91e0ccee0586f8de2d11e2611b7d3\"}, \"id\": \"bd792d8e-b392-41bd-99f0-a02bfb3fd3ce\", \"version\": 3}",
		Secret: "f80dtj5qldj0vo47eh",
	},
	{
		Key:    "{\"address\": \"eee3df9a918a97c5018e0f788c8aa6e537f88758\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0c1b8b8a0e4ccc4d7a5f8336ed8bd908\"}, \"ciphertext\": \"86f915a58742c0d20ad95a50329753a77470736d4a88c68342f1a990bdf4d17b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d2f88c953b7becd020c7dab3974388c0\"}, \"mac\": \"44a8f007b34fd4e145fa11a945416501d712fbbad6d502b0e7755be6c4ebbea8\"}, \"id\": \"e29f026f-55af-4e51-9a9d-8da231c30e39\", \"version\": 3}",
		Secret: "ejr1d7qdctynl1i55e",
	},
	{
		Key:    "{\"address\": \"d34bdd87d0477c68eb9ace4205d1d14d9d2129fb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9cd1eed4c0af88c6a820a1afcbf7acd3\"}, \"ciphertext\": \"9f982efcce1ba2573c751e7e70a5d99737c6b741dc867b033ac9a746e35b4c9c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a6396fca077e27bb9c65e671abe43385\"}, \"mac\": \"c969eee7566ec15888a92a7add64f3b506cafb36f32ce44df163149fd9bdeeb4\"}, \"id\": \"8282ae45-881a-4236-8432-8d257a4fe0d1\", \"version\": 3}",
		Secret: "udr5lrws6qajbykx3w",
	},
	{
		Key:    "{\"address\": \"e0c18cd7ad83e97c82e3bf27d82ca498d1019669\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"45fd2679fc173dedac24c5562f3f6612\"}, \"ciphertext\": \"a80c3042460ad9b49c93ef82de4f9186265222af47b1eb7466255ace3a185be1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ef3ecad1e43f053cb3dcbf9e6a755a6b\"}, \"mac\": \"2a4d5efa7241721d18335cc402e837252b0d2bd963e20845a74828970d649f23\"}, \"id\": \"bed5296b-46c3-486f-9792-fd361120e917\", \"version\": 3}",
		Secret: "8llvejfme0gti5vx1o",
	},
	{
		Key:    "{\"address\": \"6c4366ed4eea7bf2f8c75f6aa1e577a82e5f2ada\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"197159671d74ac5713eff0abfba3fe1a\"}, \"ciphertext\": \"57e18a7fedfa90f6f048648b71366edb34a6d351bc744abe270b85a7e1991d17\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e525e55005a1aee137688fbd351a27bb\"}, \"mac\": \"723663442fae4d1ebb2f2262cb8cf728dfff683a834762c0a6478237736fd115\"}, \"id\": \"ed47fa81-a9a5-4a96-93de-d0936c294ae0\", \"version\": 3}",
		Secret: "f3rbwu5mj089p9gqd5",
	},
	{
		Key:    "{\"address\": \"a624b4e451d4c87319bb2561e50e977d1a1da00f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3855d3b93e0580da42e58bf33035bf8e\"}, \"ciphertext\": \"aafcddd90186ba5e62b41c98bea01a0b69026b7882d507a0445785e08edbc73f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a5c739de8a2745553d5d7b62a21c4a0f\"}, \"mac\": \"98dd42305613e3d5e850d43da4740ce97f7101b937e888ffa3ec28dd5bd79d58\"}, \"id\": \"997de94c-0419-4562-b75b-e88000168f3c\", \"version\": 3}",
		Secret: "j2yufakggkligt8d1t",
	},
	{
		Key:    "{\"address\": \"61ee1caed7de6039648a41a8cba90a35f0a768d2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e7ab6d5b3e5db851db16a25fcda32177\"}, \"ciphertext\": \"fe275e67b80768c48f1938025355ecf79058c3e3be562d647bd45c6c375c9d7c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8591d7ac2b32ed3b038b4038a1b55fa3\"}, \"mac\": \"b3a655e65ae2d16fb8c5b407989ab8a9632f876d53e6fa76bd9c7e22dc2055a0\"}, \"id\": \"cab5e910-6be6-464a-9aef-d65650b2cf25\", \"version\": 3}",
		Secret: "sbsulw4laini1u7n46",
	},
	{
		Key:    "{\"address\": \"bd7461a40cbc63cb0929ce5975bbbe9241de0f3f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b39cd283a3c4f2c9acc778a72e662470\"}, \"ciphertext\": \"2e66bff2888208c49c91004c5c8ba5379b5a12749f1200fb426db8e42232ab29\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a97d7159979093bbcea2837c247b2707\"}, \"mac\": \"d2fe88ac21a48d434a4e405d70cdb3b8c954f890676d4bff55bbaa41275626e4\"}, \"id\": \"03720b49-fd3f-4923-8c36-61b97937bd0f\", \"version\": 3}",
		Secret: "bcm4cb3rbruavoknxy",
	},
}

var validatorKS = []*KSItem{
	{
		Key:    "{\"address\": \"859987a592576144877302c3868c5b0f96904a8e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"288191c8cb2fa91f6a2bddca00b13fb4\"}, \"ciphertext\": \"deb2facbb684bdd16186068b68d96c6a910318ef76702cdcaebc0d99de09de78\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"27ccfc92f751019d3eb0e259e10bc232\"}, \"mac\": \"ace64fabc150af9adcb3a0f45cb48c69d075ca629b0327c75b9e4ef96c5e359e\"}, \"id\": \"37f2d3b2-022f-49ae-8977-535291520f70\", \"version\": 3}",
		Secret: "jrlb3vq7rq98krwuzo",
	},
	{
		Key:    "{\"address\": \"b5899434eb21bfb2253edc48f8727266bee64e0e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e538e6759ec3c6b8864e58530c193f37\"}, \"ciphertext\": \"598e2af318a1392934bcd55dd3068c34142c0b406f21db1a24724e42412066d1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5211942db4ec04b0310272a500b22cb9\"}, \"mac\": \"c1e08f3530376b616be4b4c065b3683ff4b45d2e053c4fed0f59acbcbe7fcd36\"}, \"id\": \"753138d9-a1bd-453f-9579-7caf2c3998e9\", \"version\": 3}",
		Secret: "vmk4iezqjsb752vypm",
	},
	{
		Key:    "{\"address\": \"11be16468858152e611eb095ded4f6e0b0d104ec\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"958d5bd77154b6eb58eee3223ac0f3c1\"}, \"ciphertext\": \"04255321ae264389baa0f5fb6de9a110f7b02b205829dc78e1a3725227001cab\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0ce7a3a07fbe9956d0287463cf7e81a0\"}, \"mac\": \"a5b2f71ea483a11f78d8c50daeebac993076ef48ad9a6899f5fa4c23b4c6722f\"}, \"id\": \"5d62a6c5-15c3-470a-afa4-42808c558a03\", \"version\": 3}",
		Secret: "scsye93gpcmqsdys4f",
	},
	{
		Key:    "{\"address\": \"e56fc2c4fa1c82875f40078c7a8863bd12601ce8\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c0a50089857a29e88ca1f516af76c709\"}, \"ciphertext\": \"343796266e5cd4b9ae994300bf8b42140fe8cadca77950fbc3da5363bd04ea49\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f903ed6128e25ff7f6bfff370db44b27\"}, \"mac\": \"7a209121315b62293884ad3601a49ec12320ee1c97dac414a4bc78aa1d154302\"}, \"id\": \"ec9ebd6d-c49a-4519-9c18-8cfd4f9a84f6\", \"version\": 3}",
		Secret: "q6gu8tyyv19g8098ud",
	},
	{
		Key:    "{\"address\": \"132a4e41515d30fdf811b8635f8573781e51f481\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"97d8c8492b7b92b04a06ec6486ec4edc\"}, \"ciphertext\": \"d8beb99b3a0e66759f21cb7d2e2f9c8e3360caed7dd2e08a4b25bbdeb8b6fa31\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"fc847c98fcf8d1234e50e156cb897374\"}, \"mac\": \"32f10a6fc02b27cacff890a318f0bf1eee20ea48cce8b4abc89b0eb3181a82eb\"}, \"id\": \"7f7987d6-3ae0-4797-bb1d-3d0a943f348d\", \"version\": 3}",
		Secret: "gp5jn7nlsffe5rpctq",
	},
	{
		Key:    "{\"address\": \"2a4a1e0802e43c24225bcbecd3a395aed5ca5b6a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"eb15dc3b1c2059b9032e8cfdd11f1439\"}, \"ciphertext\": \"32d6610e2dfe2f3899d71c9cf232f928cbf6002351a9eb74be71c19e8832072f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6140f68301d13bbb1ad1bb3aec7920e1\"}, \"mac\": \"d0fab605ab8eb9c78e5d0b0277c593416295df39c775edb6b4b016c289ba73db\"}, \"id\": \"17ca467d-30e9-4fd4-b3f5-0cf413bdfd9b\", \"version\": 3}",
		Secret: "viuqbsscl5fhez62o4",
	},
	{
		Key:    "{\"address\": \"d356cdf2ce4eea071094735e8125de564233fd79\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d313ea3bd38a398470ac89b61c8c48c0\"}, \"ciphertext\": \"eb1cd19afd9deb97061c757fecba5633802771f91cc241434215aa9cefea4f8a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"126bef9e11f076b8c9bb7f6476d02b03\"}, \"mac\": \"da520ee6b351e774851537553f1c7ad2786efb65302fdba68746981d57c0dd19\"}, \"id\": \"83ca6548-bbd3-4843-820f-c766af7f5f40\", \"version\": 3}",
		Secret: "dnl7fue5ic1sd14yv1",
	},
	{
		Key:    "{\"address\": \"0657cfbac9dd7fec004a8a9f84813016e3f3b7d1\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8b91c292410240577b5acd37047d65c7\"}, \"ciphertext\": \"e3bda2cb61b25ac0de8869e8a2f6dd2e520d7c371e08ede1bb223edecfc473fb\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2971b4afde0286fbec0cc12b0940df20\"}, \"mac\": \"178683c485c783c632056c259acd76f9038e09291edf607648a4959bf5f222b6\"}, \"id\": \"4d11c050-ec3a-46ff-bcb4-51719969196f\", \"version\": 3}",
		Secret: "udefe9tmighpiyr65i",
	},
	{
		Key:    "{\"address\": \"063b93cabbf11c27648779d753fb245572e07e07\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"fa3bb04051f8aba831da413550d1eb78\"}, \"ciphertext\": \"087c928e83f14d5f2c85b6443b2735d117fdc0750d992c836530990415515765\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6cf07eb6e2ab34471d839555d5155508\"}, \"mac\": \"1fb98a1bfb5b5fb892716eadec93fe9fdfa1c163304807b9571871568423a952\"}, \"id\": \"c10e8f69-aa53-407c-9ab8-4553e1928fd2\", \"version\": 3}",
		Secret: "ih9zxvnaqk9llyz31j",
	},
	{
		Key:    "{\"address\": \"b5b9d23318dc3b0ded7a38fe220a9a99114eb45e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"341286edb66832b0e98755241205c4e7\"}, \"ciphertext\": \"1ebc43ae97bfc060f79a28021b707ece27de4e154986dc13cc4e5d9fda889c2d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a16d75b86972d14deed3cd6b83300336\"}, \"mac\": \"6b3a3cb47f7306ea7a357f738347873466c8ebc969db9536beb0be2776737209\"}, \"id\": \"bd2f0487-1926-45c9-abd3-936a5a4bd18a\", \"version\": 3}",
		Secret: "427q3m5p9pfgs44jyh",
	},
	{
		Key:    "{\"address\": \"cb83dbf0ed34d3a40918cff9c0cbe2d636a668e3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"03255619f5d5e19bb45eabbf68db15fe\"}, \"ciphertext\": \"03708272d477b773750ca4fbdf22dfc85add883e3677af98cc72f7b854ab3f2c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b5b38ba099a9af2877a22c89dcdde95f\"}, \"mac\": \"d7fc15faca73c6ea57ffe1e193ce22388430fb0642dc6db314150a79b6688f02\"}, \"id\": \"d3ca390e-c539-4617-b033-559aac18b54c\", \"version\": 3}",
		Secret: "4ki84s1x6log7wsbz4",
	},
	{
		Key:    "{\"address\": \"9ae56e7bb26e5f33339a0894dbfe69c135a5d927\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9d86914c0f7d654fef4fa54137e1671b\"}, \"ciphertext\": \"7c24eaccdd4f86773745051f12282606898ba2eb099ace17852a7aa9f0e7f785\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"df8fa32e930e98ae1251891e8c3b2ed3\"}, \"mac\": \"bb146578e55af755cfa50be5ea05d26fe852871a9ce86575c274793efff462e4\"}, \"id\": \"68a6ce99-a35f-4e9a-95c3-7e2af5e559af\", \"version\": 3}",
		Secret: "k1g59t51vrkyjaex2n",
	},
	{
		Key:    "{\"address\": \"c9601d76e2792551338b577320e9d92ba71dd3fd\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c1c394fe45d5f4d94c5dfb600838f6f1\"}, \"ciphertext\": \"fc927da542ef38630840010b11fed317ab4dc60b74d1725c23c88a12f3d8fe35\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0983fd8c4062e9516b459260f4584b58\"}, \"mac\": \"1e4bb300c6727bc957598acd303fd9a924c61c8e58c184a7ffd90c9aa51ee160\"}, \"id\": \"8699bb12-112c-4f1c-9d7b-e931a97c6b0f\", \"version\": 3}",
		Secret: "juixvjo8c14h4r58o4",
	},
	{
		Key:    "{\"address\": \"2d38b366c2c83d4bfb751939eee2e18b06d360ca\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8c8b6e5430518f05cf986c3cf688f3b2\"}, \"ciphertext\": \"3cede86d0b8a2d84205f7c715a625a8649ce56f9bb3a3642b260bd9e9355821b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"94ab43c99dfacdf5dc1adf0aba546463\"}, \"mac\": \"d14812041d7f8d9e1283ff198bc84864bbae8ff75cf66889965190808f0caa0b\"}, \"id\": \"2a95f94c-b448-4351-91d6-58c09678f90a\", \"version\": 3}",
		Secret: "7zf128dkht7mnjxjyk",
	},
	{
		Key:    "{\"address\": \"236b58a96564112c8c8707ea3fc0d771efe045eb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"254f3c5630b5d5bd2c80edd7e42ebe56\"}, \"ciphertext\": \"0840b87f4c8a5f5717492690672969d897ea54bb81ca1696ad51b67e33dac5af\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2b15d05812ffd77aa91f9bce45b9f98e\"}, \"mac\": \"7cb02622c3fdee7e96139d57b0a6286aad3ce31ac0b5fd04e33afc831c9ab089\"}, \"id\": \"86f00e9b-8af9-4d49-aab5-9a8708bd3445\", \"version\": 3}",
		Secret: "0epyeb77g5f5prk8bs",
	},
	{
		Key:    "{\"address\": \"f46b226ed4ffc4cbd237c9d54945ba1090e44aff\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e094cfc5363c751875acc7413fe6c2ab\"}, \"ciphertext\": \"b6538b30dc4ea7c8cf920294463a8c7a1698a97900ade0943df7f1f01b355388\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"01b4ddd2e403f9b1e0ab5eb59c5b10ca\"}, \"mac\": \"fd176f1163761eb2ffb1972c600a3de473c3a40a916bcaced8ce98fb17aeb0ec\"}, \"id\": \"0c71f98e-75bb-4fa3-9afe-9051877a4877\", \"version\": 3}",
		Secret: "q621hra6knnbfmkpt4",
	},
	{
		Key:    "{\"address\": \"3f90783a7ddcf86d1abb02c0f436230d66f210c4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4dee2cd9a2d0bba190adb7ca46c97d9c\"}, \"ciphertext\": \"14a5bb54398a9a773d96951e3a5d717fb69ff65e98dd92e8a54e3f99325ae71d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"17a2ceeb241824d7f928658eb5e67c37\"}, \"mac\": \"742dd1886ee0810ffc56b724ec864916648370b429d16367c84dde340b2789f1\"}, \"id\": \"14401aeb-f737-46df-8a70-ad385ec75e04\", \"version\": 3}",
		Secret: "uipacf7glmedeqdth3",
	},
	{
		Key:    "{\"address\": \"185a673c7f256fe667d80115ab28d9db92dd5f02\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"82146723a145fb6dc271e3f81b4beb9b\"}, \"ciphertext\": \"dba39cffaa9f93f02d32db005fd61a0a56d3d1a1f7b0665dacdd65d88ee64c04\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"73b9bde22c9319f0bced9f7e7bfd0829\"}, \"mac\": \"4f591376c6ae2b38117c39b2e97487cd7096e9674b99672d0a341fc04b660e89\"}, \"id\": \"1c6880e5-b9f6-4e26-bbfb-9970311b0192\", \"version\": 3}",
		Secret: "boh3zqi15uowrvg0rc",
	},
	{
		Key:    "{\"address\": \"fa0f30c6b03e161901625efb0da98b29f6e67668\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9fa049c553ec0f8a7073eca7c1c4f9d3\"}, \"ciphertext\": \"2b060d4e86155e021f73d2dde003e789f72a5270c3cdc847dcc43f9a3ab74e0e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"029eed2d7a95c25fc8e8a2573dcd00cf\"}, \"mac\": \"0b49f45653bdb5e0a182a74415aa62a9ec0ca65262f581c637b575a637679c29\"}, \"id\": \"8b8cdca6-f180-43cc-9472-4dbc1d6f91ce\", \"version\": 3}",
		Secret: "96yr7eqmgzm3if7yvj",
	},
	{
		Key:    "{\"address\": \"2dbabd6f12c9fcdd6f2754e899ee9428bf09e3ce\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d5cd8b759a81e28b252f99a3743592d9\"}, \"ciphertext\": \"7f193a34fe570f77c57a4417b2f6efd1897432d39f4b4a3c0eb986286784b1b4\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"37e3a5ed89dc0cc5b3d0ee0cb765bd1b\"}, \"mac\": \"c94fad9141dfdff5f31198768c142ae470caec47b3aeb83a6e0a7d01d61d6779\"}, \"id\": \"cfd392bc-4197-4a6f-84b3-4b40e1395b30\", \"version\": 3}",
		Secret: "vtiooyy4x18namil3p",
	},
	{
		Key:    "{\"address\": \"1365e45f6a2c61e4dc1363af5416aa068887c196\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"bb9a61f4a2be18c3e7391641d2ef98c9\"}, \"ciphertext\": \"6cba7c2b63bbaa1d2816a85a9465f2a66954161ea537cd27ae9be3bb3d9332a3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"65bbe0a79bb5c73cb7080cf4eadf171e\"}, \"mac\": \"2847720f761fce190c00d1e638679c9de2f13c7da26da5a775be5a0934670d2d\"}, \"id\": \"b644eb61-2dad-4629-a934-ad7b717c6863\", \"version\": 3}",
		Secret: "tnffq56zx8bjwuexg0",
	},
	{
		Key:    "{\"address\": \"6549c49cf970c32ee5881a04127615241707068a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2943d84adea7c636f02702ae75873f6d\"}, \"ciphertext\": \"b0213e4c57a96ff1b38a5c45903bcc5b7764603d7b9180939ab2f57da3021343\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d83d0a1d2c408d7a84638dd70c402fd9\"}, \"mac\": \"fcfdfc835b7a6796938be4147a872b4c190a47c8035029be9066a833bf9dade7\"}, \"id\": \"cc8dd95c-2803-4e5a-a33c-9b45ac99c7d2\", \"version\": 3}",
		Secret: "n9pyua7clk72j2jcc6",
	},
	{
		Key:    "{\"address\": \"ef308c99b0fbc58e3ddc66133fb73afb003ddda5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"67aad77870701c77cdc27e6330d504ca\"}, \"ciphertext\": \"a97300da053e9513900bef96ac5381da27698d706ac46f365e9496ba306b98e6\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3c0410b54bab195a7a0bd012135125f7\"}, \"mac\": \"ba51900a83bef103b9e08834f7b693910e191b1de931e997afef88f39877e022\"}, \"id\": \"2e8d90ef-6305-4153-9dba-0f1364f98cf4\", \"version\": 3}",
		Secret: "mif50asrnw14x9zpeb",
	},
	{
		Key:    "{\"address\": \"0c70bae855840eeac9b296ccc642ea41b2545d0d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ee422546980942d026cb14ce689782ae\"}, \"ciphertext\": \"98b5c2a97df3debabcf6c8bdcdc1c31f0f3d803921c21b5149bfb0bf8de8940f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e3988a27b2a076d7a14ba6cfbbeea47e\"}, \"mac\": \"bf962919751b713fd84abab7afbbce0ff2589d308289ae3a19b6b40d475f592f\"}, \"id\": \"013c47fc-950f-4f04-910a-2ba93ce0533a\", \"version\": 3}",
		Secret: "yat6wtf56tw2hkrlsp",
	},
	{
		Key:    "{\"address\": \"928077fd1b952b54427cf6c4e6d975d600272aea\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"05fc792942a4389261e7cf82e510c307\"}, \"ciphertext\": \"127806eb98c5e9f247cf92fe27f2c9fe70f271f1382070b06a93175b1e3934bf\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"78cf1adf0fce241a42b425eb30baf6f8\"}, \"mac\": \"aba6fdb0f71a63155e0add05f7e3865b45c8a81c1532d8ab59ec2824ba85dfc0\"}, \"id\": \"a0e06c72-be1f-4510-a9e6-72cad3da8502\", \"version\": 3}",
		Secret: "hg4c99utiaghfgtuva",
	},
	{
		Key:    "{\"address\": \"6e473405d86b28e5a5df07511033a7a8143eeadc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8a4ed476a02f94a8ed1e766018a6c912\"}, \"ciphertext\": \"1c93d0de08594f092dfa271562123b9d27498bf21a6ba3e781623facb9872c26\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5270888565dac1ab6fd16ef009bcfd59\"}, \"mac\": \"618fe51a6d14cdd8b13044b6c863df517f2d703c4d12caa119ea01fce76fa22f\"}, \"id\": \"01ddc9cc-fdcd-475a-be6a-19993e1e92d2\", \"version\": 3}",
		Secret: "wshetjbco4zt7q2uf7",
	},
	{
		Key:    "{\"address\": \"b20450f074a114c3ba22f76cdcf3ffe2b2460b55\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"436f347ce925e63e38af0361e45c0126\"}, \"ciphertext\": \"4c3b07d055fe0e52c137b35e8f8498047613c9b4631b470c514af4cb3abd1749\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d0850f4818544f2c4dfe9f4c9f691490\"}, \"mac\": \"3c4806a4058cd558d8e10ca7b0a735e16162c8678207614de923cde0dc3383b5\"}, \"id\": \"2bf9592d-442b-48c7-bf4c-f8cc5f1513f7\", \"version\": 3}",
		Secret: "75se4j39rgt6388ngp",
	},
	{
		Key:    "{\"address\": \"40fcaa77952c589d6042a94fc3717f475cb51405\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b746e2390e5b5a2489ee31eb7ed57988\"}, \"ciphertext\": \"5d592f788b8a9a8eca7a49b1c393e69e789070575522014f106753e4129e5362\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4cd572014b46da8b6326b4989f498aae\"}, \"mac\": \"4d1bf9397e73e05a53bf877b6f5c09dd3b76778674e9af2bfe74c7481b1165fc\"}, \"id\": \"23e61a1c-382c-4665-b347-b432abde658f\", \"version\": 3}",
		Secret: "2ujtj8aq7af3h5jzvu",
	},
	{
		Key:    "{\"address\": \"c252fb2f17ae35cb6f268b4295c50b6ce44a3c31\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3c37480de59333fa5e2a3737bf671a53\"}, \"ciphertext\": \"f6e43ff7740772bbd40b2ea1c81c842f2df1963bc12e4ee173fcb25596395756\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ccd1b5c43b65c5b9cfe2e6ca5324bc80\"}, \"mac\": \"134b529784818660286819450e93b2f27e17abd0a05b3f4fdaaa42a4199d2b08\"}, \"id\": \"0691aa07-308e-4b04-808c-c85a948b2948\", \"version\": 3}",
		Secret: "hyg2x6143416w3p2o7",
	},
	{
		Key:    "{\"address\": \"d0dc086a6fa0c8b915f75628ffa2775553d430d6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d05666459aad6ab9e2747e61a7232ebf\"}, \"ciphertext\": \"a0ba10e61b4f3e273a4d428a38fa76419a43de7c2d8f7abf0af321b6ece671b8\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b32e1ab4cf67552e442e94a0baf0d958\"}, \"mac\": \"d4d77495b4c237e7c200144cd7b2f9db905391fc652049004595a99c38f88ca9\"}, \"id\": \"dfd4c0a6-a025-4231-bacd-06e35cfedb38\", \"version\": 3}",
		Secret: "y4ksxzctqcbm8xfckk",
	},
	{
		Key:    "{\"address\": \"69db9bdc7e3b532eabdfd403a2432675c604913a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"970bf15c6a5a61f96a3f9a22be996f59\"}, \"ciphertext\": \"3b4ecfe26859ba2e031fd376373fe883a836860e807574418a947e2d7a8bb1da\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bacb2b97bdb073f75b1766f6e065a296\"}, \"mac\": \"feebcf4173d28ffae3cf33872a3a8cf6718d13c563122e155d6ec9f37f9550d1\"}, \"id\": \"9901d910-639f-48ae-9b38-c5138ebe8025\", \"version\": 3}",
		Secret: "uq4r4j6rxbae8hxvkf",
	},
	{
		Key:    "{\"address\": \"52418e23d9cdb526279ffa0501097fe15dc667a9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0aa650e1883b453d786ae185b6e680a9\"}, \"ciphertext\": \"111fa8d9356d449f17ce79b2b968e1816b68d4803bb8db0951aa39d3caad8bcd\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4c2cc12f1d93690195bf7aacf910f581\"}, \"mac\": \"895feb37acb86fc15762920396b741279677d8d4c0b0fe132bc7ff528f024f5e\"}, \"id\": \"8b828281-2568-41a4-9453-1e261ed10b6b\", \"version\": 3}",
		Secret: "30u2k6dmk3qwa35e0o",
	},
	{
		Key:    "{\"address\": \"8a7b22295125f0f5fbca7a2c2f6de92b3007f3b4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"5ec68ad753ff792c6b3857fd14c4b369\"}, \"ciphertext\": \"6ef183cd0a6f0c7373fc6eb2445199795a50d3e2d4555648b16e1b72954e47fa\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ac5581136165e64781e5d34b232c5e13\"}, \"mac\": \"6992fc868f9a7911ed91bfc5d6cac11770d4fe5983e7238f59cc4f52d26fe352\"}, \"id\": \"6d8998a2-bee5-4a10-8621-ec83c81d3566\", \"version\": 3}",
		Secret: "nm9tn78p8bkmiaxzwo",
	},
	{
		Key:    "{\"address\": \"f91dddb5ab32e82b2bb60f075620b4ef77f21446\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a282ca27ac700e6715baa975360006fb\"}, \"ciphertext\": \"a00656aca60c0baff21015650ba01144c169a52d2b0ff7a52ec8390474d85073\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d6804e105d0158420bc7010ba6255a6c\"}, \"mac\": \"aa271fa40d0f96537c10830e56be16fb1dc36080eb68ef28a7549c64a92a6e9f\"}, \"id\": \"aadc37a5-b70c-40f6-b84e-94383b77626a\", \"version\": 3}",
		Secret: "a5nwf56cgw7zkhxlo7",
	},
	{
		Key:    "{\"address\": \"232880e9e69e22163a2be497ada0acde1d527981\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"24f40cf08d7b29133d6702fbd0a2d9d9\"}, \"ciphertext\": \"42e3306fec3b9457d2b97eb88c9a45169b59adbd1917c4a4ba0c0326ee068ba2\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8791ec6a1922f48047272258732e0e6e\"}, \"mac\": \"d1bdc54a105829d584e92b59dceff7f2e3a9f4339a83edfe3ac517a2530607bb\"}, \"id\": \"e49a6332-20fc-4ac2-9f32-63f590b57f2c\", \"version\": 3}",
		Secret: "4hsi7nuywjb8qs53d6",
	},
	{
		Key:    "{\"address\": \"c501299027b70349edf4a9ffc150ff52229ea95c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f5b7a0ffd99097b392fdf89d2b7ff917\"}, \"ciphertext\": \"fb1934d8730390257317612a744c026534170a1b5ebb2fdc9d25b00af7f4d64e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"76d9cbe717b9d3feb220c6dbde55801a\"}, \"mac\": \"ff167fb3cdcc3383686d46b36c35bc0eadf92c929123b7fd6371b9e24c9e3502\"}, \"id\": \"b1400282-e236-4119-8108-7c84458296a4\", \"version\": 3}",
		Secret: "1wrvxkeerpgyieqw0c",
	},
	{
		Key:    "{\"address\": \"b05518d950c15246f1e0c239ef4d420eaee029c6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6c5ee055546b552a04f46eaab48342df\"}, \"ciphertext\": \"d3510ae17aff1729c7a10e69198df82cba431c850d727dd88da87d3f2df79001\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6387aa1d528d605d0c4cd514124a9dcc\"}, \"mac\": \"70e71cb49efebe2ae4a4835205a2fa4ec647c21147ce4a8c57bc6b113378f0c4\"}, \"id\": \"c3087c6f-1323-4c14-99eb-192629306e54\", \"version\": 3}",
		Secret: "n1gaal2l9jw8mgxcs7",
	},
	{
		Key:    "{\"address\": \"468951cec58abda1a59653e91607bb8cb204134a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6cd4961a706f01aa1d9dd645ecef2e2f\"}, \"ciphertext\": \"90164d72cca415c6a6309fd57ca0e32da4a25c7aa8c535a38ab33de7088e4c1c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d9b140c007fb25794b5903fed98e299f\"}, \"mac\": \"94f9b002378fde740ac3ad0f6024ca2ab1cd51ba62681338fcba0e37b3a27f4d\"}, \"id\": \"71ea3468-d64c-45f6-b943-0457b3840edf\", \"version\": 3}",
		Secret: "ee9gf4cn75re1ksbd8",
	},
	{
		Key:    "{\"address\": \"776328699e24ca013a61c91d602e8117df154219\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a8bf0f4b1ffd766afa5b4288a182b66d\"}, \"ciphertext\": \"9227dc48b3daece3f99de39605c19b7fb741369ae220e73e78d42f539c89bf64\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7d82669896ea2f8a12f0ae1e61c6547c\"}, \"mac\": \"43c98c061b82e86d64bcd9f689eb4820f5f8274101bcc492b445e0e4afb66b24\"}, \"id\": \"04631585-7489-4bc1-93b3-8da35b19ad29\", \"version\": 3}",
		Secret: "n8e6ypk05yy96hb8dd",
	},
	{
		Key:    "{\"address\": \"48e6b0415d40b6b80b0f768b0b8f40d704e46ccf\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0f3a34d62c72429432f0459fa5f1f31a\"}, \"ciphertext\": \"8da2ab6691f87003ab682e7568a1a67692e2edf848233d731175ec9b77cf4bdc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"09d2c7ae31e230128eae876d84d3de7f\"}, \"mac\": \"67f8c5eaa673f0a1e686c5a55d0be95c398028c0093749580cd6903b1aed36e1\"}, \"id\": \"66488434-4545-4fad-a848-66279e60c2c6\", \"version\": 3}",
		Secret: "ke59vxswxern321jdq",
	},
	{
		Key:    "{\"address\": \"6c5c76fd03c74d60a387c5b4f2d034649c438418\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d309b9ade804e5147a693c3a1f6b0ae8\"}, \"ciphertext\": \"264c8cec0923009a4c1a67b9e0f17895a42b983629882288cf5ad0393754eebc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"053d666356b8fd5ab489bf61ba6213e9\"}, \"mac\": \"0dcd864ca86abb7491f5b8fe3941d5eae8b8d98772d899e32297e4b369a04839\"}, \"id\": \"9cfd0cf4-86dc-450a-af88-0309acf7dd60\", \"version\": 3}",
		Secret: "q3xsmhtg32mpyaz6tv",
	},
	{
		Key:    "{\"address\": \"885fb2e5e351510f50ffbc19f553bd029e70d99b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"379938bda2cc4955aafd487af0099953\"}, \"ciphertext\": \"20ec4c3db464cefff60b140bff7bccc85c858ede9fa34f7999d3e1f426557940\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"abcb74c590a729e0eabc7aff6132a9a8\"}, \"mac\": \"a133bb489b68ed5da14419dd2f591d599f6a05f33788d133e823e9b73df3a4ea\"}, \"id\": \"2f0202d5-2b94-45b6-a9f5-ceffba852e63\", \"version\": 3}",
		Secret: "udi728orsh42o7n7xp",
	},
	{
		Key:    "{\"address\": \"2ea12795d0a5621bf967d031daa8e793332cd01d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"929e4e11d025062135e1e5ac1a27ce61\"}, \"ciphertext\": \"ebafe7cdf313c8b7a37d3feadff9a4a6e9e322e09f83de1a96aeb6724d3c763c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"06c92e9b2a8bc3e556a9494ebb6b71cc\"}, \"mac\": \"ab822babf4e4f55ae0f8ae93fa5c09222a10083278c1e931e853cc41ef3e3f7d\"}, \"id\": \"ec4046e7-21bc-4d21-a4b5-523400bc04e6\", \"version\": 3}",
		Secret: "goygta4rmfxzdtls6r",
	},
	{
		Key:    "{\"address\": \"aeeeacef0f2e1ff6d6f1b01f5f32d48ba5d2920d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3d141ed7618b1eed9f17ee359395ea94\"}, \"ciphertext\": \"381064b7ef6dbf5d289460f454c4f564a50661cce930fd299ff5511e00f6c0eb\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"88fac5c7b5e2123e3707d30fb1471975\"}, \"mac\": \"4ba6767a22c1c8d314255599c0e4369fdfe620edcebc75e3c7dd7179135caa42\"}, \"id\": \"7d765a0f-634f-48f4-980b-46adf36cd8b3\", \"version\": 3}",
		Secret: "20rt6xg23swg0lbwu4",
	},
	{
		Key:    "{\"address\": \"e80e3853c2715de72baebd60e7ca1671e46cb84d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"074e94ebeb00eaf36f36eea9e8c81270\"}, \"ciphertext\": \"c1bd8b9c6d5f8d18912bf10fa2f74217cdcaa6a97fba4616b6132a52cafd3866\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3272eb6150adb90e938d3b0e22a3413a\"}, \"mac\": \"20a6327dd12086dce0e689f62cda426e1a90fcf78e9f19359387794b88053f51\"}, \"id\": \"e55c2120-bca5-4b62-9f3d-cc9f56329a45\", \"version\": 3}",
		Secret: "ot29s0d94u9oo90pkf",
	},
	{
		Key:    "{\"address\": \"eaa97bc1d43da3a257069976f3c94ed6168c39ac\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2e34c32b0e19721c17e6c1b34371476f\"}, \"ciphertext\": \"58b2d856ab9da244041248770d13ac26fbff4f2668f52629d83874ce857b5f58\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c054141413a96ade928254eebaf5f771\"}, \"mac\": \"0b03d9ad8687de8b7326ceddd8b4e2a352a86fdae26303accf20526f8c129832\"}, \"id\": \"5f15dd50-265e-422e-bb32-418be5ca0a41\", \"version\": 3}",
		Secret: "s42nw8k4rx7u4mkomw",
	},
	{
		Key:    "{\"address\": \"42f196d0805eedfb63cd77ef6eef8a4356427e2c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"281838846d72f854c97872c4f15859f0\"}, \"ciphertext\": \"61174447d3ac5b3980d9e45a082458422e785c825c33b0e56f8a916ae5d1e7e3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"466b889187e56f5d88391e2464c15633\"}, \"mac\": \"126a79a4035b0d06d926737a13fed4d4ed43fa138f72c473813fc9517ad5786c\"}, \"id\": \"fc102b83-ea30-479f-a025-10950f6b2fe8\", \"version\": 3}",
		Secret: "66c8fixta9ks8p064v",
	},
	{
		Key:    "{\"address\": \"8276c96ac79c4d3423a8ef634df9d2cb1509fa6d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f9a173e00d4960119596b68fede912a8\"}, \"ciphertext\": \"54c19a89c7c7731fef3cf58f31c99257c38961189ce283a553da4ccdefd671db\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a4bff89dc6fee7b75c8964e9c42cb671\"}, \"mac\": \"e4b834bc3308fce982732ea1b0cf22fb806c517d6c890bb231995fccfa789fd4\"}, \"id\": \"3d458266-bce4-43bd-8173-143fa50c6e13\", \"version\": 3}",
		Secret: "t5wx5i1ke4lkphtpd7",
	},
	{
		Key:    "{\"address\": \"84a3e0822715b13d1a9d278a6d4f97e1e90e761d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"40b4ac8e41bac1522e4df0ee3c67c218\"}, \"ciphertext\": \"74bd785afd4627deb452fa4ec6d16163c2159622b9986169cee505f4cdac538a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"936b186dc41896b97a4dbb0b752468bf\"}, \"mac\": \"71d885ff1b18f791ed5a03f3f716458ac97f9114a5f352bdea0d88d0e2195da6\"}, \"id\": \"ee8d7ffa-5bee-4827-ad8f-36c1ce564775\", \"version\": 3}",
		Secret: "e1gmiw67e1wv8eadlo",
	},
	{
		Key:    "{\"address\": \"40a8ba52df1cadcd0e09af7b668ee823816f8c6e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f421e821278fec58a0b853c773decfec\"}, \"ciphertext\": \"d9f84c9387c1ac32ebf832e46d4788deeba6b58fea7e25cf8534d8a632a0f62c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b62370979dcfe6bb10d6566a6aaf37c0\"}, \"mac\": \"dd750d58aa8334b0d7ed0553d445d8d5d13ce849a9c529b4b6702d82b7c14aaf\"}, \"id\": \"a0692b09-400a-4d39-963f-4d2c75788bc9\", \"version\": 3}",
		Secret: "4lzm932ouip1ycitte",
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
