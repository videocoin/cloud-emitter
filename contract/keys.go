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
		Key:    "{\"address\": \"8da8bb893bb4e2af35ce00501935be8d924ff8ab\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"925e0b9bc47fb3c6d7f9a3316249203f\"}, \"ciphertext\": \"6fbed808320d704a2aae96bbfc672cb275e96b5dfc4089b3ca2859a7dacf9b9f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b2b9072ee5400d812bc2258caa765666\"}, \"mac\": \"324a35f50fc283e6977b52c646af2f34456e2c593196f479acec5719a68f6f20\"}, \"id\": \"539b20af-9a37-4f41-87e5-29735114e105\", \"version\": 3}",
		Secret: "5i0azam48jqt2gimo1",
	},
	{
		Key:    "{\"address\": \"694a4b3586442db1a348c4f12412280e3cc1c92e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c319dffe21faaf5a5e724c76037fe282\"}, \"ciphertext\": \"40db5fca3828f7a1af3a175a189487b94fedf98c6c38aab7c761098b8b6b3303\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b9a2768c42cf8424d33d43ccea89ca59\"}, \"mac\": \"0d5f3336dc8855d6e8003fdde694d0c8949747b4288be16c89c0468c0ee944fe\"}, \"id\": \"117fe7f2-fb27-4a8c-946a-e3cc1d3d325b\", \"version\": 3}",
		Secret: "b9qenfijc6770i6576",
	},
	{
		Key:    "{\"address\": \"63d0d4b3f5f6676ea36ab7e1a3bdc497bb1a3c3e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a71f8155588fc41d880141c333e61ae4\"}, \"ciphertext\": \"a766a3695466e092372e1dbc66c4c8460eb79f4880833d18edf98feefc05ffec\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6e2c9d548fb869e436b95509d22a112d\"}, \"mac\": \"50f92709d4223225ae3ac80e1303a57ec820352cf8855eca4a72ef82ad9b206b\"}, \"id\": \"3941a7cf-9e95-4fe0-9093-a57a98587bbb\", \"version\": 3}",
		Secret: "7j2de74adndwnt2u1p",
	},
	{
		Key:    "{\"address\": \"de73508d34b002a640fa16f9537bbe6ccfa8d106\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ba567084789ef727e1cffe3416e1b305\"}, \"ciphertext\": \"8af4a8dff85c82929defbff2f816021b4ab7e7f5a4679080560f85cc147ef555\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"cb279a8214635cd3064767a53e818f56\"}, \"mac\": \"4f283e38b98c0e2eec6fb21432d4cb46a6f6d8998a527d4c537bd7d0da1d83c0\"}, \"id\": \"d193c46f-711f-4627-aa9f-80a158e01f4c\", \"version\": 3}",
		Secret: "npcvlbfycej66du3dv",
	},
	{
		Key:    "{\"address\": \"868d9ce60824380e33582ac2bfd068515bfc7948\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"eb17b96b4edd0b0dbca4c537c0bceb59\"}, \"ciphertext\": \"a7095bc735b54b16f21572690bfebce87b53fdf5b0a8595e8a0b0feebc5d58ca\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f80b5cd247381383ce87f3649c9e161c\"}, \"mac\": \"013a9d8f6493a418323f65bdf819ef49946d1fe3ee80b728a41a2c6ff57f189b\"}, \"id\": \"6ee4acf3-9275-41cc-90f3-7b4912d81304\", \"version\": 3}",
		Secret: "rpxmntci0nbmdakl2o",
	},
	{
		Key:    "{\"address\": \"dbb26861daf5053da9e418ce250fb05e1ade8add\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a84a4f2920c732a29efd39cc01583eb3\"}, \"ciphertext\": \"914303006aa3e5f010f86d72fb253fd15cd5b32f5529136d4d1d70182c9c4024\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"054b421c76d93e6c5f09b91fbb50432b\"}, \"mac\": \"1738247fd44061cfea7cfe144d88ba0dac05071a7776460f3ac0cb692276e9a4\"}, \"id\": \"6ebdafe4-f4bf-4f67-ba68-9718fbaf9133\", \"version\": 3}",
		Secret: "jsqwo3e3lm1atwx1lr",
	},
	{
		Key:    "{\"address\": \"221d1124ff181470c1f1092fc90ca7a485662b51\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4a12fffa9420af6acbbe7cefb2b9c5fe\"}, \"ciphertext\": \"1a95eb97692e8235af840245340ac2f73285b685025bd321d1d03d188a85b302\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6ec703b7b7e061c1ad723cf86e4c099c\"}, \"mac\": \"4fcb469a32a6cf8789bb354aa597d1f80531a12cc2915454b43fc556e75bf2f0\"}, \"id\": \"a26f17c0-c14e-4664-887d-995916af6095\", \"version\": 3}",
		Secret: "3ft5xkwryf57wip6fy",
	},
	{
		Key:    "{\"address\": \"64105c4203f6d773926eb6be246794d85aac5e02\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"87c2ea6fcbacf0df3ef7f8689d2db37c\"}, \"ciphertext\": \"9f1dee452a1301cade0254a759abbe217bc6e2aed45aeb72749b770c3af37d27\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e1298d8c07c2c2a95d7282c585bd0754\"}, \"mac\": \"26886fb08dbf4c91fd5874cb920f4516b83b65ff4c60e15132d7b859ad302dbb\"}, \"id\": \"b5a4fbb0-7833-4594-9ee6-83ec21854c10\", \"version\": 3}",
		Secret: "mexjffj8t1fuimyqj2",
	},
	{
		Key:    "{\"address\": \"e65771a46e9440027e44c68977151a8031264847\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8b9b653299004422b84d1bbebb2e45a0\"}, \"ciphertext\": \"0eb7cd38670c01521bd907e3e30ab86d9641d09f8d6908171216bee4b591ee4c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9a096efcc67846d9b983790edff44ff3\"}, \"mac\": \"32a15c15fa45febf17bece4116308ab59480c6a3de9754ba88c497e6145e0d96\"}, \"id\": \"37b110b2-547e-4c2d-8c52-04beabe1981f\", \"version\": 3}",
		Secret: "iv7iqq82ldr1xctl9z",
	},
	{
		Key:    "{\"address\": \"9341f92ec29b2064a2c6d2f7028230dffcfe977a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"85c1c42e50e3e36623392e4854f26b9d\"}, \"ciphertext\": \"f2ce9ba59f0d6b3910177358318d5e8a7114d0c012e416afb181328e58cb9e4d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6d1ca9ed248e5c9c8e821387e59260b9\"}, \"mac\": \"4d9a40577208ec2ca6c1fbf55cac708b0d4b6d11cb87ac0281d1adf0a4200c2b\"}, \"id\": \"b9a64156-dd21-41c6-8eeb-9d27440876f5\", \"version\": 3}",
		Secret: "3yxq8lso7ppx1d1u16",
	},
	{
		Key:    "{\"address\": \"d16496626bba7c0ca87d599a25909b819607e870\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"bff58bc70823033a2b3a053e187b297b\"}, \"ciphertext\": \"6716a82ae4698b00d080ee08c8b98770b83bf4dbe4c68af38ca97050f6399761\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"fa2f3561ece9cd636184e4c5030fddfa\"}, \"mac\": \"095d29bae17e76b38f8d99f200727bdc9dac0792a7ca0b9794967f54a54a1856\"}, \"id\": \"13570a50-df4a-490a-8be4-2a6edceaf112\", \"version\": 3}",
		Secret: "xtsvkduie1c0uilo8u",
	},
	{
		Key:    "{\"address\": \"81fe9af77d14610f28c1b8c06a40c07bfdfd96a2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"904bc9feb7b75f9b1506c6531737641c\"}, \"ciphertext\": \"3ae3730c8c05545743a2f72469c928e51fae0228ad6de8c01d2c97275f87f0cf\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f1f7ded3edaf9fe3fcb92aa9d8900731\"}, \"mac\": \"466ffafe1cfb27315b72967e98bbe0e56e7d702a3f4775754f6c330c9c3147ab\"}, \"id\": \"955086fb-22d1-4447-8efd-d7d454d6d89c\", \"version\": 3}",
		Secret: "unmmq0rmr759358ian",
	},
	{
		Key:    "{\"address\": \"aa384d5024b10b8f81fe13f3e8a1ecfbe3daacb7\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f7ce1117b57b3d0ffe8fe48023424670\"}, \"ciphertext\": \"8b45500165e92e77ab7994810d62aee7599c14d113614e8ead9ff28cd769e329\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5e3678edc95623900ca92529ed26c037\"}, \"mac\": \"91afbcfb09a024c583e3c1f9205f39614b0439217d9af952f780f9a9e333ef2c\"}, \"id\": \"3b8c73ab-59ae-47ca-a1b3-33cd14560468\", \"version\": 3}",
		Secret: "p6mx70r4q7ysy4fi24",
	},
	{
		Key:    "{\"address\": \"b1147fe71240a8b89ac342622b43b582215f7e2c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"817f031d2f8f54f455792e16e4deeb7e\"}, \"ciphertext\": \"464b5e9880483d05ce9b09fd2b52a2e1c80480bc162343f6988e52f5a3f3cc79\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"885822eb52c5d9e60889c4ea7a504789\"}, \"mac\": \"5d178e379de86df30831beddb9fd9062c82bbd91453292ec144ae3aaaeb1c3d8\"}, \"id\": \"895d7c57-c815-4a5b-8060-f24581900912\", \"version\": 3}",
		Secret: "kvf4v1bo6egmv0g4sc",
	},
	{
		Key:    "{\"address\": \"39e0c40a928b309709b8062d2ab4f9f5aea37577\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"84c3dec2e66a20b76206db4b5e477f53\"}, \"ciphertext\": \"34f390ff5037d34e255561c688fec598cbae4595de2af8e995b19976220bf7f3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"34e3fd17e1c70a771b9036cff835f976\"}, \"mac\": \"9651fe97092e89931184bde2863ba12875c5052e04fadbcdc1e4c07f45a1b1d4\"}, \"id\": \"82ef0d0d-8ee1-4f8f-b4d3-3e533616b6e8\", \"version\": 3}",
		Secret: "xeup7cvm32tfg83msc",
	},
	{
		Key:    "{\"address\": \"09a42eabf5ca76eeb24b67c098f1a438357e6838\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2b078aaf7b69d3371105fe60f5ee4209\"}, \"ciphertext\": \"47d5a2b0a74db300aff49615e7f1b17349f95b48395e0a65864fd5f243fe17f1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9dacca1d022eff8f366caf0e9cd2d664\"}, \"mac\": \"632850a192fe4e87bcbb2ef1f1858fd64acf3b778436eae3232f9c49f838fed3\"}, \"id\": \"9fa2f372-7bba-4f73-b595-c848d3fa4e20\", \"version\": 3}",
		Secret: "dqzpxk1a07rinmxfts",
	},
	{
		Key:    "{\"address\": \"70bd3f9424ee6ae24412973d62d3fcf384ac1c06\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"73a28470f4190998a8ed6a4a653628ea\"}, \"ciphertext\": \"fd02232ea410de5f166a9edc862c3a34749ad2002de58286179724c715f605aa\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e34e2dda13e9f8c5c4666be0abf2c820\"}, \"mac\": \"e904d2a1a7dc5b20434a793655bbebbd73ad719ac0075a8939f3ad3ff6a45608\"}, \"id\": \"9b578a53-ac32-4b8c-ae91-5b3acfebc51c\", \"version\": 3}",
		Secret: "2ianx49y6vay3oc904",
	},
	{
		Key:    "{\"address\": \"babf0ad61211f6e494a2a69768dd077e1b70b821\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2c5ad9cfac06f2d8450e48676f0b0f66\"}, \"ciphertext\": \"5290a7ae7f3b67aaa95427282ee876b427a7e88d3db4a9f29289dcf81da64383\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4b4a037923400716d2d661788c0c774d\"}, \"mac\": \"b7c3859cab25fa3e4f42e0837ef97c4d7cd52df3224445c1f9bb2296a4973f3e\"}, \"id\": \"290d9c5a-139d-444b-b928-2ea5d613d5ca\", \"version\": 3}",
		Secret: "6mcp1xrlgeadqtksv1",
	},
	{
		Key:    "{\"address\": \"7478dd899eb1c285b0f7eeadad5d705ac2a327cf\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"18c66d463292873855cabf10768abef1\"}, \"ciphertext\": \"97403c671442c68f637ca7a6a1564ffce815c89acb1e7c579a88de1fafd26d2c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7468fb11f089c84b19383996c8c2a2c1\"}, \"mac\": \"b45f6cba624fb055de8fc9e4dbf3f7299b5cca23a65df9cd654f3c85ca6acd41\"}, \"id\": \"afc3c3f1-b0c1-4322-bc7c-f5aab262dce3\", \"version\": 3}",
		Secret: "wk2gf72erav32720e5",
	},
	{
		Key:    "{\"address\": \"c43f4d70d6a2678dfc32650655b11251eee2380e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4763f459c3518de959e79718de3933e6\"}, \"ciphertext\": \"fb6be4106ddbf127552359685c55e89330577379df2f21eb0cfef5b60abf39fc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a6330b46fd6edd85150d0c5ff403338b\"}, \"mac\": \"25d9f58a2514493105b32d4f148fa850fe8a15763d264d0e8e376b440e66fcbc\"}, \"id\": \"c68eeb23-dfd1-4f2c-93af-55857faf05e1\", \"version\": 3}",
		Secret: "nyhmcaj0qvriun0ypi",
	},
	{
		Key:    "{\"address\": \"3cd14180f25b0867c79d02e52a1fff65edb8f02e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e8bb4eb3f4f7cf0f623d82011886587b\"}, \"ciphertext\": \"557a799da206b293434ff6fef28526000f36404fe04002ed54965b4df7da3d38\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"de420e0ef515e13cc1ce5b2e58f17660\"}, \"mac\": \"3b789caf35a7c45f2dbfa005dd7791334aecd52f487913f0e6e52ad1a63e460d\"}, \"id\": \"2e688448-3358-4dcd-826c-0e6bfc745dc2\", \"version\": 3}",
		Secret: "7xmeh3izrmza04b8af",
	},
	{
		Key:    "{\"address\": \"a4cd99fdf6c5790be3e730f6aaf6b796b672da83\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"fb03c61197afc6d067cbdb19af1ae0a8\"}, \"ciphertext\": \"5667ed535266163fd1fff268374a7f3c22fc6710855ff1c060df8b0058a234a3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"60ae474dbd00cd6085a24436a03293ac\"}, \"mac\": \"63498fefa0a8da979697dab06b264c160dbe4f881cb8d2f4d9eea01ec2438d65\"}, \"id\": \"d4d4e3c8-eda6-47a2-884c-2215c0e2952f\", \"version\": 3}",
		Secret: "epekl4rrc96lgoppf5",
	},
	{
		Key:    "{\"address\": \"9edaad1f5ff683361ab221cb6c16c48888a6eb64\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"cc9fc07baa1942ba5cb377981c2d6427\"}, \"ciphertext\": \"38fa6189997e8d01a186e4dc7c45a4b387049c31aff55e1ae2bb94bf20397f65\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"26ca3475da58fa554969d5b2a16fd753\"}, \"mac\": \"d5e3dc4345ff44c558acdc9eac34ef5ca01d2f913257a171c324390c405bd10e\"}, \"id\": \"edeb76e3-8b44-4961-aa0f-66b54fc6a5af\", \"version\": 3}",
		Secret: "c4gx251yoafbik1qks",
	},
	{
		Key:    "{\"address\": \"8b8a3a31b542742e39317f47e429ed6436f263cb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e1aade0a4f3c3ae67253bf45eac3a124\"}, \"ciphertext\": \"f49cfd3edc65324404555bb92cd85489393781f76dc62bdb6a7c4eb7e548a290\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7802e43672c904e5ed80898a8cc21968\"}, \"mac\": \"8c484bd4fd4176828c80f67b97474d1b19302a0e4a858fb5901b04c9967ea5e5\"}, \"id\": \"e085e563-7f13-4fb7-b387-b21e160fb95c\", \"version\": 3}",
		Secret: "nlhb04d4awpkomndu0",
	},
	{
		Key:    "{\"address\": \"76c0339620266fb4dfe6b40302205d80d3eb713f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0683d467329ade9ea09a69e80a820981\"}, \"ciphertext\": \"e9e6c80a26dc8a6a3de25efbf159117712af3a64f929052a024256fc2a5b319e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"03a1859c01d3cd864bac775db999c958\"}, \"mac\": \"e8165dedbd875025d2681528d56fa90fc1fc3073fd155253501d5a821e2361b1\"}, \"id\": \"d34ba99d-5220-48a0-b2de-2d16c892a983\", \"version\": 3}",
		Secret: "shueohtekmvlu8w576",
	},
	{
		Key:    "{\"address\": \"410a164237dd7ae39d5c4aa284a1e99cb06b3c86\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b948b659d05034562f1ff3fc7e15b087\"}, \"ciphertext\": \"50f59b59a8af1bbed34cc407c9efbfd6d60d49ebe079d784b43b613bb430e846\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b79da84577feeb8bc2804ce6a8e0640d\"}, \"mac\": \"d278ad3233d3dc333d01977dc2bb6dd5e2a98bca2f9705512936bce875b34234\"}, \"id\": \"156d7d15-31ad-4c6b-87d5-be55f95c4d4b\", \"version\": 3}",
		Secret: "2hojsoqm8g6wecm2tm",
	},
	{
		Key:    "{\"address\": \"3ab701e126622e6620fcaf2142d8d807dc55c514\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"040a204703703934cccbca277751de03\"}, \"ciphertext\": \"85e93bc2267eb2c11bb373b77542a25e0da1a12c3efb0da2f07a622357f937b7\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5df716217a8c12d16d34315709b108a4\"}, \"mac\": \"aca495e8f2dcc6fe3cfdfae6fde2acc716fd62298299f3968260ae5f77e4318f\"}, \"id\": \"0179ac1f-056f-4b45-b374-f2eb36e5c8af\", \"version\": 3}",
		Secret: "4ragydv0qt37mwghhj",
	},
	{
		Key:    "{\"address\": \"837f843834f57956556b4d8047a15d27977ec204\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"af229a4c3c1f9e42d07bd9141139ff92\"}, \"ciphertext\": \"e4a83f9f42b6592907746dd7da667e2b65ecc1caeb56b22f7d08c515ef45fd0f\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4c16218e816f69176a372d03b58457b8\"}, \"mac\": \"143488c20c793070462c26d40cc803d2c18648bf8a53cb24314141369e629cb3\"}, \"id\": \"ed0a4738-c7d8-402f-8d13-06d578e744ed\", \"version\": 3}",
		Secret: "6in7e58xeiq5bpmkya",
	},
	{
		Key:    "{\"address\": \"f3c8b6bf86f6e8720e88d46521f5f436767756de\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6c84bf795d9f207c27b927f7f469ef77\"}, \"ciphertext\": \"30ab1bafdf193d2a6c69c1b74b80a26ffc38798ad27a56fb1dedf93c2184f4e1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e20e5428ca619b7f9a0ad5ce1c329d4e\"}, \"mac\": \"a622a20399a2689403f0851fbcababeb18589dbd811049d8e19f147dcddb9236\"}, \"id\": \"d17f5744-e5bf-45c4-866b-0bbcfc52a9d7\", \"version\": 3}",
		Secret: "7xo0c91i9fo3k9uxus",
	},
	{
		Key:    "{\"address\": \"a2735b2e8459e0f48576641869b1a918ebd68296\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6db16e6e60d530c3e1529cb11d67e590\"}, \"ciphertext\": \"e5028c7b0c1e292523c36a67252c14d941b174546fcba1ee5c5ecede3c4de3b6\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"92d6cc115e018ebc7d0fa67fc1bc97a5\"}, \"mac\": \"2b19b35d311679321375dab4f2c48cb083241d5cacdafd5754f3b439c7426597\"}, \"id\": \"a4dd235b-c0af-4420-b6e9-ff56663e1d7c\", \"version\": 3}",
		Secret: "cw1irlwj4gvm1vpxmx",
	},
	{
		Key:    "{\"address\": \"b23e484f1235e6593611e93975d2f07ef34e1bad\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3cd8a2d3b2b2e59baecf097f35867c5a\"}, \"ciphertext\": \"486e8b16f16dc64b68ada00ffcdf9def00aea72536f4403cc69a715ec8741124\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"af518e8038a82e93a879d581cd660dfa\"}, \"mac\": \"09649ff7ebc7c158225a4d670919c2eee8fd583793e1d0940a96dd0b62f40e16\"}, \"id\": \"559a124f-b0b4-498c-90d6-330c7ff4404f\", \"version\": 3}",
		Secret: "hwl1mhai9tcmlp47h4",
	},
	{
		Key:    "{\"address\": \"aba1aa15037ab380cd0c7993f243794791cb448a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e725777c653f315e92379e855fb769a8\"}, \"ciphertext\": \"2edc3730eb5ce596e09baf829f4f7f5bd6c4a38264fbfecd3b9ef8405679282c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5819047c134b70251d3ba6d495eb3ab3\"}, \"mac\": \"4081d742aac83ff0a7a9b58d85f023695fda053c8fd4fd5a7fb30550bd65df93\"}, \"id\": \"041cc154-a73b-42cc-b81c-e82226a44e85\", \"version\": 3}",
		Secret: "9af1d7kpgpeympj5bq",
	},
	{
		Key:    "{\"address\": \"d78e6a757140ffcd6134c19d5a453466f2878dec\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"4bec02c7e4efc31d71b2a6b89a2d1ddb\"}, \"ciphertext\": \"e04e28ec2ae5eb4d38023cffd1553abdc1a61d98180a0da69f12547674cfba52\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"cb5aaf3a0d3c1bfc1dd2717a822dfe8e\"}, \"mac\": \"c365377838e218faa0ad483122ecdceaa7aca34302d78857c40d7eefeb63e52b\"}, \"id\": \"8c593863-7717-4e02-b1cf-b440d786b084\", \"version\": 3}",
		Secret: "rxjz2wlyyt08s6bdrd",
	},
	{
		Key:    "{\"address\": \"39a8ae45cd0a42009d72045efd18e8d0fa476974\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3f4b0d4988658b9d06659ce4e7f9006f\"}, \"ciphertext\": \"2645c7d30578ae7463fcc7dfc125956cd13d746764e91ef922421a1290fb737e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1e3a69cb679d11b24a425d47079583cb\"}, \"mac\": \"46cb4d95ae1994be2001dd21c2b06b707e490621accd80eeea3e6addc73d69a6\"}, \"id\": \"fec70ee3-2613-419c-b9df-e229e2b030ee\", \"version\": 3}",
		Secret: "b2ug3penq9pl90k4cb",
	},
	{
		Key:    "{\"address\": \"44084dbba8a6e34112b14642f3e72cdf5ba14549\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6a5b8da98bb940c6f538571c14babbb8\"}, \"ciphertext\": \"d0e16939babc9a2c6b315cb57e03d013f44eec22dcdd975a9725000a65c19924\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"82331a3d9c22b65f868ac92eb07beb8b\"}, \"mac\": \"a4641bae4f9b84c652f4c224da4891236e68d93ba1620a72df8039f21c4c02ae\"}, \"id\": \"fbd080ad-37fd-49e8-8f95-942d6b028a3e\", \"version\": 3}",
		Secret: "7cg3jqd74y2vpiaktk",
	},
	{
		Key:    "{\"address\": \"c15944556fe392a7d6d5529550e7fe8ac93b25b1\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"62fa02bb6ac521e0bd63bd3e97566240\"}, \"ciphertext\": \"6fff8c96b5edcb61084cd9e6da9991bb29f6c536744b5ed82d2d6f37b45d1243\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"2777e8841be5cdaeac5f70b3640fe96b\"}, \"mac\": \"b2e0c9ab89763781c344bbfa19feaf1237f1fc58452dd3cb8c194022aacfae1c\"}, \"id\": \"b60f85f7-6a92-44ea-989e-ff343ae2e6b4\", \"version\": 3}",
		Secret: "v0oyjz4r5n0qacil4f",
	},
	{
		Key:    "{\"address\": \"77c4d79b468fa96a8545bedf5ef2b60ae169a9fc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"457d333c61c9f9c10b5bed07a9b6058c\"}, \"ciphertext\": \"cb0336bc427189b765a0a604545b489232a49a49c77d623ae4cb4ec7903431cb\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"200bc25cabe686608fbed87e2a5c7482\"}, \"mac\": \"71dcae099226a301eae4fdb72421f64ef052a30fa50bf358eca8b6da3ff0d6db\"}, \"id\": \"838c8434-e337-4d3c-8970-8cb85200427c\", \"version\": 3}",
		Secret: "c4tk048t2gnm0xltib",
	},
	{
		Key:    "{\"address\": \"92b49b2f3c935d0e84fb7cc64bf859b0d9402baf\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2eb94bddd01a280c0de33cf21947476d\"}, \"ciphertext\": \"3b5afd2eed85fda096c2cf41c6f6dd0be6a6908edbaa6d9875cc010b2cc60c9a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"251d7ae00f5312a094291e04ec811551\"}, \"mac\": \"7731c6dc3c8a2a63799c36b9418d6de237bb9cb224b081b9406ff50f1f159511\"}, \"id\": \"da8c0aeb-fe98-46df-9071-c6b20df52a9a\", \"version\": 3}",
		Secret: "2g7po71dd95hyfizdf",
	},
	{
		Key:    "{\"address\": \"dbae61c04c32efd9ed7ce1296f963a06a960af4f\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ac75a536c993b83ad5136ea3677648b9\"}, \"ciphertext\": \"1fe34657d938d8bad81cc7291c9aaf9a06266a0c5e42eed78d4d08e379349949\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6ea371be8d207f9d78513bf1bb2e1128\"}, \"mac\": \"86846c24b0af86d65c69e8ffcf46c8a3b126baa8ef99fa6dfc280e5a6ccc1af4\"}, \"id\": \"0ef1b3fe-d738-4861-b27d-6cc9db1d5651\", \"version\": 3}",
		Secret: "vd86b8c3jg15g1qotj",
	},
	{
		Key:    "{\"address\": \"d6a688b96177707f73146e886c7a225be8753a37\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"32ac09767447583141f308471cd0f7ed\"}, \"ciphertext\": \"d29eec561136ffe6799155a236fa0646f673053eb9839e125ec9a95d4384ad6e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5b6c4e92ca53020c61292f7ab1ce96c6\"}, \"mac\": \"2a1fc6852fdd5bb1141a737191c62f6f7db5f2be17dfe4fc54572145f8b71186\"}, \"id\": \"71c7f91d-f515-4f50-912e-b914d77b8e65\", \"version\": 3}",
		Secret: "yiszuffky4m8h0m6k9",
	},
	{
		Key:    "{\"address\": \"3ac17b64f6589e5fad4d8de50109246b5a90d9b5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dd1e38f928030e15ed5818ae0b543cff\"}, \"ciphertext\": \"432dd7f7f95881ee9b0c3839d5a5bf24f221a83ec7c472501b65952d7122b7c6\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ee534329e97ce234907e6ca26ac0071f\"}, \"mac\": \"fa416958e2e0b4c6598b851082668650713a09cc5338863ef50e4c36a3cd2102\"}, \"id\": \"a28ef276-93b8-4a11-bdd7-0daf092b715b\", \"version\": 3}",
		Secret: "od0g6bokp4zeuduizw",
	},
	{
		Key:    "{\"address\": \"864b3dcc39ecb7cc0a9236f1ac7f90023853bd9a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"59627252d11cd16e1f9f85858a3d2239\"}, \"ciphertext\": \"5dbdc11649763c076c5bcaa10fa353f98cddb31b50302c6b1b08bd27ba540121\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3acdcce2a6123adc04bdde73b890d827\"}, \"mac\": \"e4d9b1b5496f65b05c6ece311c4ea61813a350debef68e319f53651eb1a92962\"}, \"id\": \"f963497a-f546-4eab-8668-12bbbc1465c0\", \"version\": 3}",
		Secret: "gefqxrq5xi85ev2prm",
	},
	{
		Key:    "{\"address\": \"8476fe344cb1a6c6650650af8dc8650f11477e4d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e37d02bb9851ba5ad241efd461b64c45\"}, \"ciphertext\": \"bef047a389083c101a12c3c0b766c2253a9b383ff051a78f791fc683a207d7c1\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0de3f6b96f8e43c01b161e6e2729c4af\"}, \"mac\": \"c633b4bef63e3b7960468e7719459a4b337ba1df547e3cc64514d685d903cea7\"}, \"id\": \"d6abd12d-b3e2-4a56-b953-7a7cb2fed057\", \"version\": 3}",
		Secret: "yyowgvp191f2xbl0z5",
	},
	{
		Key:    "{\"address\": \"05e60d4a5143f57148abe09cfd05a048b12e9943\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0ece26891acb69b4df3a9f0ca2635834\"}, \"ciphertext\": \"a16ff7e78d8b59f4e0ebf7a8a2613ee4636e5f6915a6698d6df2ace7ee5c7ab2\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"39830ae8e1a3c3a5c913c9c831277560\"}, \"mac\": \"8ecdc05377e02a1aef42e01ba7bf8305f435e979c34fb8f5fbe4f3a51bd08fbd\"}, \"id\": \"9bb1e5af-3ff5-469e-b3a5-2fff9ebe477e\", \"version\": 3}",
		Secret: "140jg36ypvu12aid9u",
	},
	{
		Key:    "{\"address\": \"a8569c20baa4956ef61d30c13bab324d3858b6c5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6ef70720fe496a349a8571f05474c2dc\"}, \"ciphertext\": \"6a7c3329b7fa18676af9dd2056ae7da8b81c1164df524d8a6bd1778952f816df\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d8331384825f55dded2728ef3c0cf032\"}, \"mac\": \"371a241b649d8c02d4e994094a5b54af5647b6a92840194f77cf4929fab9e397\"}, \"id\": \"86d60cb4-ab07-449f-a899-c946fdc45534\", \"version\": 3}",
		Secret: "k3pfly314hheplkue8",
	},
	{
		Key:    "{\"address\": \"38894eaad27d74c47ef03a2807c97a8148e6676d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"442c688a5c845448b239842db9fc0c15\"}, \"ciphertext\": \"4cb8205b9fcf9e50bcc7c24b42dac38d28069c18d7bf0596ab32fa5ba58c9c1e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8ddb8fc527b6c42edbe8d3c0bce556b4\"}, \"mac\": \"e9e88ae3e73532762a1cd1145eee004a68003180dcc669dbd46266ecd848ff9a\"}, \"id\": \"e33db8b0-432c-431e-85d4-8fc2c332680c\", \"version\": 3}",
		Secret: "9fejqaqleksz1c4yjj",
	},
	{
		Key:    "{\"address\": \"3cb8f98b06c62e8174501d7dbab5087c734a3be1\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"14aff9883b91f14633816954e6bef1ee\"}, \"ciphertext\": \"99bfd3045b5204a93c91fdcb2ae221539cf6c58902e547db5ed058120840e634\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"08e3f3b9ecb3df2d6cd0360ccd08391e\"}, \"mac\": \"245a4b4a59c101d702d3258b4fc182d077f16a5d8f5b473fe9ec8f3d9781dc0e\"}, \"id\": \"96d048fe-4da8-4e0e-a9b4-1a3c7c8747d3\", \"version\": 3}",
		Secret: "9iqdsn2zveix052ao7",
	},
	{
		Key:    "{\"address\": \"927cce47a51fc18cffbf5336575c67f0de72b8e6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"fc31e46952416dd5bf1127a46c3d2cb5\"}, \"ciphertext\": \"1aa2b209b0afd83cdf50e95cf5951ffa69597e07561f7321582ff6c272f5d482\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"76720f59c0ae22a2155e87323b841612\"}, \"mac\": \"5f32a35ac8a5b0c9ed375fe53ed2dd67eca6881fecd809cf277042cb8e202c1c\"}, \"id\": \"45329167-5289-441f-abe1-3dd1ee81fd47\", \"version\": 3}",
		Secret: "c9qhrg47uky2ze1ded",
	},
	{
		Key:    "{\"address\": \"2c6e141e267a8339b982be8b6d180822cc3fe00c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dac176d10eba96706900f33da8eaced4\"}, \"ciphertext\": \"42760d30fa58953f0401544dfe3b6374efb809aa49efd539ebe47f87bd8ef49e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8cf20819f0b7c70e9575ddd2f4a137cd\"}, \"mac\": \"225fd26aa5f4d1563da57f3b500dccb28f582ad277c941040991b4ba5b917cfc\"}, \"id\": \"76e6dcd3-57a5-41ee-8a71-84e959c20d5c\", \"version\": 3}",
		Secret: "c4q99qtcna6uq27hv2",
	},
	{
		Key:    "{\"address\": \"01d9455666b97aad948cedfa1dfdc7321f461038\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c26199996b7f118b737ee02d52d35b71\"}, \"ciphertext\": \"45a12396239e94fefd3964d5535f9c665391b9432b59b3929a3b242c5b997809\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"81f6b6e026e8937453672792a8b47363\"}, \"mac\": \"12d99e26f407f45dac1f988afad00cd0fd14e193a405865942e0778325ec0401\"}, \"id\": \"7ce812bc-5267-4494-ba8c-51b71d1082eb\", \"version\": 3}",
		Secret: "jswxysy7ssmqjrr4b0",
	},
}

