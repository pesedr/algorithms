console.log("string has all unique chars");

const unique = "abcdef";
const notUnique = "aabbcc";

function isUnique(input) {
    let obj = {};

    for (let i=0; i<input.length; i++) {
        if (obj[input[i]]) {
            return false;
        }
        obj[input[i]] = true;
    }
    return true;
}

function test(input, expected) {
    if (isUnique(input) === expected) {
        console.log("pass");
    } else {
        console.log("fail");
    }
}

test(unique, true);
test(notUnique, false);