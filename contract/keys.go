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
		Key:    `{"address": "052e1b586d44c23739d028efad5b88cca010ac5a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "39c7336302f5c6fa99e6eb381b130ef3"}, "ciphertext": "fb33cc60f898dd1e6c1d7735ed4e3ea2365f000e5284f05f783ce24f5175fa11", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "f5665920b74febae89c2db2b1cb56228"}, "mac": "cf84492af914e3dcf2209d68572628202a3105ebe08822a68c5faeb62a56bbc2"}, "id": "6e9d267e-e96d-49aa-a43a-5e8422bf20aa", "version": 3}`,
		Secret: "35d1xll0rvvpcj5icj",
	},
	{
		Key:    `{"address": "1e1bcf371be60de590e492ecc4ae8e1122cedcc8", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f6a459448750bcc7f07aca24249dd44d"}, "ciphertext": "e12f63a31622069acf7a9b227e5c157c598d94cd181009db12a99fd985ebab29", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "df25b589d7020601c364c4423fd3475a"}, "mac": "b38a2ec812eb6de209775c0133fad6968cfc432814c6d8b6789b6f4443340371"}, "id": "64cc959b-2b4f-46a1-acfc-80b9553fd5c0", "version": 3}`,
		Secret: "j5euchb6nzf7cq7267",
	},
	{
		Key:    `{"address": "60b22116cb07c581d50e5a53710a0f69b1a6b8d2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f5cc2b8dbb9260e3c1857ad19ce8afd2"}, "ciphertext": "b82272c0601397668f97bd58df3fabd9b462265299ccd4dbd3f58f4503d864af", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "33d23a1c3f47b675e6c74ac54f590db1"}, "mac": "bbe0e647d90e593ccf40001391990e4987f75e7abaaeedb650d7ead1e1b13077"}, "id": "3529bec6-da53-477a-8cbd-624e31ae00b7", "version": 3}`,
		Secret: "9s4yts4k6izf4qx6mq",
	},
	{
		Key:    `{"address": "2aef6138985a31bda7265c937ca4bc5806aaae23", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2b5e08330533c9a04cde1e7b63fb78d5"}, "ciphertext": "211622d366b8249423bd0f427cb7077549f58c5000fe597f69193e9f48f54e59", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "57fe2f953ae545fe90380e722f717b21"}, "mac": "fff0055a628657a0ad7c378baf8f4baf1b9ca9d6e03f81cf50d52ef0c2c1423a"}, "id": "978558b6-8fe6-4de6-8374-96bf4a458805", "version": 3}`,
		Secret: "07i5j5z7vukusiohjc",
	},
	{
		Key:    `{"address": "22a00f46161164c3caf014e7a43580b48a97e122", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "37625a0401d683caaf751860143625eb"}, "ciphertext": "8ff1a05ec606051d35fb52a01846cf28d2fabe5a798c48e8e5916753f3fafc79", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "0d6d67c52d57f2e81d47f40d81434162"}, "mac": "fa1a082115f1c6a06bd658701258c27aef00e0317835c99c671126e83f1fe23d"}, "id": "4d8e9939-6082-493a-b569-8b8bff3f0743", "version": 3}`,
		Secret: "5yekuz81z1hbj1lukh",
	},
	{
		Key:    `{"address": "069a774f63ce6ed792243c540460b7c1d50081cc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "4806c32a4d7f452e96997e3bed1b8b9e"}, "ciphertext": "39c10363e79d3520d7d66fd026f8bcaa8595cf1d402b6007dc59c203a42c7be0", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "4d4c256be0e10b27c0ca30de9dedc3c2"}, "mac": "0251e39505376a68a6d3c1afc87f370474706b7818d872b870b0887701ed88f1"}, "id": "7330d428-5953-4ac4-8980-3d07cf205a8b", "version": 3}`,
		Secret: "fz013o081n165jg0gy",
	},
	{
		Key:    `{"address": "c74f27736128a5c907b37bc82ce9e67b14ba45d3", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "320fe5bc280618dac4a8a3cf75587808"}, "ciphertext": "cd9de2f97d4e96c77c1cd895be157d46ab5771d14c7a9af8f1b603cb106f9f81", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "1acde5c7e83f7d25c668711de704278c"}, "mac": "524baf29ff14d3fef89a07e1010e4fabb5ec407bc09fe746ac3af1a716fe8860"}, "id": "7623a094-20de-4226-bd5c-726aa2e6d449", "version": 3}`,
		Secret: "el9a498tk2f55pd22r",
	},
	{
		Key:    `{"address": "b3871b966b0ddd762cd208db93ff3aa390ec8c68", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "9fae1c6ef0f11894dae1f16fc6899c1f"}, "ciphertext": "6ba353fa15ad79097d05ae2df53953d041ff10e840b3bac341f4e929e504bd79", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a28645483b4865fc0dc21105fa7c2809"}, "mac": "16934116b95fc03568ac32634e7f4c32c859df853e203bafbb059c7be2dd4d36"}, "id": "ba0175b1-0c20-4e32-ae55-515e7e427910", "version": 3}`,
		Secret: "5kkojp82g3c34q82pv",
	},
	{
		Key:    `{"address": "e0980322228f7b25f2fdd46f8679b3def037a596", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f61948b4940338fc6c6c6b7ab4e5f39e"}, "ciphertext": "8ed48119e7fcc0dc2eebd9c1dd0a983cfe88cbb00078ca2d5cff51e4c6c46f8d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "abe28b02e42f8775a4bade8617555034"}, "mac": "3989b66c2e3026af975487f53b4e265d97e1c75f04103d92e3d698ac60980415"}, "id": "bb6e857a-5274-4727-96bd-1aad7a784683", "version": 3}`,
		Secret: "pjyc9dzzefbd5jwtcc",
	},
	{
		Key:    `{"address": "8e0df77eed73bee04af2940de4a16bd17d23cf2a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "cbc92216000c464096d4d804f3fdb5f5"}, "ciphertext": "a926f82ec4d0a32abd515b7c04c4729754cec262a80ae9dee1c030f8c4580124", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "31fcc84e20f7ef0fde00acf7ceb27350"}, "mac": "b4c43a25da8e174464ad4caf0c18cd32abb0530a5cc4e572b6bd9e376d096ca4"}, "id": "517cdec9-99c8-4ee1-a1cf-6c261ac27057", "version": 3}`,
		Secret: "8ehdafjrcfcmi4zrbu",
	},
	{
		Key:    `{"address": "557c09bb90fab17f2019f9885f9834aa6505bcd8", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "98f00ceb9a188cc78075970ae91a7dd3"}, "ciphertext": "f9f1510eeb013abd3fdae3249fbe4d948ce6058cf6207d63251b4a6d386d4dad", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "43c4611c3b9fd6ae79234373e5a37b20"}, "mac": "ba9c900f2ac7d3efab698e26f6a32886f4926a19868bb06107fc30103e469a3a"}, "id": "f4d85a45-1844-4822-bc41-d1e464973898", "version": 3}`,
		Secret: "4cc86uz2s5ohegiy6v",
	},
	{
		Key:    `{"address": "a8010c3d03dc4cdbd35b33cd9f313f5bb3c09702", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "b1a616c5aa9c87332851967902d3f525"}, "ciphertext": "5f37771d456be2ef24b0fe513f11a7a24dd2550847e643a19a1341ba1705b704", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "e6931bb5123c92d901fe9cca4796e266"}, "mac": "17580c55399d8aff47605d59ab38eb0a7d43e06e8cae3e6e2805c7b613e924e2"}, "id": "7301669a-8c82-46c9-b1eb-0696548da183", "version": 3}`,
		Secret: "h7x34cqjpr80itkik9",
	},
	{
		Key:    `{"address": "96beb1a46cb09799d8cffd6ef1ea4a4a3d05458b", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "6bf8211460879aad84dea70e42964cc5"}, "ciphertext": "5346f39d6eebffdb5f102c3287d5517e9b25328a32488caa7ed0c55417874b79", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "eb217e4984a68227f0845f8e353d5e5d"}, "mac": "e3cf6a66fdbb2bece17e8b7164db2de726995a17bde8a1d35462ce3a311b4cb6"}, "id": "f34624d1-2c55-4b48-8259-217842dd4db2", "version": 3}`,
		Secret: "cb8hi615u1qu2tf901",
	},
	{
		Key:    `{"address": "a2979300d7b5c4b97c660908784272408bf7c283", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "6943b3cf3b1809844be99c99987d0259"}, "ciphertext": "9b57063b7b015ae870946e20876ad685a4db35e0be5deeeab9c7038f46edfe17", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c1398ad1d1a63a313a58720d93bb562e"}, "mac": "57b04f95b2d782ca31e37b3d3e094b506c718079adc39bb36192205afbad9322"}, "id": "ac224acc-5478-4950-b11d-d97cfa5f5496", "version": 3}`,
		Secret: "ri4pbmtmvv7f5u5x6j",
	},
	{
		Key:    `{"address": "f7e7c7df5470a8384658a04dc20660c0010a38d6", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "01baa24ea86105e9b21acb57b24df15e"}, "ciphertext": "9c540d71f90c9e4a1376af4628756b6c270b7a3dce76c09da47b02c9869e80bd", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "b41f40e8daa166ae7a192b76df4da1b5"}, "mac": "46b5ce82259e44135dd69f280546bcf734259582682a433b604859d8f50e35a5"}, "id": "c046b5be-5336-4b21-b05c-056e28b3dd4a", "version": 3}`,
		Secret: "vdlv6oj2qjm519ezbv",
	},
	{
		Key:    `{"address": "151fcca277d01001a83f72893e3225db0ad47b2a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "c4482c4ab9fec53b284bc194fa4b36eb"}, "ciphertext": "887d3af6f41a4896105a8603dd62b0d9358ee40b9e8a8141dd9ca7b512a2e7e4", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "0b29a4b5fd40a3301c83cff87c888de7"}, "mac": "825f7926e35f48f7c3f553e15f1523eb5e0a5fe0ba4d62f24a6317e354fa084c"}, "id": "b4f5f51e-77c5-4eac-a758-28c4f384a619", "version": 3}`,
		Secret: "vz0qjxug9vtwn1a598",
	},
	{
		Key:    `{"address": "3af6ab6911f295ef6d7ab51c826c3f90c40a13da", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f6dec093163afafd7e561ff31cf883ec"}, "ciphertext": "dd68bf800c87543df93f0ad2a809477158d279830a3705b33217e906cc2eb513", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "01f0239f7ad50f8db4616baa2ba0027a"}, "mac": "7c036164da727e038137d263ce305a0d1fc6f99793ee20a673129c51a6cd547b"}, "id": "d07ff32d-af19-44d6-ba94-e1156247c5ca", "version": 3}`,
		Secret: "4go1h2dt6wjduwyb5l",
	},
	{
		Key:    `{"address": "712b8178ef4023ae7ba300554b0839b15531a2f5", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "86d63bf8d663dd08bad65119656657f6"}, "ciphertext": "5848e6de61413265ee275b7ba3f903c8e5b4130414a915005f9125aea4bdc969", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c64930a0027ef2be24e86dc08e26a313"}, "mac": "7e6a09a1df75217a2ed2d0897ef2ec4cfe18d4d76faa2370e80b4df325b844d7"}, "id": "953e01a3-0788-4e69-835d-5737ed2fe32e", "version": 3}`,
		Secret: "bxd4rokdo0q8ze121m",
	},
	{
		Key:    `{"address": "3a41e70f9554896e48be091d61fe6c8d0f47b870", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "c845c7e6ff0798132f1f5a038d45b19d"}, "ciphertext": "c239112905774703cfbf9a78bbf590f0253356b95d7c4636035b4ec5869b6c7c", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "9121cb5becfd25ad3b7eaf685e1e7087"}, "mac": "4526ed02459f4426c4fcc72d9b37da7b84f4cf22504675953c631098a8c77609"}, "id": "de0ced7c-5ad3-4a56-81c5-680fc620a6c6", "version": 3}`,
		Secret: "3j462zjkkk7nhkrap3",
	},
	{
		Key:    `{"address": "4f1689d078af66d504a2ff9ae75e5a0fe757dd64", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "19d82c7af1e98bea270df007793601c2"}, "ciphertext": "c28efb0c4dcce5474ce84a99cfe96f408f27d1b17106ac0ba7d2efc290538c50", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a21c1ce2b1cdf76ac7cc80b068e924d7"}, "mac": "faec3b45dba67b23764d4a66078217159b03f14632f9b13b4ca7a7932f297a46"}, "id": "1eda15df-dccc-4162-a8d4-55296e8d6a85", "version": 3}`,
		Secret: "150xbf0z1whj4nukf3",
	},
	{
		Key:    `{"address": "716e097ac85bba05e34210438d5f5d61e03f2fad", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "eb52653009e69115e687056be3cc6482"}, "ciphertext": "f55ebb2cea8dc74ede6923757ec7a029ebf6b46da13db9db0a2721733336f71e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c2e1941c904f83840b29f3190befbf79"}, "mac": "b100df21bf35f0e79e9383a74d6687d8db261a729b83962e137c4e9e6af6e505"}, "id": "87463e76-0c23-4012-9281-767eea7bfb04", "version": 3}`,
		Secret: "5tc0z6go9mag95t14x",
	},
	{
		Key:    `{"address": "5a96edd6f49af215b68b769ab2623ec6b529a53c", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "6028bfc13a2163fd42c1c6236276559c"}, "ciphertext": "e013a433eb98dfc22cbf1ce16b9064f977ec8c21a4fcd4751a356ebd6fc872e1", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "3b03681d93238ec054c6287d1b3eba99"}, "mac": "20adaf89e91f4c81fa5273dc116f4f8a87f01560845870974fd591b02df8d911"}, "id": "06abd565-8e04-4ff1-9fd9-36fbeb38dbe6", "version": 3}`,
		Secret: "wafa2oad5fepx8rxlz",
	},
	{
		Key:    `{"address": "d91986865b05f36a0e2462815f034ba9ed83abf0", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "caff5033dbf07cc30e2ad2cf762a6389"}, "ciphertext": "4ca85b3735395c560ff702acb10b88189ee71541650b2bc02f36afd0319c2c7d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "06c783031c2d36526306700424900064"}, "mac": "046742f1df15752bf9a11cdd39ff83db2fa97386b0ee68376f6c5ae71e3e2af4"}, "id": "842a306d-b2b0-4e5c-9c80-782026320f62", "version": 3}`,
		Secret: "q9863ke5qlpwzsykfl",
	},
	{
		Key:    `{"address": "5fd93fd318dbb02601d3a13bec501185a5a70ce2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "a310dc8f5ed8dfde58eca3c69f71cd4c"}, "ciphertext": "b5871073e7c5b5d600c9f72ee4d6b15e7d9859f41c181b75eb57f9675774dbea", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "4dc411fb378b218904ffc98cdb7bde8f"}, "mac": "3ddbef355c45d2aeac3dd4d4c24b0f83218aa2b25c14a7ca2c5e0ddc9fdc5d68"}, "id": "e65ee568-6e4b-4338-bd17-caf4ed00d9f7", "version": 3}`,
		Secret: "ndb1iwmu7z6edlmqym",
	},
	{
		Key:    `{"address": "fafc59e08e8e3be60724fff882330c98a86a189f", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "50dc5afe1c1b5560e51cae841013f2"}, "ciphertext": "27d45c807498b9c4519dac9ae027c250cd3c785afb09c755a6882f60ebb13a25", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "35ed4e8e18ec53bfae0ac40a6b9abe0f"}, "mac": "9b2a3b49f41863e2a7b1579affc20a3d35a17e15d9d28f38df42583e4abd3fc6"}, "id": "57eebe75-f383-4cc8-99fe-5beb01370a90", "version": 3}`,
		Secret: "pja9dpnaz0szu1a8xx",
	},
	{
		Key:    `{"address": "fc7100f811f831948770a15f48110ef064da25c8", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "4f82af7a8feeb3a5bd7ee4c4bfbabd27"}, "ciphertext": "7c4ad2378a44bca8c70629b40d4e01340cdf2f9d8fcf57650772ca8d611c80af", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "99a9632e1c7348eada6441e4c8236f97"}, "mac": "fa8ab44c766ad8a1a0f0ef8588bb3c0e8adb2dcf2ae79a07fa2687a1563490f7"}, "id": "2ba8abad-c2cb-4cfa-ae7c-80e08c12f482", "version": 3}`,
		Secret: "qykbjzvn17o7toc3fm",
	},
	{
		Key:    `{"address": "4c21d2ebdb84d3390ae37b9b4b1776e4c2808a27", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "c30bf496dc7bc9841baced5c0040f8c7"}, "ciphertext": "0982cae52a456e6a5bd69c20b27ce86a8c470b415278cec2266380946f2e79b8", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "7c110d0764bdae77e558c752455e73c0"}, "mac": "2e030596fe84ac8ce4f0db020b76f06b500eeecf131da90a1c6dc388778ae51f"}, "id": "78993afb-658e-4e7d-a5f1-d66cb2ae13c0", "version": 3}`,
		Secret: "ndm885po26nsfxkfxe",
	},
	{
		Key:    `{"address": "33af99758aed8c60db050755f044ab85604ac0dc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "e31ed85bbed4b234710f44d902d44e62"}, "ciphertext": "a949d079896846a1f561c98e1f7db36bd73579daf56c281fa944d2a1167a300b", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "444f9ee6f62ead60dd0b4f6497528fe7"}, "mac": "cd2aea8224fc15f75e8f2ed0c841e0df74890520aca3c5db2fab60e0b497fe24"}, "id": "888749e2-f219-43e4-a240-a801f9dd8f4a", "version": 3}`,
		Secret: "61cq4h4df0xt6uvvja",
	},
	{
		Key:    `{"address": "b57cd623968538ae72b03a0197cc47ec5de72a5a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "7f42cbcb577b919b00dd9dab0add2375"}, "ciphertext": "680484c2038f1cbe219dc7e4a36c1a092d525756e55c14bc0568b59294e8377e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "ada29bd3e2f2b6d582580f801a11f25f"}, "mac": "0fefc05389d6b1ba1a8dc3bb39c2be57a0a4bf09af805a1974dcf2ab578d1d92"}, "id": "d962300b-10d6-46a0-a622-b66b7054b5f6", "version": 3}`,
		Secret: "y1z7mi8yhldvakhj3p",
	},
	{
		Key:    `{"address": "74ecc0d46c2e8c6ca8fc93da018b0426a43cbdd0", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "4d68effa668572dbee76528237b86d8b"}, "ciphertext": "67b2fec81bffb881cadb67ff56266ea1e4e9c27568a946e4d64b73aaef0286ff", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "91765a5d82b960336fc55d5435c0f219"}, "mac": "c4855f10774e9f78220bd3efcbd0ba719d794883d2185d7ed4eb7ca233b51d4b"}, "id": "de026953-4947-4bb8-bd3b-3bea192fc113", "version": 3}`,
		Secret: "uw1u9laxx0gdkq4oq2",
	},
	{
		Key:    `{"address": "57af78a59d59fce9cb8966e4e70eea2a37480d29", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2f3688f46c86c41056a687748dd870c9"}, "ciphertext": "1cd13935bdee18065c3dd44e89b7f08041ff4ba257f00c4345956d08a98fbe9f", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a9fc915e5a0c7bbd80f0cd672e61fb24"}, "mac": "bb7999eda3b5e0fd63fec46449995853db5560c03a51df553dadaa4c88d38d16"}, "id": "ed7b94d4-d08d-46c0-a227-0ef3bc2c5b78", "version": 3}`,
		Secret: "d0h7o4cheiqyaruer2",
	},
	{
		Key:    `{"address": "c862e7e0432d42ea6118787826fbbe53a7c411ac", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "7922f7edbad31af7254f51af488e34f1"}, "ciphertext": "742cfa07047caa9410ed6145fcedb3d03771b1d891567a45e733c35c4042e3ae", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "ee8f092d75ee8c7de39287935931821a"}, "mac": "c55ab2ea866e9381fe07bb29c1c5ccac66ca366a3ad45a15c4d3e73ab9734a00"}, "id": "2f4ac0d8-8937-4dab-9767-37449e670fb4", "version": 3}`,
		Secret: "tad4iqw7qwcutjjy5n",
	},
	{
		Key:    `{"address": "72d15281d4656017108eb27b3577c0e245ef5311", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "e79ecaa57c61c2675abf632e2387809f"}, "ciphertext": "2dbe465c5d25d005fbe36cc6fefedc2e7969e95f881b9b0c6f153627a88039e0", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "55892f5cc5d3b6d31d01a70747fc08f8"}, "mac": "5a2cf190c3762b9a3bd50708359ed76bc9318c78e5dc98cc0782f1839f62d8db"}, "id": "5aa4b8cc-ebee-4bfb-9ed6-5da653047b42", "version": 3}`,
		Secret: "lthy4qkh6qjif877ni",
	},
	{
		Key:    `{"address": "f827a02f4833aab25dff492703d58691256c53c7", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "64456a638bdf110124c651fa2e20756f"}, "ciphertext": "b6e15cc5610d34c32ce55d11def9c57973839cb880eeb493839475c91cd5a90d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a05c92541511b582e84b1ae3cf4df958"}, "mac": "3be543727ca2dbef00246077e52dff242481149e6b4903c74eaed21ff89b706e"}, "id": "bf8cd60d-c850-4f53-9e04-5eaba8040241", "version": 3}`,
		Secret: "vv937xcys5rbizbgws",
	},
	{
		Key:    `{"address": "82f7bf819df0608d64bc9f153daa9c84de3c838d", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "0d9f26f830b789bac9669703ff574fb3"}, "ciphertext": "83989c75ada937527908d624164301a279fac71f15862e171d37756b7894de7f", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "5dcf621cb868a6edbce08c938fde4bb4"}, "mac": "23ec6511e8f889bcfacb75b312ad615bba47605a599fe0ee9a71c64d09bce706"}, "id": "232b7337-d768-4387-bfb2-761d783ea8b3", "version": 3}`,
		Secret: "v6h1fkp3bdgq2qr1oi",
	},
	{
		Key:    `{"address": "022045aa2e33773a425fdcef48c2697c6590e1db", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "532cb6cb573115ef2bd5e709ac1eb87c"}, "ciphertext": "73c704b6eaeadbe39a717cfe4dc9c1cc6071ca33ee774d82e8612149393f280c", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "f04a7962a48bd0285d48503231b7d2aa"}, "mac": "c7500b7a477ea2338865eac5ad8946d9515511246f90175da8eeafcde1048944"}, "id": "75004fed-69b5-46e0-a52e-e2020d07ce2e", "version": 3}`,
		Secret: "du2c2qqgggvq9tiig7",
	},
	{
		Key:    `{"address": "fa64d133313b844c5aa9060efc3143ff6ae2bb73", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "ace922b22da03676d10020b1aa85a537"}, "ciphertext": "f84a3fddec6eacf0f0d9e3c433ab1d95c0d7a80cf3c3393677bcfafb120cf709", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "54ee02133e822f702167dfd6bac4c075"}, "mac": "69947883ed4ee88b288143ab96a048793eb298056efdbcf5288ad49f127e6d58"}, "id": "5e226cd3-ef15-471c-8294-38a47cb25ed6", "version": 3}`,
		Secret: "9fg7jvlinhz0jc4bqu",
	},
	{
		Key:    `{"address": "cd836e52ea9408b18a9bc9efe3dbe62dedb71ca7", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "5f75410ab5244886abbb71e70ea8707f"}, "ciphertext": "6413ae56799036a042ce989e03a5c83a3669202a16b9058b8516a1a029aae325", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "aa0d7e0cd16f72bd912bb9b012b13d31"}, "mac": "b63125835532bd2fdb4f988463dd86b889c187062fa1da710c73e62268bc7591"}, "id": "1b40b082-1c02-478e-b981-f146bfd3bb42", "version": 3}`,
		Secret: "ka3rzhfkuu8b63cjcv",
	},
	{
		Key:    `{"address": "cef9d87b1df07ee40d7ae0c228abaf6c6cc12557", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "cc3ad3d9e4e2542436a3d5e470905756"}, "ciphertext": "5b005e86d3f8d0728d6e3a29572635e6b53f2b3e835b3c15189c7b40a005590d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "74dcba0965e7adb20ed8ccec674921f2"}, "mac": "c077a354a8a51c161b209b9219c67168b02a944ea5af14598255d2c805e49970"}, "id": "eaf7f293-1456-468b-ba9d-8d59e665ca2c", "version": 3}`,
		Secret: "7ubbsojkepkgjgbf27",
	},
	{
		Key:    `{"address": "79a031dea5683b53ea7e774b30c0bb215b68ce61", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "46ce9af522c98791d5934d3881c1e2c0"}, "ciphertext": "adad9a8550af853a4f4fa0f77f3c4e3ec1fda23ac8bfa8c692642a0091c76d93", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "7a61c7b74c9bd10320cf5068f5b33fdb"}, "mac": "6156684af3494ad6f2f0180fef72b0ce3cf46dba59ae831c59e1f7082bf7728e"}, "id": "b94995e2-cabe-4416-a605-92b3fa2c4cbe", "version": 3}`,
		Secret: "0e0lnf0o89qjqdlquz",
	},
	{
		Key:    `{"address": "250b20408fec0e9af2cf7b232fa8113616e4f88f", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "97f345129543ce19fa07a90b546bddd3"}, "ciphertext": "f5c6135256d848045d968bf07a1280cc2d4425c5244891ad8ba70149fd636b96", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "9d3bb986ad567833c12e3abac0e8757e"}, "mac": "baf6e24cea37719283e7f7589f3abc27de71e8de46e83a6de6311f62be46ee88"}, "id": "45085dfa-f2be-409d-802c-c97bc138f8a4", "version": 3}`,
		Secret: "yx56d22fn2jfpj0ury",
	},
	{
		Key:    `{"address": "ab6697c5a2870acb0806a050509a38b8b9a1a9c0", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "c92a01daeb07da5d2aa519cb03429877"}, "ciphertext": "8cf6b4c07f30fd6d6cb39ee05f54aa4baa267788dd1a2ca3caf155153787253c", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "124efb93c06dbf66cfcbfcb10b15201c"}, "mac": "7973b863261e05059a980ef5971c58a94d6d33057b7d2fced4a48d7cf7450047"}, "id": "dc35bfac-a4f1-4371-ad77-1607506aa9d8", "version": 3}`,
		Secret: "30isc151uiqtrpl5i1",
	},
	{
		Key:    `{"address": "90460ef310d8b328211ef4880f66eb0b59b8e4bd", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2154f0f49ef9a666203c4ec0e2594837"}, "ciphertext": "99c0685ebf8ef2ea9e36efbfe74fc32374845253c4a7bd6cbde20fb993072761", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "13a7debccc70ac06d8b461b8db65f51b"}, "mac": "1d83115f92a7edede8fa2f813dc50b755569813d0c2c74dfbae1bd326e8adeed"}, "id": "582845cc-7516-4eb2-a336-aff7ecb952cd", "version": 3}`,
		Secret: "p3qeoqyj4b79q3d1j1",
	},
	{
		Key:    `{"address": "5223c38b6efbca1d842d275fd49055687fa99868", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "711b4c0e2df1abf3c8e06501e85cba09"}, "ciphertext": "db3701ee863d5a40507d6c88f99077feac3068161d0a0b38772be4039950ae35", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "221faa2cb2983302dc258b62d9942353"}, "mac": "92f726f7589e0b5a8d99eb4a6aac3face25dcb67f657aedc6800ba3ca165132a"}, "id": "b8fd5d52-c5bf-4002-83b3-ac1ec151ddf8", "version": 3}`,
		Secret: "mz9aztakofuk7elz8j",
	},
	{
		Key:    `{"address": "227d323d453f489041f94f55c665761052df93d3", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "d3eff33f72af701a9bfc5acf778f7876"}, "ciphertext": "0ca94e57a65297b436e1a54dc02dcb8315b317f69d4809cc841df19fc5d96bfa", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "aa894fd6697cff10e5ed045bea104068"}, "mac": "04657a22f4ace04f9e755c72c39d27d80aff5608e0c566319e8b7a3ebe30ea89"}, "id": "1e1901c3-a928-4fd5-9e75-e964bedb870b", "version": 3}`,
		Secret: "8zrjy4axxmqgzebtkp",
	},
	{
		Key:    `{"address": "c0bb14800a9ecf601e3b510d39f8990abc4d3ae0", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "9e4f8128f7bab3bb63172ab851433f61"}, "ciphertext": "0c51f96d44c9e7db438f2eb84e66a7182fef8c6b10308db4384a7080410ae075", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "948995990a148eceb708d95b7e1ce33d"}, "mac": "59f1e90721e1f059dc017abb40def53cc2f3f4ecc50cce5d51ceeb98898e5b55"}, "id": "01c8f5ba-1de1-4073-9b31-2e165bea72e3", "version": 3}`,
		Secret: "o0lcl0agx4m75vjoqk",
	},
	{
		Key:    `{"address": "391b40c261585373f930145edd75243c272bf671", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "7eefcd484e010fe92cd448fba3c49145"}, "ciphertext": "47f495e82d878d7868a6d3937f0fb19e56e22d544aeb65a209ed9f889148a13e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "4e40c959aea11ad3abac131bd1049f40"}, "mac": "1df355e7973dec63f03e9fa1535e39d58a48a3c9576e2b5c6e3338e79c01e26d"}, "id": "354f1169-357d-430c-bf53-85d08050133b", "version": 3}`,
		Secret: "0c8odr6opdjpxiwe0p",
	},
	{
		Key:    `{"address": "50d194574572f3818efb20de6ebe6de20ec4fea2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "48ad924ae982dad699b37067de0cb7b9"}, "ciphertext": "3d2117a4b00272feca477cabf60a814059731f601517d9a99382c97386d440a8", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a05e265953e9f53663c149b288fd7f87"}, "mac": "1ea417e42039d91e78426cf40a63a8a7439046ccf06c5f59032e766df4c4763e"}, "id": "b94da329-e9b0-4412-9102-cdd7d9710101", "version": 3}`,
		Secret: "vzv2rs6sidr10t8u8u",
	},
	{
		Key:    `{"address": "f83c6148b8893203d43a3e4c47de7d75f89fbe43", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "fb5895ab254493001da84496ae5233e0"}, "ciphertext": "629d5b03e61e8c640ed0dcf85181093b1e03815c788a82bd33dd4fc99340f08e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "3543c5a7a673977fcae3b4f4339e7eb0"}, "mac": "b2964264a4d4f6be4de16a94dacde7987b87c9739a42466ee22d55e2ae40ba1a"}, "id": "46c27e86-cfc3-46ff-b6c1-2b7592c2c967", "version": 3}`,
		Secret: "cxu59rt4kj4z6sdgd7",
	},
	{
		Key:    `{"address": "cee037621b3332530c3654c02fa45c871d881e60", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "4edb682241cde87d387121b4873a6c36"}, "ciphertext": "30281da3fe2c96bd8b81a13cd13e9a0031d12f572b9f27e9cde1bd746305937b", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "9403ef37796cd9cc65e91e84b5823988"}, "mac": "abe65cb03f2064658b77093370461d91b6873cf3d1671a2c4b74b2bddd6df2e2"}, "id": "be6f81a8-d527-4651-8efa-73f85e1b9fcc", "version": 3}`,
		Secret: "1knk8iledpc9x9p02q",
	},
}