var validatorKS = []*KSItem{
	{
		Key:    "{\"address\": \"f9e62c77661e0d5ddd9ab46095d15e4498247215\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"d9a390ad5c51fbda9763fd19bc1df0e2\"}, \"ciphertext\": \"0f6c561b43a00504a84069b707cdca64c35787439dbb07d4aeb46ac3f436082d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b2574c5508a2a64473d733c2e7fc5809\"}, \"mac\": \"3eaa3d39ce78e89f9e0a5a950bc51faaed3b5b5b9c0309d245a00361dfccd6bf\"}, \"id\": \"708cbab4-4fc2-4941-80e4-415425fffd99\", \"version\": 3}",
		Secret: "gbfbnnpprpxy88zu46",
	},
	{
		Key:    "{\"address\": \"a2a7e1b25e0951d68579a9e3aec048331e66f188\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dfe99f9bd671ff60290da4e97162cf29\"}, \"ciphertext\": \"932e311bfc49854f39500fdf764e987f599a6344ee23019267d26ee67fb15b23\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6604b60f28fea30685593d50a909af8a\"}, \"mac\": \"1b9bddace5de592a547691274788a14f106b6ad4ea019c772b1e7e8d15849109\"}, \"id\": \"1157e194-a024-4467-9e2b-21a7d5d890ff\", \"version\": 3}",
		Secret: "b2uvgf3g3ub0hdfcal",
	},
	{
		Key:    "{\"address\": \"0a1d575ca0406f68217a7ff843a7ff28b32265a5\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"857c38c0582333922bb615b63501a082\"}, \"ciphertext\": \"b8c15e502597c2a69cb15bc9aa0c392bb1412f23731308be6535b3bec7f80737\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d9e3f1a34f41b41fb938d22ef533de63\"}, \"mac\": \"1a84d0fd8f1623b2118df3921a759664e5c34d5702d2264dd5f09b11faa961ae\"}, \"id\": \"997916c5-4fc7-4580-8706-d48510643054\", \"version\": 3}",
		Secret: "6rkt434fqedlu6qhhd",
	},
	{
		Key:    "{\"address\": \"46864e34e4055468f0280eb9546e29d156cbb4dc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"19dff932d6f0c5c118ac7520c4586687\"}, \"ciphertext\": \"16afc8e239d45b3fc129dd7d38f2f52a722a373cde23db1ad7864c2d3ff3168c\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3f4952d46b094448fc4c882c9b75c14a\"}, \"mac\": \"9842e19bd3761a8662172e240c09a47c96a1bbcde8f0e6386333618922c24ca9\"}, \"id\": \"9ba1d580-5693-4c8b-8e2e-62175a45bcd5\", \"version\": 3}",
		Secret: "xcm0xw4w4lk0pzxdio",
	},
	{
		Key:    "{\"address\": \"433212cc84cbdd061a0d7ea7deded453263b2d21\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"91237799d72536aa8cac39e1252c8fa6\"}, \"ciphertext\": \"44d7b1a935ccb5e90ae411c3648b7cb4a668ce0428f6de168ddbad6b40fc49e9\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"965a9f661e6fe78a5b43c50352204f34\"}, \"mac\": \"7f027267541fcdb9d4b73769655f9b46d0b3e37135f2ec01b0904e6515fd22d6\"}, \"id\": \"0b5d8906-763d-41d8-acf5-56efe828cb12\", \"version\": 3}",
		Secret: "gyvusy23xo5cn20ztu",
	},
	{
		Key:    "{\"address\": \"dfd7c415c4f7f1ca0ff4a82bfe7018b14653a277\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8a1fe2d366462a6166bd1e31cfabda11\"}, \"ciphertext\": \"6d24f2c0965b80596a7b2306a48771bbdfdcb9ab45ce2d9897cc8899a8dbafda\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c76ec75a669015a653f37c00770836cd\"}, \"mac\": \"6980cb742a4942b9cc9d178d51bd2ffa30560d6ea5025631205d97d43539987b\"}, \"id\": \"53060963-9267-4c5e-9ad4-3276bae0f21d\", \"version\": 3}",
		Secret: "gmq98u8huu0wouv7jd",
	},
	{
		Key:    "{\"address\": \"3067de8466f412f5c6a3c97b8ded1a9e6dd2da81\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e9e3a4e15506ba7a78f391df601f57b4\"}, \"ciphertext\": \"e0475f31b2f2ed1b7717634ee7ea3ea3d17aaa883f755873d9b16dc43d80a4bf\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1cc727fbad21bcd0b37b55b98f4affdd\"}, \"mac\": \"2f4c43fec1f37990d3caf1613994e36a17a18609fd590f12f065959460150bcc\"}, \"id\": \"bf076349-7d79-4ca2-a121-69e25d43c79d\", \"version\": 3}",
		Secret: "h8i905b49qpqwona8x",
	},
	{
		Key:    "{\"address\": \"da615666b9044896caa5a2837e9ce236fcfcbcac\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dad8522327b47aee1e3df934d00be6c4\"}, \"ciphertext\": \"01a1163adee5760c3f5c103ae30e967946ea0c9ad0d5f3cff77022f85b036ecd\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"6f414f123203098ef6b0c27fadf80c95\"}, \"mac\": \"481af55d6d38f0cdd99a1fcb18457017ac16845390b4b41f533b04bf2efdf910\"}, \"id\": \"dfd24ff8-06e4-4f28-a4f2-70b46c211c42\", \"version\": 3}",
		Secret: "myzbcn5uiu00g9wvqg",
	},
	{
		Key:    "{\"address\": \"915475c44e6783700adbff38f61c518653366c23\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"02e34dcde06c8d82954e69b1a51763a0\"}, \"ciphertext\": \"4d4bfdfc9af6c67e4b6a37a0fd99e1be1db42a9cf5646616713282a9098a65c2\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"d56a97ad68ed50bbdd6e2c2aa9481c86\"}, \"mac\": \"80d14eefc127069aa7106913ea8715a1ee44806e36f1184eda1e6761f7abb864\"}, \"id\": \"49308107-fc9e-42dd-876f-43cb8ddaf4d7\", \"version\": 3}",
		Secret: "8rhve7qgp6i0zrx3qn",
	},
	{
		Key:    "{\"address\": \"95b903d686d1231e015c91429a100afb77ef8606\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"731802130270d4d0363f185464a7b1c8\"}, \"ciphertext\": \"f7dc630f3a4e45852845332f73361ebcc409e4f24ba62521a615771793b13a37\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"4d086ce3f1bb6b7c6d47b8a33689b004\"}, \"mac\": \"efc99418bfc4aff969eccdea6602c56e1e2064dd0b0bd14469e32e3cb2445d0f\"}, \"id\": \"e1970ff4-cd2b-4563-a86f-d1b3d43f0a99\", \"version\": 3}",
		Secret: "1w4o6kv4a5uedyivrn",
	},
	{
		Key:    "{\"address\": \"4ffd0d6f7a84f02d4f5d00d1b96fc88358f819b9\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0525548fc9f54e72e7e44c4433f70d3b\"}, \"ciphertext\": \"aa8f122442960be1ac4c11e8900e15e58561def9b918a8d8b92b0ca28b7d790a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a3d2d1325df82e432620d5ecf5136ea1\"}, \"mac\": \"75d79bea187f753c6e85b268a57fe4e5eef715bd939829f96bcc9b7fcc44010d\"}, \"id\": \"e1469277-19cf-4d4e-b65f-2ea3d73c5491\", \"version\": 3}",
		Secret: "ybibc0crp38lh2c9zo",
	},
	{
		Key:    "{\"address\": \"6e2cd1cd0e77847a07805582cb1e8cac36cf5477\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f5d04c0725f4d7cad7b16458e51042c6\"}, \"ciphertext\": \"f146f32dd512a7647697d7ba342e387f79ae7cb7a62fee0f994c7303c01d97ff\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5c65eec2caaba726b111811cb507f653\"}, \"mac\": \"f81830a8bdcc7b582e0d62b09a5c059cff27cbf2f397a272d255298978d38ff5\"}, \"id\": \"98304275-a64d-42f0-a350-0b56c1f2966d\", \"version\": 3}",
		Secret: "djkh34otukcte2tcuy",
	},
	{
		Key:    "{\"address\": \"cdac636bcbaa22deebbefe28d82c3672104d2678\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c7f05df6d1f05d3eea76b4e396529dc0\"}, \"ciphertext\": \"9e6fcbe9b0b6b90509029efa0ae4be87a5b1f7f22158bdc3bf4e7666445bb885\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5ceb21a99b9c2615d00fb8db2e3c692e\"}, \"mac\": \"8344d33796ccb498ff8af8102507a2bd3ece11b9d02bdb0a05c708c8295754c4\"}, \"id\": \"ae00b020-5160-464a-9b5e-1f1c6c044a01\", \"version\": 3}",
		Secret: "f5abz3zgbaifuv8ma2",
	},
	{
		Key:    "{\"address\": \"c25d3755113dc364ca3840c9d33109dd2fd98a7d\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"23a389e4350f4db8d3afcd34d52167d0\"}, \"ciphertext\": \"0665633ab4c6ae235925807e33e6c158a79c0221a64c9b6df7b14e84226aa637\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5d80a4d33e716e8ba83d813ba6919c60\"}, \"mac\": \"2c4032a47e87b75be89698812573d8237ac5a810ca8fc1bce863c9821e675c7b\"}, \"id\": \"715f92f8-8603-4c02-bec8-f7e30200e463\", \"version\": 3}",
		Secret: "ivyubykys7xfsg3wm1",
	},
	{
		Key:    "{\"address\": \"b0b9752f08ff02f826b8b85a7ef649ebd28757ea\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"081ce676ad261fd81357018ef56a2f58\"}, \"ciphertext\": \"c613f7d76898b4566db398ec34ba610470043444a7a21033007e7d724a5d22c0\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"212e0773b6cf561b49d3dbdc729d8fa5\"}, \"mac\": \"99214be4ef082686320d0418da7bb85750c7decae0a17a8a9c92a0740593aee9\"}, \"id\": \"5e759393-3989-4709-a81c-1ace26e8b561\", \"version\": 3}",
		Secret: "einx32ih95e6c0zztn",
	},
	{
		Key:    "{\"address\": \"d2b1b1b0fa83611fe361510c90e24703dbd28f97\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"bd3589cba4061302f31a3afec7133874\"}, \"ciphertext\": \"2d2f3bdbb1bcc5df10aa89750a82a7c188d99f739d118bf4c3928a011a8478b4\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1b49dd5264090f3b5f93dd99b4227516\"}, \"mac\": \"5a1e9b04a06aacf69515dd3c190a256a66c1192d892515b733e2eaafae66d8a9\"}, \"id\": \"82fd09df-b231-47b5-8db3-509674cebade\", \"version\": 3}",
		Secret: "s9zmk0heabmbtsrsmm",
	},
	{
		Key:    "{\"address\": \"0b37bedcf653329af04cadd7aa53b1b088c24d8e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3d573c128ecb0e7120aa667310d2d761\"}, \"ciphertext\": \"7e03c7d63b62caff4f859701b360e137d2aec427d37e45f23e315d236507e683\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c249cbe24345c543e096b39a96b5cc18\"}, \"mac\": \"e45c334ac6d16fde82f06e4a44049ff7c91abc72badb8ce313830db83f0292be\"}, \"id\": \"e0467798-420d-490f-be20-9b73fe4669b6\", \"version\": 3}",
		Secret: "66yicjtje8450wfqm1",
	},
	{
		Key:    "{\"address\": \"056fc4ea250ff00d27aab9fc11436f944971b73c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1b02ba4a52bc1455ddc1296aefce8c22\"}, \"ciphertext\": \"2fb8d5d696cab667a7693de573d21b0830bc447695b6653b9ae9b256d31a0d40\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"78d6d8f8b3c84929533454b746a9dcc4\"}, \"mac\": \"ebecb4127d323a3c3f2ca788f6a157ed1de9d14e4db63c1629163de4fc33cc8b\"}, \"id\": \"fbb24328-0016-4b89-8afd-ffb67f6df551\", \"version\": 3}",
		Secret: "yvr01bjlnxi482qs7d",
	},
	{
		Key:    "{\"address\": \"daffee058b6aa0899341afe8acb1c43789e7db30\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2f121e85b75185cb1f79af8608f1b415\"}, \"ciphertext\": \"5b4d6991828c5e5e281fba3f2c40947461a9a9598dde269df359d70a20d72e55\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7d3299c804fdc3a13861bcb4a6fdf3a8\"}, \"mac\": \"dbe4e95bcd349dadacbe24f58949001262d369d1555d4b1163c20c33a1819a6d\"}, \"id\": \"9ed884b5-a3c2-4b43-a829-8cdf15e9cace\", \"version\": 3}",
		Secret: "x4tmpdsry6d3i23jzs",
	},
	{
		Key:    "{\"address\": \"476512f810dfae79bdf680422f01582c197c8e29\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"836dc6b3aaf094d659e727b2861dfb0c\"}, \"ciphertext\": \"0be8a30d94625e144dc99e2188dd56a9c942961b243cd401f8e5f9206a71a4ed\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a1f664dcd9098a5b11e8a7df6f089574\"}, \"mac\": \"de8548445b0a21e7a11710c2030b2d72afcef3667c694b39c0d0bebb8858f2c7\"}, \"id\": \"fb6351b8-87ce-4e19-a2f5-511143ae1026\", \"version\": 3}",
		Secret: "orgq6xx886q6cu5ses",
	},
	{
		Key:    "{\"address\": \"a8e002fc621a367d3e70174405c8e3141fd41090\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"1760255ad944cd83e1466332ce0b838c\"}, \"ciphertext\": \"e658f06428e0f81474b07a4422f539b29d00be19c7314741c64ee2cfb7c3d64a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"c5b8040d914a77cf4498fa8480c721ed\"}, \"mac\": \"5c448ea08a159ac6c37deaf3a2b27b22ab3c20b66c09587560d8b31a59c46b9d\"}, \"id\": \"931dd177-8e1d-407e-bde8-ff7fe5a457de\", \"version\": 3}",
		Secret: "zq5dti4injqcm0ubcu",
	},
	{
		Key:    "{\"address\": \"bafe9bcf159a539c84327e74c177f08a739a4949\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"18fb58deffa4914fb1d18c7f14a694b6\"}, \"ciphertext\": \"c498ae17b601780194125aec209e175310d8ee79510c2b49c4c0268eb7fcf949\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"44daf287466bbb41db3ca93364a2e6b3\"}, \"mac\": \"3c8359094a169fd2d6f9b9551cbfbdb0f18f1f6510ca1211e1cc884e97b49eb3\"}, \"id\": \"09b0a80e-d706-407f-9755-412ded64b8fb\", \"version\": 3}",
		Secret: "rveviqt0seedaf7k8w",
	},
	{
		Key:    "{\"address\": \"793aa071cb2eb0a85a48f45d40cb75f15ccd477b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"e3fc81c2a84a1a251d98e883a1ffb097\"}, \"ciphertext\": \"b590d3e7c62b44cbc91e81a16f48820a297936580f6d8adf18ea307aeefeb14a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b320348152fa7f631933d95096848319\"}, \"mac\": \"f90e9e058b2a066963fd1b3437e5d4ee2f5b43ab4df51791653b1d89b3abe7c4\"}, \"id\": \"34460956-f14e-4b11-8481-b465312248f3\", \"version\": 3}",
		Secret: "mnrdklye8k18l0qdr1",
	},
	{
		Key:    "{\"address\": \"dda97882cdbc9190f1452a8141cfc91a7e98a26b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2579ba1dbea9afd74f1be8a3a72ca532\"}, \"ciphertext\": \"e6745a64455b78c3ee433dfab3ea5956918fce91338ac2c6a5c003c1c427bd58\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"f061b4e4fde8f5e5d3ffd50f49a83c2b\"}, \"mac\": \"9e76c2ad709004cc350f3b366a38d80439f8fd38f3cab01bef95514bc66027dd\"}, \"id\": \"90fda6c2-e2e7-48d0-b8c7-f26d85f1d377\", \"version\": 3}",
		Secret: "x0z9u0hqflu4amior8",
	},
	{
		Key:    "{\"address\": \"fd8009ad756f7b970a49509aac529f28e27483ad\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6c92349b75e2015548ba3075919d0045\"}, \"ciphertext\": \"eb8112e8cd388314a06b42120a7370a1207ffa1ba6c154aaecb37eeb3f27414b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"e620468b52ae308bffc30f323cd85940\"}, \"mac\": \"91a153b866b600871fdd5082a6fce282818199fe569cad342cfa0e4a0322780e\"}, \"id\": \"b83b9d79-67d8-4d3a-84ce-b89b209c7b7c\", \"version\": 3}",
		Secret: "oi4roxwjyhl4x74wsc",
	},
	{
		Key:    "{\"address\": \"2c69831b9d6dae93b350381bcd0843cba2f3608b\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c06df82933409b991bf1b79c6557d227\"}, \"ciphertext\": \"277d591696bf0e4ef364557ccaa35a6299aad0e1df0c35fc1399a2816c4aa3a3\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"bad4b0a007a30f120fc88fefd1f3a9a0\"}, \"mac\": \"2649c44106473d3f4a3ad5d957fecb8b18d8e7693ba133783019014dd9a26f4a\"}, \"id\": \"8eff3ef1-d35a-4667-a9bf-7329bee2f976\", \"version\": 3}",
		Secret: "81im4e2f2emx0vq98x",
	},
	{
		Key:    "{\"address\": \"7d5e0aca0bf92d4423c23bd693b1dac03cf54fe1\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9e148131f7b87ba0e6d5c76986765427\"}, \"ciphertext\": \"3e9752f7ff8beab13749f44a65df10c92b982862ef60df567da370e20f5a6264\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"cfdcee96805885c8f5ad99e1ad97a009\"}, \"mac\": \"8ea29310c57b7c12ed248e59c8a7e48e83e75c4df307894d862651cbd7b434ba\"}, \"id\": \"d17abcb1-97d1-4ab6-b75b-d7cba83b6b56\", \"version\": 3}",
		Secret: "yyvh9uzve85p9hqnvo",
	},
	{
		Key:    "{\"address\": \"b3c8521788c0db0583183095521ef7159a1b56cc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c9cf9963dd0654d02e5ee626f7910089\"}, \"ciphertext\": \"0d939d288757e320e1a577dd064f42e21fe495dd7f7a0e4d4118212c95acacca\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"7a791b81337cafbdc20eaa7b19996825\"}, \"mac\": \"1c822a887d9bee1b05516b2113515d0acde669d1e735e0a03fd4ddcad5cad8ef\"}, \"id\": \"90539363-3429-4531-a51d-05d44a7e2a7f\", \"version\": 3}",
		Secret: "sidip6eoux6ajnnc4d",
	},
	{
		Key:    "{\"address\": \"7eb8f6a9536812a26cf84786cea2f345fb82b57a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"656549b1c4d1197473d1461cb846110d\"}, \"ciphertext\": \"0a1f10650147dc8468c96b68dbcc05e702f4b981257ac2da8785769d2ab91590\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"475352f2852bad46c3f512158aca4f48\"}, \"mac\": \"723d4bc1f95785290b6b4a9b753fd0c0e4ae8f5089f23f5247dc3b910e101c51\"}, \"id\": \"d5b583bc-8606-49e0-9490-d5269ccad5de\", \"version\": 3}",
		Secret: "cjg27rgz6um7wyf133",
	},
	{
		Key:    "{\"address\": \"105efc788ce0c5cb8f675d017dd71662e5c091c2\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"60611264f57a42b0a063ed15897565be\"}, \"ciphertext\": \"dba53c46526fc9d6642558864f0a51fbbc9c52755a53b4ef9acf684d18d5a795\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5e555fcaa67e0f8bf98c86315c257161\"}, \"mac\": \"70a33970446b7f8189b2085fd7c109938ce395d9bf3a709a6a9aa4d6280c28a0\"}, \"id\": \"f561630e-c041-4e77-bbb9-3a79927909ca\", \"version\": 3}",
		Secret: "1ndnh9pxantqa3xit2",
	},
	{
		Key:    "{\"address\": \"3312f67bf04441afd42e15d8432410c53fad02ee\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dea36b7a0abf6e4bb5fa91f8afea9aac\"}, \"ciphertext\": \"52544e34dd5d3eab9a5dd7e45c7f5cb37a7448e741a7693c18f053c850e6ef85\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9c2971699c619d45829ddb600806c2da\"}, \"mac\": \"99c4b669db259b03853470c105c45327f54d2053eebb88e5caf316e3d6921ca5\"}, \"id\": \"ee0c2cf6-ee41-4b4c-bf35-7953e201ba55\", \"version\": 3}",
		Secret: "00giqatom5obn5szvu",
	},
	{
		Key:    "{\"address\": \"35853d02c44aee2e3972ee07b9303c4bbdd52de0\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"6b10d1350dd41f748d7dd1cf982094b0\"}, \"ciphertext\": \"40ffd977462b945298b91a72c2d586e06cf2c0d0e6235caa74752ee5c07dbf9b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"3139afa9cb02b56aa1b6e82d6a2e8c34\"}, \"mac\": \"953662a3689a71250e4b1d9aa08a926c051b15d316ab65987b47506093e70d2d\"}, \"id\": \"c9df1c05-637d-42cc-b3b6-26f074308554\", \"version\": 3}",
		Secret: "4eouzqxqt2o44xdvjn",
	},
	{
		Key:    "{\"address\": \"966bd67135ef70de2255f7d43c8ecc3c76a03cf0\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"61a8ac1b7baa3dfcc3c8b7da3f6bfab9\"}, \"ciphertext\": \"25800a484e8d68ea1556627ffe6fde674417827d10ffb82a39bf3c1daea261fe\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"773c319e747ab605ba62331ee9a78e42\"}, \"mac\": \"1f3a665d77cea598afeada9c142df6912c8e698966f1c08e1b7463e328f34a16\"}, \"id\": \"bfb2acba-a8bb-4f19-bd4d-515877354ab0\", \"version\": 3}",
		Secret: "arxojju0redu4w5bgw",
	},
	{
		Key:    "{\"address\": \"07bc938ef7e905e410415532122c4e70ed9383fb\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c82f610b8fabcb675f87f6d3c5631181\"}, \"ciphertext\": \"97cb46d6b728d19a0f50456bee053ce4b56b658ee927c3bfc4f441564596c2ee\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"07518ed43e4a21235d71971f2e9829a6\"}, \"mac\": \"6b783be16e1834e466947e55c68d8bc50d61bc1bd9c4f438702003381e93787f\"}, \"id\": \"dbd8d085-099c-427b-8c2b-da172d46b492\", \"version\": 3}",
		Secret: "sfcfs2x96x7s3ipu6e",
	},
	{
		Key:    "{\"address\": \"96cab71c5747c62a894e7be65b7a2a50991bb986\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"97fd4d4762f96cf29d44674b7e9806b3\"}, \"ciphertext\": \"b5cd6f0d6bd0d0b3f80fd9826a2820394808e0b1a95d4f6db792a629946bd372\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"86b7597cc589a34c44195dd5b7448472\"}, \"mac\": \"7de3c05cf563650b900769a9c22b09f3b1790b12b972c4bd27bf954257211957\"}, \"id\": \"c2a4e7b6-c22a-4f7c-82a2-f55018d0299d\", \"version\": 3}",
		Secret: "5jf8jpztw95pbyvzen",
	},
	{
		Key:    "{\"address\": \"00bc9ee928ac0152fd8b81be84e46e8e22c6c62c\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"ab30db8714c8fc6ec390b7c33d7d5033\"}, \"ciphertext\": \"fd6e0d862f3c2b3dba2d18b6b075529801aa282c6cbb6862053a9580a1e1ba13\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"9cc99c72b267bd478d66d5079167148f\"}, \"mac\": \"371cd154327a3aa3b884725232a4d1a7372255143ba8ff0162b1d8a30b592653\"}, \"id\": \"8cda2654-708c-4972-8d4c-dfd4ed121946\", \"version\": 3}",
		Secret: "bu2r91zmbozgj1imzy",
	},
	{
		Key:    "{\"address\": \"1282cf5f24420cb0dea49b898f4de973d45e1323\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"cf138af70022e0daf248148c3de2a6c2\"}, \"ciphertext\": \"a3d1f84f3f98039e02bbe0fba0ca253d842414deb1627f1eb4f9bb4ca790900d\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"21ac7468fe2870281a96b30bb71aa7f1\"}, \"mac\": \"8078710e857b25b65fe14a04cd2106dad6d06c979be2a8327c68ad202e5fe2cf\"}, \"id\": \"f0505f02-c07a-4768-975b-8394d531da50\", \"version\": 3}",
		Secret: "9edw6i30ql1myosdcn",
	},
	{
		Key:    "{\"address\": \"b83f55c3053505a91b13cb1e14f9f32e3ef2e8af\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"86a88171c7be8cd5651872ee0bffc70f\"}, \"ciphertext\": \"776d06454f218173365c455c42c1a81806230d7885bbc9b875b2ceaaecd58e54\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"80efdb30e131248bdf372c70ddd303cd\"}, \"mac\": \"8df8efe0faa0d89fdf348ea59d457280ff74fa83856cee8d9daec8371bfc75d8\"}, \"id\": \"db29fbf0-6e9a-4486-b36b-5eb8ca8c117a\", \"version\": 3}",
		Secret: "8tucjuo3qmxss1koam",
	},
	{
		Key:    "{\"address\": \"388b3ddefc1e47189e506af3cdf5fe3e4ae2cae7\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"f618bd203de5db4718c4c15e5b021f58\"}, \"ciphertext\": \"c69de923ed3c632f935d9c6d8d8a1436dbd1360b8eff96bc6b05e1a490e0e955\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5e6580c2dcb3cfff5f96cbb09eff18fc\"}, \"mac\": \"1c1527d7329ed4c16a5b02bd06920f6fb6ff4bf15b8213d610509dafa9fb4b4f\"}, \"id\": \"b6fc2011-6602-4e19-a058-1ba0e43fee11\", \"version\": 3}",
		Secret: "oxsjbiee7yj23safwv",
	},
	{
		Key:    "{\"address\": \"8a2ee8814e803832406117d99a55bc6e1fe103cf\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"c48f4cfee008b06b9360a967829160ca\"}, \"ciphertext\": \"a014d2490e16d3ddf8d72c1d651f90629d89eda6835ec51eff7e47da9dfdcd0b\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"5311e4a5010ebead305ec069b706d347\"}, \"mac\": \"45a9418a0c2b1721653e6c7c7ddcba56caef6fe4b0127b4f55cf5447fd350ced\"}, \"id\": \"68bc83c9-a054-4210-9fef-a6df7d79f321\", \"version\": 3}",
		Secret: "u5prbxnlaee9x0x667",
	},
	{
		Key:    "{\"address\": \"97d796284e12c8aa5813d57b2fc8dd6a70aad85a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2fb4aa3237b832d04886773f00cdf1c8\"}, \"ciphertext\": \"0f8783de75b86cb01ee54bd21bbe0ea5004b800dac491bf8fd7177ce67f8f38e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8d31a04fb0c277ca3ee98a72b2c762d5\"}, \"mac\": \"f8634f0c3c398978d388a96e67ed687ef1c40a1d757a453602ea045d93f578b9\"}, \"id\": \"b1ff7f4f-76e4-444a-aed0-c2302e53431c\", \"version\": 3}",
		Secret: "qz569frhkfm1xkqiuw",
	},
	{
		Key:    "{\"address\": \"074c15fae527bd9c1d96382f01b71c6128f2767e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"3f2792fe8eee528d6477ed0666ec87b3\"}, \"ciphertext\": \"6d0cf06ef2c55e2cdf1a39a5c71bf26886902f11ee4e00f03c584dc2fbf20c3a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"0ef77b211099af8eef3fd89be1f66c05\"}, \"mac\": \"933e9e8f967152d9c18a55419d50faafa561b44bab770408faaa2a87e03b1c91\"}, \"id\": \"9813f0e8-ddc8-46f0-aa0f-829a3fbaed24\", \"version\": 3}",
		Secret: "vzpjhxf9r0i2b370f2",
	},
	{
		Key:    "{\"address\": \"14ff15a7f17f11712ebbbece911531373363a1dc\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9f257ec71bd8f9aaec51df2ea8f4b8f6\"}, \"ciphertext\": \"dc8f060b26e92e08c103873c034555a9dcdd64bf464ba74de513e981c3f0b56a\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"8e3e7464908d79f9fdb940235bc5ce2c\"}, \"mac\": \"b76e74a9cfffca3f6075577a3c2339366775ccbea89e57e81241c04c05fddc19\"}, \"id\": \"64a28282-e651-41a8-a33c-532e1370874b\", \"version\": 3}",
		Secret: "afvbtpq8twdbojg7rn",
	},
	{
		Key:    "{\"address\": \"790bf9a9441d8fbf93eaf135e0342ce07d122756\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"2933bf6da44948c7aced7ecfe3418c18\"}, \"ciphertext\": \"4d1a97a8f3ccdd8847f6e75b0b6216153836b87769fafd37da9bebda95195880\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"1438ac2f6fc06ad77ea3772543c19ba8\"}, \"mac\": \"a71fd6a6686719ff04b9c2afd57dab2b76b2ed8de9a91acc9c84c4725ed05501\"}, \"id\": \"aafd50a3-6c38-40f8-a034-a9d662b2c592\", \"version\": 3}",
		Secret: "b98btth8oufld23pjc",
	},
	{
		Key:    "{\"address\": \"3fdcfa3b5395efdbb5b5204efcde0e3c200f98e3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"a1e791ba108875235216d12c89249958\"}, \"ciphertext\": \"440a501cec4d1aeb00d516dcc3a04594bd6921959165687879a8c0a965ec7b02\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"b1f8657616ef201fead5487c2d4dddfa\"}, \"mac\": \"a9a2afd5f150934ced8aa5960c02a4d79bdbc0c436c1e987bbd5cbfe30173e0d\"}, \"id\": \"5ecea492-1b6a-44e1-ac2a-c43395c93b94\", \"version\": 3}",
		Secret: "1eee753hw8zsvif8m0",
	},
	{
		Key:    "{\"address\": \"e2d5b7a588e38493df05c7e4f668a322060c750a\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"8900267b14b6c833b5d575b4d90754ad\"}, \"ciphertext\": \"f5e7533323350546274f2bc3ba7a46c9991a8f79a34bcf96b7ec644c9f5dc564\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"df8a930ab0e2ac259bfdfd0510ddbaf9\"}, \"mac\": \"85521fc95cd3f8a904141abc08a8c8fc38538b91ee106437cc6ac1fe3ae188bf\"}, \"id\": \"4d6d9126-8fb4-41d3-8540-1e98d8852f74\", \"version\": 3}",
		Secret: "bitjbs8ymt8zk4yjkw",
	},
	{
		Key:    "{\"address\": \"662964395ab575229b5fd676338e6b75c3f68131\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"dc0c2dd6ce65da2f473f5a9c96927dac\"}, \"ciphertext\": \"0586bac9d754491f030caa5982118d4a6494f663a5aaf41566279d983f808dfc\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"00fb437630f2a785ce712168249cc58a\"}, \"mac\": \"cd4475785d0e1924ae0d4d62379c896475dc30d9b82d7a21ca35b53402146aa9\"}, \"id\": \"359c0e40-8172-491c-af83-d2c90146176d\", \"version\": 3}",
		Secret: "fb64a46e3v7xnocnw8",
	},
	{
		Key:    "{\"address\": \"27b153518d6d50cbce9125bf3468def5636ba62e\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"9364b6f77b5b5f5f41302782bbc190e6\"}, \"ciphertext\": \"d491160bf75c18dc67256f608d0075bf70f35ce102ace02ce535b027ba023033\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"a7faf8a1b8a8c86257c2412df8aa07ea\"}, \"mac\": \"18aeb700446b86aa1689f46c3d45453cc9938cec613fd4c1797de9e8d86c086e\"}, \"id\": \"a1f916bd-20da-4226-8291-4ed680646824\", \"version\": 3}",
		Secret: "bz2noz4q0bx6toxk8m",
	},
	{
		Key:    "{\"address\": \"d4de36d1ca8d53262b0244953f00bfff3afe39b3\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"b3d5d6132e9458b2eb48f4a68780d397\"}, \"ciphertext\": \"bebff2337e074725a7aba984d95b3bec3596d057898d110b134d565e15ced9ab\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"dc9369cb09b9de6f265dd34f6145a173\"}, \"mac\": \"da578872b0e32f1afbe15eb1ad1fda78668acab31ce390a0b84aa3d625f57ba9\"}, \"id\": \"c160685b-f7d7-4f9e-a9da-f8e1ea7dd8f7\", \"version\": 3}",
		Secret: "5tuhu3abqmur8qfnky",
	},
	{
		Key:    "{\"address\": \"6c788ceb01176e49a50a0d80ee4bb61e8dd054c4\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"0225391e4a681aa5fb24324bf7d8af75\"}, \"ciphertext\": \"0fc2a5631d47db60e0676be5701b640c80e8fd76c210e6cb1c69d512f48faf2e\", \"kdf\": \"scrypt\", \"kdfparams\": {\"dklen\": 32, \"n\": 262144, \"r\": 1, \"p\": 8, \"salt\": \"ad0fb294e6ad2c4ce6dd7e59b4aa9fa9\"}, \"mac\": \"9ae048f8adf0f5058d2a59fc25e549ce61094fc7f7e53915a6c7e327a9e0dd9d\"}, \"id\": \"b900403b-6dec-41f0-861d-6bce0bb49240\", \"version\": 3}",
		Secret: "83xgmjlver4ve0i3vm",
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
