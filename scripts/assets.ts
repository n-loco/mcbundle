import fs from 'node:fs';
import path from 'node:path';
import child_process from 'child_process'
import { version } from '../package.json';

const rootDir = path.resolve(path.join(import.meta.dirname, '..'));
const sourceAssetsDir = path.join(rootDir, 'assets');
const goAssetsDir = path.join(rootDir, 'internal', 'assets');

const assetsToWrite: Map<string, string> = new Map();
 
function goSource(varName: string, data: string | Buffer): string {
    const innerCode = (() => {
        if (data instanceof Buffer) {
            return `var ${varName} = [${data.length}]byte{
${(() => {
    let fi = 0, li = 16;
    const arrayLines: string[] = [];
    while (fi < data.length) {
        if (li > data.length) {
            li = data.length;
        }

        const dataSlice = data.subarray(fi, li);
        const line = (() => {
            const numbers: string[] = [];
            for (const n of dataSlice)
                numbers.push('0x' + n.toString(16).padStart(2, '0'));
            return numbers.join(', ');
        })();
        arrayLines.push(line);

        fi = li;
        li += 16;
    }
    return arrayLines.join(',\n');
})()},
}`;
        } else if (typeof data == 'string') {
            return `var ${varName} = \`${data.trim()}\``;
        }
    })();

    return `// generated code
package assets

${innerCode}
`;
}

function loadAssets() {
    assetsToWrite.set('program_version.txt.go', goSource('ProgramVersion', version));

    const dirs = fs.readdirSync(sourceAssetsDir);
    for (const file of dirs) {
        const filePath = path.join(sourceAssetsDir, file);
        const { name: fileName, ext: fileExt } = path.parse(file);
        const varName = fileName.split('_').map(w => w[0].toUpperCase() + w.slice(1)).join('');
        const fileData = fs.readFileSync(filePath, { encoding: (fileExt == '.txt') ? 'utf-8' : null });

        assetsToWrite.set(file + '.go', goSource(varName, fileData));
    }
}

function writeAssets() {
    fs.rmSync(goAssetsDir, { recursive: true, force: true });
    fs.mkdirSync(goAssetsDir);

    for (const [k, v] of assetsToWrite) {
        const filePath = path.join(goAssetsDir, k);
        fs.writeFileSync(filePath, v, 'utf-8');
    }
}

function goFMT() {
    child_process.spawnSync('go', [
        'fmt',
        './internal/assets/...'
    ], { cwd: rootDir });
}

// main
(() => {
    loadAssets();
    writeAssets();
    goFMT();
})();
