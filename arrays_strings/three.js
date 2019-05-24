console.log("palindrome permutation");

test("tactcoa", true);
test("eat", false);
test("abab", true);

function sol(a) {
    const sorted = [...a].sort();
    let odds = 0;
    let counter = 1;
    for (let i=1;i<sorted.length;i++) {
        if (sorted[i-1] === sorted[i]) {
            counter++;
        } else {
            if (counter == 1 || counter%2 != 0) {
                odds++
            }
            if (odds > 1) {
                return false;
            }
            counter = 1;
        }
    }
    return true;
}

function test(val, expected) {
    if (sol(val) == expected) {
        console.log("pass");
        return;
    }
    console.log("fail");
}