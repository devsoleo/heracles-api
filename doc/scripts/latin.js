const fs = require('fs')
const iconv = require('iconv-lite')

function convertirUtf8VersLatin1(inputFile, outputFile) {
    fs.readFile(inputFile, 'utf8', (err, data) => {
        if (err) {
            console.error(`Erreur lors de la lecture du fichier : ${err.message}`)
            return
        }

        const bufferLatin1 = iconv.encode(data, 'latin1')

        fs.writeFile(outputFile, bufferLatin1, (err) => {
            if (err) {
                console.error(`Erreur lors de l'écriture du fichier : ${err.message}`)
                return
            }

            console.log(`Conversion réussie ! Fichier enregistré sous : ${outputFile}`)
        })
    })
}

convertirUtf8VersLatin1('input.sql', 'output_latin.sql')