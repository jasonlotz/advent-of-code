import { Md5 } from "ts-md5";

const SECRET_KEY = "bgvyzdsv";

const findHash = (key: string, prefix: string) => {
    let i = 0;
    while (true) {
        const hash = Md5.hashStr(`${key}${i}`);

        if (hash.startsWith(prefix)) {
            console.log(`Found: ${hash} @ ${i}`);
            return;
        }
        i++;
    }
};

findHash(SECRET_KEY, "00000");
findHash(SECRET_KEY, "000000");
