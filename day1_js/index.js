const fs = require('fs')

function main() {
    var numArray = [];
    fs.readFile('input.txt', 'utf8', (err, data) => {
        if (err) {
            console.error(err);
            return;
        }

        var totalCal = 0;
        for (str of data.split("\n")) {
            if (str.length == 0) {
                numArray.push(totalCal);
                totalCal = 0;
                continue;
            }

            var num = parseInt(str);
            totalCal += num;
        }

        // Part one
        numArray = numArray.sort((n1, n2) => n2 - n1);
        console.log("Part one: ", numArray[0]);

        // Part two
        var topThreeCal = 0;
        for (cal of numArray.slice(0, 3)) {
            topThreeCal += cal;
        }
        console.log("Part two: ", topThreeCal);
    });

}

if (require.main == module) {
    main()
}