var validatorKS = []*KSItem{
	{
		Key:    `{"address": "148c0d7597767c0bacc36d49b8f3dffac4a0822f", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "28de6cb7d8a4eb29f5cd0eb125e36200"}, "ciphertext": "f44ae0f1507d4aed634064fc77f893d2e74d8777e846e71e68da886095193683", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "49a8d908ee909765bc342b9a6bf9166e"}, "mac": "fbba4932760ed2a2073f9734332f59eaa510b15b9308a429ff88025fc120b2f0"}, "id": "6df76ee3-8c55-43e2-a4f9-61fd95ddf170", "version": 3}`,
		Secret: "t5qqdr66wh0sv7w1ry",
	},
	{
		Key:    `{"address": "bb3f6a5f4ab8c1e3fb1042ba87af675c927f833a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "44dbb1b17da5ab2d2adb88c2d1eaa5b0"}, "ciphertext": "ca45dd17302a8c199e291daaf90df8a89a4ddcf2c71c6345ba950919b63216a6", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "128f5cf08d03cc48a2fa033b1cbf5407"}, "mac": "37c75ea7b5066cd306295bd2caa1be9adfab730afb027ea1af430350d6ee8939"}, "id": "782d8e12-781e-4b09-ae90-e2533cb0a537", "version": 3}`,
		Secret: "rmnkem6gali6730bdy",
	},
	{
		Key:    `{"address": "5e78017bc4156e8755501c2ec2f69391c70eefcf", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "12a1654201b5f888edbb0ddea8dddf62"}, "ciphertext": "17df625a1cf44cdfc7e094ff6d045598b535d106ebcd96e67992f70233a0b066", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a58ffa4a5cd517efdf51cea32cb5db96"}, "mac": "e681eec972ae68fd3eed913297d44c3bf93faa895d734ce969b5799db9d36449"}, "id": "f524c8e4-f876-4f63-b246-ba9c2ce265a4", "version": 3}`,
		Secret: "hi6msq5uypbeb3m1n9",
	},
	{
		Key:    `{"address": "8c141d6c23c5e5b82b3c1cb2f6e70b578904dc0e", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "41af2544cbfbb11744348cdd1c1f08eb"}, "ciphertext": "1ccc494c87fd35f2473479ba9d44f4a3400a91c07b00f567a8aeb15b80565ec0", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "e6c2d0423f673e7624f0bed2911c8a23"}, "mac": "b3ad393f3447f9bdefa06bf5ce7a7dcdd2475e54553d75ae3cb9b0dd67b3b6bb"}, "id": "ca9a6e89-a89f-4892-9f55-8b130ca6d16f", "version": 3}`,
		Secret: "y2jpfl9sjsl39sfae3",
	},
	{
		Key:    `{"address": "d4c423a4fa0610a3e8fbb2ccb186bf0daa0a0087", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "a51d8b330a2fdfcb9416e067172f05b3"}, "ciphertext": "e75cad836791432dfba4ab48cf704f56ee3bc81764fb97bc1fdf754be977abef", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "135ecab0e8505834467e9687dc9bfa3a"}, "mac": "2897ad3df0c7fe2ebc78420b16afddc2208fd370f0601a13803064cc5f6f7051"}, "id": "e11ca32a-2913-4aa7-bbe2-d478a934c148", "version": 3}`,
		Secret: "zdga2u6qjq1061gmiq",
	},
	{
		Key:    `{"address": "540b3a8ff43a987520901e6101c74422b88a52f2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "cfc987fc1c10e0eb9e4fb0555c75e642"}, "ciphertext": "614b1b4cf671d5f2136c645bf969027fdc0b635cbfaae6ff934ed2852954b301", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "bfe4f1e429ffcda4ceca0eeb8cdc51e9"}, "mac": "9bcf5a74c2b5c1e3e7973a81785ff38a8bbbd404e11c934e6f2455cb19cbd0bb"}, "id": "004d8141-6155-4964-9ada-8dcf0957b5ea", "version": 3}`,
		Secret: "p6xdisi2ws12f1jbgv",
	},
	{
		Key:    `{"address": "49d6e7b4bb8c6106b9385841502e25f06a3a8e72", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2f1590a2b34b87dfc59015e26db02875"}, "ciphertext": "f9a7a68b17ac552df3307eb59e2c2df9054f8cff8081070135449d64773870b2", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "fcd4be57e4f1aa40ce5f21ebabfe7659"}, "mac": "5788576ee49b7b98ac3d4154ec84a74f893e2d222aa2c5f5ef17ba21861e2d82"}, "id": "53fdc736-431e-46b6-a3a5-e5d4b612c33f", "version": 3}`,
		Secret: "x1yfunyut4qwsc5jeb",
	},
	{
		Key:    `{"address": "414f3a7c84ab72359f25d8fc52fccb0e5dbd4b22", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "6f8669f6f6f4aa301ea2d01f18d9b560"}, "ciphertext": "ce12f2a8e24d80b74fb93c877090d4db8780fdbc3e393ac5938c0b91a70212c9", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "b4c108d46495e8fc4e2460e0fd5cdb70"}, "mac": "f919aa00a00e2f1f9d1a2a438f8539024e1f601090b3ee232d27eb144e7bd2a1"}, "id": "ed227d41-2873-4097-b630-d28f82f39b0b", "version": 3}`,
		Secret: "r7c4vso4p8nmucy1tc",
	},
	{
		Key:    `{"address": "8e56d9476f00cc4e7df911b45532c5ecce152df1", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "a667060ddb3471063e3e28c00f7f7ada"}, "ciphertext": "0368ddd467f5570c7e8c555e832c3ef7fcf03b277d04454703c90c612ea77d0d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "ff6ddf2249ea999c0e2602e6d683cffc"}, "mac": "9e0b1c9987caef2e0555deb7db57a50175c31c5c08eb0af6a6602256ef82426a"}, "id": "c78be230-ed26-4780-b539-7839e1a757be", "version": 3}`,
		Secret: "qf2qrqhml8884akoc8",
	},
	{
		Key:    `{"address": "5a80d2d1fab4b0015b0a4ca4afcf254b9487204a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2bdc36e63a7826baecc7fd506366e58d"}, "ciphertext": "fcd9c6ebfd42f365b94887a59eeb894cef0636a8bfba0d5a82c00c2247a9752f", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "cfd179d806253e13c889c097dd05f9a8"}, "mac": "5b6e0bd44341199a57a14b7bdc8cd8145e27e2ccfa4699628ad8be3809b3375d"}, "id": "17183390-7439-486e-9a62-75ef9bfcf727", "version": 3}`,
		Secret: "uwzvnb1o1reg846b8y",
	},
	{
		Key:    `{"address": "8e4e540077dceb87401c82ece0f5df213483c1dc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "54824a5b74e48cc43803a41639a00bee"}, "ciphertext": "840b11f49d37c9ed2203ffadedc2d73c428d1ae294510be940b3aee363322701", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "ef5a6bfe5d040cd978a861e8cf691fde"}, "mac": "465d238e8b912e0dfda0d69440bf3d90b00822b9ae21f4cfe5b1620007ae0191"}, "id": "d1632e83-06ca-4dd4-9806-dc0eb514606f", "version": 3}`,
		Secret: "1c2wc3ginuvjkqjgg9",
	},
	{
		Key:    `{"address": "ff3b1d003dee480c5b55d87afafeb2884564cc7b", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "7aa4330f051c47e7e909152a18e977f7"}, "ciphertext": "0fe06d93f705673bd4484eff2eb6c014867652187893abca3c86599b980366a4", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "e13b1318db7d781976c4ecbb799ddf3a"}, "mac": "65e46b08b59ea764c0040480495c2469dbee4ae70d1895a024cdc6f589857534"}, "id": "54d9bbd3-97a2-4fa4-bc62-52af55de5c58", "version": 3}`,
		Secret: "03wc199pxmyhjjo41w",
	},
	{
		Key:    `{"address": "c94f9ef8ef1ab154d919835ab66d174ee7e364f7", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "88fc2aeda19cc680b29c77b74622ac70"}, "ciphertext": "42877ca3284cdfc148a72cb9b46a180d4f587c89d6039b1d283db4d5bd4533e3", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c962c6738eafd6c2a5cf5bec0b174c59"}, "mac": "136d1b4e185bb9be1953b510a03fcb053c367d9544e2e5c7211cb010adbe3c56"}, "id": "1564b2d7-2f79-4a8c-b8de-f8219a703c7c", "version": 3}`,
		Secret: "q6bua9q50r1bpd0pve",
	},
	{
		Key:    `{"address": "b944f8f416aca93a85f7f77ecf1a8eedd18739b3", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f567ba2ba0ada2a9506cae442300f125"}, "ciphertext": "c2e3b14588d9a663949bfcc6552888e7ab91c4fb6b7b13ba18b5469c5af375f2", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "967eabe152b1cd1b4de265dc1ea410b9"}, "mac": "80595d969259260d0b4cde4674432f69bc14241d4a7f8ac5ba268f9507d68a61"}, "id": "02fe0675-7747-402b-9efb-d711f7e31564", "version": 3}`,
		Secret: "67hbu8vfvw5wbpwece",
	},
	{
		Key:    `{"address": "4ae70e99d003b6ba1549a98ca88ef82b80b462dc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "de93c3aa73e13d350ebf652ba9d73a7d"}, "ciphertext": "2593e6d342db51944d9e5a208048734a2ab0ea56dd385f2019b45ca6aabdfff3", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "768a7335cd580e77abb9ce8f1fb40cf7"}, "mac": "a5dd38ca1349ad6dcc11871638b44468b3e296aaf39fe45c1f4bee767a16ac99"}, "id": "8c81a211-5fe9-42e5-a874-8699e16bac6a", "version": 3}`,
		Secret: "racoqmb9hijwsoy3bp",
	},
	{
		Key:    `{"address": "4a8f5b29ebd0ede1ab9ed102eeff433140218175", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "6342723fe0f453a0bc7ffb6a9f4274a3"}, "ciphertext": "735e02cad275818a8de30068aefaf70eb2be8b4d5f5baeb688515dd633cb26b8", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "d24d253b405850c323eaf9e4ee5fb08d"}, "mac": "80f86c495549407336814b0e0ca712f834e5dd09d46f20fce37230098d40726d"}, "id": "90583c07-1e53-4b92-8c8c-e52288eaef3b", "version": 3}`,
		Secret: "ce8calgfpeczaw9x66",
	},
	{
		Key:    `{"address": "c49e94021674bb71be155822e682f115aafc60de", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "1022acd7ade8cc5c394882c6da58d01f"}, "ciphertext": "efb7dbb811b4b3db689259de517735aebb8403f4179ef4c62058a147b97414d6", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "2ee0940061a5e038da3f24a09ea0d141"}, "mac": "e2f19077a6207c14e37a4ad57de8bcdb801bd739fd6b29a6d3dced1796d1a46d"}, "id": "b46e6f06-2cfb-407f-9987-61c630f85ffd", "version": 3}`,
		Secret: "4xpxnzh4hpw70mle90",
	},
	{
		Key:    `{"address": "00364e3e8910af8875da0b82027c2e432555af71", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "175728788e970f22ebaeecc7d5277e39"}, "ciphertext": "2df86955b9f9582351a49d54856d3e199a0110fe981a5ad4bf0d746fcaa8cd39", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a320cc30f669d083049c35f7ce1ae9bd"}, "mac": "03143f4e4156084b9a2d5e8423d27931201fb2581c9044c15acbc0810ff02caf"}, "id": "bd80a1a0-fdcb-400a-a0fb-9e14864543a7", "version": 3}`,
		Secret: "i1bkikvllz1422kgct",
	},
	{
		Key:    `{"address": "b5003bb38b57515e715deb4cf770313ae5e9fc87", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2fc8afc90aba27dc56f35a08199ea755"}, "ciphertext": "7db87fac8485885d192c0f632ed9d034905721eec5ca1b74dfc7240573aa1f49", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "0fed250d845b3dc5f1cf660eef851e3b"}, "mac": "01898674b053efbaf5e68ba5b46d78db5f783ced502c7ea0e1bd7e10320c477b"}, "id": "2a9aae2a-6aab-4100-95d0-eaa8508d9019", "version": 3}`,
		Secret: "3o6rvd2haxq9b6ncjp",
	},
	{
		Key:    `{"address": "78cfe75e812ba4f6c6d1a6332f84c2586d7465a2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f0b5fcaff9e3fa7e5e9450bff7cf9190"}, "ciphertext": "cfcf750e1a930a68794a020e2bc5b202e852989cd7f23206e1a94a8650286abb", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "8c1894de5b3ecf5b69cd7ef55aaf87bf"}, "mac": "294677688f0a2f9fe7c7200e8d942a456137aa880d33e1a9ca98bd785c85ce37"}, "id": "70b1e298-dca3-4098-9dd1-aa5bc10fdc59", "version": 3}`,
		Secret: "n9xlh762pgyozv500z",
	},
	{
		Key:    `{"address": "85b79a0551ed1f6b7b77a2488e8a9bbc71ceb9a2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "1dead234d12c84f3d3eb2c2ce00cbe64"}, "ciphertext": "67c183afcbf958c2d8d294f18a3d29f1d8f413626d1a07cd29a8b45640eeacd3", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "16af9a56eb304f83d051028f0ce48ea5"}, "mac": "0629d5fdcfe233048d086d65c35c63aace8326a174545c5745d9dcb585606d68"}, "id": "be52ff82-de51-4a8c-82dc-32e6a11ad9b5", "version": 3}`,
		Secret: "3fanazts5gn1d5s1nm",
	},
	{
		Key:    `{"address": "111869e31fbcf89fceb996b655b06f04255f11c9", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "227a2101eb4b598f9deef83ba3aa7a2c"}, "ciphertext": "9933ad60fa2575a35ce7f7aed9d2ea7d3bc9ddee11ee556f6d9bb92f93f082d2", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "ca73589577b9330ce7603974989daeaa"}, "mac": "87d879ae8d874705f3944cc2c06ce1558743afccb2c071b68cdd946709771eac"}, "id": "12fa8772-b902-41f7-adfb-b9380763d504", "version": 3}`,
		Secret: "m7d6rnrpnkg7wq6sgn",
	},
	{
		Key:    `{"address": "2e0cb271010fea4bc7f78323f18db9332bec828c", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "0bef560e6b3f0e97b74be32677e8d412"}, "ciphertext": "3a6e519ac891885db41baea8a67a10e667a370cc2f8eac9cd413a08aebd41e7e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "13edae18db5e50299089c9dcd1711539"}, "mac": "ed579a1fd059664707ba895be97e1ca3f6482e4c6c0e02f5eeca9e771d320577"}, "id": "17a14af0-a9bb-40a0-a7d9-9ac6e7bd7234", "version": 3}`,
		Secret: "5nm0gz7p43fj1334fj",
	},
	{
		Key:    `{"address": "582d31d55d3ae00094b157992ee40d92a511b9d2", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "726f2a920acdedd06cc508e0ec66c457"}, "ciphertext": "2b72f9d7cf11aca9cc125e5816c66906b35936fbc9b5b8e8509843c8421100e9", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "39ed168c58d8514c0dc5a99adc3275f6"}, "mac": "d41b82bf7b44ccc6902796d58e337941a3b0c6aac0b6afdeeb307cb6f0abc7f7"}, "id": "18aeb0c1-4398-4f50-832f-f147639143b3", "version": 3}`,
		Secret: "lwp90mefe85h5xvv9p",
	},
	{
		Key:    `{"address": "36e70f8acae930e42015fc3a3639708ce67eb2a1", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2406fe330448dabaac6b1eb5d4deafdd"}, "ciphertext": "948ea242bf3193a4aab4bd6e784078bd46f5112b832fc62677e56f9e0e4bf0a9", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "03acaf009dd7aa55422addf9cfef5b98"}, "mac": "1facc9fc893d0bd083bb334bbe3cbe50acea7c14c673a5d9593ddac098f1991b"}, "id": "b4191438-6ffc-40e8-8976-3a9fe64bbf77", "version": 3}`,
		Secret: "t73e5sov5cwqk7iy1d",
	},
	{
		Key:    `{"address": "b7d8cbc8c499536b7378071dc7fd41a7b4c5ac10", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "d2dd6b023fe0cd750ffe35fed970a612"}, "ciphertext": "3e963a97d83698f480e21f456d4018e5af05074f757b12408d8720adc357f81b", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "8e6b9836d7b41b33e59bb34a9ac31de7"}, "mac": "eee952ed48762bb4144ee2de207676d95e52a3a1440d9c642c64137a99be88ae"}, "id": "51835acc-0187-4536-a0f9-777e235744c1", "version": 3}`,
		Secret: "vjr3gm0e6m6j0zvft3",
	},
	{
		Key:    `{"address": "99160af2c8e84a2ed302a743b770c0c9839e8750", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "f924ae9c7fbf24870de865e104ff1d08"}, "ciphertext": "d4bf289ee4a6a2133d8b0ef505256537c4690f2d76f257564442f391ab50cbf3", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "d3802fd2f2c0ce4a5aa286dbca62ecc0"}, "mac": "2e0d96820d3fc2eb61f9fc6b5268973662cfcbe6db600e944447b328ad906095"}, "id": "a6876aab-8976-4092-b4f5-3b12404e065c", "version": 3}`,
		Secret: "2xcqgu3xe40kdy4aze",
	},
	{
		Key:    `{"address": "622763c8fdcb390e1c5d10956c3528d1f4fa4a71", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "dd6e5e7811975a7bbfdb514a4d991503"}, "ciphertext": "ef0c16c6d2547a1f3d255bb3ecda3d25b71e7575f6b5f5193d532946dd67ef5d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "e084f0bdea2afb9deddfb691c69fd010"}, "mac": "0e55a359da2907e7d5f3e0ed40d9db53fa06ddcd81d2944d6eda1b1fbf9f3c73"}, "id": "d70669b5-bc81-4a9b-badb-cfa81533b308", "version": 3}`,
		Secret: "du61jvdiwnjq9ierjf",
	},
	{
		Key:    `{"address": "610b890360f3aca1276c15da63fd39a9953a4081", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "a36ebc1f17af6fb86fa763a03b0e9033"}, "ciphertext": "fde1fdef76a15647406ada8cec3d177573e448c433be1e7ae7c8fd3987bc674e", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "54323778a72271caf7f2539a4ae08ad0"}, "mac": "fa1c0e0592712e8b4349efe8b0c54ce3bb386a80dbd17f7748bf9dc4b3e59b24"}, "id": "36e09c61-2cb9-434f-8b0d-d0a13b4586cd", "version": 3}`,
		Secret: "5xjgsnpyuwhhn8vqrf",
	},
	{
		Key:    `{"address": "fb4e51cd48bf4fbaa3b9b8555d517eed94f2708c", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "8d2790b695b409d5b4634e0e72deafde"}, "ciphertext": "a3a7ca057fb09b0283bd1132aae8cab3c050fd4e38bb4064c72977e5fb298773", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c466139bc6bf6cbb7451bcda0811343b"}, "mac": "2b9692af55a8a9f6415b9b2e89ccb720a99e23a62e2473b54f670e2ccf40c68a"}, "id": "c6da73ea-db26-4583-8507-12927040d7aa", "version": 3}`,
		Secret: "hm3pi51bw0foxs618y",
	},
	{
		Key:    `{"address": "39e708dc1bba9fd094254be87bbca15ca93fe1aa", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "0ef4074c0dbf7fcc350c1566b5099efa"}, "ciphertext": "e0ddc877c3eeaaf4bd165b9f7acbaf1cdfa470fcf88e85246dda08091f6f7a48", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c01344cc707930b5843285dd6dd87946"}, "mac": "eae28d8b0ee41a60f7bcca57a6213936a1285a73348442ff6af66870c6f03cd2"}, "id": "8fa23661-5355-4387-bc2f-2eae4838aead", "version": 3}`,
		Secret: "kym7kpspa1qzch7vwi",
	},
	{
		Key:    `{"address": "2e586e8f2b6e991d175e3ed9a8b79a4a7efbdcd7", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "9af9dd937f8b93fe826d972c113c21a5"}, "ciphertext": "8eb48168350fcc74e4aadabc2973a45ef1b3f77413f332b41a9c9ae48016a027", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "aa827e199eedd888cc2f7d19c57e5039"}, "mac": "334bd4fa8c365172134310f46140884a56a18ce718db6ab5ff0a39952e615b83"}, "id": "069140fe-5c6c-47ce-abd0-e6779eeb7114", "version": 3}`,
		Secret: "e4y16jrib068e65b0j",
	},
	{
		Key:    `{"address": "06f3d76ad85565cf28456db6abc52b3b75c4fbff", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "a75c9d03c98758be689ac82a4ac59ccb"}, "ciphertext": "9a4e9fb509e5c2b5fe74904258b7a2c5a65784eaf0c3744473af1ef13b29cad3", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "65b8eb225fa6ca7158e17698004a626e"}, "mac": "e4829f9222fb971e4d0fd1126a79585c17a943ffbe4c9281804f6d541a2651de"}, "id": "f0b05c66-678f-4469-962d-3ccedac519bd", "version": 3}`,
		Secret: "s72sdcqarfgboxyz6x",
	},
	{
		Key:    `{"address": "7c91f1d7ad67948a457b24e3121092db5228019a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "843b48b04d0b86c10adda17021569185"}, "ciphertext": "33508dabda944da8207b39562ed692a3d4564db1e0bbdb5705f30ef375de3bd4", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "799d451d4677fa42d91b24c34cc67e28"}, "mac": "35bda563e93c1fc27e3526403005d48a3073ce257c0c0266d1e3e1597c4f3941"}, "id": "17af0bd7-e836-4d9e-8607-0bb0567f67f2", "version": 3}`,
		Secret: "8jkaqxx9559qmx9cth",
	},
	{
		Key:    `{"address": "e90f2a0d04c42c6767dd9e5a95d8d7c9684382fd", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "e6a9b4da75ba8874c8497286cb7abc8f"}, "ciphertext": "a85bb9fb9587a7f59386b1608a35a8d2735f5985cc9a3bc78e7c4bcddcc94eee", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "0521624bdccdefd4d12d69b54fe15eec"}, "mac": "41edc5f704993e8bc84bc21728ee072b3b58665c25f003b95337482b526ee244"}, "id": "e759fc60-ceb8-4ac1-97e8-1fdc6f92459b", "version": 3}`,
		Secret: "jfsj91pyvibv7685ta",
	},
	{
		Key:    `{"address": "417fb1e6ca600de46a9a7165eac6ac7e82ef1b55", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "2d4376b1c44f773ea89bc97cb80f67c4"}, "ciphertext": "172d80b3e967bb8053eca0f522f75a7396e61e1c5bfbf3dd027ddd83494d3865", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "0143bf5c05e2af4c6176d9a7ffcaa253"}, "mac": "fb6413ade94c8e3017a994ebde9c559ad8547f2afe2fb3b8dad4f0c35c01edb0"}, "id": "bbb50e79-ec94-4963-90fd-61c81eaebf85", "version": 3}`,
		Secret: "jrawshyxuzajmhhuye",
	},
	{
		Key:    `{"address": "b38b236976b69cea332815c81d9899f1baecef89", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "1e764514a128d735f4e2a4b4d2558760"}, "ciphertext": "6dc8be51717d93ab3df6801790f5c8e1265a07ccd5f0af7625a732915869e592", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "3b6fdc3830b02417c8772a3df696a789"}, "mac": "96f0229a537e34df2ea913febae00e9e742593927b968b81c5bf1718f8156d7e"}, "id": "2e758763-fec1-493c-a649-d36657f55c94", "version": 3}`,
		Secret: "xpf3z4ayxmr2vsjrfq",
	},
	{
		Key:    `{"address": "ec24d07cc05201bc505994829b1570ec635b3b83", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "b5e584703cbd4c012014df985c591e4c"}, "ciphertext": "644d1fddb615b71c95f7e161a81dad6b61e260f5dcc75706a5a8f2adf6cfa2c8", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "b35b2e63555a1ce7bc69bdd53946351f"}, "mac": "a22884b834d12c296e65c7ca30984e2d5f4f589dbab648dab603a13f1980f22c"}, "id": "bb0bbf6f-69f3-4a21-a40d-e18211247706", "version": 3}`,
		Secret: "led1dmgsru1vtoy76k",
	},
	{
		Key:    `{"address": "51e1181c0d7897b246ac890d035eb2051173bac7", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "381acad1b83469df6301cac7c39deb51"}, "ciphertext": "ea9c8705a49b488f68cb1d5baf726c6edd896787abf8eabb67884da99681b7b6", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "fee7003fedaf18e8f275e0cccc2efbe0"}, "mac": "fdc32dfafe019e4aef579da8ab3cc040cdf2739e3d7f3b803b12ec5c295c5f61"}, "id": "a7d9fc39-80e2-4b75-b603-6e76adeb998a", "version": 3}`,
		Secret: "su9f01rm652qyz8qtu",
	},
	{
		Key:    `{"address": "d475e36f7d5b55f54a75d785960031dafd64abfc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "5f5bdb258d26ba834ed155f7acf5f865"}, "ciphertext": "eba682825acad6601bc822cdac64d1adeabaeaf32d1696b98d03fe8ab12230fa", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "3f423afa67a30445a62c6f451b6bdb5e"}, "mac": "9da75b249a058a05df3acfa0734e8b4c0fe837b5e2a56cceff23b7abfbbfc9c6"}, "id": "5feb91cc-a0f5-497e-a488-ac5f7e314b73", "version": 3}`,
		Secret: "kysc1kogacopq852dq",
	},
	{
		Key:    `{"address": "dd90743a58b6085c664d2ba4e9570b8b68d5d874", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "43cb43d8b3b3a9a69ec6864a1e77a446"}, "ciphertext": "bd4656fa613b85c8351b2d2353ae1444e7632a54c8ba8813d2b96216294f5f42", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c6fe014cb45cb6d909c285aca5c94761"}, "mac": "b2d9458391951a35d18a99f1e5a0de4f12ce46eea185d594fd8bfaf4e64b7341"}, "id": "c1c58147-8742-4a22-ac5f-007d3683c68d", "version": 3}`,
		Secret: "bzjox7kect5a3b01n1",
	},
	{
		Key:    `{"address": "523a0c1abfaf99a0433e9641378aa4a24336ec99", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "5fe9727a15d43582e5e6af47899e1458"}, "ciphertext": "6423cd0efba1ff54e5295831ee1b7667fdb708d8e4bce1e3c567b090e2fcfcf7", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "b8949e2f26057b4eb0e00b12949a7b95"}, "mac": "ca8dc020b143e272f691a98f42856d9e518ca507e6294e39dc25298eb72551d0"}, "id": "63cfd725-5727-493b-ad01-52479a3b1945", "version": 3}`,
		Secret: "cl8pljo2ti828p7533",
	},
	{
		Key:    `{"address": "af72d884fff983a784774f641632a1bb1c121bbf", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "ae72a4fa60e434c539ff3d3e3881678a"}, "ciphertext": "1802488b2357eb75a075ee0a79385c5882a684d9d18e2dd593ed19fe2d2d6310", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "617da990658455af08a466023f8b0d58"}, "mac": "99b1c4f811e087a1b908be78a0aee8a4c0d00b47eac91d64e10aed4a336717a4"}, "id": "2f3ebfbc-eb09-45fe-9f22-4a139881d358", "version": 3}`,
		Secret: "l82t8rpbyvd48oaryr",
	},
	{
		Key:    `{"address": "dbf2147a2df6a1154ccf2dde14e70ce2f0a5be14", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "ff03769496660110cf9f47cb10026fdc"}, "ciphertext": "d97258303f3721b1564c72d6d392d8057f3c66e0e8aa7692345214a9842c2f37", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "d50cfd7df7c3f8ac5805fc6562659b39"}, "mac": "8cd6971165432ef3b25c36d739bcfae2feab034b15717e778d3ee8ab6d35103b"}, "id": "65fb1c38-4e09-4bef-96ab-eb33d597ac64", "version": 3}`,
		Secret: "f3z41te96w3fyvdm7m",
	},
	{
		Key:    `{"address": "1404ba911109db883552c482412f018ddd30badc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "68c29f6b2b852152c664930a419d79d5"}, "ciphertext": "c4abf137d0e7c389aa7673697c8daf177cc56cb3206934a3c400ed915b29263d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "8d69ebeb500de04e708f69b2e125c080"}, "mac": "0772f2804f9be29edffd35cd7cbe819d9b4806f74852a6ebd5bff9e0ddc47479"}, "id": "849e35b6-7b00-4971-8bdb-7f8fdd534d3f", "version": 3}`,
		Secret: "4573ced639d5fkp6t6",
	},
	{
		Key:    `{"address": "ccc78ed3240092b5bbe88b2b05d5b061ea8ac23f", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "28e50e99874b7e322fa930f3065b7448"}, "ciphertext": "f27d81af78e01e75e7b441c2c1ccdf337fe776b38d0da0a53da9921e5264892d", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "36a422d99a5028fa941cd170911dfb78"}, "mac": "92f42497509dd7a0e57c7071fa52a8131cc24758a10ef55c6245f2af1d9fffeb"}, "id": "9b5555af-886a-454d-a463-b307ef46d399", "version": 3}`,
		Secret: "yduhj0a0oql0is6qjd",
	},
	{
		Key:    `{"address": "dcffe8e3e15c283ca3daccc1a2a024e6e8fa297a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "4d065c1846ac27cd0861a2db69607fa1"}, "ciphertext": "3ee3349715122fb924c024fbf7de3a1e2254bdc63a02814e0e3cebfdfcb02184", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "dc8ead487b2e5677db03e9c7f7e0f047"}, "mac": "3d299153a64e4d3a97443b3bc50a37b682a130a4a1befbcc0531c34b5a7275c8"}, "id": "07eed649-1159-4bd0-82bd-cde5d13405c5", "version": 3}`,
		Secret: "ricnttnf8o74cew72z",
	},
	{
		Key:    `{"address": "048e755e9c7fa4553b48271c3346e83c0d8397af", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "650b3ff26681cd767487fd2fdaddc795"}, "ciphertext": "427465a836d925e928404a028a2b832fa2ea81c5ca6ce8aca3144f2fd2fe2611", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "c402c23cef15162cca4c05dcbee0cb46"}, "mac": "5e97a9e3731debef1d8f6161e3efbd8b8a434cb88ae0edb7498dc6828aa0c607"}, "id": "7f3dd1b5-7d76-4385-adf5-c157150c6855", "version": 3}`,
		Secret: "xxcltx26y1an4fzzis",
	},
	{
		Key:    `{"address": "972a528c40d9f95967d57e2e8ba3e63ab87827da", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "8674064483f3fe544f3ef942da53d431"}, "ciphertext": "cdc6bee2ec1b595b68b7d3bcd260b5e5a33093824f56718758e009d48d00b8c5", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "e7ba48b9abd19660443371c3852f06c1"}, "mac": "77a530a5aa96cca154962f9ff908c0f0e3b0a1224d5ddfaf18ecce7539728378"}, "id": "f21a53fe-4da7-43fc-99f5-677dd561134b", "version": 3}`,
		Secret: "dmehj2s5777pbieqvp",
	},
	{
		Key:    `{"address": "5e942fca2257bb93685e7b239fe9727908892cbc", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "e9b7d3fed528047923357bf04ba4606c"}, "ciphertext": "a81c5e85e6f81be2488309c0383560e6f39dacb4392710c159bb3ea4a488f204", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "a1b124abf0f5ff1bad692a512d550151"}, "mac": "4c8b01db994480b1e4b021f10f4ac01496e4f72553fc07270b70ad112986d97c"}, "id": "2a92afc1-dc74-413d-b573-bd93c3983392", "version": 3}`,
		Secret: "8zjhexcnfwbn91d6n2",
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
	// item := managerKSPool.Next()
	item := &KSItem{
		Key:    `{"address": "052e1b586d44c23739d028efad5b88cca010ac5a", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "39c7336302f5c6fa99e6eb381b130ef3"}, "ciphertext": "fb33cc60f898dd1e6c1d7735ed4e3ea2365f000e5284f05f783ce24f5175fa11", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "f5665920b74febae89c2db2b1cb56228"}, "mac": "cf84492af914e3dcf2209d68572628202a3105ebe08822a68c5faeb62a56bbc2"}, "id": "6e9d267e-e96d-49aa-a43a-5e8422bf20aa", "version": 3}`,
		Secret: "35d1xll0rvvpcj5icj",
	}
	return []byte(item.Key), item.Secret
}

func GetValidatorKS() ([]byte, string) {
	// item := validatorKSPool.Next()
	item := &KSItem{
		Key:    `{"address": "148c0d7597767c0bacc36d49b8f3dffac4a0822f", "crypto": {"cipher": "aes-128-ctr", "cipherparams": {"iv": "28de6cb7d8a4eb29f5cd0eb125e36200"}, "ciphertext": "f44ae0f1507d4aed634064fc77f893d2e74d8777e846e71e68da886095193683", "kdf": "scrypt", "kdfparams": {"dklen": 32, "n": 262144, "r": 1, "p": 8, "salt": "49a8d908ee909765bc342b9a6bf9166e"}, "mac": "fbba4932760ed2a2073f9734332f59eaa510b15b9308a429ff88025fc120b2f0"}, "id": "6df76ee3-8c55-43e2-a4f9-61fd95ddf170", "version": 3}`,
		Secret: "t5qqdr66wh0sv7w1ry",
	}
	return []byte(item.Key), item.Secret
}
